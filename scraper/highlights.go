package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/Monkhai/kindle-highlights-exporter/shared"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var notebookURL string = "https://read.amazon.com/notebook"

func (s *Scraper) NavigateToHighlights() error {

	expiration := cdp.TimeSinceEpoch(time.Now().Add(365 * 24 * time.Hour))

	err := chromedp.Run(s.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies := []struct {
				Name, Value, Domain, Path string
			}{
				{"ubid-main", shared.UbidMain, ".amazon.com", "/"},
				{"at-main", shared.AtMain, ".amazon.com", "/"},
				{"x-main", shared.XMain, ".amazon.com", "/"},
				{"session-id", shared.SessionId, ".amazon.com", "/"},
			}
			for _, c := range cookies {
				err := network.SetCookie(c.Name, c.Value).
					WithDomain(c.Domain).
					WithPath(c.Path).
					WithExpires(&expiration).
					WithHTTPOnly(false).
					WithSecure(false).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
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
		return nil, fmt.Errorf("Not in the notebook URL. Instead in %s", currURL)
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
