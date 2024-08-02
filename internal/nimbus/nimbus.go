package nimbus

import (
	"fmt"
	"net/http"

	. "github.com/GearTech0/nimbus/pkg/raindropio"
)

// --------------CONSTANTS-------------------
const NIMBUS_CONFIG_TITLE = "nimbus.config"
const NIMBUS_CONFIG_LINK = "bot://nimbus.config"
const NIMBUS_TEST_COLLECTION = "46406303"
const NIMBUS_TAG_PREFIX = "nmbs_"

type Nimbus struct {
	Config NimbusConfig `json:"config"`
	Client *RaindropIOClient
}

// Nimbus "singleton"
var nimbus = &Nimbus{}

type NimbusConfig struct {
	Retrospan int64 `json:"retrospan"`
}

func SetupNimbus(baseurl string, bearer string) *Nimbus {
	nimbus.Client = &RaindropIOClient{}
	nimbus.Client.Baseurl = baseurl
	nimbus.Client.Bearer = bearer
	nimbus.Client.Handle = &http.Client{}
	return nimbus
}

func (n *Nimbus) RunExample() {
	// Define new Raindrop
	filter := RaindropUpdateType{
		Ids:       []int{826815716, 823844493},
		Important: true,
	}

	n.Config = NimbusConfig{}

	// Create test raindrop
	opRes := n.Client.UpdateManyRaindrops(46406303, filter)
	opRes.ExecuteOnResponse(func(jsonResponse string) {
		fmt.Print(jsonResponse)
	})
}
