package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type NimbusClient struct {
	baseurl string
	bearer  string
	handle  *http.Client
}

type CollectionType struct {
	Id string `json:"$id,omitempty"`
}

type HighlightType struct {
	Id      string   `json:"_id"`
	Text    string   `json:"text"`
	Title   string   `json:"title"`
	Color   string   `json:"color"`
	Note    string   `json:"note"`
	Created string   `json:"created"`
	Tags    []string `json:"tags"`
	Link    string   `json:"link"`
}

type ReminderType struct {
	Date string `json:"date"`
}

type RaindropType struct {
	Created    string         `json:"created,omitempty"`
	LastUpdate string         `json:"lastUpdate,omitempty"`
	Order      int64          `json:"order,omitempty,string"`
	Important  bool           `json:"important,omitempty,string"`
	Tags       []string       `json:"tags,omitempty"`
	Media      []string       `json:"media,omitempty"`
	Cover      string         `json:"cover,omitempty"`
	Collection CollectionType `json:"collection,omitempty"`
	Type       string         `json:"type,omitempty"`
	Excerpt    string         `json:"excerpt,omitempty"`
	Title      string         `json:"title,omitempty"`
	Link       string         `json:"link"`
	Highlights []string       `json:"highlights,omitempty"`
	Reminder   ReminderType   `json:"reminder,omitempty"`
}

func (r *RaindropType) ConvertToURLVals() (url.Values, error) {
	ret := url.Values{}

	ret.Set("link", r.Link)
	ret.Set("created", r.Created)
	ret.Set("lastUpdate", r.LastUpdate)
	ret.Set("order", string(r.Order))
	ret.Set("important", strconv.FormatBool(r.Important))
	ret.Set("cover", r.Cover)
	ret.Set("type", r.Type)
	ret.Set("excerpt", r.Excerpt)
	ret.Set("title", r.Title)

	return ret, nil
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

func (n *NimbusClient) CreateRaindrop(r RaindropType) (res *http.Response, e error) {
	route := "raindrop/"

	parsedVals, _ := json.Marshal(r)
	fmt.Println("urlVals: ", string(parsedVals))

	req, err := http.NewRequest("POST", n.baseurl+route, strings.NewReader(string(parsedVals)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", n.bearer)
	req.Header.Set("Content-Type", "application/json")

	return n.handle.Do(req)
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

	resp, err := c.CreateRaindrop(RaindropType{
		Link:       "https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go",
		Title:      "How do I fix this stupid bugs...",
		Collection: CollectionType{"46406303"},
	})
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
