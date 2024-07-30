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
// Client type
type RaindropIOClient struct {
	Baseurl string
	Bearer  string
	Handle  *http.Client
}

// ------------------------------------------------------------------------
// Collections types
// ------------------------------------------------------------------------

type CollectionType struct {
	Id            string               `json:"_id,omitempty"`
	Access        CollectionAccessType `json:"access,omitempty"`
	Collaborators CollaboratorsType    `json:"collaborators,omitempty"`
	Color         string               `json:"color,omitempty"`
	Count         int64                `json:"count,omitempty"`
	Cover         []string             `json:"cover,omitempty"`
	Created       string               `json:"created,omitempty"`
	Expanded      bool                 `json:"expanded,omitempty"`
	LastUpdate    string               `json:"lastUpdate,omitempty"`
	Parent        CollectionParentType `json:"parent,omitempty"`
	Public        bool                 `json:"public,omitempty"`
	Sort          int64                `json:"sort,omitempty"`
	Title         string               `json:"title,omitempty"`
	User          UserType             `json:"user,omitempty"`
	View          string               `json:"view,omitempty"`
}

type CollectionAccessType struct {
	Level     int64 `json:"level,omitempty"`
	Draggable bool  `json:"draggable,omitempty"`
}

type CollectionParentType struct {
	Id int64 `json:"$id,omitempty"`
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

// ------------------------------------------------------------------------
// Raindrop types
// ------------------------------------------------------------------------
type RaindropType struct {
	Created    string                         `json:"created,omitempty"`
	LastUpdate string                         `json:"lastUpdate,omitempty"`
	Order      int64                          `json:"order,omitempty,string"`
	Important  bool                           `json:"important,omitempty,string"`
	Tags       []string                       `json:"tags,omitempty"`
	Media      []string                       `json:"media,omitempty"`
	Cover      string                         `json:"cover,omitempty"`
	Collection RaindropCollectionPropertyType `json:"collection,omitempty"`
	Type       string                         `json:"type,omitempty"`
	Excerpt    string                         `json:"excerpt,omitempty"`
	Title      string                         `json:"title,omitempty"`
	Link       string                         `json:"link"`
	Highlights []string                       `json:"highlights,omitempty"`
	Reminder   ReminderType                   `json:"reminder,omitempty"`
}

type RaindropCollectionPropertyType struct {
	Id string `json:"$id,omitempty"`
}

// ------------------------------------------------------------------------
// General types
// ------------------------------------------------------------------------

type UserType struct {
	Id int64 `json:"$id,omitempty"`
}

type CollaboratorsType struct {
	Id       string `json:"_id,omitempty"`
	Email    string `json:"email,omitempty"`
	EmailMD5 string `json:"email_MD5,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Role     string `json:"role,omitempty"`
}

// Convert a raindrop into url.Values
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
	GetChildCollections() OperationResponseType
	GetCollectionById(id string) OperationResponseType
	GetRaindropById(id string) OperationResponseType
	CreateRaindrop(r RaindropType) OperationResponseType
}

// -------------------------------------------------------------------------
// Collections methods
// -------------------------------------------------------------------------

// Get child collections
/*
	form:
		result bool
		items Array<object>
			_id int
			access {level int, draggable bool}
			collaborators {$id string}
			color string
			cover Array<string>
			count int
			created string
			expanded bool
			lastUpdate string
			public bool
			sort int
			title string
			user {$id int}
			view string

*/
func (n *RaindropIOClient) GetChildCollections() OperationResponseType {
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

// Get collection
/*
	form:
		result bool
		item Object
			_id int
			access {level int, draggable bool, for int, root bool}
			collaborators {$id string}
			color string
			cover Array<string>
			count int
			created string
			expanded bool
			lastUpdate string
			public bool
			sort int
			title string
			user {$id int}
			view string

*/
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

// Get root collections
/*
	form:
		result bool
		items Array<object>
			_id int
			access {level int, draggable bool}
			collaborators {$id string}
			color string
			cover Array<string>
			count int
			created string
			expanded bool
			lastUpdate string
			public bool
			sort int
			title string
			user {$id int}
			view string

*/
func (n *RaindropIOClient) GetRootCollections() OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := "collections/"

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route, nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// -------------------------------------------------------------------------
// Raindrops methods
// -------------------------------------------------------------------------

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
