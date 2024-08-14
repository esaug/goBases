package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

func NewBook(Title string, Author string, Pages int) *Book {
	book := Book{
		Title,
		Author,
		Pages,
	}

	return &book
}

// Getter Setters
func (b *Book) SetTittle(title string) {
	b.Title = title
}

func (b *Book) GetTitle() string {
	return b.Title
}

type TextBook struct {
	Book
	editorial string
	level     uint
}

func NewTextBook(Title string, Author string, Pages int, editorial string, level uint) TextBook {
	return TextBook{
		Book:      Book{Title, Author, Pages},
		editorial: editorial,
		level:     level,
	}

}

func (b *TextBook) PrintTextBookInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %v, Editorial: %s, Level: %v \n",
		b.Title, b.Author, b.Pages, b.editorial, b.level)
}
