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

// checks if the PublicTimeOffset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicTimeOffset{}

// PublicTimeOffset struct for PublicTimeOffset
type PublicTimeOffset struct {
	OffsetDirection      string `json:"offsetDirection"`
	TimeUnit             string `json:"timeUnit"`
	Amount               int64  `json:"amount"`
	AdditionalProperties map[string]interface{}
}

type _PublicTimeOffset PublicTimeOffset

// NewPublicTimeOffset instantiates a new PublicTimeOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicTimeOffset(offsetDirection string, timeUnit string, amount int64) *PublicTimeOffset {
	this := PublicTimeOffset{}
	this.OffsetDirection = offsetDirection
	this.TimeUnit = timeUnit
	this.Amount = amount
	return &this
}

// NewPublicTimeOffsetWithDefaults instantiates a new PublicTimeOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicTimeOffsetWithDefaults() *PublicTimeOffset {
	this := PublicTimeOffset{}
	return &this
}

// GetOffsetDirection returns the OffsetDirection field value
func (o *PublicTimeOffset) GetOffsetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OffsetDirection
}

// GetOffsetDirectionOk returns a tuple with the OffsetDirection field value
// and a boolean to check if the value has been set.
func (o *PublicTimeOffset) GetOffsetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OffsetDirection, true
}

// SetOffsetDirection sets field value
func (o *PublicTimeOffset) SetOffsetDirection(v string) {
	o.OffsetDirection = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *PublicTimeOffset) GetTimeUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *PublicTimeOffset) GetTimeUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *PublicTimeOffset) SetTimeUnit(v string) {
	o.TimeUnit = v
}

// GetAmount returns the Amount field value
func (o *PublicTimeOffset) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PublicTimeOffset) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PublicTimeOffset) SetAmount(v int64) {
	o.Amount = v
}

func (o PublicTimeOffset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicTimeOffset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offsetDirection"] = o.OffsetDirection
	toSerialize["timeUnit"] = o.TimeUnit
	toSerialize["amount"] = o.Amount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicTimeOffset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offsetDirection",
		"timeUnit",
		"amount",
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

	varPublicTimeOffset := _PublicTimeOffset{}

	err = json.Unmarshal(data, &varPublicTimeOffset)

	if err != nil {
		return err
	}

	*o = PublicTimeOffset(varPublicTimeOffset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "offsetDirection")
		delete(additionalProperties, "timeUnit")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicTimeOffset struct {
	value *PublicTimeOffset
	isSet bool
}

func (v NullablePublicTimeOffset) Get() *PublicTimeOffset {
	return v.value
}

func (v *NullablePublicTimeOffset) Set(val *PublicTimeOffset) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicTimeOffset) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicTimeOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicTimeOffset(val *PublicTimeOffset) *NullablePublicTimeOffset {
	return &NullablePublicTimeOffset{value: val, isSet: true}
}

func (v NullablePublicTimeOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicTimeOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}