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

// checks if the ListFilterUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFilterUpdateRequest{}

// ListFilterUpdateRequest The definition of the list filter branch update request.
type ListFilterUpdateRequest struct {
	FilterBranch         PublicPropertyAssociationFilterBranchFilterBranchesInner `json:"filterBranch"`
	AdditionalProperties map[string]interface{}
}

type _ListFilterUpdateRequest ListFilterUpdateRequest

// NewListFilterUpdateRequest instantiates a new ListFilterUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFilterUpdateRequest(filterBranch PublicPropertyAssociationFilterBranchFilterBranchesInner) *ListFilterUpdateRequest {
	this := ListFilterUpdateRequest{}
	this.FilterBranch = filterBranch
	return &this
}

// NewListFilterUpdateRequestWithDefaults instantiates a new ListFilterUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFilterUpdateRequestWithDefaults() *ListFilterUpdateRequest {
	this := ListFilterUpdateRequest{}
	return &this
}

// GetFilterBranch returns the FilterBranch field value
func (o *ListFilterUpdateRequest) GetFilterBranch() PublicPropertyAssociationFilterBranchFilterBranchesInner {
	if o == nil {
		var ret PublicPropertyAssociationFilterBranchFilterBranchesInner
		return ret
	}

	return o.FilterBranch
}

// GetFilterBranchOk returns a tuple with the FilterBranch field value
// and a boolean to check if the value has been set.
func (o *ListFilterUpdateRequest) GetFilterBranchOk() (*PublicPropertyAssociationFilterBranchFilterBranchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranch, true
}

// SetFilterBranch sets field value
func (o *ListFilterUpdateRequest) SetFilterBranch(v PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	o.FilterBranch = v
}

func (o ListFilterUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFilterUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterBranch"] = o.FilterBranch

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListFilterUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterBranch",
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

	varListFilterUpdateRequest := _ListFilterUpdateRequest{}

	err = json.Unmarshal(data, &varListFilterUpdateRequest)

	if err != nil {
		return err
	}

	*o = ListFilterUpdateRequest(varListFilterUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterBranch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFilterUpdateRequest struct {
	value *ListFilterUpdateRequest
	isSet bool
}

func (v NullableListFilterUpdateRequest) Get() *ListFilterUpdateRequest {
	return v.value
}

func (v *NullableListFilterUpdateRequest) Set(val *ListFilterUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListFilterUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListFilterUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFilterUpdateRequest(val *ListFilterUpdateRequest) *NullableListFilterUpdateRequest {
	return &NullableListFilterUpdateRequest{value: val, isSet: true}
}

func (v NullableListFilterUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFilterUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}