/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"encoding/json"
)

// checks if the PublicObjectSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicObjectSearchRequest{}

// PublicObjectSearchRequest struct for PublicObjectSearchRequest
type PublicObjectSearchRequest struct {
	FilterGroups []FilterGroup `json:"filterGroups"`
	Sorts        []string      `json:"sorts"`
	Query        *string       `json:"query,omitempty"`
	Properties   []string      `json:"properties"`
	Limit        int32         `json:"limit"`
	After        int32         `json:"after"`
}

// NewPublicObjectSearchRequest instantiates a new PublicObjectSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectSearchRequest(filterGroups []FilterGroup, sorts []string, properties []string, limit int32, after int32) *PublicObjectSearchRequest {
	this := PublicObjectSearchRequest{}
	this.FilterGroups = filterGroups
	this.Sorts = sorts
	this.Properties = properties
	this.Limit = limit
	this.After = after
	return &this
}

// NewPublicObjectSearchRequestWithDefaults instantiates a new PublicObjectSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectSearchRequestWithDefaults() *PublicObjectSearchRequest {
	this := PublicObjectSearchRequest{}
	return &this
}

// GetFilterGroups returns the FilterGroups field value
func (o *PublicObjectSearchRequest) GetFilterGroups() []FilterGroup {
	if o == nil {
		var ret []FilterGroup
		return ret
	}

	return o.FilterGroups
}

// GetFilterGroupsOk returns a tuple with the FilterGroups field value
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetFilterGroupsOk() ([]FilterGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroups, true
}

// SetFilterGroups sets field value
func (o *PublicObjectSearchRequest) SetFilterGroups(v []FilterGroup) {
	o.FilterGroups = v
}

// GetSorts returns the Sorts field value
func (o *PublicObjectSearchRequest) GetSorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sorts
}

// GetSortsOk returns a tuple with the Sorts field value
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetSortsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sorts, true
}

// SetSorts sets field value
func (o *PublicObjectSearchRequest) SetSorts(v []string) {
	o.Sorts = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *PublicObjectSearchRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *PublicObjectSearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *PublicObjectSearchRequest) SetQuery(v string) {
	o.Query = &v
}

// GetProperties returns the Properties field value
func (o *PublicObjectSearchRequest) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *PublicObjectSearchRequest) SetProperties(v []string) {
	o.Properties = v
}

// GetLimit returns the Limit field value
func (o *PublicObjectSearchRequest) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PublicObjectSearchRequest) SetLimit(v int32) {
	o.Limit = v
}

// GetAfter returns the After field value
func (o *PublicObjectSearchRequest) GetAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *PublicObjectSearchRequest) GetAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *PublicObjectSearchRequest) SetAfter(v int32) {
	o.After = v
}

func (o PublicObjectSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicObjectSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterGroups"] = o.FilterGroups
	toSerialize["sorts"] = o.Sorts
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	toSerialize["properties"] = o.Properties
	toSerialize["limit"] = o.Limit
	toSerialize["after"] = o.After
	return toSerialize, nil
}

type NullablePublicObjectSearchRequest struct {
	value *PublicObjectSearchRequest
	isSet bool
}

func (v NullablePublicObjectSearchRequest) Get() *PublicObjectSearchRequest {
	return v.value
}

func (v *NullablePublicObjectSearchRequest) Set(val *PublicObjectSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectSearchRequest(val *PublicObjectSearchRequest) *NullablePublicObjectSearchRequest {
	return &NullablePublicObjectSearchRequest{value: val, isSet: true}
}

func (v NullablePublicObjectSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
