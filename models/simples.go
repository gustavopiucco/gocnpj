package models

type Simples struct {
	CnpjBasico          string `gorm:"index"`
	OpcaoSimples        string
	DataOpcaoSimples    string
	DataExclusaoSimples string
	OpcaoMei            string
	DataOpcaoMei        string
	DataExclusaoMei     string
}
