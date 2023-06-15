/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package products

import (
	"encoding/json"
)

// checks if the ErrorCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorCategory{}

// ErrorCategory struct for ErrorCategory
type ErrorCategory struct {
	Name       string `json:"name"`
	HttpStatus string `json:"httpStatus"`
}

// NewErrorCategory instantiates a new ErrorCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorCategory(name string, httpStatus string) *ErrorCategory {
	this := ErrorCategory{}
	this.Name = name
	this.HttpStatus = httpStatus
	return &this
}

// NewErrorCategoryWithDefaults instantiates a new ErrorCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorCategoryWithDefaults() *ErrorCategory {
	this := ErrorCategory{}
	return &this
}

// GetName returns the Name field value
func (o *ErrorCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ErrorCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ErrorCategory) SetName(v string) {
	o.Name = v
}

// GetHttpStatus returns the HttpStatus field value
func (o *ErrorCategory) GetHttpStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
func (o *ErrorCategory) GetHttpStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ErrorCategory) SetHttpStatus(v string) {
	o.HttpStatus = v
}

func (o ErrorCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["httpStatus"] = o.HttpStatus
	return toSerialize, nil
}

type NullableErrorCategory struct {
	value *ErrorCategory
	isSet bool
}

func (v NullableErrorCategory) Get() *ErrorCategory {
	return v.value
}

func (v *NullableErrorCategory) Set(val *ErrorCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCategory(val *ErrorCategory) *NullableErrorCategory {
	return &NullableErrorCategory{value: val, isSet: true}
}

func (v NullableErrorCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
