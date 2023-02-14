package post_utils

import (
	"errors"
	"fmt"
	"forum/config"
	postDomain "forum/posts/domain"
	"gorm.io/gorm"
)

type BadRequestError struct {
	Err error
}

const (
	listCustom = "custom"
)

func (pagination *PostList) updateData(motherPagination *config.Pagination) {
	pagination.TotalPages = motherPagination.TotalPages
	pagination.TotalRows = motherPagination.TotalRows
	pagination.Links = motherPagination.Links
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
	var paginationFactory = config.PaginationFactory(pagination.GetPage(), pagination.GetLimit(), pagination.GetSort())
	buildPagination := config.Paginate[postDomain.Post](paginationFactory, config.DbGorm)
	pagination.updateData(paginationFactory)
	if pagination.Page > pagination.TotalPages {
		buildError := errors.New(fmt.Sprintf("Out of bounds page. The highest page for your query is: %v, and you requested %v", pagination.TotalPages, pagination.Page))
		return nil, BadRequestError{Err: buildError}.Err
	}

	var result *gorm.DB
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
