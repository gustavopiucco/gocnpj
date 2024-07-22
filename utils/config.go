package utils

import "path/filepath"

const (
	BatchSize   = 5000000
	DBBatchSize = 1000
)

var CSVPath = filepath.Join("data", "csv")
var DBPath = filepath.Join("data", "cnpj.db")
