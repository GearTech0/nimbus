package raindropio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RaindropIOClient struct {
	Baseurl string
	Bearer  string
	Handle  *http.Client
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

func (n *RaindropIOClient) GetCollections() (res *http.Response, e error) {
	route := "collections/childrens"

	req, err := http.NewRequest("GET", n.Baseurl+route, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", n.Bearer)

	return n.Handle.Do(req)
}

func (n *RaindropIOClient) GetCollectionById(id string) (res *http.Response, e error) {
	route := "collection/"

	req, err := http.NewRequest("GET", n.Baseurl+route+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", n.Bearer)

	return n.Handle.Do(req)
}

func (n *RaindropIOClient) GetRaindropById(id string) (res *http.Response, e error) {
	route := "raindrop/"

	req, err := http.NewRequest("GET", n.Baseurl+route+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", n.Bearer)

	return n.Handle.Do(req)
}

func (n *RaindropIOClient) CreateRaindrop(r RaindropType) (res *http.Response, e error) {
	route := "raindrop/"

	parsedVals, _ := json.Marshal(r)
	fmt.Println("urlVals: ", string(parsedVals))

	req, err := http.NewRequest("POST", n.Baseurl+route, strings.NewReader(string(parsedVals)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	return n.Handle.Do(req)
}
