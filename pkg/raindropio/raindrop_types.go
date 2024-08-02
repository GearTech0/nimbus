package raindropio

import "net/http"

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
	Id      string   `json:"_id,omitempty"`
	Text    string   `json:"text,omitempty"`
	Title   string   `json:"title,omitempty"`
	Color   string   `json:"color,omitempty"`
	Note    string   `json:"note,omitempty"`
	Created string   `json:"created,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Link    string   `json:"link,omitempty"`
}

type ReminderType struct {
	Date string `json:"date"`
}

// ------------------------------------------------------------------------
// Raindrop types
// ------------------------------------------------------------------------
type RaindropType struct {
	Created    string               `json:"created,omitempty"`
	LastUpdate string               `json:"lastUpdate,omitempty"`
	Order      int64                `json:"order,omitempty,string"`
	Important  bool                 `json:"important,omitempty,string"`
	Tags       []string             `json:"tags,omitempty"`
	Media      []string             `json:"media,omitempty"`
	Cover      string               `json:"cover,omitempty"`
	Collection CollectionParentType `json:"collection,omitempty"`
	Type       string               `json:"type,omitempty"`
	Excerpt    string               `json:"excerpt,omitempty"`
	Title      string               `json:"title,omitempty"`
	Link       string               `json:"link,omitempty"`
	Highlights []string             `json:"highlights,omitempty"`
	Reminder   ReminderType         `json:"reminder,omitempty"`
}

type RaindropUpdateType struct {
	Ids        []int                `json:"ids,omitempty"`
	Important  bool                 `json:"important,omitempty"`
	Tags       []string             `json:"tags,omitempty"`
	Media      []any                `json:"media,omitempty"`
	Cover      string               `json:"cover,omitempty"`
	Collection CollectionParentType `json:"collection,omitempty"`
}

type FilterType struct {
	CollectionId int    `json:"collectionId,omitempty"`
	Search       string `json:"search,omitempty"`
	Sort         string `json:"sort,omitempty"`
	Page         int    `json:"page,omitempty"`
	PerPage      int    `json:"perpage,omitempty"`
	Ids          []int  `json:"ids,omitempty"`
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

type ListBody struct {
	Items []RaindropType `json:"items,omitempty"`
}

// ------------------------------------------------------------------------
// http extention
type OperationResponseType struct {
	response *http.Response
	err      error
}
