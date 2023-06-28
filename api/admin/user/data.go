package user

type (
	User struct {
		Id int 				`json:"id"`
		Uid string 			`json:"uid"`
		Name string 		`json:"name"`
		Gender string 		`json:"gender"`
		Birthday string 	`json:"birthday"`
		Height string 		`json:"height"`
		Weight string 		`json:"weight"`
		Hometown string 	`json:"hometown"`
		Education string 	`json:"education"`
		School string 		`json:"school"`
		Work string 		`json:"work"`
		CompanyType string 	`json:"co_type"`
		Income string 		`json:"income"`
		HouseCar string 	`json:"house_car"`
		MobilePhone string 	`json:"m_phone"`
		WeiXin string 		`json:"wei_xin"`
		Marriage string 	`json:"marriage"`
		Habit string 		`json:"habit"`
		Family string 		`json:"family"`
		Interest string 	`json:"interest"`
		Character string 	`json:"character"`
		FuturePlan string 	`json:"future_plan"`
		Values string 		`json:"values"`
		LoveView string 	`json:"love_view"`
		BestWish string 	`json:"best_wish"`
		BestHeight string 	`json:"best_height"`
		IsDivorce	string 	`json:"is_divorce"`
		Status int8         `json:"status"` // 0:未审核，1:已审核，2:vip
		Img   string 		`json:"img"`
		Level int8          `json:"level"`
		CreatedAt int       `json:"created_at"`
	}

	Hometown struct {
		Country string 		`json:"country"`
		Province string 	`json:"province"`
		City string 		`json:"city"`
		AddrDetail string 	`json:"addr_detail"`
	}
)