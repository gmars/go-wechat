package util

func PageCondition(page int, pageSize, limitCount int) (offset, count int) {

	if page < 1 {
		page = 1
	}

	if pageSize > limitCount {
		pageSize = limitCount
	}

	offset = (page - 1) * pageSize
	count = pageSize
	return
}
