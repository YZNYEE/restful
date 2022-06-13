package entity

type Book struct {
	Id       int
	Bookname string
	RestNum  int
	//书的状态，0：未借，1：已借
	Status int
}

type BookSlice []*Book
