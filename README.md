# Proto RPC

This project provides a simple and easy-to-use library for working
with various proto-over-http RPC protocols.

## Overview

There are many proto-over-http RPC protocols, such as `prpc` and
`Twirp`. Each protocol uses its own framework and runtime library.
However, most of these protocols differ on the URL structure, HTTP
headers, and the error format. By making these features customizable
by the application, one client library can support multiple wire
protocols with minimum complexity.

Most proto-over-http frameworks ues `protoc` plugins to generate the
client and the server libraries. While such libraries provide type
safety and ease of use, it is often not necessary for many RPC use
cases. This library uses a different API design pattern to avoid any
code generation while provide similar type safety and usability. See
the following example in Go language:

With code generated client
```
request := &FooRequest{...}
response, err := stub.Foo(ctx, request)
```

Without code generated client
```
request := &FooRequest{...}
response := &FooResponse{}
err := client.Call(ctx, "Foo", request, response)
```

Conceptually, we use `printf()` like experience for making RPC
calls. By having an extra parameter to specify the method name,
we can avoid codegen completely. While `printf()` may not provide
type safety, developers hardly ever get it wrong. This library
brings the same experience to RPC calls.

## Status

Currently, this repository only contains the Go client implementation.
It needs to be integrated with an auth library for real use cases.

Since the library has less than 200 lines of code, it should be easy
to replicate the experience to another programming languages as long
as there is proto3 library support for that language.
