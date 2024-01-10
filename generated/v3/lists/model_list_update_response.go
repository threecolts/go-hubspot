/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
)

// checks if the ListUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUpdateResponse{}

// ListUpdateResponse The updated definition of the list in response to a list update request.
type ListUpdateResponse struct {
	UpdatedList          *PublicObjectList `json:"updatedList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListUpdateResponse ListUpdateResponse

// NewListUpdateResponse instantiates a new ListUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUpdateResponse() *ListUpdateResponse {
	this := ListUpdateResponse{}
	return &this
}

// NewListUpdateResponseWithDefaults instantiates a new ListUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUpdateResponseWithDefaults() *ListUpdateResponse {
	this := ListUpdateResponse{}
	return &this
}

// GetUpdatedList returns the UpdatedList field value if set, zero value otherwise.
func (o *ListUpdateResponse) GetUpdatedList() PublicObjectList {
	if o == nil || IsNil(o.UpdatedList) {
		var ret PublicObjectList
		return ret
	}
	return *o.UpdatedList
}

// GetUpdatedListOk returns a tuple with the UpdatedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUpdateResponse) GetUpdatedListOk() (*PublicObjectList, bool) {
	if o == nil || IsNil(o.UpdatedList) {
		return nil, false
	}
	return o.UpdatedList, true
}

// HasUpdatedList returns a boolean if a field has been set.
func (o *ListUpdateResponse) HasUpdatedList() bool {
	if o != nil && !IsNil(o.UpdatedList) {
		return true
	}

	return false
}

// SetUpdatedList gets a reference to the given PublicObjectList and assigns it to the UpdatedList field.
func (o *ListUpdateResponse) SetUpdatedList(v PublicObjectList) {
	o.UpdatedList = &v
}

func (o ListUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdatedList) {
		toSerialize["updatedList"] = o.UpdatedList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varListUpdateResponse := _ListUpdateResponse{}

	err = json.Unmarshal(data, &varListUpdateResponse)

	if err != nil {
		return err
	}

	*o = ListUpdateResponse(varListUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "updatedList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListUpdateResponse struct {
	value *ListUpdateResponse
	isSet bool
}

func (v NullableListUpdateResponse) Get() *ListUpdateResponse {
	return v.value
}

func (v *NullableListUpdateResponse) Set(val *ListUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUpdateResponse(val *ListUpdateResponse) *NullableListUpdateResponse {
	return &NullableListUpdateResponse{value: val, isSet: true}
}

func (v NullableListUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
