package repository

func NewMConfigRepository() *configRepository {
	return newConfigRepository(&mConfig{})
}

func NewIntConfigRepository() *configRepository {
	return newConfigRepository(&intConfig{})
}

func NewIntDConfigRepository() *configRepository {
	return newConfigRepository(&intDConfig{})
}

//----------------------------------------------------------------------------------------------

func NewIntHisRepository() *hisRepository {
	return newHisRepository(&intHis{})
}

//----------------------------------------------------------------------------------------------

func NewDingTalkRobotRepository() *notifyRepository {
	return newNotifyRepository(&dingTalkRobotConfig{})
}
