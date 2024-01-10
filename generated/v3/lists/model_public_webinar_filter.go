/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicWebinarFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicWebinarFilter{}

// PublicWebinarFilter struct for PublicWebinarFilter
type PublicWebinarFilter struct {
	FilterType           string  `json:"filterType"`
	Operator             string  `json:"operator"`
	WebinarId            *string `json:"webinarId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicWebinarFilter PublicWebinarFilter

// NewPublicWebinarFilter instantiates a new PublicWebinarFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicWebinarFilter(filterType string, operator string) *PublicWebinarFilter {
	this := PublicWebinarFilter{}
	this.FilterType = filterType
	this.Operator = operator
	return &this
}

// NewPublicWebinarFilterWithDefaults instantiates a new PublicWebinarFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicWebinarFilterWithDefaults() *PublicWebinarFilter {
	this := PublicWebinarFilter{}
	var filterType string = "WEBINAR"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicWebinarFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicWebinarFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicWebinarFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetOperator returns the Operator field value
func (o *PublicWebinarFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicWebinarFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicWebinarFilter) SetOperator(v string) {
	o.Operator = v
}

// GetWebinarId returns the WebinarId field value if set, zero value otherwise.
func (o *PublicWebinarFilter) GetWebinarId() string {
	if o == nil || IsNil(o.WebinarId) {
		var ret string
		return ret
	}
	return *o.WebinarId
}

// GetWebinarIdOk returns a tuple with the WebinarId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicWebinarFilter) GetWebinarIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebinarId) {
		return nil, false
	}
	return o.WebinarId, true
}

// HasWebinarId returns a boolean if a field has been set.
func (o *PublicWebinarFilter) HasWebinarId() bool {
	if o != nil && !IsNil(o.WebinarId) {
		return true
	}

	return false
}

// SetWebinarId gets a reference to the given string and assigns it to the WebinarId field.
func (o *PublicWebinarFilter) SetWebinarId(v string) {
	o.WebinarId = &v
}

func (o PublicWebinarFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicWebinarFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	toSerialize["operator"] = o.Operator
	if !IsNil(o.WebinarId) {
		toSerialize["webinarId"] = o.WebinarId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicWebinarFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"operator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicWebinarFilter := _PublicWebinarFilter{}

	err = json.Unmarshal(data, &varPublicWebinarFilter)

	if err != nil {
		return err
	}

	*o = PublicWebinarFilter(varPublicWebinarFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "webinarId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicWebinarFilter struct {
	value *PublicWebinarFilter
	isSet bool
}

func (v NullablePublicWebinarFilter) Get() *PublicWebinarFilter {
	return v.value
}

func (v *NullablePublicWebinarFilter) Set(val *PublicWebinarFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicWebinarFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicWebinarFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicWebinarFilter(val *PublicWebinarFilter) *NullablePublicWebinarFilter {
	return &NullablePublicWebinarFilter{value: val, isSet: true}
}

func (v NullablePublicWebinarFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicWebinarFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}