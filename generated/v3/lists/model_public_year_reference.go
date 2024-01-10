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

// checks if the PublicYearReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicYearReference{}

// PublicYearReference struct for PublicYearReference
type PublicYearReference struct {
	ReferenceType        string `json:"referenceType"`
	Hour                 *int32 `json:"hour,omitempty"`
	Minute               *int32 `json:"minute,omitempty"`
	Second               *int32 `json:"second,omitempty"`
	Millisecond          *int32 `json:"millisecond,omitempty"`
	Month                int32  `json:"month"`
	Day                  int32  `json:"day"`
	AdditionalProperties map[string]interface{}
}

type _PublicYearReference PublicYearReference

// NewPublicYearReference instantiates a new PublicYearReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicYearReference(referenceType string, month int32, day int32) *PublicYearReference {
	this := PublicYearReference{}
	this.ReferenceType = referenceType
	this.Month = month
	this.Day = day
	return &this
}

// NewPublicYearReferenceWithDefaults instantiates a new PublicYearReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicYearReferenceWithDefaults() *PublicYearReference {
	this := PublicYearReference{}
	var referenceType string = "YEAR"
	this.ReferenceType = referenceType
	return &this
}

// GetReferenceType returns the ReferenceType field value
func (o *PublicYearReference) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *PublicYearReference) SetReferenceType(v string) {
	o.ReferenceType = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *PublicYearReference) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *PublicYearReference) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *PublicYearReference) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *PublicYearReference) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *PublicYearReference) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *PublicYearReference) SetMinute(v int32) {
	o.Minute = &v
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *PublicYearReference) GetSecond() int32 {
	if o == nil || IsNil(o.Second) {
		var ret int32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *PublicYearReference) HasSecond() bool {
	if o != nil && !IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given int32 and assigns it to the Second field.
func (o *PublicYearReference) SetSecond(v int32) {
	o.Second = &v
}

// GetMillisecond returns the Millisecond field value if set, zero value otherwise.
func (o *PublicYearReference) GetMillisecond() int32 {
	if o == nil || IsNil(o.Millisecond) {
		var ret int32
		return ret
	}
	return *o.Millisecond
}

// GetMillisecondOk returns a tuple with the Millisecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetMillisecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Millisecond) {
		return nil, false
	}
	return o.Millisecond, true
}

// HasMillisecond returns a boolean if a field has been set.
func (o *PublicYearReference) HasMillisecond() bool {
	if o != nil && !IsNil(o.Millisecond) {
		return true
	}

	return false
}

// SetMillisecond gets a reference to the given int32 and assigns it to the Millisecond field.
func (o *PublicYearReference) SetMillisecond(v int32) {
	o.Millisecond = &v
}

// GetMonth returns the Month field value
func (o *PublicYearReference) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *PublicYearReference) SetMonth(v int32) {
	o.Month = v
}

// GetDay returns the Day field value
func (o *PublicYearReference) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *PublicYearReference) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *PublicYearReference) SetDay(v int32) {
	o.Day = v
}

func (o PublicYearReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicYearReference) ToMap() (map[string]interface{}, error) {
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

func (o *PublicYearReference) UnmarshalJSON(data []byte) (err error) {
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

	varPublicYearReference := _PublicYearReference{}

	err = json.Unmarshal(data, &varPublicYearReference)

	if err != nil {
		return err
	}

	*o = PublicYearReference(varPublicYearReference)

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

type NullablePublicYearReference struct {
	value *PublicYearReference
	isSet bool
}

func (v NullablePublicYearReference) Get() *PublicYearReference {
	return v.value
}

func (v *NullablePublicYearReference) Set(val *PublicYearReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicYearReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicYearReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicYearReference(val *PublicYearReference) *NullablePublicYearReference {
	return &NullablePublicYearReference{value: val, isSet: true}
}

func (v NullablePublicYearReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicYearReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}