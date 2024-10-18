package pagination

import "math"

// PaginationData holds the pagination information
type PaginationData struct {
	Total      int         `json:"total"`      // Total number of items
	Page       int         `json:"page"`       // Current page number
	Limit      int         `json:"limit"`      // Number of items per page
	TotalPages int         `json:"totalPages"` // Total number of pages
	Data       interface{} `json:"data"`       // Paginated data
}

// Paginate calculates the offset and constructs pagination data
func Paginate(page, limit, total int, data interface{}) PaginationData {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Calculate total number of pages
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return PaginationData{
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Data:       data,
	}
}

// GetOffset calculates the offset for SQL queries
func GetOffset(page, limit int) int {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return (page - 1) * limit
}
