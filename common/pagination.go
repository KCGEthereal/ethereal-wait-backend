package common

import (
	"net/http"
	"strconv"
)

// PaginationResponse is given to the Response struct. It has every
// information a client might require to traverse through multiple
// documents. Elements are basically multiple items of data.
type PaginationResponse struct {
	TotalPages    int64 `json:"totalPages"`
	TotalElements int64 `json:"totalElements"`
	CurrentPage   int64 `json:"page"`
	Elements      int64 `json:"size"`
}

// Pagination struct is used to get query params from the URI
// and return the data as a struct to be further used by our app
//
//	/banners/byPriority?page=2&size=20
type Pagination struct {
	PageNumb int64 `json:"page"`
	PageSize int64 `json:"size"`
}

// DecodePaginationQueries is responsible to decode page and size
// query params from the URI and return a Pagination struct for
// app's consumption.
//
// We also set healthy defaults in case of the URI not having any
// pagination query parameters entered in it.
//
// By default, we will display the first page and only 10 elements
// per page.
func DecodePaginationQueries(req *http.Request) *Pagination {
	pageNumber, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil || pageNumber <= 0 {
		// If any error occurred of if pageNumber is < 0, then we
		// automatically set it to 1 so that the app does not crash
		pageNumber = 1
	}

	// Currently users can request as many items as they want in a
	// single API call. We might need to limit that.
	pageSize, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil || pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return &Pagination{
		PageNumb: int64(pageNumber),
		PageSize: int64(pageSize),
	}
}
