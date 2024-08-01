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
const NIMBUS_TAG_PREFIC = "nmbs_"

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
	changeNimbusConfig := RaindropType{
		Link: "https://www.youtube.com/watch?v=Z6grOAUEIrQ",
	}
	n.Config = NimbusConfig{}

	// Create test raindrop
	opRes := n.Client.UpdateRaindrop(823844493, changeNimbusConfig)
	opRes.ExecuteOnResponse(func(jsonResponse string) {
		fmt.Print(jsonResponse)
	})
}
