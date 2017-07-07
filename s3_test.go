package sample

import (
	"os"
	"testing"
)

func TestS3(t *testing.T) {
	table := []struct {
		file       string
		bucketName string
		key        string
	}{
		{"./README.md", "go-localstack-sample", "README.md"},
	}

	for _, row := range table {
		f, err := os.Open(row.file)
		if err != nil {
			t.Fatalf("os.Open got error: %s", err)
		}

		out := S3PutObject(row.bucketName, row.key, f)
		if *(out.ETag) == "" {
			t.Errorf("etag was empty")
		}
	}
}
