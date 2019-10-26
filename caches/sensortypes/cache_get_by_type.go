package sensortypes

import "github.com/xdevices/utilities/stringutils"

func (c cache) GetByType(sensorType string) *string {
	lock.Lock()
	defer lock.Unlock()

	value := c[sensorType]

	if stringutils.IsZero(value) {
		return nil
	}
	return &value
}
