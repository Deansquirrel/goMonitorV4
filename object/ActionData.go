package object

type IISAppPoolActionData struct {
	FId            string
	FAgentServerId string
	FName          string
	FCheckTimes    int
}

func (isd *IISAppPoolActionData) GetCheckTimes() int {
	return isd.FCheckTimes
}

type WindowsServiceActionData struct {
	FId            string
	FAgentServerId string
	FName          string
	FCheckTimes    int
}

func (wsd *WindowsServiceActionData) GetCheckTimes() int {
	return wsd.FCheckTimes
}
