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

// checks if the RmaReturnProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaReturnProduct{}

// RmaReturnProduct struct for RmaReturnProduct
type RmaReturnProduct struct {
	Grn *string `json:"grn,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Requested *RmaReturnProductProperty `json:"requested,omitempty"`
	Approved *RmaReturnProductProperty `json:"approved,omitempty"`
	Verified *RmaReturnProductProperty `json:"verified,omitempty"`
	Refunded *RmaReturnProductProperty `json:"refunded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaReturnProduct RmaReturnProduct

// NewRmaReturnProduct instantiates a new RmaReturnProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaReturnProduct() *RmaReturnProduct {
	this := RmaReturnProduct{}
	return &this
}

// NewRmaReturnProductWithDefaults instantiates a new RmaReturnProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaReturnProductWithDefaults() *RmaReturnProduct {
	this := RmaReturnProduct{}
	return &this
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// &#39;Has&#39;Grn returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Grn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *RmaReturnProduct) SetGrn(v string) {
	o.Grn = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// &#39;Has&#39;Reason returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Reason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RmaReturnProduct) SetReason(v string) {
	o.Reason = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetRequested() RmaReturnProductProperty {
	if o == nil || IsNil(o.Requested) {
		var ret RmaReturnProductProperty
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetRequestedOk() (*RmaReturnProductProperty, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// &#39;Has&#39;Requested returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Requested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given RmaReturnProductProperty and assigns it to the Requested field.
func (o *RmaReturnProduct) SetRequested(v RmaReturnProductProperty) {
	o.Requested = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetApproved() RmaReturnProductProperty {
	if o == nil || IsNil(o.Approved) {
		var ret RmaReturnProductProperty
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetApprovedOk() (*RmaReturnProductProperty, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// &#39;Has&#39;Approved returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Approved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given RmaReturnProductProperty and assigns it to the Approved field.
func (o *RmaReturnProduct) SetApproved(v RmaReturnProductProperty) {
	o.Approved = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetVerified() RmaReturnProductProperty {
	if o == nil || IsNil(o.Verified) {
		var ret RmaReturnProductProperty
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetVerifiedOk() (*RmaReturnProductProperty, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// &#39;Has&#39;Verified returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Verified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given RmaReturnProductProperty and assigns it to the Verified field.
func (o *RmaReturnProduct) SetVerified(v RmaReturnProductProperty) {
	o.Verified = &v
}

// GetRefunded returns the Refunded field value if set, zero value otherwise.
func (o *RmaReturnProduct) GetRefunded() RmaReturnProductProperty {
	if o == nil || IsNil(o.Refunded) {
		var ret RmaReturnProductProperty
		return ret
	}
	return *o.Refunded
}

// GetRefundedOk returns a tuple with the Refunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnProduct) GetRefundedOk() (*RmaReturnProductProperty, bool) {
	if o == nil || IsNil(o.Refunded) {
		return nil, false
	}
	return o.Refunded, true
}

// &#39;Has&#39;Refunded returns a boolean if a field has been set.
func (o *RmaReturnProduct) &#39;Has&#39;Refunded() bool {
	if o != nil && !IsNil(o.Refunded) {
		return true
	}

	return false
}

// SetRefunded gets a reference to the given RmaReturnProductProperty and assigns it to the Refunded field.
func (o *RmaReturnProduct) SetRefunded(v RmaReturnProductProperty) {
	o.Refunded = &v
}

func (o RmaReturnProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaReturnProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.Refunded) {
		toSerialize["refunded"] = o.Refunded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaReturnProduct) UnmarshalJSON(data []byte) (err error) {
	varRmaReturnProduct := _RmaReturnProduct{}

	err = json.Unmarshal(data, &varRmaReturnProduct)

	if err != nil {
		return err
	}

	*o = RmaReturnProduct(varRmaReturnProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grn")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "requested")
		delete(additionalProperties, "approved")
		delete(additionalProperties, "verified")
		delete(additionalProperties, "refunded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaReturnProduct) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaReturnProduct) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaReturnProduct struct {
	value *RmaReturnProduct
	isSet bool
}

func (v NullableRmaReturnProduct) Get() *RmaReturnProduct {
	return v.value
}

func (v *NullableRmaReturnProduct) Set(val *RmaReturnProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaReturnProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaReturnProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaReturnProduct(val *RmaReturnProduct) *NullableRmaReturnProduct {
	return &NullableRmaReturnProduct{value: val, isSet: true}
}

func (v NullableRmaReturnProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaReturnProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


