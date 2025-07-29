package fetcher

import (
	"log"

	"scrapping-mercadolivre-golang/src/config"

	"github.com/playwright-community/playwright-go"
)

// Inicializa o navegador e abre a página
func FetchPage() (playwright.Page, playwright.Browser, error) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
		return nil, nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // Se você quiser visualizar o navegador, defina como false
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
		return nil, nil, err
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
		return nil, nil, err
	}

	_, err = page.Goto(config.URL, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	})
	if err != nil {
		log.Fatalf("could not goto: %v", err)
		return nil, nil, err
	}

	return page, browser, nil
}
