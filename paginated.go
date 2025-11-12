package api

type Paginated[T any] struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
	Items []T `json:"items"`
}
