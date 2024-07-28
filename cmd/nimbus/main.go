package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	. "github.com/GearTech0/nimbus/internal/nimbus"
)

type InvokeResponse struct {
	Outputs     map[string]interface{}
	ReturnValue interface{}
	Logs        []string
}
type KeyChain struct {
	Bearer string `json:"bearer"`
}

func run(w http.ResponseWriter, r *http.Request) {
	url := "https://api.raindrop.io/rest/v1/"

	// Bearer token from secret folder
	contents, err := os.ReadFile("secret/keychain.json")
	if err != nil {
		fmt.Println("error: ", err)
	}

	var kc KeyChain
	err = json.Unmarshal(contents, &kc)
	if err != nil {
		fmt.Println("err: ", err)
	}
	bearer := "Bearer " + kc.Bearer

	nimbus := SetupNimbus(url, bearer)
	nimbus.RunExample()

	// TODO: See if there's a cleaner way to handle this.
	// response := &InvokeResponse{
	// 	ReturnValue: "",
	// 	Outputs: map[string]interface{}{
	// 		"output": "",
	// 	},
	// }
	// w.Header().Set("Content-Type", "application/json")
	// _ = json.NewEncoder(w).Encode(response)
}

func main() {
	run(nil, nil)
	// listenAddr := ":8080"
	// if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
	// 	listenAddr = ":" + val
	// }
	// http.HandleFunc("/trigger", run)
	// log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	// log.Fatal(http.ListenAndServe(listenAddr, nil))
}
