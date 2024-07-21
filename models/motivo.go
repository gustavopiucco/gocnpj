package models

type Motivo struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
