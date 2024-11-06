package scraper

import (
	"context"

	"github.com/chromedp/chromedp"
)

type Scraper struct {
	Ctx   context.Context
	Asins []string
}

func NewScraper() *Scraper {
	allocCtx, _ := chromedp.NewExecAllocator(context.Background())
	ctx, _ := chromedp.NewContext(allocCtx)

	return &Scraper{
		Ctx:   ctx,
		Asins: []string{},
	}
}
