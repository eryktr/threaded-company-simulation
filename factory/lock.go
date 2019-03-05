package factory

import "sync"

var list_mutex = &sync.Mutex{}
var warehouse_mutex = &sync.Mutex{}
