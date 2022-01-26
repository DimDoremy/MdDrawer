package dto

type Datas struct {
	Id        int64 `json:"id"`
	Ot        int64 `json:"ot"`
	Alias     int64 `json:"alias"`
	Setcode   int64 `json:"setcode"`
	Type      int64 `json:"type"`
	Atk       int   `json:"atk"`
	Def       int   `json:"def"`
	Level     int64 `json:"level"`
	Race      int   `json:"race"`
	Attribute int   `json:"attribute"`
	Category  int64 `json:"category"`
}

func (d *Datas) TableName() string {
	return "datas"
}
