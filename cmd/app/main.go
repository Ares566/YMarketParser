package main

import (
	"YMarketParser/internal/scraper"
)

func main() {

	productURLs := []string{
		"https://market.yandex.ru/product--weleda-detskoe-molochko-dlia-giperchuvstvitelnoi-kozhi-tela-s-alteem/1722845051?nid=18044746&show-uid=16252283348025817253216004&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100318425259&cpc=duAezd3xFTIcEnogw76L7iR_8OU-Rnayud9IefH2Wt1YLroQMpjVuv2D19Cu4te1zXfddaxjigPOQA1th-QdP3WV9GUvT72zW0qUUw50PJdJpLqB0kxiJjXwgk8WM6TTms3YQRKU9MHsEDo4Wg92IGnKlykY7IKco5iKNALdPXz4HFQbhi8hWQ%2C%2C&do-waremd5=Wesl8risSlBVIQeOcQUHug",
	}

	// парсим данные товара url
	//ProductScrapper := scraper.NewProductScraper(productURLs)

	// парсим остатки на складе для товара заданного url
	BasketScraper := scraper.NewBasketScraper(productURLs)

	scraper.NewClient(BasketScraper)

}
