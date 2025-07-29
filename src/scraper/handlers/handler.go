package handlers

import (
	"fmt"
	"log"
	"scrapping-mercadolivre-golang/src/logs"
	"scrapping-mercadolivre-golang/src/models"
	"scrapping-mercadolivre-golang/src/scraper/parser"

	"github.com/playwright-community/playwright-go"
)

// Coleta os itens da página
func GetGroups(page playwright.Page) ([]playwright.Locator, error) {
	// Crie um Locator para o seletor dos produtos
	groupLocator := page.Locator("div.items-list > div.items-with-smart-groups > .andes-card.poly-card.poly-card--grid-card.poly-card--large.andes-card--flat.andes-card--padding-0.andes-card--animated")

	// Aguarde até que o Locator esteja visível
	err := groupLocator.WaitFor(playwright.LocatorWaitForOptions{
		State: playwright.WaitForSelectorStateVisible, // Aguarda o locator ficar visível
	})
	if err != nil {
		logs.Error(fmt.Errorf("Erro ao esperar pelo seletor: %v", err))
		return nil, err
	}

	// Coletar todos os elementos encontrados pelo Locator
	group, err := groupLocator.All()
	if err != nil {
		logs.Error(fmt.Errorf("Erro ao coletar os grupos:", err))
		return nil, err
	}

	// Exibir o número de elementos encontrados
	log.Printf("Número de produtos encontrados: %d", len(group))

	// Armazenar os Locators para todos os elementos
	var locators []playwright.Locator
	for _, loc := range group {
		locators = append(locators, loc)
	}

	return locators, nil
}

func ProcessCards(group []playwright.Locator, products *[]models.Product) {
	for _, card := range group {
		var product models.Product

		// capturar o preço
		price, err := parser.ExtractPrice(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.Price = price

		// capturar título
		title, err := parser.ExtractTitle(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.Title = title

		// capturar quantidade de reviews
		quantity_reviews, err := parser.ExtractQuantityReviews(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.QuantityReviews = quantity_reviews

		// capturar a imagem do produto
		picture, err := parser.ExtractPicture(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.Picture = picture

		// capturar a URL do produto
		URL, err := parser.ExtractURL(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.URL = URL

		// capturar a nota em estrelas do produto
		stars, err := parser.ExtractStars(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.Stars = stars

		// capturar o preço de ancoragem
		anchorPrice, err := parser.ExtractAnchorPrice(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.AnchorPrice = anchorPrice

		// capturar um bool que retorna se o frete é grátis ou não
		isFreeShipping, err := parser.ExtractIsFreeShipping(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.IsFreeShipping = isFreeShipping

		// capturar a quantidade de parcelas
		installments, err := parser.ExtractInstallments(card)
		if err != nil {
			log.Println(installments)
			continue
		}
		// product.Installments = installments

		// capturar o valor de cada parcela
		installments_amount, err := parser.ExtractInstallmentsAmount(card)
		if err != nil {
			logs.Error(err)
			continue
		}
		product.InstallmentsAmount = installments_amount

		// Adiciona o produto ao slice
		*products = append(*products, product)
	}

	log.Printf("Total de produtos processados: %d", len(*products))
}
