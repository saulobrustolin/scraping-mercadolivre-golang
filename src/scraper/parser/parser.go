package parser

import (
	"errors"
	"scrapping-mercadolivre-golang/src/logs"
	"strconv"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func ExtractPrice(card playwright.Locator) (float64, error) {
	// captura o valor inteiro do preço
	capture_integer, err := card.Locator("div.poly-card__content > div.poly-component__price > div.poly-price__current > span.andes-money-amount.andes-money-amount--cents-superscript > span.andes-money-amount__fraction").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço."))
		return 0, err
	}

	// captura o valor decimal do preço
	capture_cents, err := card.Locator("div.poly-card__content > div.poly-component__price > div.poly-price__current > span.andes-money-amount.andes-money-amount--cents-superscript > span.andes-money-amount__cents andes-money-amount__cents--superscript-24").InnerText()
	if err != nil {
		capture_cents = "0"
	}

	// convertendo para float64
	integer, err := strconv.ParseFloat(capture_integer, 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de inteiros"))
		return 0, err
	}
	cents, err := strconv.ParseFloat(("0," + capture_cents), 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de centavos"))
		return 0, err
	}

	return float64(integer + cents), nil
}

func ExtractTitle(card playwright.Locator) (string, error) {
	// captura o título do produto
	title, err := card.Locator("div.poly-card__content > h3.poly-component__title-wrapper > a.poly-component__title").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração do título."))
		return "", err
	}

	return string(title), nil
}

func ExtractURL(card playwright.Locator) (string, error) {
	// captura a URL do produto
	URL, err := card.Locator("div.poly-card__content > h3.poly-component__title-wrapper > a.poly-component__title").GetAttribute("href")
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração do URL da imagem."))
		return "", err
	}

	return string(URL), nil
}

func ExtractQuantityReviews(card playwright.Locator) (int64, error) {
	// captura a quantidade de avaliações feitas no anúncio
	capture_quantity, err := card.Locator("div.poly-card__content > div.poly-component__reviews > span.poly-reviews__total").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de quantidade de reviews."))
		return 0, err
	}

	// tratamento para garantir que seja um número na hora de converter
	capture_quantity = strings.TrimSpace(capture_quantity)
	capture_quantity = strings.Trim(capture_quantity, "()")
	capture_quantity = strings.Join(strings.Fields(capture_quantity), "")

	// convertendo para int64
	quantity_reviews, err := strconv.ParseInt(capture_quantity, 10, 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de quantidade de reviews."))
		return 0, err
	}

	return int64(quantity_reviews), nil
}

func ExtractPicture(card playwright.Locator) (string, error) {
	// captura a URL da picture do produto

	URL, err := card.Locator("div.poly-card__portada > img.poly-component__picture").GetAttribute("src")
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de imagem."))

		return "", err
	}

	if strings.Contains(URL, "https://") {
		return string(URL), nil
	}

	URL, err = card.Locator("div.poly-card__portada > img.poly-component__picture").GetAttribute("data-src")
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de imagem."))
		return "", err
	}

	return string(URL), nil
}

func ExtractStars(card playwright.Locator) (float64, error) {
	capture_stars, err := card.Locator("div.poly-card__content > div.poly-component__reviews > span.poly-reviews__rating").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de notas de estrelas."))
		return 0, err
	}

	stars, err := strconv.ParseFloat(capture_stars, 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de notas de estrelas. ~ conversão de string para float64"))
		return 0, err
	}

	return float64(stars), nil
}

func ExtractAnchorPrice(card playwright.Locator) (float64, error) {
	// captura o valor inteiro do preço
	capture_integer, err := card.Locator("div.poly-card__content > div.poly-component__price > .andes-money-amount.andes-money-amount--previous.andes-money-amount--cents-comma > span.andes-money-amount__fraction").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço de ancoragem. ~ preço inteiro não encontrado"))
		return 0, err
	}

	// captura o valor decimal do preço
	capture_cents, err := card.Locator("div.poly-card__content > div.poly-component__price > .andes-money-amount.andes-money-amount--previous.andes-money-amount--cents-comma > span.andes-money-amount__cents").InnerText()
	if err != nil {
		capture_cents = "0"
	}

	// convertendo para float64
	integer, err := strconv.ParseFloat(capture_integer, 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de inteiros"))
		return 0, err
	}
	cents, err := strconv.ParseFloat(("0," + capture_cents), 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de centavos"))
		return 0, err
	}

	return float64(integer + cents), nil
}

func ExtractIsFreeShipping(card playwright.Locator) (bool, error) {
	capture_string, err := card.Locator("div.poly-component__shipping > span:first-of-type").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de verificação de frete grátis. ~ elemento não encontrado"))
	}

	if strings.Contains(capture_string, "grátis") {
		return true, nil
	}
	return false, nil
}

func ExtractInstallments(card playwright.Locator) (int64, error) {
	capture_installments, err := card.Locator("div.poly-component__price > span.poly-price__installments").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de número de parcelas. ~ elemento não encontrado"))
		return 0, err
	}

	logs.Sucess(capture_installments)
	return 0, nil
}

func ExtractInstallmentsAmount(card playwright.Locator) (float64, error) {
	// captura o valor inteiro do preço
	capture_integer, err := card.Locator("div.poly-card__content > span.poly-price__installments > span.andes-money-amount poly-phrase-price andes-money-amount--cents-comma > span.andes-money-amount__fraction").InnerText()
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço de ancoragem. ~ preço inteiro não encontrado"))
		return 0, err
	}

	// captura o valor decimal do preço
	capture_cents, err := card.Locator("div.poly-card__content > span.poly-price__installments > span.andes-money-amount poly-phrase-price andes-money-amount--cents-comma > span.andes-money-amount__cents").InnerText()
	if err != nil {
		capture_cents = "0"
	}

	// convertendo para float64
	integer, err := strconv.ParseFloat(capture_integer, 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de inteiros"))
		return 0, err
	}
	cents, err := strconv.ParseFloat(("0," + capture_cents), 64)
	if err != nil {
		logs.Error(errors.New("Erro na seção de extração de preço. ~ conversão de centavos"))
		return 0, err
	}

	return float64(integer + cents), nil
}
