// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: cnquery_explorer.proto

package explorer

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/metadata"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// service interface definition

type QueryHub interface {
	SetBundle(context.Context, *Bundle) (*Empty, error)
	DeleteQueryPack(context.Context, *Mrn) (*Empty, error)
	ValidateBundle(context.Context, *Bundle) (*Empty, error)
	GetBundle(context.Context, *Mrn) (*Bundle, error)
	GetQueryPack(context.Context, *Mrn) (*QueryPack, error)
	GetFilters(context.Context, *Mrn) (*Mqueries, error)
	List(context.Context, *ListReq) (*QueryPacks, error)
	DefaultPacks(context.Context, *DefaultPacksReq) (*URLs, error)
}

// client implementation

type QueryHubClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewQueryHubClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*QueryHubClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./QueryHub")
	if err != nil {
		return nil, err
	}

	serviceClient := &QueryHubClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *QueryHubClient) SetBundle(ctx context.Context, in *Bundle) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/SetBundle"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) DeleteQueryPack(ctx context.Context, in *Mrn) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/DeleteQueryPack"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) ValidateBundle(ctx context.Context, in *Bundle) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/ValidateBundle"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) GetBundle(ctx context.Context, in *Mrn) (*Bundle, error) {
	out := new(Bundle)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetBundle"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) GetQueryPack(ctx context.Context, in *Mrn) (*QueryPack, error) {
	out := new(QueryPack)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetQueryPack"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) GetFilters(ctx context.Context, in *Mrn) (*Mqueries, error) {
	out := new(Mqueries)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetFilters"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) List(ctx context.Context, in *ListReq) (*QueryPacks, error) {
	out := new(QueryPacks)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/List"}, ""), in, out)
	return out, err
}
func (c *QueryHubClient) DefaultPacks(ctx context.Context, in *DefaultPacksReq) (*URLs, error) {
	out := new(URLs)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/DefaultPacks"}, ""), in, out)
	return out, err
}

// server implementation

type QueryHubServerOption func(s *QueryHubServer)

func WithUnknownFieldsForQueryHubServer() QueryHubServerOption {
	return func(s *QueryHubServer) {
		s.allowUnknownFields = true
	}
}

func NewQueryHubServer(handler QueryHub, opts ...QueryHubServerOption) http.Handler {
	srv := &QueryHubServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "QueryHub",
		Methods: map[string]ranger.Method{
			"SetBundle":       srv.SetBundle,
			"DeleteQueryPack": srv.DeleteQueryPack,
			"ValidateBundle":  srv.ValidateBundle,
			"GetBundle":       srv.GetBundle,
			"GetQueryPack":    srv.GetQueryPack,
			"GetFilters":      srv.GetFilters,
			"List":            srv.List,
			"DefaultPacks":    srv.DefaultPacks,
		},
	}
	return ranger.NewRPCServer(&service)
}

type QueryHubServer struct {
	handler            QueryHub
	allowUnknownFields bool
}

func (p *QueryHubServer) SetBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Bundle
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.SetBundle(ctx, &req)
}
func (p *QueryHubServer) DeleteQueryPack(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.DeleteQueryPack(ctx, &req)
}
func (p *QueryHubServer) ValidateBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Bundle
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.ValidateBundle(ctx, &req)
}
func (p *QueryHubServer) GetBundle(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetBundle(ctx, &req)
}
func (p *QueryHubServer) GetQueryPack(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetQueryPack(ctx, &req)
}
func (p *QueryHubServer) GetFilters(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Mrn
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetFilters(ctx, &req)
}
func (p *QueryHubServer) List(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req ListReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.List(ctx, &req)
}
func (p *QueryHubServer) DefaultPacks(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req DefaultPacksReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.DefaultPacks(ctx, &req)
}

// service interface definition

type QueryConductor interface {
	Assign(context.Context, *Assignment) (*Empty, error)
	Unassign(context.Context, *Assignment) (*Empty, error)
	Resolve(context.Context, *ResolveReq) (*ResolvedPack, error)
	StoreResults(context.Context, *StoreResultsReq) (*Empty, error)
	GetReport(context.Context, *EntityDataRequest) (*Report, error)
}

// client implementation

type QueryConductorClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewQueryConductorClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*QueryConductorClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./QueryConductor")
	if err != nil {
		return nil, err
	}

	serviceClient := &QueryConductorClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *QueryConductorClient) Assign(ctx context.Context, in *Assignment) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Assign"}, ""), in, out)
	return out, err
}
func (c *QueryConductorClient) Unassign(ctx context.Context, in *Assignment) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Unassign"}, ""), in, out)
	return out, err
}
func (c *QueryConductorClient) Resolve(ctx context.Context, in *ResolveReq) (*ResolvedPack, error) {
	out := new(ResolvedPack)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Resolve"}, ""), in, out)
	return out, err
}
func (c *QueryConductorClient) StoreResults(ctx context.Context, in *StoreResultsReq) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/StoreResults"}, ""), in, out)
	return out, err
}
func (c *QueryConductorClient) GetReport(ctx context.Context, in *EntityDataRequest) (*Report, error) {
	out := new(Report)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GetReport"}, ""), in, out)
	return out, err
}

// server implementation

type QueryConductorServerOption func(s *QueryConductorServer)

func WithUnknownFieldsForQueryConductorServer() QueryConductorServerOption {
	return func(s *QueryConductorServer) {
		s.allowUnknownFields = true
	}
}

func NewQueryConductorServer(handler QueryConductor, opts ...QueryConductorServerOption) http.Handler {
	srv := &QueryConductorServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "QueryConductor",
		Methods: map[string]ranger.Method{
			"Assign":       srv.Assign,
			"Unassign":     srv.Unassign,
			"Resolve":      srv.Resolve,
			"StoreResults": srv.StoreResults,
			"GetReport":    srv.GetReport,
		},
	}
	return ranger.NewRPCServer(&service)
}

type QueryConductorServer struct {
	handler            QueryConductor
	allowUnknownFields bool
}

func (p *QueryConductorServer) Assign(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Assignment
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Assign(ctx, &req)
}
func (p *QueryConductorServer) Unassign(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Assignment
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Unassign(ctx, &req)
}
func (p *QueryConductorServer) Resolve(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req ResolveReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Resolve(ctx, &req)
}
func (p *QueryConductorServer) StoreResults(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req StoreResultsReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.StoreResults(ctx, &req)
}
func (p *QueryConductorServer) GetReport(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req EntityDataRequest
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GetReport(ctx, &req)
}