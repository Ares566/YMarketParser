package scraper

import (
	"context"

	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type Client struct {
}

func NewClient(scraper IScraper) {

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
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	// defer cancel()

	scraper.process(ctx)

}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
