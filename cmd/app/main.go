package main

import (
	"YMarketParser/internal/scraper"
)

func main() {

	ProductScraper := scraper.NewProductScraper("https://market.yandex.ru/product--sudokrem-zashchitnyi-i-uspokaivaiushchii-pod-podguznik-ot-oprelostei-i-pokrasnenii-dlia-novorozhdennykh/1728462711?nid=18044745&show-uid=16244249620444773635616002&context=search&onstock=1&sku=100235109248&cpc=_vNhijy6MiigojBVltLnIsWtN6BsXi0x4y2GxLyA-LoVMQCdlI4TdZZZmUXWxfhKVF2r8ZZgFjwqsw_U4ud4mlman6BJFIN4kh62h8hyD4IiRWAwCdMwMJ0RqTnbBIFrIbBadrZ_Fb5Fy1GIerWDNUHRiKG_UhxI0FDJCsKPx-Fu07mVpwtDpQ%2C%2C&do-waremd5=6ZRsCsU5i2tZoRvMf6Gp_A")

	scraper.NewClient(ProductScraper)

}
