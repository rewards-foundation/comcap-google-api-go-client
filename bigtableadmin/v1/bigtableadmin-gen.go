// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package bigtableadmin provides access to the Cloud Bigtable Admin API.
//
// For product documentation, see: https://cloud.google.com/bigtable/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/bigtableadmin/v1"
//   ...
//   ctx := context.Background()
//   bigtableadminService, err := bigtableadmin.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   bigtableadminService, err := bigtableadmin.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   bigtableadminService, err := bigtableadmin.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package bigtableadmin // import "google.golang.org/api/bigtableadmin/v1"

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

const apiId = "bigtableadmin:v1"
const apiName = "bigtableadmin"
const apiVersion = "v1"
const basePath = "https://bigtableadmin.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
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
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

// Cluster: A resizable group of nodes in a particular cloud location,
// capable
// of serving all Tables in the parent
// Instance.
type Cluster struct {
	// DefaultStorageType: Immutable. The type of storage used by this
	// cluster to serve its
	// parent instance's tables, unless explicitly overridden.
	//
	// Possible values:
	//   "STORAGE_TYPE_UNSPECIFIED" - The user did not specify a storage
	// type.
	//   "SSD" - Flash (SSD) storage should be used.
	//   "HDD" - Magnetic drive (HDD) storage should be used.
	DefaultStorageType string `json:"defaultStorageType,omitempty"`

	// Location: Immutable. The location where this cluster's nodes and
	// storage reside. For best
	// performance, clients should be located as close as possible to
	// this
	// cluster. Currently only zones are supported, so values should be of
	// the
	// form `projects/{project}/locations/{zone}`.
	Location string `json:"location,omitempty"`

	// Name: The unique name of the cluster. Values are of the
	// form
	// `projects/{project}/instances/{instance}/clusters/a-z*`.
	Name string `json:"name,omitempty"`

	// ServeNodes: Required. The number of nodes allocated to this cluster.
	// More nodes enable higher
	// throughput and more consistent performance.
	ServeNodes int64 `json:"serveNodes,omitempty"`

	// State: Output only. The current state of the cluster.
	//
	// Possible values:
	//   "STATE_NOT_KNOWN" - The state of the cluster could not be
	// determined.
	//   "READY" - The cluster has been successfully created and is ready to
	// serve requests.
	//   "CREATING" - The cluster is currently being created, and may be
	// destroyed
	// if the creation process encounters an error.
	// A cluster may not be able to serve requests while being created.
	//   "RESIZING" - The cluster is currently being resized, and may revert
	// to its previous
	// node count if the process encounters an error.
	// A cluster is still capable of serving requests while being
	// resized,
	// but may exhibit performance as if its number of allocated nodes
	// is
	// between the starting and requested states.
	//   "DISABLED" - The cluster has no backing nodes. The data (tables)
	// still
	// exist, but no operations can be performed on the cluster.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultStorageType")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultStorageType") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Cluster) MarshalJSON() ([]byte, error) {
	type NoMethod Cluster
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateClusterMetadata: The metadata for the Operation returned by
// CreateCluster.
type CreateClusterMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// CreateCluster operation.
	OriginalRequest *CreateClusterRequest `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// Tables: Keys: the full `name` of each table that existed in the
	// instance when
	// CreateCluster was first called,
	// i.e.
	// `projects/<project>/instances/<instance>/tables/<table>`. Any table
	// added
	// to the instance by a later API call will be created in the new
	// cluster by
	// that API call, not this one.
	//
	// Values: information on how much of a table's data has been copied to
	// the
	// newly-created cluster so far.
	Tables map[string]TableProgress `json:"tables,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FinishTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateClusterMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod CreateClusterMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateClusterRequest: Request message for
// BigtableInstanceAdmin.CreateCluster.
type CreateClusterRequest struct {
	// Cluster: Required. The cluster to be created.
	// Fields marked `OutputOnly` must be left blank.
	Cluster *Cluster `json:"cluster,omitempty"`

	// ClusterId: Required. The ID to be used when referring to the new
	// cluster within its instance,
	// e.g., just `mycluster` rather
	// than
	// `projects/myproject/instances/myinstance/clusters/mycluster`.
	ClusterId string `json:"clusterId,omitempty"`

	// Parent: Required. The unique name of the instance in which to create
	// the new cluster.
	// Values are of the form
	// `projects/{project}/instances/{instance}`.
	Parent string `json:"parent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cluster") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cluster") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateClusterRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateClusterRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateInstanceMetadata: The metadata for the Operation returned by
// CreateInstance.
type CreateInstanceMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// CreateInstance operation.
	OriginalRequest *CreateInstanceRequest `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FinishTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateInstanceMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod CreateInstanceMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateInstanceRequest: Request message for
// BigtableInstanceAdmin.CreateInstance.
type CreateInstanceRequest struct {
	// Clusters: Required. The clusters to be created within the instance,
	// mapped by desired
	// cluster ID, e.g., just `mycluster` rather
	// than
	// `projects/myproject/instances/myinstance/clusters/mycluster`.
	// Fie
	// lds marked `OutputOnly` must be left blank.
	// Currently, at most four clusters can be specified.
	Clusters map[string]Cluster `json:"clusters,omitempty"`

	// Instance: Required. The instance to create.
	// Fields marked `OutputOnly` must be left blank.
	Instance *Instance `json:"instance,omitempty"`

	// InstanceId: Required. The ID to be used when referring to the new
	// instance within its project,
	// e.g., just `myinstance` rather
	// than
	// `projects/myproject/instances/myinstance`.
	InstanceId string `json:"instanceId,omitempty"`

	// Parent: Required. The unique name of the project in which to create
	// the new instance.
	// Values are of the form `projects/{project}`.
	Parent string `json:"parent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Clusters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Clusters") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateInstanceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Instance: A collection of Bigtable Tables and
// the resources that serve them.
// All tables in an instance are served from all
// Clusters in the instance.
type Instance struct {
	// DisplayName: Required. The descriptive name for this instance as it
	// appears in UIs.
	// Can be changed at any time, but should be kept globally unique
	// to avoid confusion.
	DisplayName string `json:"displayName,omitempty"`

	// Labels: Required. Labels are a flexible and lightweight mechanism for
	// organizing cloud
	// resources into groups that reflect a customer's organizational needs
	// and
	// deployment strategies. They can be used to filter resources and
	// aggregate
	// metrics.
	//
	// * Label keys must be between 1 and 63 characters long and must
	// conform to
	//   the regular expression: `\p{Ll}\p{Lo}{0,62}`.
	// * Label values must be between 0 and 63 characters long and must
	// conform to
	//   the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`.
	// * No more than 64 labels can be associated with a given resource.
	// * Keys and values must both be under 128 bytes.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: The unique name of the instance. Values are of the
	// form
	// `projects/{project}/instances/a-z+[a-z0-9]`.
	Name string `json:"name,omitempty"`

	// State: Output only. The current state of the instance.
	//
	// Possible values:
	//   "STATE_NOT_KNOWN" - The state of the instance could not be
	// determined.
	//   "READY" - The instance has been successfully created and can serve
	// requests
	// to its tables.
	//   "CREATING" - The instance is currently being created, and may be
	// destroyed
	// if the creation process encounters an error.
	State string `json:"state,omitempty"`

	// Type: Required. The type of the instance. Defaults to `PRODUCTION`.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The type of the instance is unspecified. If
	// set when creating an
	// instance, a `PRODUCTION` instance will be created. If set when
	// updating
	// an instance, the type will be left unchanged.
	//   "PRODUCTION" - An instance meant for production use. `serve_nodes`
	// must be set
	// on the cluster.
	//   "DEVELOPMENT" - DEPRECATED: Prefer PRODUCTION for all use cases, as
	// it no longer enforces
	// a higher minimum node count than DEVELOPMENT.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Instance) MarshalJSON() ([]byte, error) {
	type NoMethod Instance
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PartialUpdateInstanceRequest: Request message for
// BigtableInstanceAdmin.PartialUpdateInstance.
type PartialUpdateInstanceRequest struct {
	// Instance: Required. The Instance which will (partially) replace the
	// current value.
	Instance *Instance `json:"instance,omitempty"`

	// UpdateMask: Required. The subset of Instance fields which should be
	// replaced.
	// Must be explicitly set.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Instance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Instance") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PartialUpdateInstanceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod PartialUpdateInstanceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TableProgress: Progress info for copying a table's data to the new
// cluster.
type TableProgress struct {
	// EstimatedCopiedBytes: Estimate of the number of bytes copied so far
	// for this table.
	// This will eventually reach 'estimated_size_bytes' unless the table
	// copy
	// is CANCELLED.
	EstimatedCopiedBytes int64 `json:"estimatedCopiedBytes,omitempty,string"`

	// EstimatedSizeBytes: Estimate of the size of the table to be copied.
	EstimatedSizeBytes int64 `json:"estimatedSizeBytes,omitempty,string"`

	// Possible values:
	//   "STATE_UNSPECIFIED"
	//   "PENDING" - The table has not yet begun copying to the new cluster.
	//   "COPYING" - The table is actively being copied to the new cluster.
	//   "COMPLETED" - The table has been fully copied to the new cluster.
	//   "CANCELLED" - The table was deleted before it finished copying to
	// the new cluster.
	// Note that tables deleted after completion will stay marked
	// as
	// COMPLETED, not CANCELLED.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EstimatedCopiedBytes") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EstimatedCopiedBytes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TableProgress) MarshalJSON() ([]byte, error) {
	type NoMethod TableProgress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateAppProfileMetadata: The metadata for the Operation returned by
// UpdateAppProfile.
type UpdateAppProfileMetadata struct {
}

// UpdateClusterMetadata: The metadata for the Operation returned by
// UpdateCluster.
type UpdateClusterMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// UpdateCluster operation.
	OriginalRequest *Cluster `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FinishTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateClusterMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateClusterMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateInstanceMetadata: The metadata for the Operation returned by
// UpdateInstance.
type UpdateInstanceMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// UpdateInstance operation.
	OriginalRequest *PartialUpdateInstanceRequest `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FinishTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateInstanceMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateInstanceMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}
