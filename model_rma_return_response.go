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
	"time"
)

// checks if the RmaReturnResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaReturnResponse{}

// RmaReturnResponse struct for RmaReturnResponse
type RmaReturnResponse struct {
	Id                    *string            `json:"id,omitempty"`
	Grn                   *string            `json:"grn,omitempty"`
	OrderGrn              *string            `json:"orderGrn,omitempty"`
	Status                *string            `json:"status,omitempty"`
	Products              []RmaReturnProduct `json:"products,omitempty"`
	PreferredRefundMethod *RmaRefundMethod   `json:"preferredRefundMethod,omitempty"`
	RefundShippingCost    *bool              `json:"refundShippingCost,omitempty"`
	RefundPaymentCost     *bool              `json:"refundPaymentCost,omitempty"`
	CustomerInfo          *RmaCustomerInfo   `json:"customerInfo,omitempty"`
	ReturnAddress         *RmaPostalAddress  `json:"returnAddress,omitempty"`
	Note                  *string            `json:"note,omitempty"`
	History               []RmaReturnHistory `json:"history,omitempty"`
	CreatedAt             *time.Time         `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time         `json:"updatedAt,omitempty"`
	OrderData             *RmaOrderData      `json:"orderData,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _RmaReturnResponse RmaReturnResponse

// NewRmaReturnResponse instantiates a new RmaReturnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaReturnResponse() *RmaReturnResponse {
	this := RmaReturnResponse{}
	var preferredRefundMethod RmaRefundMethod = RMAREFUNDMETHOD_UNKNOWN
	this.PreferredRefundMethod = &preferredRefundMethod
	return &this
}

// NewRmaReturnResponseWithDefaults instantiates a new RmaReturnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaReturnResponseWithDefaults() *RmaReturnResponse {
	this := RmaReturnResponse{}
	var preferredRefundMethod RmaRefundMethod = RMAREFUNDMETHOD_UNKNOWN
	this.PreferredRefundMethod = &preferredRefundMethod
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RmaReturnResponse) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *RmaReturnResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetOrderGrn returns the OrderGrn field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetOrderGrn() string {
	if o == nil || IsNil(o.OrderGrn) {
		var ret string
		return ret
	}
	return *o.OrderGrn
}

// GetOrderGrnOk returns a tuple with the OrderGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetOrderGrnOk() (*string, bool) {
	if o == nil || IsNil(o.OrderGrn) {
		return nil, false
	}
	return o.OrderGrn, true
}

// HasOrderGrn returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasOrderGrn() bool {
	if o != nil && !IsNil(o.OrderGrn) {
		return true
	}

	return false
}

// SetOrderGrn gets a reference to the given string and assigns it to the OrderGrn field.
func (o *RmaReturnResponse) SetOrderGrn(v string) {
	o.OrderGrn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RmaReturnResponse) SetStatus(v string) {
	o.Status = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetProducts() []RmaReturnProduct {
	if o == nil || IsNil(o.Products) {
		var ret []RmaReturnProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetProductsOk() ([]RmaReturnProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []RmaReturnProduct and assigns it to the Products field.
func (o *RmaReturnResponse) SetProducts(v []RmaReturnProduct) {
	o.Products = v
}

// GetPreferredRefundMethod returns the PreferredRefundMethod field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetPreferredRefundMethod() RmaRefundMethod {
	if o == nil || IsNil(o.PreferredRefundMethod) {
		var ret RmaRefundMethod
		return ret
	}
	return *o.PreferredRefundMethod
}

// GetPreferredRefundMethodOk returns a tuple with the PreferredRefundMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetPreferredRefundMethodOk() (*RmaRefundMethod, bool) {
	if o == nil || IsNil(o.PreferredRefundMethod) {
		return nil, false
	}
	return o.PreferredRefundMethod, true
}

// HasPreferredRefundMethod returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasPreferredRefundMethod() bool {
	if o != nil && !IsNil(o.PreferredRefundMethod) {
		return true
	}

	return false
}

// SetPreferredRefundMethod gets a reference to the given RmaRefundMethod and assigns it to the PreferredRefundMethod field.
func (o *RmaReturnResponse) SetPreferredRefundMethod(v RmaRefundMethod) {
	o.PreferredRefundMethod = &v
}

// GetRefundShippingCost returns the RefundShippingCost field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetRefundShippingCost() bool {
	if o == nil || IsNil(o.RefundShippingCost) {
		var ret bool
		return ret
	}
	return *o.RefundShippingCost
}

// GetRefundShippingCostOk returns a tuple with the RefundShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetRefundShippingCostOk() (*bool, bool) {
	if o == nil || IsNil(o.RefundShippingCost) {
		return nil, false
	}
	return o.RefundShippingCost, true
}

// HasRefundShippingCost returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasRefundShippingCost() bool {
	if o != nil && !IsNil(o.RefundShippingCost) {
		return true
	}

	return false
}

// SetRefundShippingCost gets a reference to the given bool and assigns it to the RefundShippingCost field.
func (o *RmaReturnResponse) SetRefundShippingCost(v bool) {
	o.RefundShippingCost = &v
}

// GetRefundPaymentCost returns the RefundPaymentCost field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetRefundPaymentCost() bool {
	if o == nil || IsNil(o.RefundPaymentCost) {
		var ret bool
		return ret
	}
	return *o.RefundPaymentCost
}

