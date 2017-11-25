// Package service must implement the generated proto's server interface
package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/carousell/Orion/example/echo/echo_proto"
	"github.com/carousell/Orion/interceptors"
	"github.com/carousell/Orion/utils/headers"
	"google.golang.org/grpc"
)

type svc struct {
	appendText string
	debug      bool
}

func (s *svc) GetRequestHeaders() []string {
	return []string{}
}

func (s *svc) GetResponseHeaders() []string {
	return []string{"Original-Msg"}
}

func GetService(config Config) echo_proto.EchoServiceServer {
	s := new(svc)
	s.appendText = config.AppendText
	s.debug = config.Debug
	return s
}

func (s *svc) Echo(ctx context.Context, req *echo_proto.EchoRequest) (*echo_proto.EchoResponse, error) {
	resp := new(echo_proto.EchoResponse)
	resp.Msg = s.appendText + req.GetMsg()
	return resp, nil
}

func (s *svc) Upper(ctx context.Context, req *echo_proto.UpperRequest) (*echo_proto.UpperResponse, error) {
	resp := new(echo_proto.UpperResponse)
	resp.Msg = strings.ToUpper(s.appendText + req.GetMsg())
	hdrs := headers.RequestHeadersFromContext(ctx)
	if hdrs != nil {
		fmt.Println("All request headers", hdrs.GetAll())
	}
	headers.AddToResponseHeaders(ctx, "original-msg", req.GetMsg())
	return resp, nil
}

func (s *svc) GetInterceptors() []grpc.UnaryServerInterceptor {
	icpt := []grpc.UnaryServerInterceptor{}
	if s.debug {
		icpt = append(icpt, interceptors.DebugLoggingInterceptor())
	}
	return icpt
}
