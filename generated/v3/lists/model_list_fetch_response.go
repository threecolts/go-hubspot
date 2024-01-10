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

// checks if the ListFetchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFetchResponse{}

// ListFetchResponse The response for a list fetch request.
type ListFetchResponse struct {
	List                 PublicObjectList `json:"list"`
	AdditionalProperties map[string]interface{}
}

type _ListFetchResponse ListFetchResponse

// NewListFetchResponse instantiates a new ListFetchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFetchResponse(list PublicObjectList) *ListFetchResponse {
	this := ListFetchResponse{}
	this.List = list
	return &this
}

// NewListFetchResponseWithDefaults instantiates a new ListFetchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFetchResponseWithDefaults() *ListFetchResponse {
	this := ListFetchResponse{}
	return &this
}

// GetList returns the List field value
func (o *ListFetchResponse) GetList() PublicObjectList {
	if o == nil {
		var ret PublicObjectList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ListFetchResponse) GetListOk() (*PublicObjectList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *ListFetchResponse) SetList(v PublicObjectList) {
	o.List = v
}

func (o ListFetchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFetchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListFetchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list",
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

	varListFetchResponse := _ListFetchResponse{}

	err = json.Unmarshal(data, &varListFetchResponse)

	if err != nil {
		return err
	}

	*o = ListFetchResponse(varListFetchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFetchResponse struct {
	value *ListFetchResponse
	isSet bool
}

func (v NullableListFetchResponse) Get() *ListFetchResponse {
	return v.value
}

func (v *NullableListFetchResponse) Set(val *ListFetchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFetchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFetchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFetchResponse(val *ListFetchResponse) *NullableListFetchResponse {
	return &NullableListFetchResponse{value: val, isSet: true}
}

func (v NullableListFetchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFetchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
