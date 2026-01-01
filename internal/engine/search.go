package engine

import (
	"os"
	"path/filepath"
	"strings"
)

// FindFiles scans a directory for files matching a specific name or extension
func FindFiles(root string, searchTerm string, extension string) ([]string, error) {
	var matches []string

	// Walk through the folder structure
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip folders we don't have permission to access
		}

		if !info.IsDir() {
			fileName := info.Name()

			// Check if file contains the search term AND has the right extension
			if strings.Contains(strings.ToLower(fileName), strings.ToLower(searchTerm)) {
				if extension == "" || filepath.Ext(fileName) == extension {
					matches = append(matches, path)
				}
			}
		}
		return nil
	})

	return matches, err
}
