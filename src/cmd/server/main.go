package main

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.mercadolivre.com.br/ofertas?container_id=MLB779362-1&promotion_type=lightning"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	group, err := page.Locator("div.items-list > div.items-with-smart-groups").InnerHTML()
	if err != nil {
		log.Fatalf("could not get div groups: %v", err)
	}

	// for i, entry := range group {
	// 	card, err := entry.Locator("td.title > span > a")
	// 	if err != nil {
	// 		log.Fatalf("could not get text content: %v", err)
	// 	}
	// }
	// if err = browser.Close(); err != nil {
	// 	log.Fatalf("could not close browser: %v", err)
	// }
	// if err = pw.Stop(); err != nil {
	// 	log.Fatalf("could not stop Playwright: %v", err)
	// }
}
