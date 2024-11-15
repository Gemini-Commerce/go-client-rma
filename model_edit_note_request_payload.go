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

// checks if the EditNoteRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditNoteRequestPayload{}

// EditNoteRequestPayload struct for EditNoteRequestPayload
type EditNoteRequestPayload struct {
	Author *string `json:"author,omitempty"`
	Note *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EditNoteRequestPayload EditNoteRequestPayload

// NewEditNoteRequestPayload instantiates a new EditNoteRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditNoteRequestPayload() *EditNoteRequestPayload {
	this := EditNoteRequestPayload{}
	return &this
}

// NewEditNoteRequestPayloadWithDefaults instantiates a new EditNoteRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditNoteRequestPayloadWithDefaults() *EditNoteRequestPayload {
	this := EditNoteRequestPayload{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *EditNoteRequestPayload) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditNoteRequestPayload) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *EditNoteRequestPayload) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *EditNoteRequestPayload) SetAuthor(v string) {
	o.Author = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *EditNoteRequestPayload) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditNoteRequestPayload) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *EditNoteRequestPayload) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *EditNoteRequestPayload) SetNote(v string) {
	o.Note = &v
}

func (o EditNoteRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditNoteRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EditNoteRequestPayload) UnmarshalJSON(data []byte) (err error) {
	varEditNoteRequestPayload := _EditNoteRequestPayload{}

	err = json.Unmarshal(data, &varEditNoteRequestPayload)

	if err != nil {
		return err
	}

	*o = EditNoteRequestPayload(varEditNoteRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "author")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EditNoteRequestPayload) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *EditNoteRequestPayload) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableEditNoteRequestPayload struct {
	value *EditNoteRequestPayload
	isSet bool
}

func (v NullableEditNoteRequestPayload) Get() *EditNoteRequestPayload {
	return v.value
}

func (v *NullableEditNoteRequestPayload) Set(val *EditNoteRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEditNoteRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEditNoteRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditNoteRequestPayload(val *EditNoteRequestPayload) *NullableEditNoteRequestPayload {
	return &NullableEditNoteRequestPayload{value: val, isSet: true}
}

func (v NullableEditNoteRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditNoteRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


