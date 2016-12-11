package checkparam

import (
	"container/list"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/server"
)

type CheckHandler func(*CheckItem, *server.HttpContext) error

type CheckItem struct {
	Name      string
	ErrorCode string
	Handler   CheckHandler
}

func GenCheckItem(name string, errorCode string, handler CheckHandler) *CheckItem {
	return &CheckItem{name, errorCode, handler}
}

type tCheckItemsMap map[string]*list.List

var checkItemsMap tCheckItemsMap

func init() {
	checkItemsMap = tCheckItemsMap{}
}

func getOrCreateCheckItemMap(path string) *list.List {
	checkItemList, ok := checkItemsMap[path]
	if !ok {
		checkItemList = list.New()
		checkItemsMap[path] = checkItemList
		logger.Debug("Create check item map for %s", path)
	}
	return checkItemList
}

func Add(path string, item *CheckItem) {
	checkItemList := getOrCreateCheckItemMap(path)
	checkItemList.PushBack(item)
}

func MultiAdd(path string, items []*CheckItem) {
	checkItemList := getOrCreateCheckItemMap(path)
	for _, item := range items {
		checkItemList.PushBack(item)
	}
}
