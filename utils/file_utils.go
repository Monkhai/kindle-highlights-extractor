package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Monkhai/kindle-highlights-exporter/scraper"
)

func WriteBookToMarkdown(book scraper.Book, outputDir string) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := filepath.Join(outputDir, book.Title+".md")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "# %s\n\n", book.Title)
	if err != nil {
		return err
	}
	for _, highlight := range book.Highlights {
		_, err = fmt.Fprintf(file, "- %s\n", highlight)
		if err != nil {
			return err
		}
	}
	return nil
}
