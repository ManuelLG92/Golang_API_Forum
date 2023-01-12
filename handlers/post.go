package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.com/forum/auth"
	"golang.com/forum/config"
	"golang.com/forum/helpers"
	"golang.com/forum/models"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostUpdateOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Posts []Post

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	post, err := helpers.DecodeBody[Post](r.Body, "Error")
	if err != nil {
		models.SendUnprocessableEntity(w, fmt.Sprintf("Error trying fit the body to struct. %v", err.Error()))
		return
	}

	createdPost, err := createPost(*userId, post.Title, post.Content)
	if err != nil {
		fmt.Println("Error trying to create a user: ", err)
		models.SendNoContent(w)
		return
	}
	fmt.Println("Post created")
	models.SendCreated(w, createdPost.Id)

}

// End CreatePost public function

// Init createPost private function
func createPost(userId string, title, content string) (*Post, error) {
	post, errorValidPost := newPost(userId, title, content)
	if errorValidPost != nil {
		return nil, errorValidPost
	}
	return post, nil

}

func newPost(userId string, title string, content string) (*Post, error) {

	post := &Post{UserId: userId, Title: title, Content: content}
	errNewPost := post.ValidPost()
	if errNewPost != nil {
		return nil, errNewPost
	}
	return post, nil

}

func (post *Post) editPost(data *PostUpdateOptions) (*Post, error) {
	post.Title = data.Title
	post.Content = data.Content
	post.UpdatedAt = time.Now().String()
	errNewPost := post.ValidPost()
	if errNewPost != nil {
		return nil, errNewPost
	}
	return post, nil

}

func (post *Post) ValidPost() error {
	if validPostData(post.Title, post.Content) != nil {
		return models.ErrorInvalidPost
	}
	return nil
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postId := getIdByUrl(r)
	if postId == "" {
		http.Error(w, "No valid post Id", http.StatusNotAcceptable)
		return
	}
	var userId string = *auth.GetUserIdFromContext(r.Context())
	post, err := helpers.DecodeBody[PostUpdateOptions](r.Body, "Error")
	if err != nil {
		models.SendUnprocessableEntity(w, "Body couldn't be parsed body to &PostForOptions")
	}
	postModel, err := getPostById(postId)
	if !Equals(userId, postModel.UserId) {
		http.Error(w, "You onyl can eddit your own posts.", http.StatusForbidden)
		return
	}
	updatedPost, errw := postModel.editPost(post)
	if errw != nil {
		http.Error(w, "No valid post Id", http.StatusNotAcceptable)
		return
	}
	if updatedPost != nil {
		fmt.Println("Error trying to Update post: ", updatedPost)
		http.Error(w, "Error trying to Update post.", http.StatusNoContent)
		return
	}
	fmt.Println("Post Updated")
	w.WriteHeader(http.StatusAccepted)

}

// End Update Post

func editPost() {

}

// Init DeletePost
func DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := getIdByUrl(r)
	//fmt.Println("post de param: ", postId)
	if postId == "" {
		http.Error(w, "No valid post Id for Delete", http.StatusNotAcceptable)
		return
	}
	userId := auth.GetUserIdFromContext(r.Context())
	if userId != nil {
		models.SendNotAuth(w)
		return
	}
	if *userId == postId {
		//http.Error(w, "You only can delete your own posts.", http.StatusForbidden) }
		deleteStatus := true
		if deleteStatus != false {
			http.Error(w, "Has been a mistake trying yo delete your post", http.StatusNotAcceptable)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

// End editPost private function

func GetPosts(w http.ResponseWriter, _ *http.Request) {
	allPosts := getPosts()
	if allPosts == nil {
		models.SendNoContent(w)
		return
	}

	models.SendData(w, allPosts)
}
func getPosts() *Posts {
	posts := Posts{}
	sql := "SElECT id, user_id, title, content,created_at, updated_at from POSTS"
	rowsPost, err := config.Query(sql)
	if err != nil {
		return nil
	}
	for rowsPost.Next() {
		post := Post{}
		err := rowsPost.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Could not catch any row")
		}
		//fmt.Println(post.CreatedAt)
		createdPostDate, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		createdPostDateTime := createdPostDate.Format("Jan 2 2006 at 15:04")
		//fmt.Println(createdPostDateTime)
		post.CreatedAt = createdPostDateTime
		//fmt.Println(post.CreatedAt)
		//fmt.Println("Post id pasado: ", post.UserId)

		posts = append(posts, post)
	}
	return &posts
}

func getNameAndLastNameUser(userId int) (string, string, string) {
	var username, lastname, country string
	//fmt.Println("Post id en la funcion: ", userId)
	sqlUsers := "SElECT name,last_name, country from USERS where id=?"
	rowUsernames, err := config.Query(sqlUsers, userId)
	if err != nil {
		fmt.Println("Error nil ", err)
		return "", "", ""
	}
	for rowUsernames.Next() {
		err := rowUsernames.Scan(&username, &lastname, &country)
		//fmt.Println("username a devolver en rows: ", username)
		//fmt.Println("Lastusername a devolver en rows: ", lastname)
		if err != nil {
			fmt.Println("Getting name and lastname")
		}
	}
	//fmt.Println("username a devolver: ", username)
	//fmt.Println("Lastusername a devolver: ", lastname)
	return username, lastname, country
}

func GetPostsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, errConv := vars["id"]
	if !errConv {
		fmt.Println("The ID couldn't been catch")
	}
	post, errGetPost := getPostById(postId)
	if errGetPost != nil {
		//models.SendNoContent(w)
		http.Error(w, "Post not found.", http.StatusNotFound)
	}
	_ = json.NewEncoder(w).Encode(post)

}
func getPostById(postId string) (*Post, error) {
	post := &Post{}
	sql := "SElECT id, user_id, title, content, created_at, updated_at FROM posts where id=?"
	rowsPost, err := config.Query(sql, postId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rowsPost.Next() {
		err := rowsPost.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Could not catch any row")
			return nil, errors.New("Post not found")
		}
		//fmt.Println(post)
	}
	return post, nil
}

func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("The ID couldn't been catch")
		models.SendCustom(w, "The <*/id> parameter is required", 400)
		return
	}

	posts, err := getPostsByUser(userId)
	usernameIdString := strconv.Itoa(userId)
	if err != nil {
		_, _ = fmt.Fprintf(w, "El usuario "+usernameIdString+" no ha creado ningun post.")
		models.SendNoContent(w)
		return
	}
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		models.SendInternalServerError(w)
		return
	}
}
func getPostsByUser(userId int) (Posts, error) {
	posts := Posts{}
	sql := "SElECT id, user_id, title, content FROM posts where user_id=?"
	rowsPost, err := config.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	for rowsPost.Next() {
		post := Post{}
		err := rowsPost.Scan(&post.Id, &post.UserId, &post.Title, &post.Content)
		if err != nil {
			fmt.Println("Could not catch any row")
		}
		posts = append(posts, post)
	}
	if len(posts) < 1 {
		return nil, errors.New("The user hasn't created any posts yet")
	}
	return posts, nil
}
