package main

import (
	"gocnpj/models"
	"gocnpj/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(utils.DBPath), &gorm.Config{})
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

	utils.ProcessFiles(db, "MUNIC", &[]models.Municipio{})
	utils.ProcessFiles(db, "MOTI", &[]models.Motivo{})
	utils.ProcessFiles(db, "NATJU", &[]models.NaturezaJuridica{})
	utils.ProcessFiles(db, "QUALS", &[]models.QualificacaoSocio{})
	utils.ProcessFiles(db, "SIMPLES", &[]models.Simples{})
	utils.ProcessFiles(db, "SOCIO", &[]models.Socios{})
	utils.ProcessFiles(db, "CNAE", &[]models.Cnae{})
	utils.ProcessFiles(db, "PAIS", &[]models.Pais{})
	utils.ProcessFiles(db, "EMPRE", &[]models.Empresa{})
	utils.ProcessFiles(db, "ESTABELE", &[]models.Estabelecimento{})
}
