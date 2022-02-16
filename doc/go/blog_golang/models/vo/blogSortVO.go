package vo

type BlogSortVO struct {
	Keyword     string `json:"keyword"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	Uid         string `json:"uid"`
	Status      int    `json:"status"`
	SortName    string `json:"sortName"`
	Content     string `json:"content"`
	Sort        int    `json:"sort"`
}
