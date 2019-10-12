package sensors

import "sync"

type cache map[string]RegisteredSensor

var (
	lock sync.Mutex
)

var Cache *cache
