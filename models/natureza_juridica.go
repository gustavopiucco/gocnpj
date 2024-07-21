package models

type NaturezaJuridica struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
