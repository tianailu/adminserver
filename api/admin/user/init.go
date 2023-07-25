package user

func init() {
	go dealUserData()
}

func InitTable() {
	go createTable()
}
