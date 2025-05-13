package utils

import "go-fiber-app/core/dto"

func Pagination(page int, limit int) (int, int) {
	return (page - 1) * limit, limit
}

// tatal 30     limit 10     skip 0     page 1
// tatal 30     limit 10     skip 10     page 2
// tatal 30     limit 10     skip 20     page 3
func CalculatePagination(totalItems int64, page dto.QueryPage) dto.Pagination {
	if page.Limit <= 0 {
		page.Limit = 10 // Default to 10 items per page if limit is invalid
	}
	if page.Skip < 0 {
		page.Skip = 0 // Prevent negative skip values
	}
	currentPage := int64((page.Skip + page.Limit) / page.Limit)
	if page.Page > 0 {
		currentPage = int64(page.Page)
	}
	totalPages := int64((totalItems + int64(page.Limit) - 1) / int64(page.Limit)) // Calculate total pages (round up)
	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = 0 // Set nextPage to 0 if it exceeds totalPages
	}
	previousPage := currentPage - 1
	if previousPage < 1 {
		previousPage = 0 // Set previousPage to 0 if it is less than 1
	}

	return dto.Pagination{
		CurrentPage:  int64(currentPage),
		TotalPage:    int64(totalPages),
		TotalItem:    totalItems,
		PageSize:     int64(page.Limit),
		NextPage:     &nextPage,
		PreviousPage: &previousPage,
	}
}

func CalculatePaginationFromPage(totalItems int64, page int, limit int) dto.Pagination {
	if limit <= 0 {
		limit = 10 // Default to 10 items per page if limit is invalid
	}
	if page < 1 {
		page = 1 // Prevent negative skip values
	}

	totalPages := int64((totalItems + int64(limit) - 1) / int64(limit)) // Calculate total pages (round up)
	nextPage := int64(page + 1)
	if nextPage > totalPages {
		nextPage = 0 // Set nextPage to 0 if it exceeds totalPages
	}
	previousPage := int64(page - 1)
	if previousPage < 1 {
		previousPage = 0 // Set previousPage to 0 if it is less than 1
	}

	return dto.Pagination{
		CurrentPage:  int64(page),
		TotalPage:    int64(totalPages),
		TotalItem:    totalItems,
		PageSize:     int64(limit),
		NextPage:     &nextPage,
		PreviousPage: &previousPage,
	}
}
