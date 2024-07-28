package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type NimbusClient struct {
	baseurl string
	bearer  string
	handle  *http.Client
}

func (n *NimbusClient) GetCollections() (res *http.Response, e error) {
	route := "collections/childrens"

	req, err := http.NewRequest("GET", n.baseurl+route, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", n.bearer)

	return n.handle.Do(req)
}

func (n *NimbusClient) GetCollectionById(id string) (res *http.Response, e error) {
	route := "collection/"

	req, err := http.NewRequest("GET", n.baseurl+route+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", n.bearer)

	return n.handle.Do(req)
}

func (n *NimbusClient) GetRaindropById(id string) (res *http.Response, e error) {
	route := "raindrop/"

	req, err := http.NewRequest("GET", n.baseurl+route+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", n.bearer)

	return n.handle.Do(req)
}

func (n *NimbusClient) CreateRaindrop(v url.Values) (res *http.Response, e error) {
	route := "raindrop/"

	req, err := http.NewRequest("POST", n.baseurl+route, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", n.bearer)

	// example, should be replaced by parameters
	vals := url.Values{}
	vals.Set("link", "")
	//req.PostForm.Add()
	return nil, nil
}

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

	c := &NimbusClient{bearer: bearer, handle: &http.Client{}, baseurl: url}

	resp, err := c.GetRaindropById("823418374")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Print(string([]byte(body)))

	// TODO: See if there's a cleaner way to handle this.
	response := &InvokeResponse{
		ReturnValue: "",
		Outputs: map[string]interface{}{
			"output": "",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/trigger", run)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
