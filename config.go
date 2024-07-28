package gocnpj

import "path/filepath"

const (
	BatchSize   = 250000
	DBBatchSize = 1000
)

var CSVPath = filepath.Join("data", "csv")
var DBPath = filepath.Join("data", "cnpj.db")
