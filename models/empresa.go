package models

type Empresa struct {
	CnpjBasico                string `gorm:"index"`
	RazaoSocial               string `gorm:"index"`
	NaturezaJuridica          string
	QualificacaoResponsavel   string
	CapitalSocial             string
	PorteEmpresa              string
	EnteFederativoResponsavel string
}
