package main

import "client"
import "net/http"
import "fmt"
import "context"

import sm "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
import proto "github.com/golang/protobuf/proto"
import status "google.golang.org/genproto/googleapis/rpc/status"

func main() {
	c := &client.Client{
		BaseURL: "...",
		HTTP: http.DefaultClient,
		ContentType: "application/x-protobuf",
		UserAgent: "listservices/0.1",
		ApiKey: "...",
		Status: &status.Status{},
	}
	request := &sm.ListServicesRequest{}
	response := &sm.ListServicesResponse{}
	err := c.Call(context.Background(), "ListServices", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
