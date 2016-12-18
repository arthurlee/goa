package checkparam

import ()

type CheckItem struct {
	name      string
	errorCode string
	required  bool
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

func (me *CheckItem) IsRequired() bool {
	return me.required
}

func GenCheckItem(name string, errorCode string, required bool, handler CheckHandler) CheckBase {
	return &CheckItem{name, errorCode, required, handler}
}
