package scraper

import (
	"errors"
	"time"

	"github.com/chromedp/chromedp"
)

var notebookURL string = "https://read.amazon.com/notebook"

func (s *Scraper) NavigateToHighlights() error {
	err := chromedp.Run(s.Ctx,
		chromedp.Navigate(notebookURL),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) GetAsins() ([]string, error) {
	var asins []string
	var currURL string
	err := chromedp.Run(s.Ctx,
		chromedp.Location(&currURL),
	)
	if err != nil {
		return nil, err
	}

	if currURL != notebookURL {
		return nil, errors.New("Not in the notebook URL")
	}

	err = chromedp.Run(s.Ctx,
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.kp-notebook-library-each-book span[data-get-annotations-for-asin]')).map(el => el.getAttribute('data-get-annotations-for-asin')).map(data => JSON.parse(data).asin)`, &asins),
	)
	if err != nil {
		return nil, err
	}
	return asins, nil
}

func (s *Scraper) NextBook(asin string) error {
	err := chromedp.Run(s.Ctx,
		chromedp.Click(`span[data-get-annotations-for-asin='{"asin":"`+asin+`"}'] a`, chromedp.NodeVisible),
		chromedp.Sleep(4*time.Second), // Adjust wait time as necessary
	)
	if err != nil {
		return err
	}
	return nil
}
