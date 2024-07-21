package models

type Cnae struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
