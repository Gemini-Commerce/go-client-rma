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

// checks if the ListReturnsRequestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReturnsRequestFilter{}

// ListReturnsRequestFilter struct for ListReturnsRequestFilter
type ListReturnsRequestFilter struct {
	SearchTerms []string `json:"searchTerms,omitempty"`
	ReturnIds []string `json:"returnIds,omitempty"`
	ReturnStatus []string `json:"returnStatus,omitempty"`
	OrderGrns []string `json:"orderGrns,omitempty"`
	OrderNumbers []string `json:"orderNumbers,omitempty"`
	OrderItemGrns []string `json:"orderItemGrns,omitempty"`
	OrderItemSkus []string `json:"orderItemSkus,omitempty"`
	CustomerGrns []string `json:"customerGrns,omitempty"`
	CustomerEmails []string `json:"customerEmails,omitempty"`
	CreatedAtStart *time.Time `json:"createdAtStart,omitempty"`
	CreatedAtEnd *time.Time `json:"createdAtEnd,omitempty"`
}

// NewListReturnsRequestFilter instantiates a new ListReturnsRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReturnsRequestFilter() *ListReturnsRequestFilter {
	this := ListReturnsRequestFilter{}
	return &this
}

// NewListReturnsRequestFilterWithDefaults instantiates a new ListReturnsRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReturnsRequestFilterWithDefaults() *ListReturnsRequestFilter {
	this := ListReturnsRequestFilter{}
	return &this
}

// GetSearchTerms returns the SearchTerms field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetSearchTerms() []string {
	if o == nil || IsNil(o.SearchTerms) {
		var ret []string
		return ret
	}
	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetSearchTermsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchTerms) {
		return nil, false
	}
	return o.SearchTerms, true
}

// HasSearchTerms returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasSearchTerms() bool {
	if o != nil && !IsNil(o.SearchTerms) {
		return true
	}

	return false
}

// SetSearchTerms gets a reference to the given []string and assigns it to the SearchTerms field.
func (o *ListReturnsRequestFilter) SetSearchTerms(v []string) {
	o.SearchTerms = v
}

// GetReturnIds returns the ReturnIds field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetReturnIds() []string {
	if o == nil || IsNil(o.ReturnIds) {
		var ret []string
		return ret
	}
	return o.ReturnIds
}

// GetReturnIdsOk returns a tuple with the ReturnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetReturnIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReturnIds) {
		return nil, false
	}
	return o.ReturnIds, true
}

// HasReturnIds returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasReturnIds() bool {
	if o != nil && !IsNil(o.ReturnIds) {
		return true
	}

	return false
}

// SetReturnIds gets a reference to the given []string and assigns it to the ReturnIds field.
func (o *ListReturnsRequestFilter) SetReturnIds(v []string) {
	o.ReturnIds = v
}

// GetReturnStatus returns the ReturnStatus field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetReturnStatus() []string {
	if o == nil || IsNil(o.ReturnStatus) {
		var ret []string
		return ret
	}
	return o.ReturnStatus
}

// GetReturnStatusOk returns a tuple with the ReturnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetReturnStatusOk() ([]string, bool) {
	if o == nil || IsNil(o.ReturnStatus) {
		return nil, false
	}
	return o.ReturnStatus, true
}

// HasReturnStatus returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasReturnStatus() bool {
	if o != nil && !IsNil(o.ReturnStatus) {
		return true
	}

	return false
}

// SetReturnStatus gets a reference to the given []string and assigns it to the ReturnStatus field.
func (o *ListReturnsRequestFilter) SetReturnStatus(v []string) {
	o.ReturnStatus = v
}

// GetOrderGrns returns the OrderGrns field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetOrderGrns() []string {
	if o == nil || IsNil(o.OrderGrns) {
		var ret []string
		return ret
	}
	return o.OrderGrns
}

// GetOrderGrnsOk returns a tuple with the OrderGrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetOrderGrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderGrns) {
		return nil, false
	}
	return o.OrderGrns, true
}

