package user

type (
	User struct {
		Id               int    `json:"id"`               // ID
		AccountId        string `json:"account_id"`       // 账号ID
		Uid              string `json:"uid"`              // 用户ID
		Name             string `json:"name"`             // 用户名
		Img              string `json:"img"`              // 个人头像
		Gender           string `json:"gender"`           // 性别
		Constellation    string `json:"constellation"`    // 星座
		Birthday         int64  `json:"birthday"`         // 生日
		Height           string `json:"height"`           // 身高
		Weight           string `json:"weight"`           // 体重
		Hometown         string `json:"hometown"`         // 家乡
		Education        string `json:"education"`        // 最高学历
		School           string `json:"school"`           // 毕业院校
		Work             string `json:"work"`             // 职业
		Company          string `json:"company"`          // 公司
		Income           string `json:"income"`           // 收入
		MobilePhone      string `json:"m_phone"`          // 手机号码
		WeChat           string `json:"wechat"`           // 微信
		Marriage         string `json:"marriage"`         // 婚姻状态
		Habit            string `json:"habit"`            // 生活习惯
		ConsumptionView  string `json:"consumption_view"` // 消费观
		Family           string `json:"family"`           // 家庭背景
		Interest         string `json:"interest"`         // 兴趣爱好
		LoveView         string `json:"love_view"`        // 爱情观
		TargetAppearance string `json:"ta_appearance"`    // 希望另一半的样子
		BeImpressed      string `json:"be_impressed"`     // 对方什么最能打动自己
		TargetAge        string `json:"ta_age"`           // 希望另一半年龄范围
		TargetHeight     string `json:"ta_height"`        // 希望另一半身高范围
		TargetCity       string `json:"ta_city"`          // 希望另一半所在城市
		TargetHometown   string `json:"ta_hometown"`      // 希望另一半的家乡
		TargetEducation  string `json:"ta_education"`     // 希望另一半最低学历
		TargetMarriage   string `json:"ta_marriage"`      // 希望另一半婚姻状态
		VipLevel         int8   `json:"vip_level"`        // 会员等级
		IdCard           string `json:"id_card"`          // 身份证
		RealName         string `json:"real_name"`        // 真实姓名
		IsRealNameAuth   int8   `json:"is_rn_auth"`       // 是否完成实名认证，0:未认证，1:已通过认证
		WorkAuthImg      string `json:"work_auth_img"`    // 工作认证图片
		IsWorkAuth       int8   `json:"is_work_auth"`     // 是否完成工作认证，0:未认证，1:已通过认证
		EduAuthImg       string `json:"edu_auth_img"`     // 学历认证图片
		IsEduAuth        int8   `json:"is_edu_auth"`      // 是否完成学历认证，0:未认证，1:已通过认证
		Status           int8   `json:"status"`           // 0:未审核，1:已审核，2:vip
		CreatePlace      string `json:"create_place"`     // 注册地
		CreatedAt        int64  `json:"created_at"`       // 注册时间
		UpdateAt         int64  `json:"update_at"`        // 修改时间
	}

	Hometown struct {
		Country    string `json:"country"`
		Province   string `json:"province"`
		City       string `json:"city"`
		AddrDetail string `json:"addr_detail"`
	}

	VipLevel struct {
		Id                int     `json:"id"`                 // ID
		Level             string  `json:"level"`              // 等级
		Name              string  `json:"name"`               // 名称
		Desc              string  `json:"desc"`               // 描述
		Price             float64 `json:"price"`              // 价格
		Discount          float32 `json:"discount"`           // 折扣
		DowngradeStrategy int8    `json:"downgrade_strategy"` // 降级策略
		Status            int8    `json:"status"`             // 状态，0:禁用，1:启用
		CreatedAt         int64   `json:"created_at"`         // 创建时间
	}

	VipTag struct {
		Id        int    `json:"id"`         // ID
		Name      string `json:"name"`       // 标签名
		CreatedAt int64  `json:"created_at"` // 创建时间
	}
)

func (m *User) TableName() string {
	return "tb_user"
}
