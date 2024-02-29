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
	"bytes"
	"fmt"
)

// checks if the RmaListReturnsRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListReturnsRequestSort{}

// RmaListReturnsRequestSort struct for RmaListReturnsRequestSort
type RmaListReturnsRequestSort struct {
	// evaluation_order is the order in which the sort will be applied. The lower the number, the earlier the sort will be applied.
	EvaluationOrder int64 `json:"evaluationOrder"`
	Field RmaListReturnsRequestSortSortField `json:"field"`
	Order *RmaSortOrder `json:"order,omitempty"`
}

type _RmaListReturnsRequestSort RmaListReturnsRequestSort

// NewRmaListReturnsRequestSort instantiates a new RmaListReturnsRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListReturnsRequestSort(evaluationOrder int64, field RmaListReturnsRequestSortSortField) *RmaListReturnsRequestSort {
	this := RmaListReturnsRequestSort{}
	this.EvaluationOrder = evaluationOrder
	this.Field = field
	var order RmaSortOrder = DESC
	this.Order = &order
	return &this
}

// NewRmaListReturnsRequestSortWithDefaults instantiates a new RmaListReturnsRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListReturnsRequestSortWithDefaults() *RmaListReturnsRequestSort {
	this := RmaListReturnsRequestSort{}
	var field RmaListReturnsRequestSortSortField = UNKNOWN
	this.Field = field
	var order RmaSortOrder = DESC
	this.Order = &order
	return &this
}

// GetEvaluationOrder returns the EvaluationOrder field value
func (o *RmaListReturnsRequestSort) GetEvaluationOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EvaluationOrder
}

// GetEvaluationOrderOk returns a tuple with the EvaluationOrder field value
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequestSort) GetEvaluationOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrder, true
}

// SetEvaluationOrder sets field value
func (o *RmaListReturnsRequestSort) SetEvaluationOrder(v int64) {
	o.EvaluationOrder = v
}

// GetField returns the Field field value
func (o *RmaListReturnsRequestSort) GetField() RmaListReturnsRequestSortSortField {
	if o == nil {
		var ret RmaListReturnsRequestSortSortField
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequestSort) GetFieldOk() (*RmaListReturnsRequestSortSortField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *RmaListReturnsRequestSort) SetField(v RmaListReturnsRequestSortSortField) {
	o.Field = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RmaListReturnsRequestSort) GetOrder() RmaSortOrder {
	if o == nil || IsNil(o.Order) {
		var ret RmaSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListReturnsRequestSort) GetOrderOk() (*RmaSortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RmaListReturnsRequestSort) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given RmaSortOrder and assigns it to the Order field.
func (o *RmaListReturnsRequestSort) SetOrder(v RmaSortOrder) {
	o.Order = &v
}

func (o RmaListReturnsRequestSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListReturnsRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluationOrder"] = o.EvaluationOrder
	toSerialize["field"] = o.Field
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

func (o *RmaListReturnsRequestSort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"evaluationOrder",
		"field",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRmaListReturnsRequestSort := _RmaListReturnsRequestSort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRmaListReturnsRequestSort)

	if err != nil {
		return err
	}

	*o = RmaListReturnsRequestSort(varRmaListReturnsRequestSort)

	return err
}

type NullableRmaListReturnsRequestSort struct {
	value *RmaListReturnsRequestSort
	isSet bool
}

func (v NullableRmaListReturnsRequestSort) Get() *RmaListReturnsRequestSort {
	return v.value
}

func (v *NullableRmaListReturnsRequestSort) Set(val *RmaListReturnsRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListReturnsRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListReturnsRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListReturnsRequestSort(val *RmaListReturnsRequestSort) *NullableRmaListReturnsRequestSort {
	return &NullableRmaListReturnsRequestSort{value: val, isSet: true}
}

func (v NullableRmaListReturnsRequestSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListReturnsRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


