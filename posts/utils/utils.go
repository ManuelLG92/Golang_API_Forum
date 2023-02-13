package post_utils

import (
	"errors"
	"fmt"
	"forum/config"
	"forum/helpers"
	postDomain "forum/posts/domain"
	"forum/storage"
	"gorm.io/gorm"
	"math"
)

type PostList struct {
	config.Pagination
	Data []postDomain.Post `json:"data"`
}

type BadRequestError struct {
	Err error
}

const (
	listCustom = "custom"
)

var cache = storage.NewCache[postDomain.Post]()

func GetPostByIdAndUserId(postId, userId string) (*postDomain.Post, error) {
	var composedCacheKey = postId + "-" + userId

	if result := cache.Get(composedCacheKey); result != nil {
		return result, nil
	}
	fmt.Println("GetPostByIdAndUserId init", postId, userId)

	var post = &postDomain.Post{}

	if err := config.DbGorm.Where("id = ? AND user_id = ?", postId, userId).First(&post); err.Error != nil {
		fmt.Println("error quering", err)
		return nil, err.Error
	}
	fmt.Println("GetPostByIdAndUserId after where", post)

	if post.Id == "" {
		return nil, errors.New(fmt.Sprintf("Post %v not found", postId))
	}
	cache.Set(postId, *post)
	return post, nil
}

func GetPostsByUser(userId string) (*PostList, error) {

	var posts []postDomain.Post
	pagination := config.Pagination{Limit: 10}
	postsList := PostList{pagination, posts}
	values, err := listByMe(&postsList, userId)
	if err != nil {
		fmt.Printf("error %v \n", err)
		return nil, err
	}

	return values, nil

}

func GetPosts(pagination config.Pagination) (*PostList, error) {
	var posts []postDomain.Post
	postsList := PostList{pagination, posts}

	values, err := list(&postsList)
	if err != nil {
		fmt.Printf("error %v \n", err)
		return nil, err
	}

	return values, nil

}

func listByMe(pagination *PostList, userId string) (*PostList, error) {
	var values, err = listFactory(pagination, listCustom, "user_id = ?", userId)
	var posts []postDomain.Post
	if err != nil {
		return nil, err
	}
	pagination.Data = posts
	return values, nil
}

func list(pagination *PostList) (*PostList, error) {
	var values, err = listFactory(pagination, "all")
	if err != nil {
		return nil, err
	}
	return values, nil
}

func listFactory(pagination *PostList, kind string, data ...interface{}) (*PostList, error) {
	var posts []postDomain.Post
	var result *gorm.DB
	buildPagination := paginate(pagination, config.DbGorm)
	if pagination.Page > pagination.TotalPages {
		buildError := errors.New(fmt.Sprintf("Out of bounds page. The highest page for your query is: %v, and you requested %v", pagination.TotalPages, pagination.Page))
		return nil, BadRequestError{Err: buildError}.Err
	}
	switch {
	case kind == listCustom:
		result = config.DbGorm.Scopes(buildPagination).Where(data).Find(&posts)
	default:
		result = config.DbGorm.Scopes(buildPagination).Find(&posts)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	pagination.Data = posts
	return pagination, nil

}

func paginate(pagination *PostList, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	host := helpers.Get("HOST")
	var totalRows int64
	db.Model(postDomain.Post{}).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	for i := 0; i < totalPages; i++ {
		builder := fmt.Sprintf("%s/posts/latest?page=%v&limit=%v", host, i+1, pagination.Limit)
		pagination.Links = append(pagination.Links, builder)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func Save(post *postDomain.Post) error {
	result := config.DbGorm.Save(&post)
	if result.Error != nil {
		fmt.Printf("error saving post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Printf("saved posts %v.", post.Id)
	return nil
}

func Delete(id string) error {
	fmt.Println("init post_utils.Delete")

	result := config.DbGorm.Where("id = ?", id).Delete(&postDomain.Post{})
	fmt.Println("after where init post_utils.Delete")

	if result.Error != nil {
		fmt.Println("after where inside error != nil post_utils.Delete")

		fmt.Printf("error deleting post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Println("after where no error post_utils.Delete")

	fmt.Printf("Deleted post %v.", id)

	cache.Delete(id)
	return nil
}
