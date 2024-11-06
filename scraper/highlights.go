package scraper

import (
	"time"

	"github.com/chromedp/chromedp"
)

func (s *Scraper) NavigateToHighlights() error {
	err := chromedp.Run(s.Ctx,
		chromedp.Navigate("https://read.amazon.com/notebook"),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) GetAsins() ([]string, error) {
	var asins []string
	err := chromedp.Run(s.Ctx,
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
