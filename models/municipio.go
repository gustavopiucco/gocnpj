package models

type Municipio struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
