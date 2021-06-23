package scraper

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type ProductCard struct {
	url      string
	bcrumbs  []*cdp.Node
	name     string
	label    string
	mark     string
	rpl      string
	oldPrice string
	discount string
	curPrice string
}

func NewProductScraper(_url string) *ProductCard {
	return &ProductCard{url: _url}
}

func (product *ProductCard) process(ctx context.Context) {

	err := chromedp.Run(
		ctx,

		chromedp.Navigate(product.url),
		chromedp.Sleep(555*time.Millisecond),
		chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),
		//chromedp.Nodes(`//li[@itemtype='https://schema.org/ListItem']`, &bcrumbs),
		chromedp.Nodes(`//span[@itemprop='name']/text()`, &product.bcrumbs),
		chromedp.InnerHTML(`//div[@class='x__tNeZtug']/div/h1`, &product.name),
		chromedp.InnerHTML(`//div[@class='QldO3ndWeZ']`, &product.label),
		chromedp.InnerHTML(`//span[@class='_3nFvoU2Uov']`, &product.mark),
		chromedp.InnerHTML(`//a[@class='_27nuSZ19h7 _2J5l3Ahj0X _15qW8ohHT4 cia-cs']`, &product.rpl),
		chromedp.InnerHTML(`//div[@class="h2EBR70V3s _34FT3corIW"]/span[@class='_5cNkUrAI2M']`, &product.oldPrice),
		chromedp.InnerHTML(`//div[@class='_1xSlMTtCCt _1emLjCV7YT']/span[@class='_1L-3GDwnCL']`, &product.discount),
		chromedp.InnerHTML(`//div[@class='_3NaXxl-HYN _3kWlKUNlTg _3WgcBT2Eyd']/span/span[1]`, &product.curPrice),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Breadcrumbs: ")
	for _, item := range product.bcrumbs {
		fmt.Print("  ► ")
		fmt.Println(item.NodeValue)
	}

	fmt.Printf("Название: %s\n", product.name)
	fmt.Printf("Метка: %s\n", product.label)
	fmt.Printf("Рейтинг: %s\n", product.mark)
	fmt.Printf("Отзывы: %s\n", product.rpl)
	fmt.Printf("Цена без скидки: %s\n", product.oldPrice)
	fmt.Printf("Цена со скидкой: %s\n", product.curPrice)
	fmt.Printf("Дисконт: %s\n", product.discount)

}
