package scraper

import "github.com/chromedp/chromedp"

type Book struct {
	Title      string
	Highlights []string
}

func (s *Scraper) GetBook() (Book, error) {
	var title string
	var highlights []string
	err := chromedp.Run(s.Ctx,
		// Select all <span id="highlight"> elements and extract their text content
		chromedp.Evaluate(`Array.from(document.querySelectorAll('span#highlight')).map(el => el.textContent)`, &highlights),
		chromedp.Evaluate(`document.querySelector("#annotation-scroller > div > div.a-row.a-spacing-base > div.a-column.a-span5 > h3").textContent`, &title),
	)
	if err != nil {
		return Book{}, err
	}
	return Book{Title: title, Highlights: highlights}, nil

}
