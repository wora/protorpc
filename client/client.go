
package client

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import context "golang.org/x/net/context"

import "bytes"
import "io"
import "io/ioutil"
import "net/http"
import "strings"

// A generic client for performing proto-over-http RPCs. This client
// lets the application control the RPC endpoints, error format,
// and other settings, so it can work with different wire protocols,
//
// The main functionality of the client is to marshal the request
// message and unmarshals the response message, and use the provided
// http transport to handle the http request.
type Client struct {
	// REQUIRED. The base url used for this client.
	BaseURL string
	// REQUIRED. The http client used for this client.
	HTTP *http.Client
	// REQUIRED. The content type used for the proto encoding.  It must
	// start with "application/json" for JSON encoding or contain the
	// word "protobuf" for proto binary encoding. Other encodings can be
	// supported in the future.
	ContentType string
	// REQUIRED. The user agent string for sending request.
	UserAgent string
	// OPTIONAL. Google API Key.
	ApiKey string
	// REQUIRED. The proto message for the error payload.
	Status proto.Message
}

// Represent an error, supporting both client error and server error.
type Error struct {
	// The http status code. For client errors, the `Code` is 900.
	Code int
	// The http status message or local error message if the Code is 900.
	Message string
	// The error payload.
	Status proto.Message
}

// Implement error interface for Status.
func (err *Error) Error() string {
	if err.Status != nil {
		return proto.MarshalTextString(err.Status)
	}
	return err.Message
}

// Make a RPC call and return an `Error` if any.
//
// The method name will be appended to the BaseURL to form the
// full URL for making the RPC call. The method name may contain
// URL query parameter(s), so it can address arbitrary RPC call
// that can be expressed as an HTTP URL.
//
// The `req` and `res` are the request and the response message
// as defined in the proto file.
func (c *Client) Call(ctx context.Context, method string, req proto.Message, res proto.Message) error {
	request, err := c.createRequest(ctx, c.BaseURL+method, req)
	if err != nil {
		return &Error{900, err.Error(), nil}
	}
	response, err := c.sendRequest(ctx, request)
	if err != nil {
		return &Error{900, err.Error(), nil}
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Handle successful response.
		err = c.handleResponse(ctx, response, res)
		if err != nil {
			return &Error{900, err.Error(), nil}
		}
		return nil
	} else {
		// Handle error response.
		status := proto.Clone(c.Status)
		if c.handleResponse(ctx, response, status) != nil {
			status = nil
		}
		return &Error{response.StatusCode, response.Status, status}
	}
}

func (c *Client) createRequest(ctx context.Context, url string, req proto.Message) (*http.Request, error) {
	var body io.Reader
	if strings.Contains(c.ContentType, "protobuf") {
		data, err := proto.Marshal(req)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(data)
	} else {
		data, err := (&jsonpb.Marshaler{}).MarshalToString(req)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBufferString(data)
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", c.ContentType)
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
	if strings.Contains(ct, "protobuf") {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return proto.Unmarshal(data, res)
	}
	if strings.HasPrefix(ct, "application/json") {
		return jsonpb.Unmarshal(response.Body, res)
	}
	return &Error{900, "Unsupported content type '" + ct + "'.", nil}
}
