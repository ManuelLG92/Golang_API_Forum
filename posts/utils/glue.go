package post_utils

import (
	"errors"
	"fmt"
	"forum/config"
	postDomain "forum/posts/domain"
	"forum/storage"
)

type PostList struct {
	config.Pagination
	Data []postDomain.Post `json:"data"`
}

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

func GetPostsByUser(userId string, motherPagination config.Pagination) (*PostList, error) {

	var posts []postDomain.Post
	//pagination := config.Pagination{Limit: 10}
	//pagination := config.PaginationFactory(motherPagination.GetPage(), motherPagination.GetLimit(), motherPagination.GetSort())
	postsList := PostList{motherPagination, posts}
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

//func paginate2(pagination *PostList, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
//
//	host := helpers.Get("HOST")
//	var totalRows int64
//	db.Model(postDomain.Post{}).Count(&totalRows)
//	pagination.TotalRows = totalRows
//	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
//	pagination.TotalPages = totalPages
//	for i := 0; i < totalPages; i++ {
//		builder := fmt.Sprintf("%s/posts/latest?page=%v&limit=%v", host, i+1, pagination.Limit)
//		pagination.Links = append(pagination.Links, builder)
//	}
//	return func(db *gorm.DB) *gorm.DB {
//		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
//	}
//}
