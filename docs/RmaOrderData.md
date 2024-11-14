# # RmaOrderData


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**Grn**| **string** |   | [optional]
**Number**| **string** |   | [optional]
**Status**| **string** |   | [optional]
**Channel**| **string** |   | [optional]
**Market**| **string** |   | [optional]
**Items**| [**[]RmaOrderDataItem**](RmaOrderDataItem.md) |   | [optional]
**Currency**| [**RmaCurrency**](RmaCurrency.md) |  for more information please, see Model/RmaCurrency.php  | [optional] [default to RMACURRENCY_XXX]
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   | [optional]
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

