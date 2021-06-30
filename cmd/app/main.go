package main

import (
	"YMarketParser/internal/scraper"
)

func main() {

	productURL := []string{
		"https://market.yandex.ru/product--weleda-maslo-dlia-mladentsev-s-kalenduloi-s-nezhnym-aromatom/1722845046?nid=18044746&show-uid=16250239923955073468616003&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100242020836&cpc=Yx9sPGU0iDU7srEKm7Q9niTzkEBn8l2H9XQzDp0-XZvZyuN9sGwr93EAFakavWsPknbiMgPlV5kP66oHh696GhaBqpirMeknp60A1tmpaMD2otpS9hRRDV0QpG4VRshFWzClueCC9AHXZngPwpRt5eSd93ajwACHficVcA4IfoWw05tQlbBb9A%2C%2C&do-waremd5=MM5ZEPEDe0pFL2iJNNnGTQ",
		"https://market.yandex.ru/product--weleda-detskoe-molochko-dlia-giperchuvstvitelnoi-kozhi-tela-s-alteem/1722845051?nid=18044746&show-uid=16250239923955073468616004&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100318425259&cpc=Yx9sPGU0iDVfmcrLvxRn7_kea6YJUHAmdoN9Glw9gAgrzs_52ajQLBGQJ2i7Lnx156JE2Ts7tbzXBHaUZRJFnfNc3QVo74Q6AiEQpRJLrtiadENmDXNsQpunyQ5EPmFvG-FBx0nKWUJplZFw-BW_UAFFZm4BB-5i0iusYByRFk8_q2is3eI9Fw%2C%2C&do-waremd5=Wesl8risSlBVIQeOcQUHug",
		"https://market.yandex.ru/product--weleda-detskii-balzam-zashchitnyi-ot-vetra-i-kholoda-s-kalenduloi/1722845037?nid=18044746&show-uid=16250239923955073468616009&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100234383551&cpc=uaV08jcA5Qbk4txyDVQcJRFQ2dTNLzWLld_I4n69z0Jq1QC1o_RK5dTmUgD_vk31FrpQWHoNg1dS_CWVxVI5PTCs9nmRROKkDZLmF-IAtQK4tPeZaj9BGnGmdtfgBJi-RPzw8SWsez6e1A5w2d0C2Z6JAzLazubNc3KP4j5yQzJyV7Rpdlsicg%2C%2C&do-waremd5=n2JLc8YVR23LxNrTklrFIQ",
	}

	// парсим данные товара url
	ProductScrapper := scraper.NewProductScraper(productURL)

	// парсим остатки на складе для товара заданного url
	//Scraper := scraper.NewBasketScraper(url)

	scraper.NewClient(ProductScrapper)

}
