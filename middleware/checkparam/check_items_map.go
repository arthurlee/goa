package checkparam

import (
	"container/list"
	"github.com/arthurlee/goa/logger"
)

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

func Add(path string, item CheckBase) {
	checkItemList := getOrCreateCheckItemMap(path)
	checkItemList.PushBack(item)
}

func MultiAdd(path string, items []CheckBase) {
	checkItemList := getOrCreateCheckItemMap(path)
	for _, item := range items {
		checkItemList.PushBack(item)
	}
}
