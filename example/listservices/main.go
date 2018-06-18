package main

import "net/http"
import "fmt"
import "context"
import "os"

import client "github.com/wora/protorpc/client"
import proto "github.com/golang/protobuf/proto"
import servicemanagement "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Usage: cmd baseUrl apiKey")
		return
	}
	c := &client.Client{
		HTTP: http.DefaultClient,
		BaseURL: os.Args[1],
		UserAgent: "listservices/0.1",
		ApiKey: os.Args[2],
	}
	request := &servicemanagement.ListServicesRequest{}
	response := &servicemanagement.ListServicesResponse{}
	err := c.Call(context.Background(), "ListServices", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
