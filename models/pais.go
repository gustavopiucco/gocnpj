package models

type Pais struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
