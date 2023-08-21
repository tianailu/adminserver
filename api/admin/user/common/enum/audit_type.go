package enum

type AuditType int32

const (
	UnknownAuditType AuditType = iota
	UserBaseInfoAudit
	WorkAuth
	EduAuth
	RealNameAuth
)

func GetAuditTypeByValue(v int32) AuditType {
	switch v {
	case int32(UserBaseInfoAudit):
		return UserBaseInfoAudit
	case int32(WorkAuth):
		return WorkAuth
	case int32(EduAuth):
		return EduAuth
	case int32(RealNameAuth):
		return RealNameAuth
	default:
		return UnknownAuditType
	}
}

func (e AuditType) Verify() bool {
	return e != UnknownAuditType
}
