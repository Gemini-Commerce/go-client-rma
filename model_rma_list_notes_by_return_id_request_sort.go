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
	"fmt"
)

// checks if the RmaListNotesByReturnIdRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaListNotesByReturnIdRequestSort{}

// RmaListNotesByReturnIdRequestSort struct for RmaListNotesByReturnIdRequestSort
type RmaListNotesByReturnIdRequestSort struct {
	// evaluation_order is the order in which the sort will be applied. The lower the number, the earlier the sort will be applied.
	EvaluationOrder int64 `json:"evaluationOrder"`
	Field RmaListNotesByReturnIdRequestSortSortField `json:"field"`
	Order *RmaSortOrder `json:"order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaListNotesByReturnIdRequestSort RmaListNotesByReturnIdRequestSort

// NewRmaListNotesByReturnIdRequestSort instantiates a new RmaListNotesByReturnIdRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaListNotesByReturnIdRequestSort(evaluationOrder int64, field RmaListNotesByReturnIdRequestSortSortField) *RmaListNotesByReturnIdRequestSort {
	this := RmaListNotesByReturnIdRequestSort{}
	this.EvaluationOrder = evaluationOrder
	this.Field = field
	var order RmaSortOrder = RMASORTORDER_DESC
	this.Order = &order
	return &this
}

// NewRmaListNotesByReturnIdRequestSortWithDefaults instantiates a new RmaListNotesByReturnIdRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaListNotesByReturnIdRequestSortWithDefaults() *RmaListNotesByReturnIdRequestSort {
	this := RmaListNotesByReturnIdRequestSort{}
	var field RmaListNotesByReturnIdRequestSortSortField = RMALISTNOTESBYRETURNIDREQUESTSORTSORTFIELD_UNKNOWN
	this.Field = field
	var order RmaSortOrder = RMASORTORDER_DESC
	this.Order = &order
	return &this
}

// GetEvaluationOrder returns the EvaluationOrder field value
func (o *RmaListNotesByReturnIdRequestSort) GetEvaluationOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EvaluationOrder
}

// GetEvaluationOrderOk returns a tuple with the EvaluationOrder field value
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequestSort) GetEvaluationOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrder, true
}

// SetEvaluationOrder sets field value
func (o *RmaListNotesByReturnIdRequestSort) SetEvaluationOrder(v int64) {
	o.EvaluationOrder = v
}

// GetField returns the Field field value
func (o *RmaListNotesByReturnIdRequestSort) GetField() RmaListNotesByReturnIdRequestSortSortField {
	if o == nil {
		var ret RmaListNotesByReturnIdRequestSortSortField
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequestSort) GetFieldOk() (*RmaListNotesByReturnIdRequestSortSortField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *RmaListNotesByReturnIdRequestSort) SetField(v RmaListNotesByReturnIdRequestSortSortField) {
	o.Field = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RmaListNotesByReturnIdRequestSort) GetOrder() RmaSortOrder {
	if o == nil || IsNil(o.Order) {
		var ret RmaSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaListNotesByReturnIdRequestSort) GetOrderOk() (*RmaSortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// &#39;Has&#39;Order returns a boolean if a field has been set.
func (o *RmaListNotesByReturnIdRequestSort) &#39;Has&#39;Order() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given RmaSortOrder and assigns it to the Order field.
func (o *RmaListNotesByReturnIdRequestSort) SetOrder(v RmaSortOrder) {
	o.Order = &v
}

func (o RmaListNotesByReturnIdRequestSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaListNotesByReturnIdRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluationOrder"] = o.EvaluationOrder
	toSerialize["field"] = o.Field
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaListNotesByReturnIdRequestSort) UnmarshalJSON(data []byte) (err error) {
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

	varRmaListNotesByReturnIdRequestSort := _RmaListNotesByReturnIdRequestSort{}

	err = json.Unmarshal(data, &varRmaListNotesByReturnIdRequestSort)

	if err != nil {
		return err
	}

	*o = RmaListNotesByReturnIdRequestSort(varRmaListNotesByReturnIdRequestSort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "evaluationOrder")
		delete(additionalProperties, "field")
		delete(additionalProperties, "order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaListNotesByReturnIdRequestSort) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaListNotesByReturnIdRequestSort) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaListNotesByReturnIdRequestSort struct {
	value *RmaListNotesByReturnIdRequestSort
	isSet bool
}

func (v NullableRmaListNotesByReturnIdRequestSort) Get() *RmaListNotesByReturnIdRequestSort {
	return v.value
}

func (v *NullableRmaListNotesByReturnIdRequestSort) Set(val *RmaListNotesByReturnIdRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaListNotesByReturnIdRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaListNotesByReturnIdRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaListNotesByReturnIdRequestSort(val *RmaListNotesByReturnIdRequestSort) *NullableRmaListNotesByReturnIdRequestSort {
	return &NullableRmaListNotesByReturnIdRequestSort{value: val, isSet: true}
}

func (v NullableRmaListNotesByReturnIdRequestSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaListNotesByReturnIdRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


