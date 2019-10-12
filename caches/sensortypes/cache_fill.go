package sensortypes

func (c cache) Fill(cachedTypes []CachedTypes) {
	lock.Lock()
	defer lock.Unlock()
	for _, item := range cachedTypes {
		c[item.Type] = item.Topic
	}
}
