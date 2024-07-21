package models

type Socios struct {
	CnpjBasico                     string `gorm:"index"`
	IdentificadorSocio             string
	NomeSocio                      string `gorm:"index"`
	CnpjCpfSocio                   string `gorm:"index"`
	QualificacaoSocio              string
	DataEntradaSociedade           string
	Pais                           string
	RepresentanteLegal             string `gorm:"index"`
	NomeRepresentanteLegal         string `gorm:"index"`
	QualificacaoRepresentanteLegal string
	FaixaEtaria                    string
}
