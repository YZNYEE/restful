package entity

type Book struct {
	Id       int
	Bookname string
	//书的状态，false：还可以借，true：已借完
	Status bool
}

type BookSlice []*Book
