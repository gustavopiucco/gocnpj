package models

type Estabelecimento struct {
	CnpjBasico              string `gorm:"index"`
	CnpjOrdem               string
	CnpjDv                  string
	MatrizFilial            string
	NomeFantasia            string `gorm:"index"`
	SituacaoCadastral       string
	DataSituacaoCadastral   string
	MotivoSituacaoCadastral string
	NomeCidadeExterior      string
	Pais                    string
	DataInicioAtividade     string
	CnaeFiscal              string
	CnaeFiscalSecundaria    string
	TipoLogradouro          string
	Logradouro              string
	Numero                  string
	Complemento             string
	Bairro                  string
	Cep                     string
	UF                      string
	Municipio               string
	DDD1                    string
	Telefone1               string
	DDD2                    string
	Telefone2               string
	DDD_fax                 string
	Fax                     string
	Email                   string
	SituacaoEspecial        string
	DataSituacaoEspecial    string
}
