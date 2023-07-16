package page

func CalPageOffset(pageNum, pageSize int) (int, int) {
	if pageNum <= 0 {
		pageNum = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (pageNum - 1) * pageSize

	return offset, pageSize
}
