package main

import (
	"scrapping-mercadolivre-golang/src/logs"
	"scrapping-mercadolivre-golang/src/models"
	"scrapping-mercadolivre-golang/src/scraper/fetcher"
	"scrapping-mercadolivre-golang/src/scraper/handlers"
	"scrapping-mercadolivre-golang/src/scraper/storage"
)

func main() {
	page, browser, err := fetcher.FetchPage()
	if err != nil {
		logs.Error(err)
	}
	defer browser.Close()

	group, err := handlers.GetGroups(page)
	if err != nil {
		logs.Error(err)
	}

	var products []models.Product

	handlers.ProcessCards(group, &products)
	storage.SaveData(&products)
}
