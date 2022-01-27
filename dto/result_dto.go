package dto

type ResultDto struct {
	RtnCode int         `json:"rtnCode"`
	RtnMsg  string      `json:"rtnMsg"`
	RtnData interface{} `json:"rtnData"`
}

type ResultData struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}
