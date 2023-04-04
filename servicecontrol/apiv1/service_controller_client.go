// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package servicecontrol

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	servicecontrolpb "cloud.google.com/go/servicecontrol/apiv1/servicecontrolpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newServiceControllerClientHook clientHook

// ServiceControllerCallOptions contains the retry settings for each method of ServiceControllerClient.
type ServiceControllerCallOptions struct {
	Check  []gax.CallOption
	Report []gax.CallOption
}

func defaultServiceControllerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("servicecontrol.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("servicecontrol.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://servicecontrol.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServiceControllerCallOptions() *ServiceControllerCallOptions {
	return &ServiceControllerCallOptions{
		Check: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		Report: []gax.CallOption{},
	}
}

func defaultServiceControllerRESTCallOptions() *ServiceControllerCallOptions {
	return &ServiceControllerCallOptions{
		Check: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		Report: []gax.CallOption{},
	}
}

// internalServiceControllerClient is an interface that defines the methods available from Service Control API.
type internalServiceControllerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Check(context.Context, *servicecontrolpb.CheckRequest, ...gax.CallOption) (*servicecontrolpb.CheckResponse, error)
	Report(context.Context, *servicecontrolpb.ReportRequest, ...gax.CallOption) (*servicecontrolpb.ReportResponse, error)
}

