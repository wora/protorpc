# Proto RPC

This project provides a simple and easy-to-use library for working
with multiple proto-over-http RPC protocols.

## Introduction

There are many proto-over-http RPC protocols, such as `prpc` and
`Twirp`. Each protocol uses its own framework and runtime library.
However, most of these protocols only differ on the URL structure,
HTTP headers, and the error format. By making these features
customizable by the application, a single library can support
multiple wire protocols with great usability and minimum complexity.

## Design

Most proto-over-http frameworks ues `protoc` plugins to generate the
client and the server libraries. While such libraries provide type
safety and ease of use, they also put restrictions on many use cases,
such as mandating the URL structure. The `protoc` plugins also add
complexity to the build system, and introduce compatibility issues.

This library uses a different *API design pattern* to avoid the
code generation, while provide good type safety, flexibility and
usability, see the following below:

With code generation
```Go
stub := NewStub(...)
request := &FooRequest{...}
response, err := stub.Foo(ctx, request)
```

Without code generation
```Go
client := NewClient(...)
request := &FooRequest{...}
response := &FooResponse{}
err := client.Call(ctx, "Foo", request, response)
```

Conceptually, we use `printf()` like experience for making RPC
calls. By having an extra parameter to specify the method name,
we can avoid codegen completely. While `printf()` may not provide
type safety, developers rarely use it wrong. This library brings
the same experience to RPC calls.

## Protocol Support

**URL Request -> Response | Error**

This library supports any proto-over-http protocol that meets
these requirements:

* Each RPC endpoint is identified by a unique HTTP URL, which is
  published by API documentation or API service discovery. This
  library does not impose any restriction on the URL structure.

* Each RPC request is sent using an HTTP POST method. Other HTTP
  methods can be added later if needed.

* The request and response messages are passed via HTTP bodies.

* The error response is passed via HTTP response body.

* Only proto JSON and proto binary encodings are supported, but
  proto text and proto yaml can be added in the future. The
  encoding format must be specified by the `Content-Type` HTTP
  header.

## Status

Currently, this repository only contains the Go client implementation.
It can be seamlessly integrated with an auth library for real use
cases. Please see the `example` directory on how to use the library.

Since the library has much less than 200 lines of code, it can be
easily transcoded into other programming languages where proto is
supported.
