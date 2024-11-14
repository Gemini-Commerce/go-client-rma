# # RmaReturnResponse


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** |   | [optional]
**Grn**| **string** |   | [optional]
**OrderGrn**| **string** |   | [optional]
**Status**| **string** |   | [optional]
**Products**| [**[]RmaReturnProduct**](RmaReturnProduct.md) |   | [optional]
**PreferredRefundMethod**| [**RmaRefundMethod**](RmaRefundMethod.md) |  for more information please, see Model/RmaRefundMethod.php  | [optional] [default to RMAREFUNDMETHOD_UNKNOWN]
**RefundShippingCost**| **bool** |   | [optional]
**RefundPaymentCost**| **bool** |   | [optional]
**CustomerInfo**| [**RmaCustomerInfo**](RmaCustomerInfo.md) |   | [optional]
**ReturnAddress**| [**RmaPostalAddress**](RmaPostalAddress.md) |   | [optional]
**Note**| **string** |   | [optional]
**History**| [**[]RmaReturnHistory**](RmaReturnHistory.md) |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**OrderData**| [**RmaOrderData**](RmaOrderData.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

