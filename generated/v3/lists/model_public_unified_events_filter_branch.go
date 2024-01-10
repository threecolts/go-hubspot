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

// checks if the PublicUnifiedEventsFilterBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicUnifiedEventsFilterBranch{}

// PublicUnifiedEventsFilterBranch struct for PublicUnifiedEventsFilterBranch
type PublicUnifiedEventsFilterBranch struct {
	FilterBranchType     string                                                     `json:"filterBranchType"`
	FilterBranches       []PublicPropertyAssociationFilterBranchFilterBranchesInner `json:"filterBranches"`
	Filters              []PublicPropertyAssociationFilterBranchFiltersInner        `json:"filters"`
	EventTypeId          string                                                     `json:"eventTypeId"`
	Operator             string                                                     `json:"operator"`
	FilterBranchOperator string                                                     `json:"filterBranchOperator"`
	AdditionalProperties map[string]interface{}
}

type _PublicUnifiedEventsFilterBranch PublicUnifiedEventsFilterBranch

// NewPublicUnifiedEventsFilterBranch instantiates a new PublicUnifiedEventsFilterBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicUnifiedEventsFilterBranch(filterBranchType string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, eventTypeId string, operator string, filterBranchOperator string) *PublicUnifiedEventsFilterBranch {
	this := PublicUnifiedEventsFilterBranch{}
	this.FilterBranchType = filterBranchType
	this.FilterBranches = filterBranches
	this.Filters = filters
	this.EventTypeId = eventTypeId
	this.Operator = operator
	this.FilterBranchOperator = filterBranchOperator
	return &this
}

// NewPublicUnifiedEventsFilterBranchWithDefaults instantiates a new PublicUnifiedEventsFilterBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicUnifiedEventsFilterBranchWithDefaults() *PublicUnifiedEventsFilterBranch {
	this := PublicUnifiedEventsFilterBranch{}
	var filterBranchType string = "UNIFIED_EVENTS"
	this.FilterBranchType = filterBranchType
	return &this
}

// GetFilterBranchType returns the FilterBranchType field value
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchType
}

// GetFilterBranchTypeOk returns a tuple with the FilterBranchType field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchType, true
}

// SetFilterBranchType sets field value
func (o *PublicUnifiedEventsFilterBranch) SetFilterBranchType(v string) {
	o.FilterBranchType = v
}

// GetFilterBranches returns the FilterBranches field value
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner {
	if o == nil {
		var ret []PublicPropertyAssociationFilterBranchFilterBranchesInner
		return ret
	}

	return o.FilterBranches
}

// GetFilterBranchesOk returns a tuple with the FilterBranches field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranchesOk() ([]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterBranches, true
}

// SetFilterBranches sets field value
func (o *PublicUnifiedEventsFilterBranch) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	o.FilterBranches = v
}

// GetFilters returns the Filters field value
func (o *PublicUnifiedEventsFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner {
	if o == nil {
		var ret []PublicPropertyAssociationFilterBranchFiltersInner
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetFiltersOk() ([]PublicPropertyAssociationFilterBranchFiltersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *PublicUnifiedEventsFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner) {
	o.Filters = v
}

// GetEventTypeId returns the EventTypeId field value
func (o *PublicUnifiedEventsFilterBranch) GetEventTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTypeId
}

// GetEventTypeIdOk returns a tuple with the EventTypeId field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetEventTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeId, true
}

// SetEventTypeId sets field value
func (o *PublicUnifiedEventsFilterBranch) SetEventTypeId(v string) {
	o.EventTypeId = v
}

// GetOperator returns the Operator field value
func (o *PublicUnifiedEventsFilterBranch) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicUnifiedEventsFilterBranch) SetOperator(v string) {
	o.Operator = v
}

// GetFilterBranchOperator returns the FilterBranchOperator field value
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranchOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchOperator
}

// GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilterBranch) GetFilterBranchOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchOperator, true
}

// SetFilterBranchOperator sets field value
func (o *PublicUnifiedEventsFilterBranch) SetFilterBranchOperator(v string) {
	o.FilterBranchOperator = v
}

func (o PublicUnifiedEventsFilterBranch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicUnifiedEventsFilterBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterBranchType"] = o.FilterBranchType
	toSerialize["filterBranches"] = o.FilterBranches
	toSerialize["filters"] = o.Filters
	toSerialize["eventTypeId"] = o.EventTypeId
	toSerialize["operator"] = o.Operator
	toSerialize["filterBranchOperator"] = o.FilterBranchOperator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicUnifiedEventsFilterBranch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterBranchType",
		"filterBranches",
		"filters",
		"eventTypeId",
		"operator",
		"filterBranchOperator",
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

	varPublicUnifiedEventsFilterBranch := _PublicUnifiedEventsFilterBranch{}

	err = json.Unmarshal(data, &varPublicUnifiedEventsFilterBranch)

	if err != nil {
		return err
	}

	*o = PublicUnifiedEventsFilterBranch(varPublicUnifiedEventsFilterBranch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterBranchType")
		delete(additionalProperties, "filterBranches")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "eventTypeId")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "filterBranchOperator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicUnifiedEventsFilterBranch struct {
	value *PublicUnifiedEventsFilterBranch
	isSet bool
}

func (v NullablePublicUnifiedEventsFilterBranch) Get() *PublicUnifiedEventsFilterBranch {
	return v.value
}

func (v *NullablePublicUnifiedEventsFilterBranch) Set(val *PublicUnifiedEventsFilterBranch) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicUnifiedEventsFilterBranch) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicUnifiedEventsFilterBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicUnifiedEventsFilterBranch(val *PublicUnifiedEventsFilterBranch) *NullablePublicUnifiedEventsFilterBranch {
	return &NullablePublicUnifiedEventsFilterBranch{value: val, isSet: true}
}

func (v NullablePublicUnifiedEventsFilterBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicUnifiedEventsFilterBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}