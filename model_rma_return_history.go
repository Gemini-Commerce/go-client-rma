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

// checks if the RmaReturnHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaReturnHistory{}

// RmaReturnHistory struct for RmaReturnHistory
type RmaReturnHistory struct {
	Date                 *time.Time `json:"date,omitempty"`
	Status               *string    `json:"status,omitempty"`
	Note                 *string    `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaReturnHistory RmaReturnHistory

// NewRmaReturnHistory instantiates a new RmaReturnHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaReturnHistory() *RmaReturnHistory {
	this := RmaReturnHistory{}
	return &this
}

// NewRmaReturnHistoryWithDefaults instantiates a new RmaReturnHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaReturnHistoryWithDefaults() *RmaReturnHistory {
	this := RmaReturnHistory{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RmaReturnHistory) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnHistory) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RmaReturnHistory) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *RmaReturnHistory) SetDate(v time.Time) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RmaReturnHistory) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnHistory) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RmaReturnHistory) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RmaReturnHistory) SetStatus(v string) {
	o.Status = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *RmaReturnHistory) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaReturnHistory) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *RmaReturnHistory) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *RmaReturnHistory) SetNote(v string) {
	o.Note = &v
}

func (o RmaReturnHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaReturnHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaReturnHistory) UnmarshalJSON(data []byte) (err error) {
	varRmaReturnHistory := _RmaReturnHistory{}

	err = json.Unmarshal(data, &varRmaReturnHistory)

	if err != nil {
		return err
	}

	*o = RmaReturnHistory(varRmaReturnHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "status")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaReturnHistory) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaReturnHistory) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaReturnHistory struct {
	value *RmaReturnHistory
	isSet bool
}

func (v NullableRmaReturnHistory) Get() *RmaReturnHistory {
	return v.value
}

func (v *NullableRmaReturnHistory) Set(val *RmaReturnHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaReturnHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaReturnHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaReturnHistory(val *RmaReturnHistory) *NullableRmaReturnHistory {
	return &NullableRmaReturnHistory{value: val, isSet: true}
}

func (v NullableRmaReturnHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaReturnHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
