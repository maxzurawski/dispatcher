package sensortypes

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
	url := fmt.Sprintf("%s/api/sensortypes/cachetypes/", proxy)
	response, err := http.Get(url)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not obtain sensors types", err.Error())
		return err
	}

	var cachedTypes []CachedTypes
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not read response body", err.Error())
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		publishers.Logger().Info(processId, "", "there are no sensors types present yet")
		TypesCache = &cache{}
		return nil
	}
	err = json.Unmarshal(body, &cachedTypes)
	if err != nil {
		publishers.Logger().Error(processId, "", "could not decode cached sensors types", err.Error())
		return err
	}

	publishers.Logger().Info(processId, "", "resetting sensors types cache")
	TypesCache = &cache{}
	TypesCache.Fill(cachedTypes)
	return nil
}