// ServiceControllerClient is a client for interacting with Service Control API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Google Service Control API (at /service-control/overview)
//
// Lets clients check and report operations against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
type ServiceControllerClient struct {
	// The internal transport-dependent client.
	internalClient internalServiceControllerClient

	// The call options for this service.
	CallOptions *ServiceControllerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServiceControllerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServiceControllerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ServiceControllerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Check checks whether an operation on a service should be allowed to proceed
// based on the configuration of the service and related policies. It must be
// called before the operation is executed.
//
// If feasible, the client should cache the check results and reuse them for
// 60 seconds. In case of any server errors, the client should rely on the
// cached results for much longer time to avoid outage.
// WARNING: There is general 60s delay for the configuration and policy
// propagation, therefore callers MUST NOT depend on the Check method having
// the latest policy information.
//
// NOTE: the CheckRequest has
// the size limit (wire-format byte size) of 1MB.
//
// This method requires the servicemanagement.services.check permission
// on the specified service. For more information, see
// Cloud IAM (at https://cloud.google.com/iam).
func (c *ServiceControllerClient) Check(ctx context.Context, req *servicecontrolpb.CheckRequest, opts ...gax.CallOption) (*servicecontrolpb.CheckResponse, error) {
	return c.internalClient.Check(ctx, req, opts...)
}

// Report reports operation results to Google Service Control, such as logs and
// metrics. It should be called after an operation is completed.
//
// If feasible, the client should aggregate reporting data for up to 5
// seconds to reduce API traffic. Limiting aggregation to 5 seconds is to
// reduce data loss during client crashes. Clients should carefully choose
// the aggregation time window to avoid data loss risk more than 0.01%
// for business and compliance reasons.
//
// NOTE: the ReportRequest has
// the size limit (wire-format byte size) of 1MB.
//
// This method requires the servicemanagement.services.report permission
// on the specified service. For more information, see
// Google Cloud IAM (at https://cloud.google.com/iam).
func (c *ServiceControllerClient) Report(ctx context.Context, req *servicecontrolpb.ReportRequest, opts ...gax.CallOption) (*servicecontrolpb.ReportResponse, error) {
	return c.internalClient.Report(ctx, req, opts...)
}

// serviceControllerGRPCClient is a client for interacting with Service Control API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type serviceControllerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ServiceControllerClient
	CallOptions **ServiceControllerCallOptions

	// The gRPC API client.
	serviceControllerClient servicecontrolpb.ServiceControllerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServiceControllerClient creates a new service controller client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Google Service Control API (at /service-control/overview)
//
// Lets clients check and report operations against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
func NewServiceControllerClient(ctx context.Context, opts ...option.ClientOption) (*ServiceControllerClient, error) {
	clientOpts := defaultServiceControllerGRPCClientOptions()
	if newServiceControllerClientHook != nil {
		hookOpts, err := newServiceControllerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ServiceControllerClient{CallOptions: defaultServiceControllerCallOptions()}

	c := &serviceControllerGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		serviceControllerClient: servicecontrolpb.NewServiceControllerClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *serviceControllerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *serviceControllerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *serviceControllerGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type serviceControllerRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ServiceControllerClient
	CallOptions **ServiceControllerCallOptions
}

// NewServiceControllerRESTClient creates a new service controller rest client.
//
// Google Service Control API (at /service-control/overview)
//
// Lets clients check and report operations against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
func NewServiceControllerRESTClient(ctx context.Context, opts ...option.ClientOption) (*ServiceControllerClient, error) {
	clientOpts := append(defaultServiceControllerRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultServiceControllerRESTCallOptions()
	c := &serviceControllerRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ServiceControllerClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultServiceControllerRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://servicecontrol.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://servicecontrol.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://servicecontrol.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *serviceControllerRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *serviceControllerRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *serviceControllerRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *serviceControllerGRPCClient) Check(ctx context.Context, req *servicecontrolpb.CheckRequest, opts ...gax.CallOption) (*servicecontrolpb.CheckResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Check[0:len((*c.CallOptions).Check):len((*c.CallOptions).Check)], opts...)
	var resp *servicecontrolpb.CheckResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceControllerClient.Check(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceControllerGRPCClient) Report(ctx context.Context, req *servicecontrolpb.ReportRequest, opts ...gax.CallOption) (*servicecontrolpb.ReportResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 16000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Report[0:len((*c.CallOptions).Report):len((*c.CallOptions).Report)], opts...)
	var resp *servicecontrolpb.ReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceControllerClient.Report(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Check checks whether an operation on a service should be allowed to proceed
// based on the configuration of the service and related policies. It must be
// called before the operation is executed.
//
// If feasible, the client should cache the check results and reuse them for
// 60 seconds. In case of any server errors, the client should rely on the
// cached results for much longer time to avoid outage.
// WARNING: There is general 60s delay for the configuration and policy
// propagation, therefore callers MUST NOT depend on the Check method having
// the latest policy information.
//
// NOTE: the CheckRequest has
// the size limit (wire-format byte size) of 1MB.
//
// This method requires the servicemanagement.services.check permission
// on the specified service. For more information, see
// Cloud IAM (at https://cloud.google.com/iam).
func (c *serviceControllerRESTClient) Check(ctx context.Context, req *servicecontrolpb.CheckRequest, opts ...gax.CallOption) (*servicecontrolpb.CheckResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/services/%v:check", req.GetServiceName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Check[0:len((*c.CallOptions).Check):len((*c.CallOptions).Check)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &servicecontrolpb.CheckResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Report reports operation results to Google Service Control, such as logs and
// metrics. It should be called after an operation is completed.
//
// If feasible, the client should aggregate reporting data for up to 5
// seconds to reduce API traffic. Limiting aggregation to 5 seconds is to
// reduce data loss during client crashes. Clients should carefully choose
// the aggregation time window to avoid data loss risk more than 0.01%
// for business and compliance reasons.
//
// NOTE: the ReportRequest has
// the size limit (wire-format byte size) of 1MB.
//
// This method requires the servicemanagement.services.report permission
// on the specified service. For more information, see
// Google Cloud IAM (at https://cloud.google.com/iam).
func (c *serviceControllerRESTClient) Report(ctx context.Context, req *servicecontrolpb.ReportRequest, opts ...gax.CallOption) (*servicecontrolpb.ReportResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/services/%v:report", req.GetServiceName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Report[0:len((*c.CallOptions).Report):len((*c.CallOptions).Report)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &servicecontrolpb.ReportResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
