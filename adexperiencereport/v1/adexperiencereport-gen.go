// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package adexperiencereport provides access to the Ad Experience Report API.
//
// For product documentation, see: https://developers.google.com/ad-experience-report/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/adexperiencereport/v1"
//   ...
//   ctx := context.Background()
//   adexperiencereportService, err := adexperiencereport.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   adexperiencereportService, err := adexperiencereport.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   adexperiencereportService, err := adexperiencereport.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package adexperiencereport // import "google.golang.org/api/adexperiencereport/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "adexperiencereport:v1"
const apiName = "adexperiencereport"
const apiVersion = "v1"
const basePath = "https://adexperiencereport.googleapis.com/"
const mtlsBasePath = "https://adexperiencereport.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Sites = NewSitesService(s)
	s.ViolatingSites = NewViolatingSitesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Sites *SitesService

	ViolatingSites *ViolatingSitesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	return rs
}

type SitesService struct {
	s *Service
}

func NewViolatingSitesService(s *Service) *ViolatingSitesService {
	rs := &ViolatingSitesService{s: s}
	return rs
}

type ViolatingSitesService struct {
	s *Service
}

// PlatformSummary: A site's Ad Experience Report summary on a single
// platform.
type PlatformSummary struct {
	// BetterAdsStatus: The site's Ad Experience Report status on this
	// platform.
	//
	// Possible values:
	//   "UNKNOWN" - Not reviewed.
	//   "PASSING" - Passing.
	//   "WARNING" - Warning. No longer a possible status.
	//   "FAILING" - Failing.
	BetterAdsStatus string `json:"betterAdsStatus,omitempty"`

	// EnforcementTime: The time at
	// which
	// [enforcement](https://support.google.com/webtools/answer/7308033
	// ) against
	// the site began or will begin on this platform.
	//
	// Not set when the
	// filter_status
	// is OFF.
	EnforcementTime string `json:"enforcementTime,omitempty"`

	// FilterStatus: The site's
	// [enforcement
	// status](https://support.google.com/webtools/answer/730803
	// 3) on this
	// platform.
	//
	// Possible values:
	//   "UNKNOWN" - N/A.
	//   "ON" - Ad filtering is on.
	//   "OFF" - Ad filtering is off.
	//   "PAUSED" - Ad filtering is paused.
	//   "PENDING" - Ad filtering is pending.
	FilterStatus string `json:"filterStatus,omitempty"`

	// LastChangeTime: The time at which the site's status last changed on
	// this platform.
	LastChangeTime string `json:"lastChangeTime,omitempty"`

	// Region: The site's regions on this platform.
	//
	// No longer populated, because there is no longer any semantic
	// difference
	// between sites in different regions.
	//
	// Possible values:
	//   "REGION_UNKNOWN" - Ad standard not yet defined for your region.
	//   "REGION_A" - Region A.
	//   "REGION_B" - Region B.
	//   "REGION_C" - Region C.
	Region []string `json:"region,omitempty"`

	// ReportUrl: A link to the full Ad Experience Report for the site on
	// this platform..
	//
	// Not set in
	// ViolatingSitesResponse.
	//
	// Note that you must complete the [Search Console
	// verification
	// process](https://support.google.com/webmasters/answer/900
	// 8080) for the site
	// before you can access the full report.
	ReportUrl string `json:"reportUrl,omitempty"`

	// UnderReview: Whether the site is currently under review on this
	// platform.
	UnderReview bool `json:"underReview,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BetterAdsStatus") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BetterAdsStatus") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PlatformSummary) MarshalJSON() ([]byte, error) {
	type NoMethod PlatformSummary
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SiteSummaryResponse: Response message for GetSiteSummary.
type SiteSummaryResponse struct {
	// DesktopSummary: The site's Ad Experience Report summary on desktop.
	DesktopSummary *PlatformSummary `json:"desktopSummary,omitempty"`

	// MobileSummary: The site's Ad Experience Report summary on mobile.
	MobileSummary *PlatformSummary `json:"mobileSummary,omitempty"`

	// ReviewedSite: The name of the reviewed site, e.g. `google.com`.
	ReviewedSite string `json:"reviewedSite,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DesktopSummary") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DesktopSummary") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SiteSummaryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SiteSummaryResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ViolatingSitesResponse: Response message for ListViolatingSites.
type ViolatingSitesResponse struct {
	// ViolatingSites: The list of violating sites.
	ViolatingSites []*SiteSummaryResponse `json:"violatingSites,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ViolatingSites") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ViolatingSites") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ViolatingSitesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ViolatingSitesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "adexperiencereport.sites.get":

type SitesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a site's Ad Experience Report summary.
func (r *SitesService) Get(name string) *SitesGetCall {
	c := &SitesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesGetCall) Fields(s ...googleapi.Field) *SitesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SitesGetCall) IfNoneMatch(entityTag string) *SitesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SitesGetCall) Context(ctx context.Context) *SitesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SitesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200711")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexperiencereport.sites.get" call.
// Exactly one of *SiteSummaryResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SiteSummaryResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SitesGetCall) Do(opts ...googleapi.CallOption) (*SiteSummaryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SiteSummaryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a site's Ad Experience Report summary.",
	//   "flatPath": "v1/sites/{sitesId}",
	//   "httpMethod": "GET",
	//   "id": "adexperiencereport.sites.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the site whose summary to get, e.g.\n`sites/http%3A%2F%2Fwww.google.com%2F`.\n\nFormat: `sites/{site}`",
	//       "location": "path",
	//       "pattern": "^sites/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "SiteSummaryResponse"
	//   }
	// }

}

// method id "adexperiencereport.violatingSites.list":

type ViolatingSitesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists sites that are failing in the Ad Experience Report on at
// least one
// platform.
func (r *ViolatingSitesService) List() *ViolatingSitesListCall {
	c := &ViolatingSitesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ViolatingSitesListCall) Fields(s ...googleapi.Field) *ViolatingSitesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ViolatingSitesListCall) IfNoneMatch(entityTag string) *ViolatingSitesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ViolatingSitesListCall) Context(ctx context.Context) *ViolatingSitesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ViolatingSitesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ViolatingSitesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200711")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/violatingSites")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "adexperiencereport.violatingSites.list" call.
// Exactly one of *ViolatingSitesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ViolatingSitesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ViolatingSitesListCall) Do(opts ...googleapi.CallOption) (*ViolatingSitesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ViolatingSitesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists sites that are failing in the Ad Experience Report on at least one\nplatform.",
	//   "flatPath": "v1/violatingSites",
	//   "httpMethod": "GET",
	//   "id": "adexperiencereport.violatingSites.list",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/violatingSites",
	//   "response": {
	//     "$ref": "ViolatingSitesResponse"
	//   }
	// }

}
