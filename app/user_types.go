// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "viron": Application User Types
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
	"unicode/utf8"
)

// postPayload user type.
type postPayload struct {
	// contents
	Contents *string `form:"contents,omitempty" json:"contents,omitempty" xml:"contents,omitempty"`
	// published_at
	PublishedAt *time.Time `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
	// status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// url name
	URLName *string `form:"url_name,omitempty" json:"url_name,omitempty" xml:"url_name,omitempty"`
}

// Validate validates the postPayload type instance.
func (ut *postPayload) Validate() (err error) {
	if ut.URLName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url_name"))
	}
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.Contents == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "contents"))
	}
	if ut.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "status"))
	}
	if ut.Contents != nil {
		if utf8.RuneCountInString(*ut.Contents) > 120 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.contents`, *ut.Contents, utf8.RuneCountInString(*ut.Contents), 120, false))
		}
	}
	if ut.Status != nil {
		if !(*ut.Status == "draft" || *ut.Status == "published") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.status`, *ut.Status, []interface{}{"draft", "published"}))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 120 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 120, false))
		}
	}
	return
}

// Publicize creates PostPayload from postPayload
func (ut *postPayload) Publicize() *PostPayload {
	var pub PostPayload
	if ut.Contents != nil {
		pub.Contents = *ut.Contents
	}
	if ut.PublishedAt != nil {
		pub.PublishedAt = ut.PublishedAt
	}
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	if ut.URLName != nil {
		pub.URLName = *ut.URLName
	}
	return &pub
}

// PostPayload user type.
type PostPayload struct {
	// contents
	Contents string `form:"contents" json:"contents" xml:"contents"`
	// published_at
	PublishedAt *time.Time `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
	// status
	Status string `form:"status" json:"status" xml:"status"`
	// title
	Title string `form:"title" json:"title" xml:"title"`
	// url name
	URLName string `form:"url_name" json:"url_name" xml:"url_name"`
}

// Validate validates the PostPayload type instance.
func (ut *PostPayload) Validate() (err error) {
	if ut.URLName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url_name"))
	}
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.Contents == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "contents"))
	}
	if ut.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "status"))
	}
	if utf8.RuneCountInString(ut.Contents) > 120 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.contents`, ut.Contents, utf8.RuneCountInString(ut.Contents), 120, false))
	}
	if !(ut.Status == "draft" || ut.Status == "published") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.status`, ut.Status, []interface{}{"draft", "published"}))
	}
	if utf8.RuneCountInString(ut.Title) > 120 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.title`, ut.Title, utf8.RuneCountInString(ut.Title), 120, false))
	}
	return
}

// signinPayload user type.
type signinPayload struct {
	// ID or Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the signinPayload type instance.
func (ut *signinPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 256, false))
		}
	}
	return
}

// Publicize creates SigninPayload from signinPayload
func (ut *signinPayload) Publicize() *SigninPayload {
	var pub SigninPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// SigninPayload user type.
type SigninPayload struct {
	// ID or Email
	Email string `form:"email" json:"email" xml:"email"`
	// Password
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the SigninPayload type instance.
func (ut *SigninPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if utf8.RuneCountInString(ut.Password) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 256, false))
	}
	return
}
