package main

import (
	"YMarketParser/internal/scraper"
)

func main() {

	url := "https://market.yandex.ru/product--weleda-maslo-dlia-massazha-zhivotika-mladentsev/1722845054?cpa=1&cpc=jicXAlT-HIQ9e34t44x25QFcUTndCaki88Tm10jzI54uc2j9xhEgiKuSBt5SccV9nuUpLAknGx_wrL1Dtq18pJT4iqmyALWzzKyKF8QO0G1EhGsgVMpFN-F3yBNGzLdijxsUDwYP0V-CeR2zkoG7OIWrtiUQlFZfWlBVZ0t0edK2uN8GsW08DQ%2C%2C&sku=100237420679&do-waremd5=0qOLKtWpzaxppJIIortPFQ"

	// парсим данные товара url
	//Scraper := scraper.NewProductScraper(url)

	// парсим остатки на складе для товара заданного url
	Scraper := scraper.NewBasketScraper(url)

	scraper.NewClient(Scraper)

}
