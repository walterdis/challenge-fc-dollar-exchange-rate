package src

type Exchange struct {
	Id int `gorm:"primary_key"`

	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code       string
	CodeIn     string
	Name       string
	High       float64 `json:",string"`
	Low        float64 `json:",string"`
	VarBid     float64 `json:",string"`
	PctChange  float64 `json:",string"`
	Bid        float64 `json:",string"`
	Ask        float64 `json:",string"`
	Timestamp  int     `json:",string"`
	CreateDate string  `json:"create_date"`
}
