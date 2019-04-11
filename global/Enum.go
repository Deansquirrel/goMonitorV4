package global

type ConfigType int

const (
	CM ConfigType = iota
	CInt
	CIntD
	CHealth
	CWebState
	CCrmDzXfTest
)

type NotifyType int

const (
	NDingTalkRobot NotifyType = iota
)

type ConfigHisType int

const (
	HInt ConfigHisType = iota
	HWebState
	HCrmDzXfTest
)

type ActionType int

const (
	AWindowsService ActionType = iota
	AIISAppPool
)

type ActionHisType int

const (
	HWindowsService ActionHisType = iota
	HIISAppPool
)
