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

// checks if the RmaCustomerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaCustomerInfo{}

// RmaCustomerInfo struct for RmaCustomerInfo
type RmaCustomerInfo struct {
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Grn *string `json:"grn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaCustomerInfo RmaCustomerInfo

// NewRmaCustomerInfo instantiates a new RmaCustomerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaCustomerInfo() *RmaCustomerInfo {
	this := RmaCustomerInfo{}
	return &this
}

// NewRmaCustomerInfoWithDefaults instantiates a new RmaCustomerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaCustomerInfoWithDefaults() *RmaCustomerInfo {
	this := RmaCustomerInfo{}
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *RmaCustomerInfo) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaCustomerInfo) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// &#39;Has&#39;Firstname returns a boolean if a field has been set.
func (o *RmaCustomerInfo) &#39;Has&#39;Firstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *RmaCustomerInfo) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *RmaCustomerInfo) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaCustomerInfo) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// &#39;Has&#39;Lastname returns a boolean if a field has been set.
func (o *RmaCustomerInfo) &#39;Has&#39;Lastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *RmaCustomerInfo) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RmaCustomerInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaCustomerInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// &#39;Has&#39;Email returns a boolean if a field has been set.
func (o *RmaCustomerInfo) &#39;Has&#39;Email() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RmaCustomerInfo) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *RmaCustomerInfo) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaCustomerInfo) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// &#39;Has&#39;Phone returns a boolean if a field has been set.
func (o *RmaCustomerInfo) &#39;Has&#39;Phone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *RmaCustomerInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *RmaCustomerInfo) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaCustomerInfo) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// &#39;Has&#39;Grn returns a boolean if a field has been set.
func (o *RmaCustomerInfo) &#39;Has&#39;Grn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *RmaCustomerInfo) SetGrn(v string) {
	o.Grn = &v
}

func (o RmaCustomerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaCustomerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaCustomerInfo) UnmarshalJSON(data []byte) (err error) {
	varRmaCustomerInfo := _RmaCustomerInfo{}

	err = json.Unmarshal(data, &varRmaCustomerInfo)

	if err != nil {
		return err
	}

	*o = RmaCustomerInfo(varRmaCustomerInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "grn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaCustomerInfo) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *RmaCustomerInfo) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableRmaCustomerInfo struct {
	value *RmaCustomerInfo
	isSet bool
}

func (v NullableRmaCustomerInfo) Get() *RmaCustomerInfo {
	return v.value
}

func (v *NullableRmaCustomerInfo) Set(val *RmaCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaCustomerInfo(val *RmaCustomerInfo) *NullableRmaCustomerInfo {
	return &NullableRmaCustomerInfo{value: val, isSet: true}
}

func (v NullableRmaCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


