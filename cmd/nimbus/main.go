package main

import (
	"encoding/json"
	"fmt"
	"io"
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

type KeyChain struct {
	Bearer string `json:"bearer"`
}

func main() {
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
}
