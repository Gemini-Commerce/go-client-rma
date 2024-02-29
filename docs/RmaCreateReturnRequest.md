# # RmaCreateReturnRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**OrderGrn**| **string** |   |
**Products**| [**[]CreateReturnRequestProduct**](CreateReturnRequestProduct.md) |   |
**PreferredRefundMethod**| [**RmaRefundMethod**](RmaRefundMethod.md) |  for more information please, see Model/RmaRefundMethod.php  | [default to UNKNOWN]
**RefundShippingCost**| **bool** |   | [optional]
**RefundPaymentCost**| **bool** |   | [optional]
**CustomerInfo**| [**RmaCustomerInfo**](RmaCustomerInfo.md) |   | [optional]
**ReturnAddress**| [**RmaPostalAddress**](RmaPostalAddress.md) |   | [optional]
**Note**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

