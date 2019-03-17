// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa-sample": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hiroykam/goa-sample/design
// --out=$(GOPATH)/src/github.com/hiroykam/goa-sample
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// LoginAuthContext provides the auth login action context.
type LoginAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *LoginAuthPayload
}

// NewLoginAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller login action.
func NewLoginAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// loginAuthPayload is the auth login action payload.
type loginAuthPayload struct {
	// name of sample
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// detail of sample
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *loginAuthPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates LoginAuthPayload from loginAuthPayload
func (payload *loginAuthPayload) Publicize() *LoginAuthPayload {
	var pub LoginAuthPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// LoginAuthPayload is the auth login action payload.
type LoginAuthPayload struct {
	// name of sample
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// detail of sample
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *LoginAuthPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginAuthContext) OK(r *Auth) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.auth+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginAuthContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LoginAuthContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ReauthenticateAuthContext provides the auth reauthenticate action context.
type ReauthenticateAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ReauthenticateAuthPayload
}

// NewReauthenticateAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller reauthenticate action.
func NewReauthenticateAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*ReauthenticateAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ReauthenticateAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// reauthenticateAuthPayload is the auth reauthenticate action payload.
type reauthenticateAuthPayload struct {
	// refresh token
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" yaml:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *reauthenticateAuthPayload) Validate() (err error) {
	if payload.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "refresh_token"))
	}
	return
}

// Publicize creates ReauthenticateAuthPayload from reauthenticateAuthPayload
func (payload *reauthenticateAuthPayload) Publicize() *ReauthenticateAuthPayload {
	var pub ReauthenticateAuthPayload
	if payload.RefreshToken != nil {
		pub.RefreshToken = *payload.RefreshToken
	}
	return &pub
}

// ReauthenticateAuthPayload is the auth reauthenticate action payload.
type ReauthenticateAuthPayload struct {
	// refresh token
	RefreshToken string `form:"refresh_token" json:"refresh_token" yaml:"refresh_token" xml:"refresh_token"`
}

// Validate runs the validation rules defined in the design.
func (payload *ReauthenticateAuthPayload) Validate() (err error) {
	if payload.RefreshToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "refresh_token"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ReauthenticateAuthContext) OK(r *Auth) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.auth+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ReauthenticateAuthContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ReauthenticateAuthContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// AddSamplesContext provides the samples add action context.
type AddSamplesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AddSamplesPayload
}

// NewAddSamplesContext parses the incoming request URL and body, performs validations and creates the
// context used by the samples controller add action.
func NewAddSamplesContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddSamplesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddSamplesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// addSamplesPayload is the samples add action payload.
type addSamplesPayload struct {
	// detail of sample
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" yaml:"detail,omitempty" xml:"detail,omitempty"`
	// name of sample
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *addSamplesPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}
	return
}

// Publicize creates AddSamplesPayload from addSamplesPayload
func (payload *addSamplesPayload) Publicize() *AddSamplesPayload {
	var pub AddSamplesPayload
	if payload.Detail != nil {
		pub.Detail = *payload.Detail
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// AddSamplesPayload is the samples add action payload.
type AddSamplesPayload struct {
	// detail of sample
	Detail string `form:"detail" json:"detail" yaml:"detail" xml:"detail"`
	// name of sample
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *AddSamplesPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *AddSamplesContext) OK(r *Sample) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.sample+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AddSamplesContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AddSamplesContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *AddSamplesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// DeleteSamplesContext provides the samples delete action context.
type DeleteSamplesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeleteSamplesContext parses the incoming request URL and body, performs validations and creates the
// context used by the samples controller delete action.
func NewDeleteSamplesContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteSamplesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteSamplesContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteSamplesContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteSamplesContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *DeleteSamplesContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSamplesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListSamplesContext provides the samples list action context.
type ListSamplesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListSamplesContext parses the incoming request URL and body, performs validations and creates the
// context used by the samples controller list action.
func NewListSamplesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListSamplesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListSamplesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSamplesContext) OK(r SamplesCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.samples+json; type=collection")
	}
	if r == nil {
		r = SamplesCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListSamplesContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListSamplesContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListSamplesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowSamplesContext provides the samples show action context.
type ShowSamplesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowSamplesContext parses the incoming request URL and body, performs validations and creates the
// context used by the samples controller show action.
func NewShowSamplesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowSamplesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowSamplesContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSamplesContext) OK(r *Sample) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.sample+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowSamplesContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ShowSamplesContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSamplesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateSamplesContext provides the samples update action context.
type UpdateSamplesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *UpdateSamplesPayload
}

// NewUpdateSamplesContext parses the incoming request URL and body, performs validations and creates the
// context used by the samples controller update action.
func NewUpdateSamplesContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateSamplesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateSamplesContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// updateSamplesPayload is the samples update action payload.
type updateSamplesPayload struct {
	// detail of sample
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" yaml:"detail,omitempty" xml:"detail,omitempty"`
	// name of sample
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateSamplesPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}
	return
}

// Publicize creates UpdateSamplesPayload from updateSamplesPayload
func (payload *updateSamplesPayload) Publicize() *UpdateSamplesPayload {
	var pub UpdateSamplesPayload
	if payload.Detail != nil {
		pub.Detail = *payload.Detail
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateSamplesPayload is the samples update action payload.
type UpdateSamplesPayload struct {
	// detail of sample
	Detail string `form:"detail" json:"detail" yaml:"detail" xml:"detail"`
	// name of sample
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateSamplesPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateSamplesContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateSamplesContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateSamplesContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSamplesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
