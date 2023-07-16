package auth

type AccountInfo struct {
	Id          uint   `json:"id"`
	UserId      string `json:"user_id"`
	MobilePhone string `json:"mobile_phone"`
	Account     string `json:"account"`
	AccountType string `json:"account_type"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Avatar      string `json:"avatar"`
	Status      int8   `json:"status"`
	LastLoginAt int64  `json:"last_login_at"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	Remark      string `json:"remark"`
}
