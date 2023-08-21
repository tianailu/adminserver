package enum

type UserStatus int8

const (
	UnknownUserStatus UserStatus = iota // 未知状态
	ActiveStatus                        // 活跃状态
	SuspensionStatus                    // 封号状态
	MuteStatus                          // 禁言状态
	LogoutStatus                        // 注销状态
)

func GetUserStatusWithValue(value int8) UserStatus {
	switch value {
	case int8(ActiveStatus):
		return ActiveStatus
	case int8(SuspensionStatus):
		return SuspensionStatus
	case int8(MuteStatus):
		return MuteStatus
	case int8(LogoutStatus):
		return LogoutStatus
	default:
		return UnknownUserStatus
	}
}

func (e UserStatus) IsNormalUser() bool {
	return e == ActiveStatus || e == MuteStatus
}

func (e UserStatus) Value() int8 {
	return int8(e)
}
