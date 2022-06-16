package entity

import "time"

type Record struct {
	Bookid  int
	Userid  int
	Borrow  time.Time
	Expired time.Time
}

type RecordSlice []*Record
