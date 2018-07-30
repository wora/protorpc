package main

import (
	"os"
	"fmt"
	"context"
	"github.com/wora/protorpc/client"
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/api/servicemanagement/v1"
	"github.com/shinfan/sgauth"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: cmd aud baseUrl")
		return
	}
	http, err := sgauth.NewHTTPClient(context.Background(), &sgauth.Settings{
		Audience: os.Args[1],
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	c:= &client.Client{
		HTTP:        http,
		BaseURL:     os.Args[2],
		UserAgent:   "protorpc/0.1",
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
