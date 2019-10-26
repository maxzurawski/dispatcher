package sensors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/maxzurawski/dispatcher/publishers"

	"github.com/maxzurawski/dispatcher/config"
)

func Init(processId string) error {
	proxy := config.DispatcherConfig().ProxyService()
	url := fmt.Sprintf("%s/api/register/cachesensors/", proxy)
	response, err := http.Get(url)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not obtain sensors", err.Error())
		return err
	}

	var cachedSensors []RegisteredSensor
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not read response body", err.Error())
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		publishers.Logger().Info(processId, "", "there are no sensors present yet")
		Cache = &cache{}
		return nil
	}

	err = json.Unmarshal(body, &cachedSensors)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not decode cached sensors", err.Error())
		return err
	}

	publishers.Logger().Info(processId, "", "resetting sensors register cache")
	Cache = &cache{}
	Cache.FillCache(cachedSensors)
	return nil
}
