package common

type ListResult[T any] struct {
	Data    []T   `json:"data"`
	EndPage int64 `json:"end_page"`
}
