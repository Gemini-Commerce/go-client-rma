# Go API client for rma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import rma "github.com/Gemini-Commerce/go-client-rma"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `rma.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), rma.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `rma.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), rma.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `rma.ContextOperationServerIndices` and `rma.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), rma.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), rma.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://rma.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*RmaAPI* | [**AddNote**](docs/RmaAPI.md#addnote) | **Post** /rma.Rma/AddNote | Add a note to an RMA
*RmaAPI* | [**ApproveReturn**](docs/RmaAPI.md#approvereturn) | **Post** /rma.Rma/ApproveReturn | Approve a return
*RmaAPI* | [**CancelReturn**](docs/RmaAPI.md#cancelreturn) | **Post** /rma.Rma/CancelReturn | Cancel a return
*RmaAPI* | [**ConfirmReturnApproveItems**](docs/RmaAPI.md#confirmreturnapproveitems) | **Post** /rma.Rma/ConfirmReturnApproveItems | Confirm return approve items
*RmaAPI* | [**CreateReturn**](docs/RmaAPI.md#createreturn) | **Post** /rma.Rma/CreateReturn | Create a return
*RmaAPI* | [**DeleteNote**](docs/RmaAPI.md#deletenote) | **Post** /rma.Rma/DeleteNote | Delete a note from an RMA
*RmaAPI* | [**EditNote**](docs/RmaAPI.md#editnote) | **Post** /rma.Rma/EditNote | Edit a note on an RMA
*RmaAPI* | [**GetReturn**](docs/RmaAPI.md#getreturn) | **Post** /rma.Rma/GetReturn | Get a return
*RmaAPI* | [**ListNotesByReturnId**](docs/RmaAPI.md#listnotesbyreturnid) | **Post** /rma.Rma/ListNotesByReturnId | List notes by return id
*RmaAPI* | [**ListReturns**](docs/RmaAPI.md#listreturns) | **Post** /rma.Rma/ListReturns | List returns
*RmaAPI* | [**RefundReturn**](docs/RmaAPI.md#refundreturn) | **Post** /rma.Rma/RefundReturn | Refund a return
*RmaAPI* | [**RejectReturn**](docs/RmaAPI.md#rejectreturn) | **Post** /rma.Rma/RejectReturn | Reject a return
*RmaAPI* | [**SetReceivedItems**](docs/RmaAPI.md#setreceiveditems) | **Post** /rma.Rma/SetReceivedItems | Set received items
*RmaAPI* | [**SkipReturnStatus**](docs/RmaAPI.md#skipreturnstatus) | **Post** /rma.Rma/SkipReturnStatus | Skip a return status


## Documentation For Models

 - [CreateReturnRequestProduct](docs/CreateReturnRequestProduct.md)
 - [EditNoteRequestPayload](docs/EditNoteRequestPayload.md)
 - [ListReturnsRequestFilter](docs/ListReturnsRequestFilter.md)
 - [OrderDataSubtotal](docs/OrderDataSubtotal.md)
 - [OrderDataSubtotalCode](docs/OrderDataSubtotalCode.md)
 - [OrderDataTotal](docs/OrderDataTotal.md)
 - [OrderDataTotalCode](docs/OrderDataTotalCode.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [ProtobufNullValue](docs/ProtobufNullValue.md)
 - [RmaAddNoteRequest](docs/RmaAddNoteRequest.md)
 - [RmaApproveReturnRequest](docs/RmaApproveReturnRequest.md)
 - [RmaApproveReturnRequestItem](docs/RmaApproveReturnRequestItem.md)
 - [RmaCancelReturnRequest](docs/RmaCancelReturnRequest.md)
 - [RmaConfirmReturnApproveItemsRequest](docs/RmaConfirmReturnApproveItemsRequest.md)
 - [RmaConfirmReturnApproveItemsRequestItem](docs/RmaConfirmReturnApproveItemsRequestItem.md)
 - [RmaCreateReturnRequest](docs/RmaCreateReturnRequest.md)
 - [RmaCurrency](docs/RmaCurrency.md)
 - [RmaCustomerInfo](docs/RmaCustomerInfo.md)
 - [RmaDeleteNoteRequest](docs/RmaDeleteNoteRequest.md)
 - [RmaEditNoteRequest](docs/RmaEditNoteRequest.md)
 - [RmaGetReturnRequest](docs/RmaGetReturnRequest.md)
 - [RmaListNotesByReturnIdRequest](docs/RmaListNotesByReturnIdRequest.md)
 - [RmaListNotesByReturnIdRequestSort](docs/RmaListNotesByReturnIdRequestSort.md)
 - [RmaListNotesByReturnIdRequestSortSortField](docs/RmaListNotesByReturnIdRequestSortSortField.md)
 - [RmaListNotesByReturnIdResponse](docs/RmaListNotesByReturnIdResponse.md)
 - [RmaListReturnsRequest](docs/RmaListReturnsRequest.md)
 - [RmaListReturnsRequestSort](docs/RmaListReturnsRequestSort.md)
 - [RmaListReturnsRequestSortSortField](docs/RmaListReturnsRequestSortSortField.md)
 - [RmaListReturnsResponse](docs/RmaListReturnsResponse.md)
 - [RmaMoney](docs/RmaMoney.md)
 - [RmaNoteResponse](docs/RmaNoteResponse.md)
 - [RmaOrderData](docs/RmaOrderData.md)
 - [RmaOrderDataItem](docs/RmaOrderDataItem.md)
 - [RmaPostalAddress](docs/RmaPostalAddress.md)
 - [RmaRefundMethod](docs/RmaRefundMethod.md)
 - [RmaRefundReturnRequest](docs/RmaRefundReturnRequest.md)
 - [RmaRejectReturnRequest](docs/RmaRejectReturnRequest.md)
 - [RmaReturnHistory](docs/RmaReturnHistory.md)
 - [RmaReturnProduct](docs/RmaReturnProduct.md)
 - [RmaReturnProductProperty](docs/RmaReturnProductProperty.md)
 - [RmaReturnResponse](docs/RmaReturnResponse.md)
 - [RmaSetReceivedItemsRequest](docs/RmaSetReceivedItemsRequest.md)
 - [RmaSetReceivedItemsRequestItem](docs/RmaSetReceivedItemsRequestItem.md)
 - [RmaSkipReturnStatusRequest](docs/RmaSkipReturnStatusRequest.md)
 - [RmaSortOrder](docs/RmaSortOrder.md)
 - [RpcStatus](docs/RpcStatus.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

