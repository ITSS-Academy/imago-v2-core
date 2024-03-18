package common

type QueryOpts struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	OrderBy string `json:"orderBy"`
	Order   string `json:"order"`
	Query   string `json:"query"`
}
