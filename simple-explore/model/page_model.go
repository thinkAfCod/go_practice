package model

type PageModel struct {
	PageSize  int64       `json:"pageSize" form:"pageSize"`
	PageNo    int64       `json:"pageNo" form:"pageNo"`
	PageTotal int64       `json:"pageTotal" form:"pageTotal"`
	Count     int64       `json:"count" form:"count"`
	PageData  interface{} `json:"data" form:"data"`
}
