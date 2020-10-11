package model

type ExploreFile struct {
	Id           int64  `json:"id" form:"id"`
	TabId        int64  `json:"tabId" form:"tabId"`
	FileId       int64  `json:"fileId" form:"fileId"`
	MediaType    string `json:"mediaType" form:"mediaType"`
	ParentItemId int64  `json:"parentItemId" form:"parentItemId"`
	Name         string `json:"name" form:"name"`
	Cover        string `json:"cover" form:"cover"`
	Description  string `json:"description" form:"description"`
	Keyword      string `json:"keyword" form:"keyword"`
}
