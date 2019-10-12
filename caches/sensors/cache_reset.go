package sensors

func (c *cache) Reset() {
	*c = make(map[string]RegisteredSensor)
}
