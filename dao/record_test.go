package dao

import (
	"fmt"
	"testing"
)

func TestAddRecord(t *testing.T) {
	//AddRecord(1, 1, 10)
}

func TestFindlatestamonth(t *testing.T) {
	r := Findlatestamonth()
	for _, v := range r {
		fmt.Println(v)
	}
}
func TestFindbyuser(t *testing.T) {
	r := Findbyuser(1)
	for _, v := range r {
		fmt.Println(v)
	}
}
