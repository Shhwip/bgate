package model

import (
	"fmt"

	"github.com/Shhwip/bgate-scraper/reader/style"
)

type Verse struct {
	Book    string  `db:"book"`
	Chapter int     `db:"chapter"`
	Number  int     `db:"number"`
	Part    int     `db:"part"`
	Text    string  `db:"text"`
	Title   *string `db:"title"`
}

type Footnote struct {
	Book    string `db:"book"`
	Chapter int    `db:"chapter"`
	Number  int    `db:"number"`
	Text    string `db:"text"`
	Verse   int    `db:"verse_id"`
}

type Book struct {
	Name     string `db:"name"`
	Chapters int    `db:"chapters"`
}

func (v Verse) HasTitle() bool {
	return v.Title != nil
}

func (v Verse) TitleString() string {
	return style.TitleStyle.Render(*v.Title)
}

func (v Verse) ChapterString() string {
	text := fmt.Sprintf(" %s: %d ", v.Book, v.Chapter)
	return style.ChapterStyle.Render(text)
}

func (v Verse) NumberString() string {
	text := fmt.Sprintf("%d ", v.Number)
	return style.NumberStyle.Render(text)
}

func (b Book) String() string {
	return fmt.Sprintf("%s (%d)", style.BookStyle.Render(b.Name), b.Chapters)
}
