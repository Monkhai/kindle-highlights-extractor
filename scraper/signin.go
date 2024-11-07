package scraper

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	"github.com/Monkhai/kindle-highlights-exporter/shared"
	"github.com/chromedp/chromedp"
)

var signinURL string = "https://www.amazon.com/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https://www.amazon.com/?ref_%3Dnav_signin&openid.identity=http://specs.openid.net/auth/2.0/identifier_select&openid.assoc_handle=usflex&openid.mode=checkid_setup&openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select&openid.ns=http://specs.openid.net/auth/2.0"
var targetURL string = "https://www.amazon.com/?ref_=nav_signin"

func (s *Scraper) Signin() error {
	err := chromedp.Run(s.Ctx,
		chromedp.Navigate(signinURL),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		return err
	}
	reader := bufio.NewScanner(os.Stdin)
	email := shared.GetInput(reader, "Enter your email: ")
	password := shared.GetInput(reader, "Enter your password: ")
	var currentURL string
	err = chromedp.Run(s.Ctx,
		chromedp.SendKeys(`#ap_email`, email),
		chromedp.Click(`#continue`),
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys(`#ap_password`, password),
		chromedp.Click(`#signInSubmit`),
		chromedp.Sleep(3*time.Second),
		chromedp.Location(&currentURL),
	)
	if err != nil {
		return err
	}

	if currentURL != targetURL {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("disable-gpu", true),
		)
		allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
		ctx, _ := chromedp.NewContext(allocCtx)
		s.Ctx = ctx

		log.Println("You need to complete the signin manually please")

		err := chromedp.Run(s.Ctx,
			chromedp.Navigate(signinURL),
		)
		if err != nil {
			return err
		}

		shared.GetInput(reader, "Once you are done press any key. Ensure that you are logged in!")

		opts = append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", false),
		)
		allocCtx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
		ctx, _ = chromedp.NewContext(allocCtx)
		s.Ctx = ctx
	}

	return nil
}
