package sensors

func (c cache) GetByUuid(uuid string) *RegisteredSensor {
	lock.Lock()
	defer lock.Unlock()

	if sensor, ok := c[uuid]; !ok {
		return nil
	} else {
		return &sensor
	}
}
