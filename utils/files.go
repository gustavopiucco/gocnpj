package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

func ProcessFiles(db *gorm.DB, contains string, model any) {
	for _, fileName := range getCSVFiles(contains) {
		fmt.Println("Processing file", fileName)

		file, err := os.Open(filepath.Join(CSVPath, fileName))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := transform.NewReader(file, charmap.ISO8859_1.NewDecoder())

		csvReader := csv.NewReader(bufio.NewReader(reader))
		csvReader.Comma = ';'

		var lines [][]string
		totalLines := 0

		for {
			line, err := csvReader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				panic(err)
			}

			lines = append(lines, line)
			totalLines++

			if len(lines) == BatchSize {
				fmt.Println("Processing batch of", BatchSize, "lines")
				fmt.Println("Total lines read:", totalLines)
				processBatch(db, lines, model)
				// Reseta o buffer de linhas para a próxima batch
				lines = nil
			}
		}

		// Processa o restante das linhas se houver
		if len(lines) > 0 {
			fmt.Println("Processing batch of", len(lines), "remaining lines")
			fmt.Println("Total lines read:", totalLines)
			processBatch(db, lines, model)
		}
	}
}

func processBatch(db *gorm.DB, lines [][]string, model any) {
	if err := AppendLinesToSlice(lines, model); err != nil {
		panic(err)
	}

	fmt.Printf("Creating records in database... ")
	db.CreateInBatches(model, DBBatchSize)
	fmt.Println("Done")
}

func getCSVFiles(contains string) []string {
	var fileNames []string

	err := filepath.WalkDir(CSVPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(d.Name(), contains) {
			fileNames = append(fileNames, d.Name())
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return fileNames
}
