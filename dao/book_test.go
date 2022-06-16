package dao

import (
	"fmt"
	"testing"
)
import "restful/utils"

func TestAddBook(t *testing.T) {
	//AddBook("testbook")
	for i := 0; i < 100; i++ {
		n := utils.RandString(4)
		AddBook(n)
	}
}
func TestSelectbyOffetandlimit(t *testing.T) {
	b := SelectbyOffetandlimit(0, 5)
	for _, v := range b {
		fmt.Println(v)
	}
	//fmt.Println(b)

}

func TestUpdate(t *testing.T) {
	b := FindbyId(1)
	fmt.Println(b)
	UpdateBookStatus(1)
	fmt.Println(FindbyId(1))
}
func TestFindbyId(t *testing.T) {
	b := FindbyId(1)
	fmt.Println(b)
}

func TestUpdateBookStatus(t *testing.T) {
	UpdateBookStatus(1)
}

func TestFindbyname(t *testing.T) {
	b := Findbyname("Y%1Y878")
	fmt.Println(len(b))

}
