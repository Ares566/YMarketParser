package scraper

import (
	"context"
	"fmt"
	"math/rand"

	"time"

	"github.com/chromedp/chromedp"
)

type ProductBalance struct {
	Url     string
	Balance string
}

type BasketScraiper struct {
	urls     []string
	Balances []ProductBalance
}

func NewBasketScraper(_urls []string) *BasketScraiper {
	return &BasketScraiper{urls: _urls}
}

func (basketscraper *BasketScraiper) process(ctx context.Context) {

	basketCnt := "1000"

	rand.Seed(time.Now().UnixNano())

	for _, producturl := range basketscraper.urls {

		// задержка для совести
		randomSleep := 5 + rand.Intn(20)
		time.Sleep(time.Duration(randomSleep) * time.Second)

		balance := ProductBalance{Url: producturl}

		// ждем загрузки страницы
		err := chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 20, chromedp.Tasks{
				chromedp.Navigate(balance.Url),
				chromedp.Sleep(555 * time.Millisecond),
				chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),

				// добавляем в корзину товар
				chromedp.Click(`//button[@class='zsSJkfeAPw _16jABpOZ2- LS3-2-cZ2Z _2Sz75Y384m _3iugb2hKWO _2VlTHnWxF8 _1jnLi6H271']`),
				chromedp.Sleep(2000 * time.Millisecond),
			}),
		)

		// TODO нужно добавить лог ошибок и его анализатор чтоб поймать момент когда парсинг уже в принципе не работает
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
			continue
		}

		err = chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 20, chromedp.Tasks{
				// переходим в корзину
				chromedp.Navigate("https://pokupki.market.yandex.ru/my/cart"),
				chromedp.Sleep(5000 * time.Millisecond),
				chromedp.WaitVisible("div.b_3bUkPPXxr5", chromedp.ByQuery),
				chromedp.Sleep(2000 * time.Millisecond),

				// увеличиваем счетчик товаров до 10000
				chromedp.SendKeys(`//div[@class='b_3I6VbdL-46']//input[@class='b_3qxDpnyrNf']`, basketCnt),
				chromedp.Sleep(3000 * time.Millisecond),

				// проверяем остатки на складе
				chromedp.InnerHTML(`//div[@class='b_iVxPKotZvY']//div[@class='cia-vs']`, &balance.Balance),
			}),
		)

		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
			continue
		}

		// TODO добавить условия достаточности информации для вставки в БД
		basketscraper.Balances = append(basketscraper.Balances, balance)

	}

	for _, balance := range basketscraper.Balances {

		fmt.Printf("URL: %s\n", balance.Url)
		fmt.Printf("Остаток: %s\n", balance.Balance)

	}

}
