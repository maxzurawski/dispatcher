package sensortypes

func (c *cache) Reset() {
	*c = make(map[string]string)
}
