/*
Package orion is a small lightweight framework written around grpc with the aim to shorten time to build microservices

Source code for Orion can be found at https://github.com/carousell/Orion

It is derived from 'Framework' a small microservices framework written and used inside https://carousell.com, It comes with a number of sensible defaults such as zipkin tracing, hystrix, live reload of configuration, etc.


Why Orion

orion uses protocol-buffers definitions (https://developers.google.com/protocol-buffers/docs/proto3) using gRPC and orion proto plugin as base for building services.

using proto definitions as our service base allows us to define clean contracts that everyone can understand and enables auto generation of client code.

you define your services as a proto definition, for example
	service SimpleService{
		rpc Echo (EchoRequest) returns (EchoResponse){
		}
	}
	message EchoRequest {
		string msg = 1;
	}
	message EchoResponse {
		string msg = 1;
	}
The above definition represents a sample service named 'SimpleService' which accepts 'EchoRequest' and returns 'EchoResponse' at 'Echo' endpoint.

after you have generated the code from protoc (using grpc and orion plugin), you need to implement the server interface generated by gRPC, orion uses this service definition and enables HTTP/gRPC calls to be made to the same service implementation.


Whats Incuded

it comes included with.

	Hystrix (http://github.com/afex/hystrix-go/hystrix)
	Zipkin (http://github.com/opentracing/opentracing-go)
	NewRelic (http://github.com/newrelic/go-agent)
	Prometheus (http://github.com/grpc-ecosystem/go-grpc-prometheus)
	Pprof (https://golang.org/pkg/net/http/pprof/))
	Configuration (http://github.com/spf13/viper)
	Live Configuration Reload (http://github.com/carousell/Orion/utils/listenerutils)
	And much more...
Getting Started
First follow the install guide at https://github.com/carousell/Orion/blob/master/README.md
*/
package orion

// This comment block (re)generates the documentation.
//go:generate godoc2ghmd -ex -file=README.md github.com/carousell/Orion/orion
//go:generate godoc2ghmd -ex -file=handlers/README.md github.com/carousell/Orion/orion/handlers
