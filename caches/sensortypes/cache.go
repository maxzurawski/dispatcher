package sensortypes

import "sync"

type cache map[string]string

var (
	lock = sync.Mutex{}
)

var TypesCache *cache