// GetRefundPaymentCostOk returns a tuple with the RefundPaymentCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetRefundPaymentCostOk() (*bool, bool) {
	if o == nil || IsNil(o.RefundPaymentCost) {
		return nil, false
	}
	return o.RefundPaymentCost, true
}

// HasRefundPaymentCost returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasRefundPaymentCost() bool {
	if o != nil && !IsNil(o.RefundPaymentCost) {
		return true
	}

	return false
}

// SetRefundPaymentCost gets a reference to the given bool and assigns it to the RefundPaymentCost field.
func (o *RmaReturnResponse) SetRefundPaymentCost(v bool) {
	o.RefundPaymentCost = &v
}

// GetCustomerInfo returns the CustomerInfo field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetCustomerInfo() RmaCustomerInfo {
	if o == nil || IsNil(o.CustomerInfo) {
		var ret RmaCustomerInfo
		return ret
	}
	return *o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetCustomerInfoOk() (*RmaCustomerInfo, bool) {
	if o == nil || IsNil(o.CustomerInfo) {
		return nil, false
	}
	return o.CustomerInfo, true
}

// HasCustomerInfo returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasCustomerInfo() bool {
	if o != nil && !IsNil(o.CustomerInfo) {
		return true
	}

	return false
}

// SetCustomerInfo gets a reference to the given RmaCustomerInfo and assigns it to the CustomerInfo field.
func (o *RmaReturnResponse) SetCustomerInfo(v RmaCustomerInfo) {
	o.CustomerInfo = &v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetReturnAddress() RmaPostalAddress {
	if o == nil || IsNil(o.ReturnAddress) {
		var ret RmaPostalAddress
		return ret
	}
	return *o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetReturnAddressOk() (*RmaPostalAddress, bool) {
	if o == nil || IsNil(o.ReturnAddress) {
		return nil, false
	}
	return o.ReturnAddress, true
}

// HasReturnAddress returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasReturnAddress() bool {
	if o != nil && !IsNil(o.ReturnAddress) {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given RmaPostalAddress and assigns it to the ReturnAddress field.
func (o *RmaReturnResponse) SetReturnAddress(v RmaPostalAddress) {
	o.ReturnAddress = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *RmaReturnResponse) SetNote(v string) {
	o.Note = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetHistory() []RmaReturnHistory {
	if o == nil || IsNil(o.History) {
		var ret []RmaReturnHistory
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetHistoryOk() ([]RmaReturnHistory, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []RmaReturnHistory and assigns it to the History field.
func (o *RmaReturnResponse) SetHistory(v []RmaReturnHistory) {
	o.History = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RmaReturnResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RmaReturnResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrderData returns the OrderData field value if set, zero value otherwise.
func (o *RmaReturnResponse) GetOrderData() RmaOrderData {
	if o == nil || IsNil(o.OrderData) {
		var ret RmaOrderData
		return ret
	}
	return *o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnResponse) GetOrderDataOk() (*RmaOrderData, bool) {
	if o == nil || IsNil(o.OrderData) {
		return nil, false
	}
	return o.OrderData, true
}

// HasOrderData returns a boolean if a field has been set.
func (o *RmaReturnResponse) HasOrderData() bool {
	if o != nil && !IsNil(o.OrderData) {
		return true
	}

	return false
}

// SetOrderData gets a reference to the given RmaOrderData and assigns it to the OrderData field.
func (o *RmaReturnResponse) SetOrderData(v RmaOrderData) {
	o.OrderData = &v
}

func (o RmaReturnResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaReturnResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.OrderGrn) {
		toSerialize["orderGrn"] = o.OrderGrn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.PreferredRefundMethod) {
		toSerialize["preferredRefundMethod"] = o.PreferredRefundMethod
	}
	if !IsNil(o.RefundShippingCost) {
		toSerialize["refundShippingCost"] = o.RefundShippingCost
	}
	if !IsNil(o.RefundPaymentCost) {
		toSerialize["refundPaymentCost"] = o.RefundPaymentCost
	}
	if !IsNil(o.CustomerInfo) {
		toSerialize["customerInfo"] = o.CustomerInfo
	}
	if !IsNil(o.ReturnAddress) {
		toSerialize["returnAddress"] = o.ReturnAddress
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.OrderData) {
		toSerialize["orderData"] = o.OrderData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaReturnResponse) UnmarshalJSON(data []byte) (err error) {
	varRmaReturnResponse := _RmaReturnResponse{}

	err = json.Unmarshal(data, &varRmaReturnResponse)

	if err != nil {
		return err
	}

	*o = RmaReturnResponse(varRmaReturnResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "orderGrn")
		delete(additionalProperties, "status")
		delete(additionalProperties, "products")
		delete(additionalProperties, "preferredRefundMethod")
		delete(additionalProperties, "refundShippingCost")
		delete(additionalProperties, "refundPaymentCost")
		delete(additionalProperties, "customerInfo")
		delete(additionalProperties, "returnAddress")
		delete(additionalProperties, "note")
		delete(additionalProperties, "history")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "orderData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaReturnResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaReturnResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaReturnResponse struct {
	value *RmaReturnResponse
	isSet bool
}

func (v NullableRmaReturnResponse) Get() *RmaReturnResponse {
	return v.value
}

func (v *NullableRmaReturnResponse) Set(val *RmaReturnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaReturnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaReturnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaReturnResponse(val *RmaReturnResponse) *NullableRmaReturnResponse {
	return &NullableRmaReturnResponse{value: val, isSet: true}
}

func (v NullableRmaReturnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaReturnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
