package user

func init() {
	go createTable()
	go dealUserData()
}
