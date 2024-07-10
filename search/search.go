package search

import (
	"github.com/Shhwip/bgate-scraper/reader/model"
)

type Searcher interface {
	Query(query string) ([]model.Verse, []model.Footnote, error)
	Booklist() ([]model.Book, error)
	Translation() string
}
