package structs

// Keyword ...
type Keyword struct {
	KeywordID   int64  `json:"keywordid"`
	KaKeyword   string `json:"kakeyword"`
	KaAttribute string `json:"kaattribute"`
	KaLongDescr string `json:"kalongdescr"`
	KaHide      int64  `json:"kahide"`
	KeyTypesID  int64  `json:"keytypesid"`
}

//ListKeyword for list
type ListKeyword struct {
	Keywords []Keyword `json:"keyword"`
}

//GetListKeywordResult obj
type GetListKeywordResult struct {
	MyListKeyword ListKeyword `json:"getlistkeywordresult"`
	ErrorCode     int         `json:"errorcode"`
	ErrorDesc     string      `json:"errordesc"`
}

// NewGetListKeywordResult Obj
func NewGetListKeywordResult() *GetListKeywordResult {
	return &GetListKeywordResult{
		ErrorCode: 1,
		ErrorDesc: "Unexpected Error",
	}
}
