// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa-sample": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hiroykam/goa-sample/design
// --out=$(GOPATH)/src/github.com/hiroykam/goa-sample
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// sample detail (default view)
//
// Identifier: application/vnd.sample+json; view=default
type Sample struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// 詳細
	Detail string `form:"detail" json:"detail" yaml:"detail" xml:"detail"`
	// sample id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
	// user id
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the Sample media type instance.
func (mt *Sample) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "detail"))
	}

	return
}

// DecodeSample decodes the Sample instance encoded in resp body.
func (c *Client) DecodeSample(resp *http.Response) (*Sample, error) {
	var decoded Sample
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// SampleCollection is the media type for an array of Sample (default view)
//
// Identifier: application/vnd.sample+json; type=collection; view=default
type SampleCollection []*Sample

// Validate validates the SampleCollection media type instance.
func (mt SampleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeSampleCollection decodes the SampleCollection instance encoded in resp body.
func (c *Client) DecodeSampleCollection(resp *http.Response) (SampleCollection, error) {
	var decoded SampleCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// sample list (default view)
//
// Identifier: application/vnd.samples+json; view=default
type Samples struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the Samples media type instance.
func (mt *Samples) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeSamples decodes the Samples instance encoded in resp body.
func (c *Client) DecodeSamples(resp *http.Response) (*Samples, error) {
	var decoded Samples
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// SamplesCollection is the media type for an array of Samples (default view)
//
// Identifier: application/vnd.samples+json; type=collection; view=default
type SamplesCollection []*Samples

// Validate validates the SamplesCollection media type instance.
func (mt SamplesCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeSamplesCollection decodes the SamplesCollection instance encoded in resp body.
func (c *Client) DecodeSamplesCollection(resp *http.Response) (SamplesCollection, error) {
	var decoded SamplesCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
