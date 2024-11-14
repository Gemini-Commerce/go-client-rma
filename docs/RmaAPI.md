# \RmaAPI

All URIs are relative to *https://rma.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNote**](RmaAPI.md#AddNote) | **Post** /rma.Rma/AddNote | Add a note to an RMA
[**ApproveReturn**](RmaAPI.md#ApproveReturn) | **Post** /rma.Rma/ApproveReturn | Approve a return
[**CancelReturn**](RmaAPI.md#CancelReturn) | **Post** /rma.Rma/CancelReturn | Cancel a return
[**ConfirmReturnApproveItems**](RmaAPI.md#ConfirmReturnApproveItems) | **Post** /rma.Rma/ConfirmReturnApproveItems | Confirm return approve items
[**CreateReturn**](RmaAPI.md#CreateReturn) | **Post** /rma.Rma/CreateReturn | Create a return
[**DeleteNote**](RmaAPI.md#DeleteNote) | **Post** /rma.Rma/DeleteNote | Delete a note from an RMA
[**EditNote**](RmaAPI.md#EditNote) | **Post** /rma.Rma/EditNote | Edit a note on an RMA
[**GetReturn**](RmaAPI.md#GetReturn) | **Post** /rma.Rma/GetReturn | Get a return
[**ListNotesByReturnId**](RmaAPI.md#ListNotesByReturnId) | **Post** /rma.Rma/ListNotesByReturnId | List notes by return id
[**ListReturns**](RmaAPI.md#ListReturns) | **Post** /rma.Rma/ListReturns | List returns
[**RefundReturn**](RmaAPI.md#RefundReturn) | **Post** /rma.Rma/RefundReturn | Refund a return
[**RejectReturn**](RmaAPI.md#RejectReturn) | **Post** /rma.Rma/RejectReturn | Reject a return
[**SetReceivedItems**](RmaAPI.md#SetReceivedItems) | **Post** /rma.Rma/SetReceivedItems | Set received items
[**SkipReturnStatus**](RmaAPI.md#SkipReturnStatus) | **Post** /rma.Rma/SkipReturnStatus | Skip a return status



## AddNote

> RmaNoteResponse AddNote(ctx).Body(body).Execute()

Add a note to an RMA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaAddNoteRequest("TenantId_example", "ReturnId_example", "Author_example", "Note_example") // RmaAddNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.AddNote(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.AddNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNote`: RmaNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.AddNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaAddNoteRequest**](RmaAddNoteRequest.md) |  | 

### Return type

[**RmaNoteResponse**](RmaNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveReturn

> RpcStatus ApproveReturn(ctx).Body(body).Execute()

Approve a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaApproveReturnRequest("TenantId_example", "Id_example", false, false, []openapiclient.RmaApproveReturnRequestItem{*openapiclient.NewRmaApproveReturnRequestItem("Grn_example", "Quantity_example")}) // RmaApproveReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.ApproveReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.ApproveReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveReturn`: RpcStatus
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.ApproveReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaApproveReturnRequest**](RmaApproveReturnRequest.md) |  | 

### Return type

[**RpcStatus**](RpcStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelReturn

> map[string]interface{} CancelReturn(ctx).Body(body).Execute()

Cancel a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaCancelReturnRequest("TenantId_example", "Id_example") // RmaCancelReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.CancelReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.CancelReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelReturn`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.CancelReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaCancelReturnRequest**](RmaCancelReturnRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmReturnApproveItems

> map[string]interface{} ConfirmReturnApproveItems(ctx).Body(body).Execute()

Confirm return approve items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaConfirmReturnApproveItemsRequest("TenantId_example", "Id_example", []openapiclient.RmaConfirmReturnApproveItemsRequestItem{*openapiclient.NewRmaConfirmReturnApproveItemsRequestItem("Grn_example", "Quantity_example")}) // RmaConfirmReturnApproveItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.ConfirmReturnApproveItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.ConfirmReturnApproveItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmReturnApproveItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.ConfirmReturnApproveItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmReturnApproveItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaConfirmReturnApproveItemsRequest**](RmaConfirmReturnApproveItemsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReturn

> RmaReturnResponse CreateReturn(ctx).Body(body).Execute()

Create a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaCreateReturnRequest("TenantId_example", "OrderGrn_example", []openapiclient.CreateReturnRequestProduct{*openapiclient.NewCreateReturnRequestProduct("Grn_example", "Quantity_example")}, openapiclient.rmaRefundMethod("REFUND_METHOD_UNKNOWN")) // RmaCreateReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.CreateReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.CreateReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReturn`: RmaReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.CreateReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaCreateReturnRequest**](RmaCreateReturnRequest.md) |  | 

### Return type

[**RmaReturnResponse**](RmaReturnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNote

> map[string]interface{} DeleteNote(ctx).Body(body).Execute()

Delete a note from an RMA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaDeleteNoteRequest("TenantId_example", "Id_example") // RmaDeleteNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.DeleteNote(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.DeleteNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.DeleteNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaDeleteNoteRequest**](RmaDeleteNoteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditNote

> RmaNoteResponse EditNote(ctx).Body(body).Execute()

Edit a note on an RMA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaEditNoteRequest("TenantId_example", "Id_example") // RmaEditNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.EditNote(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.EditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditNote`: RmaNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.EditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaEditNoteRequest**](RmaEditNoteRequest.md) |  | 

### Return type

[**RmaNoteResponse**](RmaNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturn

> RmaReturnResponse GetReturn(ctx).Body(body).Execute()

Get a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaGetReturnRequest("TenantId_example", "Id_example") // RmaGetReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.GetReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.GetReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturn`: RmaReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.GetReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaGetReturnRequest**](RmaGetReturnRequest.md) |  | 

### Return type

[**RmaReturnResponse**](RmaReturnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotesByReturnId

> RmaListNotesByReturnIdResponse ListNotesByReturnId(ctx).Body(body).Execute()

List notes by return id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaListNotesByReturnIdRequest("TenantId_example", "ReturnId_example") // RmaListNotesByReturnIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.ListNotesByReturnId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.ListNotesByReturnId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotesByReturnId`: RmaListNotesByReturnIdResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.ListNotesByReturnId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotesByReturnIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaListNotesByReturnIdRequest**](RmaListNotesByReturnIdRequest.md) |  | 

### Return type

[**RmaListNotesByReturnIdResponse**](RmaListNotesByReturnIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReturns

> RmaListReturnsResponse ListReturns(ctx).Body(body).Execute()

List returns

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaListReturnsRequest("TenantId_example") // RmaListReturnsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.ListReturns(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.ListReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReturns`: RmaListReturnsResponse
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.ListReturns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaListReturnsRequest**](RmaListReturnsRequest.md) |  | 

### Return type

[**RmaListReturnsResponse**](RmaListReturnsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundReturn

> map[string]interface{} RefundReturn(ctx).Body(body).Execute()

Refund a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaRefundReturnRequest("TenantId_example", "Id_example") // RmaRefundReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.RefundReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.RefundReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundReturn`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.RefundReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaRefundReturnRequest**](RmaRefundReturnRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectReturn

> map[string]interface{} RejectReturn(ctx).Body(body).Execute()

Reject a return

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaRejectReturnRequest("TenantId_example", "Id_example") // RmaRejectReturnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.RejectReturn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.RejectReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectReturn`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.RejectReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaRejectReturnRequest**](RmaRejectReturnRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReceivedItems

> map[string]interface{} SetReceivedItems(ctx).Body(body).Execute()

Set received items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaSetReceivedItemsRequest("TenantId_example", "Id_example", []openapiclient.RmaSetReceivedItemsRequestItem{*openapiclient.NewRmaSetReceivedItemsRequestItem("Grn_example", "Quantity_example")}) // RmaSetReceivedItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.SetReceivedItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.SetReceivedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetReceivedItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.SetReceivedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetReceivedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaSetReceivedItemsRequest**](RmaSetReceivedItemsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipReturnStatus

> map[string]interface{} SkipReturnStatus(ctx).Body(body).Execute()

Skip a return status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-rma"
)

func main() {
	body := *openapiclient.NewRmaSkipReturnStatusRequest("TenantId_example", "Id_example") // RmaSkipReturnStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaAPI.SkipReturnStatus(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaAPI.SkipReturnStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkipReturnStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RmaAPI.SkipReturnStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkipReturnStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RmaSkipReturnStatusRequest**](RmaSkipReturnStatusRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

