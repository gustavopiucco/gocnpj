package main

import (
	"github.com/gustavopiucco/gocnpj"
	"github.com/gustavopiucco/gocnpj/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(gocnpj.DBPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Municipio{},
		&models.Motivo{},
		&models.NaturezaJuridica{},
		&models.QualificacaoSocio{},
		&models.Simples{},
		&models.Socios{},
		&models.Cnae{},
		&models.Pais{},
		&models.Empresa{},
		&models.Estabelecimento{},
	)

	gocnpj.ProcessFiles(db, "MUNIC", &[]models.Municipio{})
	gocnpj.ProcessFiles(db, "MOTI", &[]models.Motivo{})
	gocnpj.ProcessFiles(db, "NATJU", &[]models.NaturezaJuridica{})
	gocnpj.ProcessFiles(db, "QUALS", &[]models.QualificacaoSocio{})
	gocnpj.ProcessFiles(db, "SIMPLES", &[]models.Simples{})
	gocnpj.ProcessFiles(db, "SOCIO", &[]models.Socios{})
	gocnpj.ProcessFiles(db, "CNAE", &[]models.Cnae{})
	gocnpj.ProcessFiles(db, "PAIS", &[]models.Pais{})
	gocnpj.ProcessFiles(db, "EMPRE", &[]models.Empresa{})
	gocnpj.ProcessFiles(db, "ESTABELE", &[]models.Estabelecimento{})
}
