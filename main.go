package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Monkhai/kindle-highlights-exporter/scraper"
	"github.com/Monkhai/kindle-highlights-exporter/shared"
	"github.com/Monkhai/kindle-highlights-exporter/utils"
)

func main() {

	s := scraper.NewScraper()
	if err := s.Signin(); err != nil {
		log.Fatalf("fuck signing %v", err)
	}
	if err := s.NavigateToHighlights(); err != nil {
		log.Fatalf("fuck navigating to highlights %v", err)
	}
	asins, err := s.GetAsins()
	if err != nil {
		log.Fatalf("fuck getting asins %v", err)
	}
	s.Asins = asins

	books := []scraper.Book{}

	for i := range s.Asins {
		book, err := s.GetBook()
		if err != nil {
			s.NextBook(s.Asins[i])
			book, err = s.GetBook()
			if err != nil {
				log.Println("still not letting me read this wow")
			}
		}
		books = append(books, book)
		if i < len(s.Asins)-1 {
			s.NextBook(s.Asins[i+1])
		}
	}

	reader := bufio.NewScanner(os.Stdin)
	outputDir := shared.GetInput(reader, "Where should we save this file to? (should be a path to a dir)")
	for _, book := range books {
		err := utils.WriteBookToMarkdown(book, outputDir)
		if err != nil {
			log.Fatalf("fuck creating files %v", err)
		}
	}
}
