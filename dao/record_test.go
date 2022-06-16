package dao

import (
	"fmt"
	"testing"
)

func TestAddRecord(t *testing.T) {
	//AddRecord(1, 1, 10)
}

func TestFindbyday(t *testing.T) {
	r := Findbyday(2)
	for _, v := range r {
		fmt.Println(v)
	}
}
func TestFindbyuser(t *testing.T) {
	r := Findbyuser(4)
	for _, v := range r {
		fmt.Println(v)
	}
}
func TestFindRecordstoReturn(t *testing.T) {
	r := FindRecordstoReturn(1)
	for _, v := range r {
		fmt.Println(v)
	}
}

func TestFindAllRecordstoReturns(t *testing.T) {
	r := FindAllRecordstoReturns()
	for _, v := range r {
		fmt.Println(v)
	}

}
