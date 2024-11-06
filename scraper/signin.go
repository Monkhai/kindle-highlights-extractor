package scraper

import (
	"bufio"
	"os"
	"time"

	"github.com/Monkhai/kindle-highlights-exporter/shared"
	"github.com/chromedp/chromedp"
)

var signinUrl string = "https://www.amazon.com/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https://www.amazon.com/?ref_%3Dnav_signin&openid.identity=http://specs.openid.net/auth/2.0/identifier_select&openid.assoc_handle=usflex&openid.mode=checkid_setup&openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select&openid.ns=http://specs.openid.net/auth/2.0"

func (s *Scraper) Signin() error {
	err := chromedp.Run(s.Ctx,
		chromedp.Navigate(signinUrl),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		return err
	}
	reader := bufio.NewScanner(os.Stdin)
	email := shared.GetInput(reader, "Enter your email: ")
	password := shared.GetInput(reader, "Enter your password: ")
	err = chromedp.Run(s.Ctx,
		chromedp.SendKeys(`#ap_email`, email),
		chromedp.Click(`#continue`),
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys(`#ap_password`, password),
		chromedp.Click(`#signInSubmit`),
		chromedp.Sleep(3*time.Second), // Wait for login to process
	)
	if err != nil {
		return err
	}
	return nil
}
