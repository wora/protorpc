# Proto RPC

This project provides a simple and easy-to-use library for handling
proto-over-http RPC protocols.

## Overview

There are many proto-over-http RPC protocols, such as `prpc` and
`Twirp`, and each of them requires its own framework and runtime
library. However, most of them only differ on the URL structure, HTTP
headers, and the error payload format. By making these features
customizable by the application, a single client library can directly
support multiple wire protocols.

Another issue is the existing frameworks typically use `protoc`
plugins to generate the client and the server stubs. While such stubs
provide type safety and ease of use, it may not be necessary for many
RPC use cases. This library uses a different API design to avoid any
code generation while provide similar type safety and usability, see
the following example in Go language:

With code generated stub
```
request := &FooRequest{...}
response, err := stub.Foo(ctx, request)
```

Without code generated stub
```
request := &FooRequest{...}
response := &FooResponse{}
err := client.Call(ctx, "Foo", request, response)
```

## Status

Currently, this repository only contains the Go client implementation.
It needs to be integrated with an auth library for real use cases.

Since the library has less than 200 lines of code, it should be easy
to replicate the experience to another programming languages as long
as there is proto3 library support for that language.
