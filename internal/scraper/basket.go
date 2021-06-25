package scraper

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Basket struct {
	url          string // это не URL корзинки, это URL товара который нужно положить в корзину
	stockBalance string
	bcrumbs      []*cdp.Node
	name         string
	label        string
	mark         string
	rpl          string
	oldPrice     string
	discount     string
	curPrice     string
}

func NewBasketScraper(_url string) *Basket {
	return &Basket{url: _url}
}

func (basket *Basket) process(ctx context.Context) {

	basketCnt := "1000"

	err := chromedp.Run(
		ctx,

		chromedp.Navigate(basket.url),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),
		chromedp.Sleep(7000*time.Millisecond),

		// берем данные товара, раз уж загрузили
		chromedp.Nodes(`//span[@itemprop='name']/text()`, &basket.bcrumbs),
		chromedp.InnerHTML(`//div[@class='x__tNeZtug']/div/h1`, &basket.name),
		chromedp.InnerHTML(`//div[@class='QldO3ndWeZ']`, &basket.label),
		chromedp.InnerHTML(`//span[@class='_3nFvoU2Uov']`, &basket.mark),
		chromedp.InnerHTML(`//a[@class='_27nuSZ19h7 _2J5l3Ahj0X _15qW8ohHT4 cia-cs']`, &basket.rpl),
		chromedp.InnerHTML(`//div[@class="h2EBR70V3s _34FT3corIW"]/span[@class='_5cNkUrAI2M']`, &basket.oldPrice),
		chromedp.InnerHTML(`//div[@class='_1xSlMTtCCt _1emLjCV7YT']/span[@class='_1L-3GDwnCL']`, &basket.discount),
		chromedp.InnerHTML(`//div[@class='_3NaXxl-HYN _3kWlKUNlTg _3WgcBT2Eyd']/span/span[1]`, &basket.curPrice),

		// добавляем в корзину товар
		chromedp.Click(`//button[@class='zsSJkfeAPw _16jABpOZ2- LS3-2-cZ2Z _2Sz75Y384m _3iugb2hKWO _2VlTHnWxF8 _1jnLi6H271']`),
		chromedp.Sleep(2000*time.Millisecond),

		// переходим в корзину
		chromedp.Navigate("https://pokupki.market.yandex.ru/my/cart"),
		chromedp.Sleep(5000*time.Millisecond),
		chromedp.WaitVisible("div.b_3bUkPPXxr5", chromedp.ByQuery),
		chromedp.Sleep(2000*time.Millisecond),

		// увеличиваем счетчик товаров до 10000
		chromedp.SendKeys(`//div[@class='b_3I6VbdL-46']//input[@class='b_3qxDpnyrNf']`, basketCnt),
		chromedp.Sleep(3000*time.Millisecond),

		// проверяем остатки на складе
		chromedp.InnerHTML(`//div[@class='b_iVxPKotZvY']//div[@class='cia-vs']`, &basket.stockBalance),
	)

	fmt.Printf("URL: %s\n", basket.url)
	fmt.Print("Breadcrumbs: ")
	for _, item := range basket.bcrumbs {
		fmt.Print(" ► ")
		fmt.Print(item.NodeValue)
	}
	fmt.Println("")
	fmt.Printf("Название: %s\n", basket.name)
	fmt.Printf("Метка: %s\n", basket.label)
	fmt.Printf("Рейтинг: %s\n", basket.mark)
	fmt.Printf("Отзывы: %s\n", basket.rpl)
	fmt.Printf("Цена без скидки: %s\n", basket.oldPrice)
	fmt.Printf("Цена со скидкой: %s\n", basket.curPrice)
	fmt.Printf("Дисконт: %s\n", basket.discount)
	fmt.Printf("Остатки на складе: %s \n", basket.stockBalance)

	if err != nil {
		log.Fatal(err)
	}

}
