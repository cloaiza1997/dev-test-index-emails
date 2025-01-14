package main

import (
	"flag"

	upload "github.com/cloaiza1997/dev-test-tr-emails/functions/upload"
)

func main() {
	mailDir := flag.String("emails", "./mock/maildir", "Mail directory path")
	indexByBatch := flag.Bool("batch", true, "Index by batch")

	flag.Parse()

	upload.InitUpload(*mailDir, *indexByBatch)
}
