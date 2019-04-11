package object

import "time"

type IISAppPoolActionHisData struct {
	FId            string
	FConfigId      string
	FAgentServerId string
	FName          string
	FCheckTimes    int
	FContent       string
	FOprTime       time.Time
}

type WindowsServiceHisData struct {
	FId            string
	FConfigId      string
	FAgentServerId string
	FName          string
	FCheckTimes    int
	FContent       string
	FOprTime       time.Time
}
