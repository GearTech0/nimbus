// Package raindropio contains api wrappers for the Raindrop IO API
package raindropio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------
// Raindrop types
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

// ------------------------------------------------------------------------
// http extention
type OperationResponseType struct {
	response *http.Response
	err      error
}

// ------------------------------------------------------------------------
// Operations [TODO: Needs implementation]
type Operation interface {
	GetCollections() OperationResponseType
	GetCollectionById(id string) OperationResponseType
	GetRaindropById(id string) OperationResponseType
	CreateRaindrop(r RaindropType) OperationResponseType
}

func (n *RaindropIOClient) GetCollections() OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := "collections/childrens"

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route, nil)
	if opRes.err != nil {
		return opRes
	}

	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

func (n *RaindropIOClient) GetCollectionById(id string) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := "collection/"

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+id, nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

func (n *RaindropIOClient) GetRaindropById(id string) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := "raindrop/"

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+id, nil)
	if opRes.err != nil {
		return opRes
	}

	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

func (n *RaindropIOClient) CreateRaindrop(r RaindropType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := "raindrop/"

	parsedVals, _ := json.Marshal(r)
	fmt.Println("urlVals: ", string(parsedVals))

	var req *http.Request
	req, opRes.err = http.NewRequest("POST", n.Baseurl+route, strings.NewReader(string(parsedVals)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

func (opRes *OperationResponseType) ExecuteOnResponse(callback func(jsonResponse string)) {
	// if errored, panic
	if opRes.err != nil {
		panic(opRes.err)
	}

	defer opRes.response.Body.Close()

	body, err := io.ReadAll(opRes.response.Body)
	if err != nil {
		panic(err)
	}

	callback(string(body))
}
