/*
RMA Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rma

import (
	"encoding/json"
)

// checks if the RmaOrderDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaOrderDataItem{}

// RmaOrderDataItem struct for RmaOrderDataItem
type RmaOrderDataItem struct {
	Id *string `json:"id,omitempty"`
	ProductGrn *string `json:"productGrn,omitempty"`
	QtyOrdered *int64 `json:"qtyOrdered,omitempty"`
	QtyCommitted *int64 `json:"qtyCommitted,omitempty"`
	QtyShipped *int64 `json:"qtyShipped,omitempty"`
	UnitSalePrice *RmaMoney `json:"unitSalePrice,omitempty"`
	UnitListPrice *RmaMoney `json:"unitListPrice,omitempty"`
	UnitVatAmount *RmaMoney `json:"unitVatAmount,omitempty"`
	RowSalePrice *RmaMoney `json:"rowSalePrice,omitempty"`
	RowListPrice *RmaMoney `json:"rowListPrice,omitempty"`
	RowVatAmount *RmaMoney `json:"rowVatAmount,omitempty"`
	VatPercentage *float32 `json:"vatPercentage,omitempty"`
	VatInaccurate *bool `json:"vatInaccurate,omitempty"`
	VatCalculated *bool `json:"vatCalculated,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	ProductCode *string `json:"productCode,omitempty"`
	ProductSku *string `json:"productSku,omitempty"`
	ProductOptions *string `json:"productOptions,omitempty"`
	ProductImg *string `json:"productImg,omitempty"`
	ProductData *string `json:"productData,omitempty"`
	ShipmentInfoReference *string `json:"shipmentInfoReference,omitempty"`
	PromotionGrn []string `json:"promotionGrn,omitempty"`
	ProductIsVirtual *bool `json:"productIsVirtual,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaOrderDataItem RmaOrderDataItem

// NewRmaOrderDataItem instantiates a new RmaOrderDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaOrderDataItem() *RmaOrderDataItem {
	this := RmaOrderDataItem{}
	return &this
}

// NewRmaOrderDataItemWithDefaults instantiates a new RmaOrderDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaOrderDataItemWithDefaults() *RmaOrderDataItem {
	this := RmaOrderDataItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// &#39;Has&#39;Id returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;Id() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RmaOrderDataItem) SetId(v string) {
	o.Id = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// &#39;Has&#39;ProductGrn returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *RmaOrderDataItem) SetProductGrn(v string) {
	o.ProductGrn = &v
}

// GetQtyOrdered returns the QtyOrdered field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetQtyOrdered() int64 {
	if o == nil || IsNil(o.QtyOrdered) {
		var ret int64
		return ret
	}
	return *o.QtyOrdered
}

// GetQtyOrderedOk returns a tuple with the QtyOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetQtyOrderedOk() (*int64, bool) {
	if o == nil || IsNil(o.QtyOrdered) {
		return nil, false
	}
	return o.QtyOrdered, true
}

// &#39;Has&#39;QtyOrdered returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;QtyOrdered() bool {
	if o != nil && !IsNil(o.QtyOrdered) {
		return true
	}

	return false
}

// SetQtyOrdered gets a reference to the given int64 and assigns it to the QtyOrdered field.
func (o *RmaOrderDataItem) SetQtyOrdered(v int64) {
	o.QtyOrdered = &v
}

// GetQtyCommitted returns the QtyCommitted field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetQtyCommitted() int64 {
	if o == nil || IsNil(o.QtyCommitted) {
		var ret int64
		return ret
	}
	return *o.QtyCommitted
}

// GetQtyCommittedOk returns a tuple with the QtyCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetQtyCommittedOk() (*int64, bool) {
	if o == nil || IsNil(o.QtyCommitted) {
		return nil, false
	}
	return o.QtyCommitted, true
}

// &#39;Has&#39;QtyCommitted returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;QtyCommitted() bool {
	if o != nil && !IsNil(o.QtyCommitted) {
		return true
	}

	return false
}

// SetQtyCommitted gets a reference to the given int64 and assigns it to the QtyCommitted field.
func (o *RmaOrderDataItem) SetQtyCommitted(v int64) {
	o.QtyCommitted = &v
}

// GetQtyShipped returns the QtyShipped field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetQtyShipped() int64 {
	if o == nil || IsNil(o.QtyShipped) {
		var ret int64
		return ret
	}
	return *o.QtyShipped
}

// GetQtyShippedOk returns a tuple with the QtyShipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetQtyShippedOk() (*int64, bool) {
	if o == nil || IsNil(o.QtyShipped) {
		return nil, false
	}
	return o.QtyShipped, true
}

// &#39;Has&#39;QtyShipped returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;QtyShipped() bool {
	if o != nil && !IsNil(o.QtyShipped) {
		return true
	}

	return false
}

// SetQtyShipped gets a reference to the given int64 and assigns it to the QtyShipped field.
func (o *RmaOrderDataItem) SetQtyShipped(v int64) {
	o.QtyShipped = &v
}

// GetUnitSalePrice returns the UnitSalePrice field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetUnitSalePrice() RmaMoney {
	if o == nil || IsNil(o.UnitSalePrice) {
		var ret RmaMoney
		return ret
	}
	return *o.UnitSalePrice
}

// GetUnitSalePriceOk returns a tuple with the UnitSalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetUnitSalePriceOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.UnitSalePrice) {
		return nil, false
	}
	return o.UnitSalePrice, true
}

// &#39;Has&#39;UnitSalePrice returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;UnitSalePrice() bool {
	if o != nil && !IsNil(o.UnitSalePrice) {
		return true
	}

	return false
}

// SetUnitSalePrice gets a reference to the given RmaMoney and assigns it to the UnitSalePrice field.
func (o *RmaOrderDataItem) SetUnitSalePrice(v RmaMoney) {
	o.UnitSalePrice = &v
}

// GetUnitListPrice returns the UnitListPrice field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetUnitListPrice() RmaMoney {
	if o == nil || IsNil(o.UnitListPrice) {
		var ret RmaMoney
		return ret
	}
	return *o.UnitListPrice
}

// GetUnitListPriceOk returns a tuple with the UnitListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetUnitListPriceOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.UnitListPrice) {
		return nil, false
	}
	return o.UnitListPrice, true
}

// &#39;Has&#39;UnitListPrice returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;UnitListPrice() bool {
	if o != nil && !IsNil(o.UnitListPrice) {
		return true
	}

	return false
}

// SetUnitListPrice gets a reference to the given RmaMoney and assigns it to the UnitListPrice field.
func (o *RmaOrderDataItem) SetUnitListPrice(v RmaMoney) {
	o.UnitListPrice = &v
}

// GetUnitVatAmount returns the UnitVatAmount field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetUnitVatAmount() RmaMoney {
	if o == nil || IsNil(o.UnitVatAmount) {
		var ret RmaMoney
		return ret
	}
	return *o.UnitVatAmount
}

// GetUnitVatAmountOk returns a tuple with the UnitVatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetUnitVatAmountOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.UnitVatAmount) {
		return nil, false
	}
	return o.UnitVatAmount, true
}

// &#39;Has&#39;UnitVatAmount returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;UnitVatAmount() bool {
	if o != nil && !IsNil(o.UnitVatAmount) {
		return true
	}

	return false
}

// SetUnitVatAmount gets a reference to the given RmaMoney and assigns it to the UnitVatAmount field.
func (o *RmaOrderDataItem) SetUnitVatAmount(v RmaMoney) {
	o.UnitVatAmount = &v
}

// GetRowSalePrice returns the RowSalePrice field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetRowSalePrice() RmaMoney {
	if o == nil || IsNil(o.RowSalePrice) {
		var ret RmaMoney
		return ret
	}
	return *o.RowSalePrice
}

// GetRowSalePriceOk returns a tuple with the RowSalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetRowSalePriceOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.RowSalePrice) {
		return nil, false
	}
	return o.RowSalePrice, true
}

// &#39;Has&#39;RowSalePrice returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;RowSalePrice() bool {
	if o != nil && !IsNil(o.RowSalePrice) {
		return true
	}

	return false
}

// SetRowSalePrice gets a reference to the given RmaMoney and assigns it to the RowSalePrice field.
func (o *RmaOrderDataItem) SetRowSalePrice(v RmaMoney) {
	o.RowSalePrice = &v
}

// GetRowListPrice returns the RowListPrice field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetRowListPrice() RmaMoney {
	if o == nil || IsNil(o.RowListPrice) {
		var ret RmaMoney
		return ret
	}
	return *o.RowListPrice
}

// GetRowListPriceOk returns a tuple with the RowListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetRowListPriceOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.RowListPrice) {
		return nil, false
	}
	return o.RowListPrice, true
}

// &#39;Has&#39;RowListPrice returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;RowListPrice() bool {
	if o != nil && !IsNil(o.RowListPrice) {
		return true
	}

	return false
}

// SetRowListPrice gets a reference to the given RmaMoney and assigns it to the RowListPrice field.
func (o *RmaOrderDataItem) SetRowListPrice(v RmaMoney) {
	o.RowListPrice = &v
}

// GetRowVatAmount returns the RowVatAmount field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetRowVatAmount() RmaMoney {
	if o == nil || IsNil(o.RowVatAmount) {
		var ret RmaMoney
		return ret
	}
	return *o.RowVatAmount
}

// GetRowVatAmountOk returns a tuple with the RowVatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetRowVatAmountOk() (*RmaMoney, bool) {
	if o == nil || IsNil(o.RowVatAmount) {
		return nil, false
	}
	return o.RowVatAmount, true
}

// &#39;Has&#39;RowVatAmount returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;RowVatAmount() bool {
	if o != nil && !IsNil(o.RowVatAmount) {
		return true
	}

	return false
}

// SetRowVatAmount gets a reference to the given RmaMoney and assigns it to the RowVatAmount field.
func (o *RmaOrderDataItem) SetRowVatAmount(v RmaMoney) {
	o.RowVatAmount = &v
}

// GetVatPercentage returns the VatPercentage field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetVatPercentage() float32 {
	if o == nil || IsNil(o.VatPercentage) {
		var ret float32
		return ret
	}
	return *o.VatPercentage
}

// GetVatPercentageOk returns a tuple with the VatPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetVatPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.VatPercentage) {
		return nil, false
	}
	return o.VatPercentage, true
}

// &#39;Has&#39;VatPercentage returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;VatPercentage() bool {
	if o != nil && !IsNil(o.VatPercentage) {
		return true
	}

	return false
}

// SetVatPercentage gets a reference to the given float32 and assigns it to the VatPercentage field.
func (o *RmaOrderDataItem) SetVatPercentage(v float32) {
	o.VatPercentage = &v
}

// GetVatInaccurate returns the VatInaccurate field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetVatInaccurate() bool {
	if o == nil || IsNil(o.VatInaccurate) {
		var ret bool
		return ret
	}
	return *o.VatInaccurate
}

// GetVatInaccurateOk returns a tuple with the VatInaccurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetVatInaccurateOk() (*bool, bool) {
	if o == nil || IsNil(o.VatInaccurate) {
		return nil, false
	}
	return o.VatInaccurate, true
}

// &#39;Has&#39;VatInaccurate returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;VatInaccurate() bool {
	if o != nil && !IsNil(o.VatInaccurate) {
		return true
	}

	return false
}

// SetVatInaccurate gets a reference to the given bool and assigns it to the VatInaccurate field.
func (o *RmaOrderDataItem) SetVatInaccurate(v bool) {
	o.VatInaccurate = &v
}

// GetVatCalculated returns the VatCalculated field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetVatCalculated() bool {
	if o == nil || IsNil(o.VatCalculated) {
		var ret bool
		return ret
	}
	return *o.VatCalculated
}

// GetVatCalculatedOk returns a tuple with the VatCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetVatCalculatedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatCalculated) {
		return nil, false
	}
	return o.VatCalculated, true
}

// &#39;Has&#39;VatCalculated returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;VatCalculated() bool {
	if o != nil && !IsNil(o.VatCalculated) {
		return true
	}

	return false
}

// SetVatCalculated gets a reference to the given bool and assigns it to the VatCalculated field.
func (o *RmaOrderDataItem) SetVatCalculated(v bool) {
	o.VatCalculated = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// &#39;Has&#39;ProductName returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *RmaOrderDataItem) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// &#39;Has&#39;ProductCode returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *RmaOrderDataItem) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductSku returns the ProductSku field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductSku() string {
	if o == nil || IsNil(o.ProductSku) {
		var ret string
		return ret
	}
	return *o.ProductSku
}

// GetProductSkuOk returns a tuple with the ProductSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ProductSku) {
		return nil, false
	}
	return o.ProductSku, true
}

// &#39;Has&#39;ProductSku returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductSku() bool {
	if o != nil && !IsNil(o.ProductSku) {
		return true
	}

	return false
}

// SetProductSku gets a reference to the given string and assigns it to the ProductSku field.
func (o *RmaOrderDataItem) SetProductSku(v string) {
	o.ProductSku = &v
}

// GetProductOptions returns the ProductOptions field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductOptions() string {
	if o == nil || IsNil(o.ProductOptions) {
		var ret string
		return ret
	}
	return *o.ProductOptions
}

// GetProductOptionsOk returns a tuple with the ProductOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.ProductOptions) {
		return nil, false
	}
	return o.ProductOptions, true
}

// &#39;Has&#39;ProductOptions returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductOptions() bool {
	if o != nil && !IsNil(o.ProductOptions) {
		return true
	}

	return false
}

// SetProductOptions gets a reference to the given string and assigns it to the ProductOptions field.
func (o *RmaOrderDataItem) SetProductOptions(v string) {
	o.ProductOptions = &v
}

// GetProductImg returns the ProductImg field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductImg() string {
	if o == nil || IsNil(o.ProductImg) {
		var ret string
		return ret
	}
	return *o.ProductImg
}

// GetProductImgOk returns a tuple with the ProductImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductImgOk() (*string, bool) {
	if o == nil || IsNil(o.ProductImg) {
		return nil, false
	}
	return o.ProductImg, true
}

// &#39;Has&#39;ProductImg returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductImg() bool {
	if o != nil && !IsNil(o.ProductImg) {
		return true
	}

	return false
}

// SetProductImg gets a reference to the given string and assigns it to the ProductImg field.
func (o *RmaOrderDataItem) SetProductImg(v string) {
	o.ProductImg = &v
}

// GetProductData returns the ProductData field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductData() string {
	if o == nil || IsNil(o.ProductData) {
		var ret string
		return ret
	}
	return *o.ProductData
}

// GetProductDataOk returns a tuple with the ProductData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProductData) {
		return nil, false
	}
	return o.ProductData, true
}

// &#39;Has&#39;ProductData returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductData() bool {
	if o != nil && !IsNil(o.ProductData) {
		return true
	}

	return false
}

// SetProductData gets a reference to the given string and assigns it to the ProductData field.
func (o *RmaOrderDataItem) SetProductData(v string) {
	o.ProductData = &v
}

// GetShipmentInfoReference returns the ShipmentInfoReference field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetShipmentInfoReference() string {
	if o == nil || IsNil(o.ShipmentInfoReference) {
		var ret string
		return ret
	}
	return *o.ShipmentInfoReference
}

// GetShipmentInfoReferenceOk returns a tuple with the ShipmentInfoReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetShipmentInfoReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentInfoReference) {
		return nil, false
	}
	return o.ShipmentInfoReference, true
}

// &#39;Has&#39;ShipmentInfoReference returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ShipmentInfoReference() bool {
	if o != nil && !IsNil(o.ShipmentInfoReference) {
		return true
	}

	return false
}

// SetShipmentInfoReference gets a reference to the given string and assigns it to the ShipmentInfoReference field.
func (o *RmaOrderDataItem) SetShipmentInfoReference(v string) {
	o.ShipmentInfoReference = &v
}

// GetPromotionGrn returns the PromotionGrn field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetPromotionGrn() []string {
	if o == nil || IsNil(o.PromotionGrn) {
		var ret []string
		return ret
	}
	return o.PromotionGrn
}

// GetPromotionGrnOk returns a tuple with the PromotionGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetPromotionGrnOk() ([]string, bool) {
	if o == nil || IsNil(o.PromotionGrn) {
		return nil, false
	}
	return o.PromotionGrn, true
}

// &#39;Has&#39;PromotionGrn returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;PromotionGrn() bool {
	if o != nil && !IsNil(o.PromotionGrn) {
		return true
	}

	return false
}

// SetPromotionGrn gets a reference to the given []string and assigns it to the PromotionGrn field.
func (o *RmaOrderDataItem) SetPromotionGrn(v []string) {
	o.PromotionGrn = v
}

// GetProductIsVirtual returns the ProductIsVirtual field value if set, zero value otherwise.
func (o *RmaOrderDataItem) GetProductIsVirtual() bool {
	if o == nil || IsNil(o.ProductIsVirtual) {
		var ret bool
		return ret
	}
	return *o.ProductIsVirtual
}

// GetProductIsVirtualOk returns a tuple with the ProductIsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaOrderDataItem) GetProductIsVirtualOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductIsVirtual) {
		return nil, false
	}
	return o.ProductIsVirtual, true
}

// &#39;Has&#39;ProductIsVirtual returns a boolean if a field has been set.
func (o *RmaOrderDataItem) &#39;Has&#39;ProductIsVirtual() bool {
	if o != nil && !IsNil(o.ProductIsVirtual) {
		return true
	}

	return false
}

// SetProductIsVirtual gets a reference to the given bool and assigns it to the ProductIsVirtual field.
func (o *RmaOrderDataItem) SetProductIsVirtual(v bool) {
	o.ProductIsVirtual = &v
}

func (o RmaOrderDataItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaOrderDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}
	if !IsNil(o.QtyOrdered) {
		toSerialize["qtyOrdered"] = o.QtyOrdered
	}
	if !IsNil(o.QtyCommitted) {
		toSerialize["qtyCommitted"] = o.QtyCommitted
	}
	if !IsNil(o.QtyShipped) {
		toSerialize["qtyShipped"] = o.QtyShipped
	}
	if !IsNil(o.UnitSalePrice) {
		toSerialize["unitSalePrice"] = o.UnitSalePrice
	}
	if !IsNil(o.UnitListPrice) {
		toSerialize["unitListPrice"] = o.UnitListPrice
	}
	if !IsNil(o.UnitVatAmount) {
		toSerialize["unitVatAmount"] = o.UnitVatAmount
	}
	if !IsNil(o.RowSalePrice) {
		toSerialize["rowSalePrice"] = o.RowSalePrice
	}
	if !IsNil(o.RowListPrice) {
		toSerialize["rowListPrice"] = o.RowListPrice
	}
	if !IsNil(o.RowVatAmount) {
		toSerialize["rowVatAmount"] = o.RowVatAmount
	}
	if !IsNil(o.VatPercentage) {
		toSerialize["vatPercentage"] = o.VatPercentage
	}
	if !IsNil(o.VatInaccurate) {
		toSerialize["vatInaccurate"] = o.VatInaccurate
	}
	if !IsNil(o.VatCalculated) {
		toSerialize["vatCalculated"] = o.VatCalculated
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.ProductCode) {
		toSerialize["productCode"] = o.ProductCode
	}
	if !IsNil(o.ProductSku) {
		toSerialize["productSku"] = o.ProductSku
	}
	if !IsNil(o.ProductOptions) {
		toSerialize["productOptions"] = o.ProductOptions
	}
	if !IsNil(o.ProductImg) {
		toSerialize["productImg"] = o.ProductImg
	}
	if !IsNil(o.ProductData) {
		toSerialize["productData"] = o.ProductData
	}
	if !IsNil(o.ShipmentInfoReference) {
		toSerialize["shipmentInfoReference"] = o.ShipmentInfoReference
	}
	if !IsNil(o.PromotionGrn) {
		toSerialize["promotionGrn"] = o.PromotionGrn
	}
	if !IsNil(o.ProductIsVirtual) {
		toSerialize["productIsVirtual"] = o.ProductIsVirtual
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaOrderDataItem) UnmarshalJSON(data []byte) (err error) {
	varRmaOrderDataItem := _RmaOrderDataItem{}

	err = json.Unmarshal(data, &varRmaOrderDataItem)

	if err != nil {
		return err
	}

	*o = RmaOrderDataItem(varRmaOrderDataItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "productGrn")
		delete(additionalProperties, "qtyOrdered")
		delete(additionalProperties, "qtyCommitted")
		delete(additionalProperties, "qtyShipped")
		delete(additionalProperties, "unitSalePrice")
		delete(additionalProperties, "unitListPrice")
		delete(additionalProperties, "unitVatAmount")
		delete(additionalProperties, "rowSalePrice")
		delete(additionalProperties, "rowListPrice")
		delete(additionalProperties, "rowVatAmount")
		delete(additionalProperties, "vatPercentage")
		delete(additionalProperties, "vatInaccurate")
		delete(additionalProperties, "vatCalculated")
		delete(additionalProperties, "productName")
		delete(additionalProperties, "productCode")
		delete(additionalProperties, "productSku")
		delete(additionalProperties, "productOptions")
		delete(additionalProperties, "productImg")
		delete(additionalProperties, "productData")
		delete(additionalProperties, "shipmentInfoReference")
		delete(additionalProperties, "promotionGrn")
		delete(additionalProperties, "productIsVirtual")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaOrderDataItem) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaOrderDataItem) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaOrderDataItem struct {
	value *RmaOrderDataItem
	isSet bool
}

func (v NullableRmaOrderDataItem) Get() *RmaOrderDataItem {
	return v.value
}

func (v *NullableRmaOrderDataItem) Set(val *RmaOrderDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaOrderDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaOrderDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaOrderDataItem(val *RmaOrderDataItem) *NullableRmaOrderDataItem {
	return &NullableRmaOrderDataItem{value: val, isSet: true}
}

func (v NullableRmaOrderDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaOrderDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


