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

// checks if the PublicQuarterReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicQuarterReference{}

// PublicQuarterReference struct for PublicQuarterReference
type PublicQuarterReference struct {
	ReferenceType        string `json:"referenceType"`
	Hour                 *int32 `json:"hour,omitempty"`
	Minute               *int32 `json:"minute,omitempty"`
	Second               *int32 `json:"second,omitempty"`
	Millisecond          *int32 `json:"millisecond,omitempty"`
	Month                int32  `json:"month"`
	Day                  int32  `json:"day"`
	AdditionalProperties map[string]interface{}
}

type _PublicQuarterReference PublicQuarterReference

// NewPublicQuarterReference instantiates a new PublicQuarterReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicQuarterReference(referenceType string, month int32, day int32) *PublicQuarterReference {
	this := PublicQuarterReference{}
	this.ReferenceType = referenceType
	this.Month = month
	this.Day = day
	return &this
}

// NewPublicQuarterReferenceWithDefaults instantiates a new PublicQuarterReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicQuarterReferenceWithDefaults() *PublicQuarterReference {
	this := PublicQuarterReference{}
	var referenceType string = "QUARTER"
	this.ReferenceType = referenceType
	return &this
}

// GetReferenceType returns the ReferenceType field value
func (o *PublicQuarterReference) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *PublicQuarterReference) SetReferenceType(v string) {
	o.ReferenceType = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *PublicQuarterReference) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *PublicQuarterReference) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *PublicQuarterReference) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *PublicQuarterReference) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *PublicQuarterReference) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *PublicQuarterReference) SetMinute(v int32) {
	o.Minute = &v
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *PublicQuarterReference) GetSecond() int32 {
	if o == nil || IsNil(o.Second) {
		var ret int32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *PublicQuarterReference) HasSecond() bool {
	if o != nil && !IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given int32 and assigns it to the Second field.
func (o *PublicQuarterReference) SetSecond(v int32) {
	o.Second = &v
}

// GetMillisecond returns the Millisecond field value if set, zero value otherwise.
func (o *PublicQuarterReference) GetMillisecond() int32 {
	if o == nil || IsNil(o.Millisecond) {
		var ret int32
		return ret
	}
	return *o.Millisecond
}

// GetMillisecondOk returns a tuple with the Millisecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetMillisecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Millisecond) {
		return nil, false
	}
	return o.Millisecond, true
}

// HasMillisecond returns a boolean if a field has been set.
func (o *PublicQuarterReference) HasMillisecond() bool {
	if o != nil && !IsNil(o.Millisecond) {
		return true
	}

	return false
}

// SetMillisecond gets a reference to the given int32 and assigns it to the Millisecond field.
func (o *PublicQuarterReference) SetMillisecond(v int32) {
	o.Millisecond = &v
}

// GetMonth returns the Month field value
func (o *PublicQuarterReference) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *PublicQuarterReference) SetMonth(v int32) {
	o.Month = v
}

// GetDay returns the Day field value
func (o *PublicQuarterReference) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *PublicQuarterReference) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *PublicQuarterReference) SetDay(v int32) {
	o.Day = v
}

func (o PublicQuarterReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicQuarterReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceType"] = o.ReferenceType
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Second) {
		toSerialize["second"] = o.Second
	}
	if !IsNil(o.Millisecond) {
		toSerialize["millisecond"] = o.Millisecond
	}
	toSerialize["month"] = o.Month
	toSerialize["day"] = o.Day

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicQuarterReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceType",
		"month",
		"day",
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

	varPublicQuarterReference := _PublicQuarterReference{}

	err = json.Unmarshal(data, &varPublicQuarterReference)

	if err != nil {
		return err
	}

	*o = PublicQuarterReference(varPublicQuarterReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referenceType")
		delete(additionalProperties, "hour")
		delete(additionalProperties, "minute")
		delete(additionalProperties, "second")
		delete(additionalProperties, "millisecond")
		delete(additionalProperties, "month")
		delete(additionalProperties, "day")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicQuarterReference struct {
	value *PublicQuarterReference
	isSet bool
}

func (v NullablePublicQuarterReference) Get() *PublicQuarterReference {
	return v.value
}

func (v *NullablePublicQuarterReference) Set(val *PublicQuarterReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicQuarterReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicQuarterReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicQuarterReference(val *PublicQuarterReference) *NullablePublicQuarterReference {
	return &NullablePublicQuarterReference{value: val, isSet: true}
}

func (v NullablePublicQuarterReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicQuarterReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}