// HasOrderGrns returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasOrderGrns() bool {
	if o != nil && !IsNil(o.OrderGrns) {
		return true
	}

	return false
}

// SetOrderGrns gets a reference to the given []string and assigns it to the OrderGrns field.
func (o *ListReturnsRequestFilter) SetOrderGrns(v []string) {
	o.OrderGrns = v
}

// GetOrderNumbers returns the OrderNumbers field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetOrderNumbers() []string {
	if o == nil || IsNil(o.OrderNumbers) {
		var ret []string
		return ret
	}
	return o.OrderNumbers
}

// GetOrderNumbersOk returns a tuple with the OrderNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetOrderNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderNumbers) {
		return nil, false
	}
	return o.OrderNumbers, true
}

// HasOrderNumbers returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasOrderNumbers() bool {
	if o != nil && !IsNil(o.OrderNumbers) {
		return true
	}

	return false
}

// SetOrderNumbers gets a reference to the given []string and assigns it to the OrderNumbers field.
func (o *ListReturnsRequestFilter) SetOrderNumbers(v []string) {
	o.OrderNumbers = v
}

// GetOrderItemGrns returns the OrderItemGrns field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetOrderItemGrns() []string {
	if o == nil || IsNil(o.OrderItemGrns) {
		var ret []string
		return ret
	}
	return o.OrderItemGrns
}

// GetOrderItemGrnsOk returns a tuple with the OrderItemGrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetOrderItemGrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderItemGrns) {
		return nil, false
	}
	return o.OrderItemGrns, true
}

// HasOrderItemGrns returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasOrderItemGrns() bool {
	if o != nil && !IsNil(o.OrderItemGrns) {
		return true
	}

	return false
}

// SetOrderItemGrns gets a reference to the given []string and assigns it to the OrderItemGrns field.
func (o *ListReturnsRequestFilter) SetOrderItemGrns(v []string) {
	o.OrderItemGrns = v
}

// GetOrderItemSkus returns the OrderItemSkus field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetOrderItemSkus() []string {
	if o == nil || IsNil(o.OrderItemSkus) {
		var ret []string
		return ret
	}
	return o.OrderItemSkus
}

// GetOrderItemSkusOk returns a tuple with the OrderItemSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetOrderItemSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderItemSkus) {
		return nil, false
	}
	return o.OrderItemSkus, true
}

// HasOrderItemSkus returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasOrderItemSkus() bool {
	if o != nil && !IsNil(o.OrderItemSkus) {
		return true
	}

	return false
}

// SetOrderItemSkus gets a reference to the given []string and assigns it to the OrderItemSkus field.
func (o *ListReturnsRequestFilter) SetOrderItemSkus(v []string) {
	o.OrderItemSkus = v
}

// GetCustomerGrns returns the CustomerGrns field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetCustomerGrns() []string {
	if o == nil || IsNil(o.CustomerGrns) {
		var ret []string
		return ret
	}
	return o.CustomerGrns
}

// GetCustomerGrnsOk returns a tuple with the CustomerGrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetCustomerGrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerGrns) {
		return nil, false
	}
	return o.CustomerGrns, true
}

// HasCustomerGrns returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasCustomerGrns() bool {
	if o != nil && !IsNil(o.CustomerGrns) {
		return true
	}

	return false
}

// SetCustomerGrns gets a reference to the given []string and assigns it to the CustomerGrns field.
func (o *ListReturnsRequestFilter) SetCustomerGrns(v []string) {
	o.CustomerGrns = v
}

// GetCustomerEmails returns the CustomerEmails field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetCustomerEmails() []string {
	if o == nil || IsNil(o.CustomerEmails) {
		var ret []string
		return ret
	}
	return o.CustomerEmails
}

// GetCustomerEmailsOk returns a tuple with the CustomerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetCustomerEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerEmails) {
		return nil, false
	}
	return o.CustomerEmails, true
}

// HasCustomerEmails returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasCustomerEmails() bool {
	if o != nil && !IsNil(o.CustomerEmails) {
		return true
	}

	return false
}

