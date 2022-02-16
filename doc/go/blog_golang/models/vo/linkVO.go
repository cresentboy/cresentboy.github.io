package vo

type LinkVO struct {
	Keyword     string      `json:"keyword"`
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Uid         string      `json:"uid"`
	Status      int         `json:"status"`
	Title       string      `json:"title"`
	Summary     string      `json:"summary"`
	Url         string      `json:"url"`
	Sort        int         `json:"sort"`
	LinkStatus  interface{} `json:"linkStatus"`
	Email       string      `json:"email"`
	FileUid     string      `json:"fileUid"`
}
