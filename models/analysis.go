package models

type Analysis struct {
	ID   int    `json:"id" db:"ID"`
	Date string `json:"date" db:"Date"`
	Bld  string `json:"bld" db:"Bld"`
	Ubg  string `json:"ubg" db:"Ubg"`
	Bil  string `json:"bil" db:"Bil"`
	Pro  string `json:"pro" db:"Pro"`
	Nit  string `json:"nit" db:"Nit"`
	Ket  string `json:"ket" db:"Ket"`
	Glu  string `json:"glu" db:"Glu"`
	PH   string `json:"ph" db:"pH"`
	SG   string `json:"sh" db:"SG"`
	Leu  string `json:"leu" db:"Leu"`
}
