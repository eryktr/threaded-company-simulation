package factory

import "sync"

var list_mutex = &sync.Mutex{}
var warehouse_mutex = &sync.Mutex{}

func lock_list() {
	list_mutex.Lock()
}

func unlock_list() {
	list_mutex.Unlock()
}

func lock_warehouse() {
	warehouse_mutex.Lock()
}

func unlock_warehouse() {
	warehouse_mutex.Unlock()
}
