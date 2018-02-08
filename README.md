# Proto RPC

This project provides a simple and easy-to-use library for working
with multiple proto-over-http RPC protocols.

## Overview

There are many proto-over-http RPC protocols, such as `prpc` and
`Twirp`. Each protocol uses its own framework and runtime library.
However, most of these protocols only differ on the URL structure,
HTTP headers, and the error format. By making these features
customizable by the application, a single library can support
multiple protocols with great usability and minimum complexity.

## Usage

Most proto-over-http frameworks ues `protoc` plugins to generate the
client and the server libraries. While such libraries provide type
safety and ease of use, they are not strictly necessary for many RPC
use cases. The plugins add complexity to the build system, and often
have compatibility issues.

This library uses a different *API design pattern* to avoid the
code generation, while provide similar type safety and usability,
see below:

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

## Status

Currently, this repository only contains the Go client implementation.
It can be seamlessly integrated with an auth library for real use
cases. Please see the `example` directory on how to use the library.

Since the library has much less than 200 lines of code, it can be
easily transcoded into other programming languages where proto3 is
supported.
