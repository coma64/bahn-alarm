package handlers

import "github.com/coma64/bahn-alarm-backend/squeries"

type BahnAlarmApi struct {
	queries *squeries.Queries
}

func NewBahnAlarmApi(db squeries.DBTX) *BahnAlarmApi {
	return &BahnAlarmApi{queries: squeries.New(db)}
}

func defaultPagination(userPage, userSize *int) (offset, size int) {
	size = 50
	if userSize != nil {
		size = *userSize
	}

	offset = 0
	if userPage != nil {
		offset = *userPage * size
	}

	return offset, size
}
