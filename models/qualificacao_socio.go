package models

type QualificacaoSocio struct {
	Codigo    string `gorm:"index"`
	Descricao string
}
