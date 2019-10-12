package sensors

func (c cache) FillCache(sensors []RegisteredSensor) {
	lock.Lock()
	defer lock.Unlock()
	for _, item := range sensors {
		c[item.Uuid] = item
	}
}
