package handlers

type BahnAlarmApi struct{}

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
