package main

import (
	upload "github.com/cloaiza1997/dev-test-tr-emails/functions/upload"
)

func main() {
	mailDir := "./mock/maildir"
	indexByBatch := true

	upload.InitUpload(mailDir, indexByBatch)
}
