package service

import (
	"restful/dao"
	"restful/entity"
)

func Searchbymonthandday(day, month int) entity.RecordSlice {
	r := dao.Findbydayandmonth(day, month)
	return r
}

func Searchbyusertoreturn(userid int) entity.RecordSlice {
	return dao.FindRecordstoReturn(userid)
}
func SearchtoReturn() entity.RecordSlice {
	return dao.FindAllRecordstoReturns()
}
