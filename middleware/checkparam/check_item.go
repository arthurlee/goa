package checkparam

import ()

type CheckItem struct {
	name      string
	errorCode string
	handler   CheckHandler
}

func (me *CheckItem) GetName() string {
	return me.name
}

func (me *CheckItem) GetErrorCode() string {
	return me.errorCode
}

func (me *CheckItem) GetHandler() CheckHandler {
	return me.handler
}

func GenCheckItem(name string, errorCode string, handler CheckHandler) CheckBase {
	return &CheckItem{name, errorCode, handler}
}
