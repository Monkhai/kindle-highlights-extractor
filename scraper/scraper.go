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
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx)

	return &Scraper{
		Ctx:   ctx,
		Asins: []string{},
	}
}
