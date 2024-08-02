// Package raindropio contains api wrappers for the Raindrop IO API
package raindropio

import (
	"encoding/json"
	"io"
	"net/http"
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
const ROUTE_RAINDROPS string = "raindrops/"
const ROUTE_SUGGEST string = "suggest/"

// ------------------------------------------------------------------------
// Operations
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
		items Object
			_id int
			access {level int, draggable bool}
			collaborators {$id string}
			color string
			cover []string
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
			cover []string
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
func (n *RaindropIOClient) GetCollection(id int) OperationResponseType {
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
		items []object
			_id int
			access {level int, draggable bool}
			collaborators {$id string}
			color string
			cover []string
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
		cover []string

	[OUT] form:
		result bool
		item JSON['CollectionType']
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
		cover []string

	[OUT] form:
		result bool
		item JSON['CollectionType']
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
/*
	[OUT] form:
		result bool
*/
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
/*
	[IN] form:
		ids []string
*/
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
/*
	[OUT] form:
		result bool
		item RaindropType
*/
func (n *RaindropIOClient) GetRaindrop(id int) OperationResponseType {
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
		tags []string
		media []string
		cover string
		collection {id: int}
		type string
		excerpt string
		title string
		link string
		highlights []Highlight(?)
		reminder

	[OUT] form:
		result bool
		item RaindropType
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
/*
	[IN] form:
		created string
		lastUpdate string
		order int
		important bool
		tags []string
		media []string
		cover string
		collection {$id int}
		type string
		excerpt string
		title string
		link string
		highlights []HighlightType
		reminder {}

	[OUT] form:
		result bool
		item RaindropType
*/
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
/*
	[OUT] form:
		result bool
*/
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
/*
	[OUT] form:
		result bool
		item Object
			collections []{$id int}
		tags []string
*/
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
/*
	[OUT] form:
		result bool
		item Object
			collections []{$id int}
		tags []string
*/
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

// Get Raindrops
/*
	[OUT] form:
		result bool
		items []RaindropType
*/
func (n *RaindropIOClient) GetRaindrops(collectionId int, filter FilterType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROPS

	parsedFilter := CreateFilterQuery(&filter)

	var req *http.Request
	req, opRes.err = http.NewRequest("GET", n.Baseurl+route+strconv.Itoa(collectionId)+parsedFilter, nil)
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Create many raindrops
/*
	[IN] form:
		items []RaindropType

	[OUT] form:
		result bool
		items []RaindropType
*/
func (n *RaindropIOClient) CreateManyRaindrops(in ListBody) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROPS

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

// Update many raindrops
/*
	[IN] form:
		ids []string
		important bool
		tags []string
		media []string
		cover string
		collection CollectionParentType
*/
func (n *RaindropIOClient) UpdateManyRaindrops(collectionId int, updates RaindropUpdateType) OperationResponseType {
	opRes := OperationResponseType{response: nil, err: nil}
	route := ROUTE_RAINDROPS

	var parsed []byte
	parsed, opRes.err = json.Marshal(updates)
	if opRes.err != nil {
		return opRes
	}

	var req *http.Request
	req, opRes.err = http.NewRequest("PUT", n.Baseurl+route+strconv.Itoa(collectionId), strings.NewReader(string(parsed)))
	if opRes.err != nil {
		return opRes
	}
	req.Header.Set("Authorization", n.Bearer)
	req.Header.Set("Content-Type", "application/json")

	opRes.response, opRes.err = n.Handle.Do(req)
	return opRes
}

// Build a filter query string
func CreateFilterQuery(filter *FilterType) string {
	queryString := "?"

	if filter != nil {
		var q []string
		if filter.Sort != "" {
			q = append(q, "sort="+filter.Sort)
		}
		if filter.Page >= 0 {
			q = append(q, "page="+strconv.Itoa(filter.Page))
		}
		if filter.PerPage >= 0 {
			q = append(q, "perpage="+strconv.Itoa(filter.PerPage))
		}
		if filter.Search != "" {
			q = append(q, "search="+filter.Search)
		}

		// build query string
		for i, s := range q {
			if i > 0 {
				queryString += "&"
			}
			queryString += s
		}
	}

	return queryString
}

// Add callback for output of an operation.
// JSON form located above wrapper methods.
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
