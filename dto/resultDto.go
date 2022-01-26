package dto

type ResultDto struct {
	RtnCode int         `json:"rtnCode"`
	RtnMsg  string      `json:"rtnMsg"`
	RtnData interface{} `json:"rtnData"`
}
