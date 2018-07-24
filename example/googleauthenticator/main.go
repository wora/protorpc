package main

import (
	"os"
	"fmt"
	"context"
	"github.com/shinfan/sgauth"

	client "github.com/wora/protorpc/client"
	proto "github.com/golang/protobuf/proto"
	servicemanagement "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
)

func NewClient(ctx context.Context, service_name string, api_name string) (*client.Client, error) {
	var credentials = &sgauth.Credentials{
		ServiceAccount: &sgauth.ServiceAccount{
			ServiceName: service_name,
			APIName: api_name,
		},
	}
	return client.NewClient(ctx, credentials)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: cmd service_name api_name")
		fmt.Println("Sample cmd: go run main.go " +
			"servicemanagement.googleapis.com google.api.servicemanagement.v1.ServiceManager")
		return
	}
	c, err := NewClient(context.Background(), os.Args[1], os.Args[2])
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	request := &servicemanagement.ListServicesRequest{}
	response := &servicemanagement.ListServicesResponse{}
	err = c.Call(context.Background(), "ListServices", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
