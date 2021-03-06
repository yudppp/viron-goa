// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "viron": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/yudppp/viron-goa/design
// --out=$(GOPATH)/src/github.com/yudppp/viron-goa
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// blog post (default view)
//
// Identifier: application/vnd.post+json; view=default
type Post struct {
	// contents
	Contents string `form:"contents" json:"contents" xml:"contents"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// published_at
	PublishedAt *time.Time `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
	// status
	Status string `form:"status" json:"status" xml:"status"`
	// name
	Title string `form:"title" json:"title" xml:"title"`
	// url name
	URLName string `form:"url_name" json:"url_name" xml:"url_name"`
}

// Validate validates the Post media type instance.
func (mt *Post) Validate() (err error) {

	if mt.URLName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url_name"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Contents == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "contents"))
	}
	if mt.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}
	if !(mt.Status == "draft" || mt.Status == "published") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, mt.Status, []interface{}{"draft", "published"}))
	}
	return
}

// PostCollection is the media type for an array of Post (default view)
//
// Identifier: application/vnd.post+json; type=collection; view=default
type PostCollection []*Post

// Validate validates the PostCollection media type instance.
func (mt PostCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// viron query (default view)
//
// Identifier: application/vnd.query+json; view=default
type Query struct {
	// key
	Key string `form:"key" json:"key" xml:"key"`
	// type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the Query media type instance.
func (mt *Query) Validate() (err error) {
	if mt.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "key"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// viron api (default view)
//
// Identifier: application/vnd.vironapi+json; view=default
type Vironapi struct {
	// name
	Method string `form:"method" json:"method" xml:"method"`
	// path
	Path string `form:"path" json:"path" xml:"path"`
}

// Validate validates the Vironapi media type instance.
func (mt *Vironapi) Validate() (err error) {
	if mt.Method == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "method"))
	}
	if mt.Path == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "path"))
	}
	return
}

// viron authtype media (default view)
//
// Identifier: application/vnd.vironauthtype+json; view=default
type Vironauthtype struct {
	// method
	Method string `form:"method" json:"method" xml:"method"`
	// auth provider
	Provider string `form:"provider" json:"provider" xml:"provider"`
	// auth type
	Type string `form:"type" json:"type" xml:"type"`
	// url
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the Vironauthtype media type instance.
func (mt *Vironauthtype) Validate() (err error) {
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.Provider == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "provider"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Method == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "method"))
	}
	return
}

// VironauthtypeCollection is the media type for an array of Vironauthtype (default view)
//
// Identifier: application/vnd.vironauthtype+json; type=collection; view=default
type VironauthtypeCollection []*Vironauthtype

// Validate validates the VironauthtypeCollection media type instance.
func (mt VironauthtypeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// viron component (default view)
//
// Identifier: application/vnd.vironcomponent+json; view=default
type Vironcomponent struct {
	// api
	API *Vironapi `form:"api" json:"api" xml:"api"`
	// name
	Name string `form:"name" json:"name" xml:"name"`
	// pagination
	Pagination *bool `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
	// primary key
	Primary *string `form:"primary,omitempty" json:"primary,omitempty" xml:"primary,omitempty"`
	// query
	Query []*Query `form:"query,omitempty" json:"query,omitempty" xml:"query,omitempty"`
	// style
	Style string `form:"style" json:"style" xml:"style"`
	// table label
	TableLabels []string `form:"table_labels,omitempty" json:"table_labels,omitempty" xml:"table_labels,omitempty"`
}

// Validate validates the Vironcomponent media type instance.
func (mt *Vironcomponent) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Style == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "style"))
	}
	if mt.API == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "api"))
	}
	if mt.API != nil {
		if err2 := mt.API.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range mt.Query {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// viron page (default view)
//
// Identifier: application/vnd.vironpage+json; view=default
type Vironpage struct {
	// pages
	Components []*Vironcomponent `form:"components" json:"components" xml:"components"`
	// group
	Group string `form:"group" json:"group" xml:"group"`
	// id
	ID string `form:"id" json:"id" xml:"id"`
	// name
	Name string `form:"name" json:"name" xml:"name"`
	// section
	Section string `form:"section" json:"section" xml:"section"`
}

// Validate validates the Vironpage media type instance.
func (mt *Vironpage) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Section == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "section"))
	}
	if mt.Group == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "group"))
	}
	if mt.Components == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "components"))
	}
	for _, e := range mt.Components {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// viron setting (default view)
//
// Identifier: application/vnd.vironsetting+json; view=default
type Vironsetting struct {
	// color
	Color string `form:"color" json:"color" xml:"color"`
	// name
	Name string `form:"name" json:"name" xml:"name"`
	// pages
	Pages []*Vironpage `form:"pages" json:"pages" xml:"pages"`
	// tags
	Tags []string `form:"tags" json:"tags" xml:"tags"`
	// theme
	Theme string `form:"theme" json:"theme" xml:"theme"`
	// thumbnail
	Thumbnail string `form:"thumbnail" json:"thumbnail" xml:"thumbnail"`
}

// Validate validates the Vironsetting media type instance.
func (mt *Vironsetting) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Color == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "color"))
	}
	if mt.Theme == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "theme"))
	}
	if mt.Pages == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "pages"))
	}
	if mt.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tags"))
	}
	if mt.Thumbnail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thumbnail"))
	}
	for _, e := range mt.Pages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
