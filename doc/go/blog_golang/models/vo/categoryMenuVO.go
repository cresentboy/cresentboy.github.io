package vo

type CategoryMenuVO struct {
	Keyword           string `json:"keyword"`
	CurrentPage       int    `json:"currentPage"`
	PageSize          int    `json:"pageSize"`
	Uid               string `json:"uid"`
	Status            int    `json:"status"`
	Name              string `json:"Name"`
	MenuLevel         int    `json:"menuLevel"`
	MenuType          int    `json:"menuType"`
	Summary           string `json:"summary"`
	Icon              string `json:"icon"`
	ParentUid         string `json:"parentUid"`
	Url               string `json:"url"`
	Sort              int    `json:"sort"`
	IsShow            int    `json:"isShow"`
	IsJumpExternalUrl int    `json:"isJumpExternalUrl"`
}
