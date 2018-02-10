package main

import "net/http"
import "fmt"
import "context"

import client "github.com/wora/protorpc/client"
import proto "github.com/golang/protobuf/proto"

// The following function provides equivalent functionality as the
// Twirp example without using any code generator. The main difference
// is the returned error type is client.Error, which wraps the Twirp
// error payload.
func main() {
	c := &client.Client{
		BaseURL: "http://localhost:8080/twirp/twitch.twirp.example.Haberdasher/",
		HTTP: &http.Client{},
		ContentType: "application/protobuf",
		Status: &Error{},
	}
	request := &Size{Inches: 12}
	response := &Hat{}
	err := c.Call(context.Background(), "MakeHat", request, response)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(proto.MarshalTextString(response))
	}
}
