package migrations

import "embed"

// fs stands for f=> find all the sql files in the directory pathilestructure
//
//go:embed *.sql
var FS embed.FS
