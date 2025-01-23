package main

import (
	"testing"

	upload "github.com/cloaiza1997/dev-test-index-emails/functions/upload"
)

func TestProcessData(t *testing.T) {
	upload.InitUpload(upload.UploadOptions{
		Index:        "test-emails",
		MailDir:      "../mock/email-data/maildir",
		IndexJson:    "../data/index-structure.json",
		Routines:     10,
		IndexByBatch: true,
		BatchSize:    10000,
	})
}
