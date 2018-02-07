package main

import "client"
import "fmt"
import "context"

import sm "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
import proto "github.com/golang/protobuf/proto"
import status "google.golang.org/genproto/googleapis/rpc/status"
import google "golang.org/x/oauth2/google"

func main() {
	ctx := context.Background()
	http, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	c := &client.Client{
		BaseURL: "...",
		HTTP: http,
		ContentType: "application/x-protobuf",
		UserAgent: "listservices/0.1",
		Status: &status.Status{},
	}
	request := &sm.ListServicesRequest{}
	response := &sm.ListServicesResponse{}
	err = c.Call(ctx, "ListServices", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
