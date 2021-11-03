/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"encoding/json"
)

// PreviousPage struct for PreviousPage
type PreviousPage struct {
	Before string  `json:"before"`
	Link   *string `json:"link,omitempty"`
}

// NewPreviousPage instantiates a new PreviousPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviousPage(before string) *PreviousPage {
	this := PreviousPage{}
	this.Before = before
	return &this
}

// NewPreviousPageWithDefaults instantiates a new PreviousPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviousPageWithDefaults() *PreviousPage {
	this := PreviousPage{}
	return &this
}

// GetBefore returns the Before field value
func (o *PreviousPage) GetBefore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Before
}

// GetBeforeOk returns a tuple with the Before field value
// and a boolean to check if the value has been set.
func (o *PreviousPage) GetBeforeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Before, true
}

// SetBefore sets field value
func (o *PreviousPage) SetBefore(v string) {
	o.Before = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *PreviousPage) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousPage) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *PreviousPage) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *PreviousPage) SetLink(v string) {
	o.Link = &v
}

func (o PreviousPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["before"] = o.Before
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	return json.Marshal(toSerialize)
}

type NullablePreviousPage struct {
	value *PreviousPage
	isSet bool
}

func (v NullablePreviousPage) Get() *PreviousPage {
	return v.value
}

func (v *NullablePreviousPage) Set(val *PreviousPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviousPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviousPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviousPage(val *PreviousPage) *NullablePreviousPage {
	return &NullablePreviousPage{value: val, isSet: true}
}

func (v NullablePreviousPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviousPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
