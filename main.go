package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"gocnpj/models"
	"gocnpj/utils"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const batchSize = 500000
const dbBatchSize = 1000

var csvPath = filepath.Join("data", "csv")
var dbPath = filepath.Join("data", "cnpj.db")

func main() {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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

	processFiles(db, "MUNIC", &[]models.Municipio{})
	processFiles(db, "MOTI", &[]models.Motivo{})
	processFiles(db, "NATJU", &[]models.NaturezaJuridica{})
	processFiles(db, "QUALS", &[]models.QualificacaoSocio{})
	processFiles(db, "SIMPLES", &[]models.Simples{})
	processFiles(db, "SOCIO", &[]models.Socios{})
	processFiles(db, "CNAE", &[]models.Cnae{})
	processFiles(db, "PAIS", &[]models.Pais{})
	processFiles(db, "EMPRE", &[]models.Empresa{})
	processFiles(db, "ESTABELE", &[]models.Estabelecimento{})
}

func processFiles(db *gorm.DB, contains string, model any) {
	for _, fileName := range getCSVFiles(contains) {
		fmt.Println("Processing file", fileName)

		file, err := os.Open(filepath.Join(csvPath, fileName))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := transform.NewReader(file, charmap.ISO8859_1.NewDecoder())

		csvReader := csv.NewReader(bufio.NewReader(reader))
		csvReader.Comma = ';'

		var lines [][]string

		for {
			line, err := csvReader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				panic(err)
			}

			lines = append(lines, line)

			if len(lines) == batchSize {
				fmt.Println("Processing batch of", batchSize, "lines")
				processBatch(db, lines, model)
				// Reseta o buffer de linhas e a slice para a próxima batch
				utils.ClearSlice(model)
				lines = nil
			}
		}

		// Processa o restante das linhas se houver
		if len(lines) > 0 {
			fmt.Println("Processing remaining", len(lines), "lines")
			processBatch(db, lines, model)
		}
	}
}

func processBatch(db *gorm.DB, lines [][]string, model any) {
	if err := utils.AppendToSlice(model, lines); err != nil {
		panic(err)
	}

	fmt.Printf("Creating records in database... ")
	db.CreateInBatches(model, dbBatchSize)
	fmt.Println("Done")
}

func getCSVFiles(contains string) []string {
	var fileNames []string

	err := filepath.WalkDir(csvPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(d.Name(), contains) {
			fileNames = append(fileNames, d.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fileNames
}
