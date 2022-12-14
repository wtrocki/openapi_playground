/*
 * API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contenttype

import (
	"encoding/json"
)

// DataObj struct for DataObj
type DataObj struct {
	Test *string `json:"test,omitempty"`
}

// NewDataObj instantiates a new DataObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataObj() *DataObj {
	this := DataObj{}
	return &this
}

// NewDataObjWithDefaults instantiates a new DataObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataObjWithDefaults() *DataObj {
	this := DataObj{}
	return &this
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *DataObj) GetTest() string {
	if o == nil || o.Test == nil {
		var ret string
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObj) GetTestOk() (*string, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *DataObj) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given string and assigns it to the Test field.
func (o *DataObj) SetTest(v string) {
	o.Test = &v
}

func (o DataObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	return json.Marshal(toSerialize)
}

type NullableDataObj struct {
	value *DataObj
	isSet bool
}

func (v NullableDataObj) Get() *DataObj {
	return v.value
}

func (v *NullableDataObj) Set(val *DataObj) {
	v.value = val
	v.isSet = true
}

func (v NullableDataObj) IsSet() bool {
	return v.isSet
}

func (v *NullableDataObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataObj(val *DataObj) *NullableDataObj {
	return &NullableDataObj{value: val, isSet: true}
}

func (v NullableDataObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


