package main

import "fmt"
import "context"

import sm "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
import client "github.com/wora/protorpc/client"
import proto "github.com/golang/protobuf/proto"
import status "google.golang.org/genproto/googleapis/rpc/status"
import google "golang.org/x/oauth2/google"

func NewClient(ctx context.Context, baseUrl string) (*client.Client, error) {
	http, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, err
	}
	c := &client.Client{
		BaseURL:     baseUrl,
		HTTP:        http,
		ContentType: "application/x-protobuf",
		UserAgent:   "protorpc/0.1",
		Status:      &status.Status{},
	}
	return c, nil
}

func main() {
	c, err := NewClient(context.Background(), "...")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	request := &sm.ListServicesRequest{}
	response := &sm.ListServicesResponse{}
	err = c.Call(context.Background(), "ListServices", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
