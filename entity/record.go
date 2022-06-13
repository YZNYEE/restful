package entity

import "time"

type Record struct {
	Bookid     int
	Userid     int
	Borrowtime time.Time
	expired    time.Time
}

type RecordSlice []*Record
