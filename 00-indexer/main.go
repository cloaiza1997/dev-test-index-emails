package main

import (
	"flag"

	upload "github.com/cloaiza1997/dev-test-tr-emails/functions/upload"
)

func main() {
	index := flag.String("i", "emails", "Index emails")
	mailDir := flag.String("p", "./mock/maildir", "Mail directory path")
	routines := flag.Int("r", 10, "Routines to use")
	indexByBatch := flag.Bool("b", true, "Index by batch")
	batchSize := flag.Int("s", 10000, "Batch size")

	flag.Parse()

	upload.InitUpload(upload.UploadOptions{
		Index:        *index,
		MailDir:      *mailDir,
		Routines:     *routines,
		IndexByBatch: *indexByBatch,
		BatchSize:    *batchSize,
	})
}
