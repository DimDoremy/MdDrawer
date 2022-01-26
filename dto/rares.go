package dto

type Rares struct {
	Id      int64  `json:"id"`
	Rare    int    `json:"rare"`
	Package string `json:"package"`
	Url     string `json:"url"`
}

func (r *Rares) TableName() string {
	return "rares"
}
