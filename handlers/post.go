package handlers

import (
	"../config"
	"../models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Uuid string `json:"uuid"`
	//UserName string `json:"user_name"`
	//LastName string `json:"last_name"`
	//Country string `json:"country"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostForOptions struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	//UserName string `json:"user_name"`
	//LastName string `json:"last_name"`
	//Country string `json:"country"`
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostWithUserData struct {
	//PostInto Post
	Id int `json:"id"`
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	LastName string `json:"last_name"`
	Country string `json:"country"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type Posts []PostWithUserData

// Init CreatePost public function
func CreatePost (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	post := Post{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	errorSession, userId := getSessionByUuid(post.Uuid)
	if errorSession != nil {
		fmt.Println("WE couldnt catch any session by uuid: ", err)
		http.Error(w,"You should be authenticated to create posts. ", http.StatusForbidden)
		return
	}
	if userId != 0 {
		createdPost, err := createPost(post.UserId, post.Title, post.Content)
		if err != nil {
			fmt.Println("Error trying to create a user: ", err)
			models.SendNoContent(w)
		} else {
			fmt.Println("Post created")
			_ : json.NewEncoder(w).Encode(createdPost)
			models.SendData(w,createdPost)
		}
	}


}
// End CreatePost public function

// Init createPost private function
func createPost (userId int, title, content string)  (*Post, error ){
	post, errorValidPost := newPost(userId, title, content)
	if errorValidPost != nil {
		return nil, errorValidPost
	} else {
		errorInsertPost := post.insertPost()
		if errorInsertPost != nil {
			return nil,errorInsertPost
		}
	}

	return post,nil

}
// End createPost private function

func newPost (userId int, title, content string) (*Post, error) {

	post := &Post{UserId: userId, Title: title, Content: content}
	errNewPost := post.ValidPost()
	if errNewPost != nil {
		 return nil, errNewPost
	}
	return post, nil

}
// Init validPost Public function
func (post *Post) ValidPost () error  {
	if validUserId(post.UserId) != nil && validPostData(post.Title, post.Content) != nil {
		return models.ErrorInvalidPost
	}
	return nil
}
// End validPost Public function
/*
{
	"user_id": 25,
	"title": "test post,
	"content": "Description del post de test"

}
*/

// Init  Update Post
func UpdatePost (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	postId := getIdByUrl(r)
	if postId == 0 {
		http.Error(w,"No valid post Id", http.StatusNotAcceptable)
		return
	}
	post := PostForOptions{}
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&post)
	//fmt.Println("UUid de la cookie: ", post.Uuid )
	err, userId := getSessionByUuid(post.Uuid)
	if err != nil {
		fmt.Println("WE couldnt catch any session by uuid: ", err)
	}
	//fmt.Println("USer id by post: ", post.UserId)
	//fmt.Println("USer id by sql : ", userId)
	sameUser := validateUserOptions(userId,post.UserId)
	if sameUser {
		updatedPost:= editPost (postId, post.UserId, post.Title, post.Content)
		if updatedPost != nil {
			fmt.Println("Error trying to Update post: ", updatedPost)
			http.Error(w, "Error trying to Update post.", http.StatusNoContent)
		} else {
			fmt.Println("Post Updated")
			w.WriteHeader(http.StatusAccepted)
			//_ = json.NewEncoder(w).Encode(updatedPost)
			//models.SendData(w,createdPost)
		}
	} else {
		http.Error(w, "You onyl can eddit your own posts.", http.StatusForbidden)
	}
}

// End Update Post

//Init DeletePost
func DeletePost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	postId := getIdByUrl(r)
	//fmt.Println("post de param: ", postId)
	if postId == 0 {
		http.Error(w,"No valid post Id for Delete", http.StatusNotAcceptable)
		return
	}
	post := PostForOptions{}
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&post)
	post.Id = postId
	//fmt.Println("Boidy request: : ", r.Body )
	//fmt.Println("UUid de la cookie: ", post.Uuid )
	//fmt.Println("UUid de la cookie: ", post)

	err, userId := getSessionByUuid(post.Uuid)
	if err != nil {
		fmt.Println("WE couldnt catch any session by uuid: ", err)
	}
	//fmt.Println("user id 1 : ", userId )
	//fmt.Println("posrt id: ", post.UserId)
	sameUser := validateUserOptions(userId,post.UserId)
	//fmt.Println("Validacion userID post userID: ", sameUser)
	if sameUser {
	deleteStatus := deletePost(post.Id)
	if deleteStatus != nil {
		http.Error(w,"Has been a mistake trying yo delete your post", http.StatusNotAcceptable)
	}
	w.WriteHeader(http.StatusAccepted)
	} else {
		http.Error(w, "You only can delete your own posts.", http.StatusForbidden)
	}


}

func validateUserOptions (userId, postUserId int) bool  {
	if  userId == postUserId  {
		return true
	} else {
		return false
	}
}
//End DeletePost


// Init editPost private function
func editPost (postId , userId int, title, content string) error {
	/*post, errorValidPost := newPost(userId, title, content)
	if errorValidPost != nil {
		return nil, errorValidPost
	} else {}*/
		EditedPost := updatePostSql(postId, userId, title, content)
		if EditedPost != nil {
			return EditedPost
		}
	return nil
}
// End editPost private function


func GetPosts (w http.ResponseWriter, _ *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	allPosts := getPosts()
	if allPosts == nil {
		models.SendNoContent(w)
	} else {
		models.SendData(w, allPosts)
	}

}
func getPosts () *Posts  {
	syncSession.Lock()
	defer syncSession.Unlock()
	posts := Posts{}
	sql := "SElECT id, user_id, title, content,created_at, updated_at from POSTS"
	rowsPost, err := config.Query(sql)
	if err != nil {
		return nil
	}
	for rowsPost.Next() {
		post := PostWithUserData{}
		err := rowsPost.Scan(&post.Id,&post.UserId,&post.Title,&post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Could not catch any row")
		}
		//fmt.Println(post.CreatedAt)
		createdPostDate, _ :=  time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		createdPostDateTime := createdPostDate.Format("Jan 2 2006 at 15:04")
		//fmt.Println(createdPostDateTime)
		post.CreatedAt = createdPostDateTime
		//fmt.Println(post.CreatedAt)
		//fmt.Println("Post id pasado: ", post.UserId)

		username, lastname, country := getNameAndLastNameUser(post.UserId)
		if len(username) > 0 && len(lastname)>0 && len(country)>0  {
			post.UserName = username
			post.LastName = lastname
			post.Country = country
		}

		posts = append(posts, post)
	}
	return &posts
}

func getNameAndLastNameUser (userId int) (string, string, string) {
		var username, lastname, country string
		//fmt.Println("Post id en la funcion: ", userId)
		sqlUsers := "SElECT name,last_name, country from USERS where id=?"
		rowUsernames , err := config.Query(sqlUsers,userId)
		if err != nil {
			fmt.Println("Error nil ", err)
			return "","",""
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

func GetPostsById (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	postId, errConv := strconv.Atoi(vars["id"])
	if errConv != nil {
		fmt.Println("The ID couldn't been catch")
	}
	post, errGetPost := getPostById(postId)
	if errGetPost != nil {
		//models.SendNoContent(w)
		http.Error(w,"Post not found.", http.StatusNotFound)
	}
		_ = json.NewEncoder(w).Encode(post)

}
func getPostById(postId int) (*PostWithUserData, error) {
	syncSession.Lock()
	defer syncSession.Unlock()
	post := &PostWithUserData{}
	sql := "SElECT id, user_id, title, content, created_at, updated_at FROM posts where id=?"
	rowsPost, err := config.Query(sql, postId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rowsPost.Next() {
		err := rowsPost.Scan(&post.Id,&post.UserId,&post.Title,&post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Could not catch any row")
			return nil, errors.New("Post not found")
		} else {
			username, lastname, country := getNameAndLastNameUser(post.UserId)
			if len(username) > 0 && len(lastname)>0 && len(country)>0  {
				post.UserName = username
				post.LastName = lastname
				post.Country = country
			}
		}
		//fmt.Println(post)
	}
	return post, nil
}

func GetPostsByUser (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("The ID couldn't been catch")
	} else {
		posts, err := getPostsByUser(userId)
		usernameIdString := strconv.Itoa(userId)
		if err != nil {
			_,_ = fmt.Fprintf(w,"El usuario " + usernameIdString + " no ha creado ningun post.")
			models.SendNoContent(w)
		} else {
			_ : json.NewEncoder(w).Encode(posts)
		}
	}


}
func getPostsByUser (userId int) (Posts, error) {
	syncSession.Lock()
	defer syncSession.Unlock()
	posts := Posts{}
	sql := "SElECT id, user_id, title, content FROM posts where user_id=?"
	rowsPost, err := config.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	for rowsPost.Next() {
		post := PostWithUserData{}
		err := rowsPost.Scan(&post.Id,&post.UserId,&post.Title,&post.Content)
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





