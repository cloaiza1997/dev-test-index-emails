package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkFilePath(root string, walk func(path string)) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the path (%s): %v", root, err)
		}

		if !info.IsDir() {
			walk(path)
		}

		return nil
	})

	return err
}
