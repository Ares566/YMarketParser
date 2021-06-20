package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"),
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel = chromedp.NewContext(
		ctx,
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 40*time.Second)
	defer cancel()

	var bcrumbs []*cdp.Node
	var productname string
	var productlabel string
	var productmark string
	var productrpl string
	var productOldPrice string
	var productDiscount string
	var productCurPrice string

	err := chromedp.Run(
		ctx,

		chromedp.Navigate("https://market.yandex.ru/product--weleda-maslo-dlia-mladentsev-s-kalenduloi-s-nezhnym-aromatom/1722845046?nid=18044746&show-uid=16241930550221652064616007&context=search&glfilter=7893318%3A8513199&onstock=1&sku=100242020836&cpc=aqIWvBZI52-Rmes1bmE-m7-vpAC0mGb6RK30-W1r7DxZBgwESAIeL4gLBUhip-6kUvaIhQjauStXBAg8SGFpmht7TDQ-UH65pweDw0fbsGw-GbfDt-rLiBmsJy5ojcVoTe6r2ZPmciZ0I21IqWD2LjQ6cvf6zXK61_Owi3jj7T0DnjOKqPQrog%2C%2C&do-waremd5=lKaY33YMvA80DANTDVUCaA"),
		chromedp.Sleep(555*time.Millisecond),
		chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),
		//chromedp.Nodes(`//li[@itemtype='https://schema.org/ListItem']`, &bcrumbs), //span[@itemprop='name']
		chromedp.Nodes(`//span[@itemprop='name']/text()`, &bcrumbs), //span[@itemprop='name']
		chromedp.InnerHTML(`//div[@class='x__tNeZtug']/div/h1`, &productname),
		chromedp.InnerHTML(`//div[@class='QldO3ndWeZ']`, &productlabel),
		chromedp.InnerHTML(`//span[@class='_3nFvoU2Uov']`, &productmark),
		chromedp.InnerHTML(`//a[@class='_27nuSZ19h7 _2J5l3Ahj0X _15qW8ohHT4 cia-cs']`, &productrpl),
		chromedp.InnerHTML(`//div[@class="h2EBR70V3s _34FT3corIW"]/span[@class='_5cNkUrAI2M']`, &productOldPrice),
		chromedp.InnerHTML(`//div[@class='_1xSlMTtCCt _1emLjCV7YT']/span[@class='_1L-3GDwnCL']`, &productDiscount),
		chromedp.InnerHTML(`//div[@class='_3NaXxl-HYN _3kWlKUNlTg _3WgcBT2Eyd']/span/span[1]`, &productCurPrice),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Breadcrumbs: ")
	for _, item := range bcrumbs {
		fmt.Print("  ► ")
		fmt.Println(item.NodeValue)
	}

	fmt.Printf("Название: %s\n", productname)
	fmt.Printf("Метка: %s\n", productlabel)
	fmt.Printf("Рейтинг: %s\n", productmark)
	fmt.Printf("Отзывы: %s\n", productrpl)
	fmt.Printf("Цена без скидки: %s\n", productOldPrice)
	fmt.Printf("Цена со скидкой: %s\n", productCurPrice)
	fmt.Printf("Дисконт: %s\n", productDiscount)

}

func NodeValues(nodes []*cdp.Node) []string {
	var vs []string
	for _, n := range nodes {
		vs = append(vs, n.NodeValue)
	}
	return vs
}
