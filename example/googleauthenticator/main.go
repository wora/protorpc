package main

import (
	"os"
	"fmt"
	"context"
	"github.com/wora/protorpc/client"
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/api/servicemanagement/v1"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cmd baseUrl")
		fmt.Println("Sample cmd: go run main.go " +
			"https://servicemanagement.googleapis.com/$rpc/google.api.servicemanagement.v1.ServiceManager/")
		return
	}

	c, err := client.NewClient(context.Background(), nil, os.Args[1])
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