// SetCustomerEmails gets a reference to the given []string and assigns it to the CustomerEmails field.
func (o *ListReturnsRequestFilter) SetCustomerEmails(v []string) {
	o.CustomerEmails = v
}

// GetCreatedAtStart returns the CreatedAtStart field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetCreatedAtStart() time.Time {
	if o == nil || IsNil(o.CreatedAtStart) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtStart
}

// GetCreatedAtStartOk returns a tuple with the CreatedAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetCreatedAtStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtStart) {
		return nil, false
	}
	return o.CreatedAtStart, true
}

// HasCreatedAtStart returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasCreatedAtStart() bool {
	if o != nil && !IsNil(o.CreatedAtStart) {
		return true
	}

	return false
}

// SetCreatedAtStart gets a reference to the given time.Time and assigns it to the CreatedAtStart field.
func (o *ListReturnsRequestFilter) SetCreatedAtStart(v time.Time) {
	o.CreatedAtStart = &v
}

// GetCreatedAtEnd returns the CreatedAtEnd field value if set, zero value otherwise.
func (o *ListReturnsRequestFilter) GetCreatedAtEnd() time.Time {
	if o == nil || IsNil(o.CreatedAtEnd) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtEnd
}

// GetCreatedAtEndOk returns a tuple with the CreatedAtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnsRequestFilter) GetCreatedAtEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtEnd) {
		return nil, false
	}
	return o.CreatedAtEnd, true
}

// HasCreatedAtEnd returns a boolean if a field has been set.
func (o *ListReturnsRequestFilter) HasCreatedAtEnd() bool {
	if o != nil && !IsNil(o.CreatedAtEnd) {
		return true
	}

	return false
}

// SetCreatedAtEnd gets a reference to the given time.Time and assigns it to the CreatedAtEnd field.
func (o *ListReturnsRequestFilter) SetCreatedAtEnd(v time.Time) {
	o.CreatedAtEnd = &v
}

func (o ListReturnsRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReturnsRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchTerms) {
		toSerialize["searchTerms"] = o.SearchTerms
	}
	if !IsNil(o.ReturnIds) {
		toSerialize["returnIds"] = o.ReturnIds
	}
	if !IsNil(o.ReturnStatus) {
		toSerialize["returnStatus"] = o.ReturnStatus
	}
	if !IsNil(o.OrderGrns) {
		toSerialize["orderGrns"] = o.OrderGrns
	}
	if !IsNil(o.OrderNumbers) {
		toSerialize["orderNumbers"] = o.OrderNumbers
	}
	if !IsNil(o.OrderItemGrns) {
		toSerialize["orderItemGrns"] = o.OrderItemGrns
	}
	if !IsNil(o.OrderItemSkus) {
		toSerialize["orderItemSkus"] = o.OrderItemSkus
	}
	if !IsNil(o.CustomerGrns) {
		toSerialize["customerGrns"] = o.CustomerGrns
	}
	if !IsNil(o.CustomerEmails) {
		toSerialize["customerEmails"] = o.CustomerEmails
	}
	if !IsNil(o.CreatedAtStart) {
		toSerialize["createdAtStart"] = o.CreatedAtStart
	}
	if !IsNil(o.CreatedAtEnd) {
		toSerialize["createdAtEnd"] = o.CreatedAtEnd
	}
	return toSerialize, nil
}

type NullableListReturnsRequestFilter struct {
	value *ListReturnsRequestFilter
	isSet bool
}

func (v NullableListReturnsRequestFilter) Get() *ListReturnsRequestFilter {
	return v.value
}

func (v *NullableListReturnsRequestFilter) Set(val *ListReturnsRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableListReturnsRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableListReturnsRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReturnsRequestFilter(val *ListReturnsRequestFilter) *NullableListReturnsRequestFilter {
	return &NullableListReturnsRequestFilter{value: val, isSet: true}
}

func (v NullableListReturnsRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReturnsRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


