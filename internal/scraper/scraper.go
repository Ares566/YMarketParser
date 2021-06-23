package scraper

import "context"

type IScraper interface {
	process(ctx context.Context)
}
