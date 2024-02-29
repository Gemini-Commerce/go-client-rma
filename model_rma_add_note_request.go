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

// checks if the RmaAddNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaAddNoteRequest{}

// RmaAddNoteRequest struct for RmaAddNoteRequest
type RmaAddNoteRequest struct {
	TenantId string `json:"tenantId"`
	ReturnId string `json:"returnId"`
	Author string `json:"author"`
	Note string `json:"note"`
}

type _RmaAddNoteRequest RmaAddNoteRequest

// NewRmaAddNoteRequest instantiates a new RmaAddNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaAddNoteRequest(tenantId string, returnId string, author string, note string) *RmaAddNoteRequest {
	this := RmaAddNoteRequest{}
	this.TenantId = tenantId
	this.ReturnId = returnId
	this.Author = author
	this.Note = note
	return &this
}

// NewRmaAddNoteRequestWithDefaults instantiates a new RmaAddNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaAddNoteRequestWithDefaults() *RmaAddNoteRequest {
	this := RmaAddNoteRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *RmaAddNoteRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RmaAddNoteRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RmaAddNoteRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetReturnId returns the ReturnId field value
func (o *RmaAddNoteRequest) GetReturnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value
// and a boolean to check if the value has been set.
func (o *RmaAddNoteRequest) GetReturnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnId, true
}

// SetReturnId sets field value
func (o *RmaAddNoteRequest) SetReturnId(v string) {
	o.ReturnId = v
}

// GetAuthor returns the Author field value
func (o *RmaAddNoteRequest) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *RmaAddNoteRequest) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *RmaAddNoteRequest) SetAuthor(v string) {
	o.Author = v
}

// GetNote returns the Note field value
func (o *RmaAddNoteRequest) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *RmaAddNoteRequest) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *RmaAddNoteRequest) SetNote(v string) {
	o.Note = v
}

func (o RmaAddNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaAddNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["returnId"] = o.ReturnId
	toSerialize["author"] = o.Author
	toSerialize["note"] = o.Note
	return toSerialize, nil
}

func (o *RmaAddNoteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"returnId",
		"author",
		"note",
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

	varRmaAddNoteRequest := _RmaAddNoteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRmaAddNoteRequest)

	if err != nil {
		return err
	}

	*o = RmaAddNoteRequest(varRmaAddNoteRequest)

	return err
}

type NullableRmaAddNoteRequest struct {
	value *RmaAddNoteRequest
	isSet bool
}

func (v NullableRmaAddNoteRequest) Get() *RmaAddNoteRequest {
	return v.value
}

func (v *NullableRmaAddNoteRequest) Set(val *RmaAddNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaAddNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaAddNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaAddNoteRequest(val *RmaAddNoteRequest) *NullableRmaAddNoteRequest {
	return &NullableRmaAddNoteRequest{value: val, isSet: true}
}

func (v NullableRmaAddNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaAddNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


