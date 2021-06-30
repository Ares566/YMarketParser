package scraper

import (
	"context"
	"fmt"
	"math/rand"
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
type ProductScraper struct {
	urls  []string
	Сards []ProductCard // пул для вставки в БД
}

func NewProductScraper(_urls []string) *ProductScraper {
	return &ProductScraper{urls: _urls}
}

func (productscrapper *ProductScraper) process(ctx context.Context) {

	rand.Seed(time.Now().UnixNano())

	for _, producturl := range productscrapper.urls {

		// задержка для совести
		randomSleep := 5 + rand.Intn(20)
		time.Sleep(time.Duration(randomSleep) * time.Second)

		product := ProductCard{url: producturl}

		// ждем загрузки страницы
		err := chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 20, chromedp.Tasks{
				chromedp.Navigate(product.url),
				chromedp.Sleep(555 * time.Millisecond),
				chromedp.WaitVisible("div.x__tNeZtug", chromedp.ByQuery),
			}),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		// забираем категорию товара, подкатегорию и название, на это даем 3 сек
		err = chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 3, chromedp.Tasks{
				chromedp.Nodes(`//span[@itemprop='name']/text()`, &product.bcrumbs),
				chromedp.InnerHTML(`//div[@class='x__tNeZtug']/div/h1`, &product.name),
			}),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		// забираем метку, рейтинг и отзывы
		err = chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 3, chromedp.Tasks{
				chromedp.InnerHTML(`//div[@class='QldO3ndWeZ']`, &product.label),
				chromedp.InnerHTML(`//span[@class='_3nFvoU2Uov']`, &product.mark),
				chromedp.InnerHTML(`//a[@class='_27nuSZ19h7 _2J5l3Ahj0X _15qW8ohHT4 cia-cs']`, &product.rpl),
			}),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		// старой цены может и не быть
		err = chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 2, chromedp.Tasks{
				chromedp.InnerHTML(`//div[@class="h2EBR70V3s _34FT3corIW"]/span[@class='_5cNkUrAI2M']`, &product.oldPrice),
			}),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		// дисконта тоже может не быть
		err = chromedp.Run(
			ctx,
			RunWithTimeOut(&ctx, 2, chromedp.Tasks{
				chromedp.InnerHTML(`//div[@class='_1xSlMTtCCt _1emLjCV7YT']/span[@class='_1L-3GDwnCL']`, &product.discount),
			}),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		//  забираем цену
		err = chromedp.Run(
			ctx,
			chromedp.InnerHTML(`//div[@class='_3NaXxl-HYN _3kWlKUNlTg _3WgcBT2Eyd']/span/span[1]`, &product.curPrice),
		)
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			fmt.Printf("%s", err)
		}

		// TODO добавить условия достаточности информации для вставки в БД
		productscrapper.Сards = append(productscrapper.Сards, product)

	}


	for _, product := range productscrapper.Сards {

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
		fmt.Printf("Дисконт: %s\n\n\n\n", product.discount)

	}
}
