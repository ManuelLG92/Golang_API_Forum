package config

import (
	"fmt"
	"net/url"
	"strconv"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

var errorParsing = func(field string) string {
	return fmt.Sprintf("error parsing data from query field <%v>", field)
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func (p *Pagination) SetPageFromString(page string) error {
	if page == "" {
		p.Page = 1
		return nil
	}
	convert, err := strconv.Atoi(page)
	if err != nil {
		return err
	}
	p.Page = convert
	return nil
}
func (p *Pagination) SetLimitFromString(limit string) error {
	if limit == "" {
		p.Limit = 10
		return nil
	}
	convert, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	p.Page = convert
	return nil
}

func (p *Pagination) PaginateFromUrlQueryParams(query url.Values) *[]string {
	page := query.Get("page")
	limit := query.Get("limit")
	sort := query.Get("sort")
	return p.FactoryFromPrimitives(page, limit, sort)
}
func (p *Pagination) FactoryFromPrimitives(page, limit, sort string) *[]string {
	var errorCollection []string
	err := p.SetPageFromString(page)
	if err != nil {
		errorCollection = append(errorCollection, errorParsing("Page"))
	}
	err = p.SetLimitFromString(limit)
	if err != nil {
		errorCollection = append(errorCollection, errorParsing("Limit"))
	}
	if len(errorCollection) > 0 {
		return &errorCollection
	}
	p.Sort = sort
	if validations := p.CheckValidData(); validations != nil {
		return validations
	}
	return nil
}

func (p *Pagination) CheckValidData() *[]string {
	var errorCollection []string
	if p.Limit == 0 || p.Limit > 100 {
		errorCollection = append(errorCollection, fmt.Sprint("The limit must be between 0 and 100."))
	}
	if p.Sort != "" {
		errorCollection = append(errorCollection, fmt.Sprint("Sort not implemented yet. The default sort is by id and desc"))
		p.Sort = "Id desc"
	}
	p.Sort = "Id desc"
	if p.Page < 0 {
		errorCollection = append(errorCollection, fmt.Sprint("Just are allowed pages with positive number."))
	}

	if len(errorCollection) == 0 {
		return nil
	}
	return &errorCollection
}
