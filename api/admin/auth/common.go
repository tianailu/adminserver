package auth

type AccountInfo struct {
	Id          uint     `json:"id"`
	AccountId   string   `json:"account_id"`
	MobilePhone string   `json:"mobile_phone"`
	Account     string   `json:"account"`
	AccountType string   `json:"account_type"`
	Name        string   `json:"name"`
	Avatar      string   `json:"avatar"`
	Roles       []string `json:"roles"`
	Status      int8     `json:"status"`
	LoginCount  uint     `json:"login_count"`
	LastLoginIp string   `json:"last_login_ip"`
	LastLoginAt int64    `json:"last_login_at"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
	Remark      string   `json:"remark"`
}
