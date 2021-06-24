package scraper

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/chromedp/chromedp"
)

type Basket struct {
	url          string // это не URL корзинки, это URL товара который нужно положить в корзину
	innerHTML    string
	stockBalance string
}

func NewBasketScraper(_url string) *Basket {
	return &Basket{url: _url}
}

func (basket *Basket) process(ctx context.Context) {

	basketCnt := "1000"

	err := chromedp.Run(
		ctx,

		chromedp.Navigate(basket.url),
		chromedp.Sleep(2000*time.Millisecond),
		chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),
		chromedp.Sleep(2000*time.Millisecond),

		// добавляем в корзину товар
		chromedp.Click(`//button[@class='zsSJkfeAPw _16jABpOZ2- LS3-2-cZ2Z _2Sz75Y384m _3iugb2hKWO _2VlTHnWxF8 _1jnLi6H271']`),
		chromedp.Sleep(2000*time.Millisecond),

		// переходим в корзину
		chromedp.Navigate("https://pokupki.market.yandex.ru/my/cart"),
		chromedp.Sleep(5000*time.Millisecond),
		chromedp.WaitVisible("div.b_3bUkPPXxr5", chromedp.ByQuery),
		chromedp.Sleep(2000*time.Millisecond),

		// TODO для проверки, потом убрать
		chromedp.InnerHTML(`div.b_3bUkPPXxr5`, &basket.innerHTML, chromedp.ByQuery),

		// увеличиваем счетчик товаров до 10000
		chromedp.Value(`//div[@class='b_3I6VbdL-46']//input[@class='b_3qxDpnyrNf']`, &basketCnt),

		// проверяем остатки на складе
		// <div class="cia-vs" data-zone-name="bad-count-notification">В наличии 5 штук</div>
		// TODO нужно разобраться, не всегда попадают нужные данные
		chromedp.InnerHTML(`div.cia-vs`, &basket.stockBalance, chromedp.ByQuery),
	)

	fmt.Printf("Basket innerHTML: %s \n", basket.innerHTML)
	fmt.Printf("Stock Balance: %s \n", basket.stockBalance)

	if err != nil {
		log.Fatal(err)
	}

}
