package dto

type Texts struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Str3  string `json:"str3"`
	Str4  string `json:"str4"`
	Str5  string `json:"str5"`
	Str6  string `json:"str6"`
	Str7  string `json:"str7"`
	Str8  string `json:"str8"`
	Str9  string `json:"str9"`
	Str10 string `json:"str10"`
	Str11 string `json:"str11"`
	Str12 string `json:"str12"`
	Str13 string `json:"str13"`
	Str14 string `json:"str14"`
	Str15 string `json:"str15"`
	Str16 string `json:"str16"`
}

func (t *Texts) TableName() string {
	return "texts"
}
