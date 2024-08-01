// Package raindropio contains api wrappers for the Raindrop IO API
package raindropio

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------
// Constants
// ------------------------------------------------------------------------
const ROUTE_COLLECTIONS string = "collections/"
const ROUTE_COLLECTION string = "collection/"
const ROUTE_CHILDRENS string = "childrens/"
const ROUTE_MERGE string = "merge/"
const ROUTE_RAINDROP string = "raindrop/"
const ROUTE_SUGGEST string = "suggest/"

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

type IDList struct {
	Ids []int `json:"ids"`
}

type LinkBody struct {
	Link string `json:"link"`
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
	[OUT] form:
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
	route := ROUTE_COLLECTIONS + ROUTE_CHILDRENS

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
	[OUT] form:
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
func (n *RaindropIOClient) GetCollectionById(id int) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_COLLECTION

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+strconv.Itoa(id), nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Get root collections
/*
	[OUT] form:
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
	route := ROUTE_COLLECTIONS

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route, nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Create Collection
/*
	[IN] form:
		view string
		title string
		sort int
		public bool
		parent {$id int}
		cover Array<string>
*/
func (n *RaindropIOClient) CreateCollection(in CollectionType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_COLLECTION

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("POST", n.Baseurl+route, strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)
	req.Header.Add("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Update an existing collection
/*
	[IN] form:
		view string
		title string
		sort int
		public bool
		parent {$id int}
		cover Array<string>
*/
func (n *RaindropIOClient) UpdateCollection(id int, in CollectionType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_COLLECTION

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("PUT", n.Baseurl+route+strconv.Itoa(id), strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)
	req.Header.Add("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Remove collection
func (n *RaindropIOClient) RemoveCollection(id int) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_COLLECTION

	var req *http.Request
	req, opRes.err = http.NewRequest("DELETE", n.Baseurl+route+strconv.Itoa(id), nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Remove multiple collections
func (n *RaindropIOClient) RemoveMultipleCollections(in IDList) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_COLLECTIONS

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("DELETE", n.Baseurl+route, strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Add("Authorization", n.Bearer)
	req.Header.Add("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// -------------------------------------------------------------------------
// Raindrops methods
// -------------------------------------------------------------------------

// Get raindrop
func (n *RaindropIOClient) GetRaindropById(id int) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+strconv.Itoa(id), nil)
	if opRes.err != nil {
		return opRes
	}

	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Create raindrop
/*
	[IN] form:
		created string
		lastUpdate string
		order int
		important boolean
		tags Array<string>
		media Array<string>
		cover string
		collection {id: int}
		type string
		excerpt string
		title string
		link string
		highlights Array<Highlight(?)>
		reminder
*/
func (n *RaindropIOClient) CreateRaindrop(in RaindropType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("POST", n.Baseurl+route, strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Update raindrop
func (n *RaindropIOClient) UpdateRaindrop(id int, in RaindropType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("PUT", n.Baseurl+route+strconv.Itoa(id), strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Remove raindrop
func (n *RaindropIOClient) RemoveRaindrop(id int) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP

	var req *http.Request
	req, opRes.err = http.NewRequest("DELETE", n.Baseurl+route+strconv.Itoa(id), nil)
	if opRes.err != nil {
		return opRes
	}

	req.Header.Add("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Suggest collection and tags for new bookmark
func (n *RaindropIOClient) NewBookmarkSuggestions(in LinkBody) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP + ROUTE_SUGGEST

	var parsed []byte
	parsed, opRes.err = json.Marshal(in)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("POST", n.Baseurl+route, strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Suggest collection and tags for new bookmark
func (n *RaindropIOClient) ExistingBookmarkSuggestions(id int) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROP

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+strconv.Itoa(id)+ROUTE_SUGGEST, nil)
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
