package main

import (
	"YMarketParser/internal/scraper"
)

func main() {

	url := "https://market.yandex.ru/product--weleda-detskii-krem-dlia-litsa-s-kalenduloi/1722845053?nid=18044746&show-uid=16245982200210986714116006&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100237415624&cpc=XqNhJZ6_HpdF_tnO2f2KNeHgF0KlmQ3bQEBFiSoBBCJlIlp7ijuQAL4Z2BgPQ1o3DqCqJAcgA7HAhTKa78dj5Y0tGhQp992D41suc9v0x-UFyZeqm7o86hA7flKN8D4XgQuHmdvrMcMn3L31B-s8dqlX9-ybLR5AOBulRKMtG0oR8woUoUeUdA%2C%2C&do-waremd5=Ez6gyjjEMoevdQa1OaVc_g"

	// парсим данные товара url
	//Scraper := scraper.NewProductScraper(url)

	// парсим остатки на складе для товара заданного url
	Scraper := scraper.NewBasketScraper(url)

	scraper.NewClient(Scraper)

}
