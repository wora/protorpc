package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
	"github.com/shinfan/sgauth"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/status"
)

// A generic client for performing proto-over-http RPCs. This client
// lets the application control the RPC endpoints, error format,
// and other settings, so it can work with different wire protocols,
//
// The main functionality of the client is to marshal the request
// message and unmarshals the response message, and use the provided
// http transport to handle the http request.
type Client struct {
	// REQUIRED. The HTTP client used for this API client stub.
	// The client library assumes the HTTP client also handles
	// API authentication by attaching the correct Authorization
	// header to the request.
	HTTP *http.Client
	// REQUIRED. The base URL used for this client stub.
	BaseURL string
	// REQUIRED. The user agent string for sending the request.
	UserAgent string
	// OPTIONAL. The Google API Key used for sending the request.
	ApiKey string

}

// Creates a client with the given credentials. Application default credentials
// will be used if no credentials are provided.
// For more information about application default credentials please read:
// https://cloud.google.com/docs/authentication/production
func NewClient(ctx context.Context, settings *sgauth.Settings, baseUrl string) (*Client, error) {
	if settings == nil {
		settings = &sgauth.Settings{}
	}
	if settings.AuthMethod() == sgauth.MethodJWT && settings.Audience == "" {
		// Derive audience from url if using JWT and aud is not provided explicitly.
		settings.Audience = createAudience(baseUrl)
	}
	return createClient(ctx, settings, baseUrl)
}

func createClient(ctx context.Context, settings *sgauth.Settings, baseUrl string) (*Client, error) {
	http, err := sgauth.NewHTTPClient(ctx, settings)
	if err != nil {
		return nil, err
	}
	return &Client{
		HTTP:        http,
		BaseURL:     baseUrl,
		UserAgent:   "protorpc/0.1",
		ApiKey: settings.APIKey,
	}, nil
}

// Defines `google.rpc.Status` as an error type.
type Error status.Status

func (e *Error) Error() string {
     return fmt.Sprintf("gRPC error: code %d, message %q", e.Code, e.Message)
}

// Make a RPC call and return an `Error` if any.
//
// The method name will be appended to the BaseURL to form the
// full URL for making the RPC call. The method name may contain
// URL query parameter(s), so it can address arbitrary RPC call
// that can be expressed as an HTTP URL.
//
// The `req` and `res` are the request and the response message.
// For RPC errors, the returned error will be `google.rpc.Status`.
func (c *Client) Call(ctx context.Context, method string, req proto.Message, res proto.Message) error {
	request, err := c.createRequest(ctx, c.BaseURL+method, req)
	if err != nil {
		return err
	}
	response, err := c.sendRequest(ctx, request)
	if err != nil {
		return err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Handle OK response.
		return c.handleResponse(ctx, response, res)
	} else {
		// Handle error response.
		s := &status.Status{}
		return c.handleResponse(ctx, response, s)
	}
}

func (c *Client) createRequest(ctx context.Context, url string, req proto.Message) (*http.Request, error) {
	var body io.Reader
	// Marshalls request message into bytes.
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	body = bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-protobuf")
	request.Header.Set("Accept", "application/x-protobuf")
	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}
	if c.ApiKey != "" {
		request.Header.Set("X-Goog-Api-Key", c.ApiKey)
	}
	return request.WithContext(ctx), nil
}

func (c *Client) sendRequest(ctx context.Context, request *http.Request) (*http.Response, error) {
	return c.HTTP.Do(request)
}

func (c *Client) handleResponse(ctx context.Context, response *http.Response, res proto.Message) error {
	defer response.Body.Close()
	ct := response.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/x-protobuf") {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		fmt.Println(proto.Unmarshal(data, res))
		fmt.Println(res)
		return proto.Unmarshal(data, res)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	return &Error{Code: 2, Message: "Unsupported content type '" + ct + "'."}
}
