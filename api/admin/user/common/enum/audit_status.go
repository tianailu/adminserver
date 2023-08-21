package enum

type AuditStatus int8

const (
	UnknownAuditStatus AuditStatus = iota
	NoAudit
	AuditPass
	AuditFailed
	AuditAgain
)

func GetAuditStatusWithValue(value int8) AuditStatus {
	switch value {
	case int8(NoAudit):
		return NoAudit
	case int8(AuditPass):
		return AuditPass
	case int8(AuditFailed):
		return AuditFailed
	case int8(AuditAgain):
		return AuditAgain
	default:
		return UnknownAuditStatus
	}
}

func (e AuditStatus) Verify() bool {
	return e != UnknownAuditStatus
}

func (e AuditStatus) Value() int8 {
	return int8(e)
}
