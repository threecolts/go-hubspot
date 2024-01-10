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

// checks if the PublicDatePropertyOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicDatePropertyOperation{}

// PublicDatePropertyOperation struct for PublicDatePropertyOperation
type PublicDatePropertyOperation struct {
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	Year                         int32  `json:"year"`
	Month                        string `json:"month"`
	Day                          int32  `json:"day"`
	AdditionalProperties         map[string]interface{}
}

type _PublicDatePropertyOperation PublicDatePropertyOperation

// NewPublicDatePropertyOperation instantiates a new PublicDatePropertyOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDatePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, year int32, month string, day int32) *PublicDatePropertyOperation {
	this := PublicDatePropertyOperation{}
	this.OperationType = operationType
	this.Operator = operator
	this.IncludeObjectsWithNoValueSet = includeObjectsWithNoValueSet
	this.Year = year
	this.Month = month
	this.Day = day
	return &this
}

// NewPublicDatePropertyOperationWithDefaults instantiates a new PublicDatePropertyOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDatePropertyOperationWithDefaults() *PublicDatePropertyOperation {
	this := PublicDatePropertyOperation{}
	var operationType string = "DATE"
	this.OperationType = operationType
	return &this
}

// GetOperationType returns the OperationType field value
func (o *PublicDatePropertyOperation) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *PublicDatePropertyOperation) SetOperationType(v string) {
	o.OperationType = v
}

// GetOperator returns the Operator field value
func (o *PublicDatePropertyOperation) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicDatePropertyOperation) SetOperator(v string) {
	o.Operator = v
}

// GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field value
func (o *PublicDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeObjectsWithNoValueSet
}

// GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeObjectsWithNoValueSet, true
}

// SetIncludeObjectsWithNoValueSet sets field value
func (o *PublicDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool) {
	o.IncludeObjectsWithNoValueSet = v
}

// GetYear returns the Year field value
func (o *PublicDatePropertyOperation) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *PublicDatePropertyOperation) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *PublicDatePropertyOperation) GetMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *PublicDatePropertyOperation) SetMonth(v string) {
	o.Month = v
}

// GetDay returns the Day field value
func (o *PublicDatePropertyOperation) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *PublicDatePropertyOperation) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *PublicDatePropertyOperation) SetDay(v int32) {
	o.Day = v
}

func (o PublicDatePropertyOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicDatePropertyOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationType"] = o.OperationType
	toSerialize["operator"] = o.Operator
	toSerialize["includeObjectsWithNoValueSet"] = o.IncludeObjectsWithNoValueSet
	toSerialize["year"] = o.Year
	toSerialize["month"] = o.Month
	toSerialize["day"] = o.Day

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicDatePropertyOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operationType",
		"operator",
		"includeObjectsWithNoValueSet",
		"year",
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

	varPublicDatePropertyOperation := _PublicDatePropertyOperation{}

	err = json.Unmarshal(data, &varPublicDatePropertyOperation)

	if err != nil {
		return err
	}

	*o = PublicDatePropertyOperation(varPublicDatePropertyOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operationType")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "includeObjectsWithNoValueSet")
		delete(additionalProperties, "year")
		delete(additionalProperties, "month")
		delete(additionalProperties, "day")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicDatePropertyOperation struct {
	value *PublicDatePropertyOperation
	isSet bool
}

func (v NullablePublicDatePropertyOperation) Get() *PublicDatePropertyOperation {
	return v.value
}

func (v *NullablePublicDatePropertyOperation) Set(val *PublicDatePropertyOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDatePropertyOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDatePropertyOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDatePropertyOperation(val *PublicDatePropertyOperation) *NullablePublicDatePropertyOperation {
	return &NullablePublicDatePropertyOperation{value: val, isSet: true}
}

func (v NullablePublicDatePropertyOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDatePropertyOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
