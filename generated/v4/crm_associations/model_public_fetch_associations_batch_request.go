/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"encoding/json"
)

// checks if the PublicFetchAssociationsBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicFetchAssociationsBatchRequest{}

// PublicFetchAssociationsBatchRequest struct for PublicFetchAssociationsBatchRequest
type PublicFetchAssociationsBatchRequest struct {
	Id    string  `json:"id"`
	After *string `json:"after,omitempty"`
}

// NewPublicFetchAssociationsBatchRequest instantiates a new PublicFetchAssociationsBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicFetchAssociationsBatchRequest(id string) *PublicFetchAssociationsBatchRequest {
	this := PublicFetchAssociationsBatchRequest{}
	this.Id = id
	return &this
}

// NewPublicFetchAssociationsBatchRequestWithDefaults instantiates a new PublicFetchAssociationsBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicFetchAssociationsBatchRequestWithDefaults() *PublicFetchAssociationsBatchRequest {
	this := PublicFetchAssociationsBatchRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PublicFetchAssociationsBatchRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicFetchAssociationsBatchRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicFetchAssociationsBatchRequest) SetId(v string) {
	o.Id = v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *PublicFetchAssociationsBatchRequest) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicFetchAssociationsBatchRequest) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *PublicFetchAssociationsBatchRequest) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *PublicFetchAssociationsBatchRequest) SetAfter(v string) {
	o.After = &v
}

func (o PublicFetchAssociationsBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicFetchAssociationsBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	return toSerialize, nil
}

type NullablePublicFetchAssociationsBatchRequest struct {
	value *PublicFetchAssociationsBatchRequest
	isSet bool
}

func (v NullablePublicFetchAssociationsBatchRequest) Get() *PublicFetchAssociationsBatchRequest {
	return v.value
}

func (v *NullablePublicFetchAssociationsBatchRequest) Set(val *PublicFetchAssociationsBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicFetchAssociationsBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicFetchAssociationsBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicFetchAssociationsBatchRequest(val *PublicFetchAssociationsBatchRequest) *NullablePublicFetchAssociationsBatchRequest {
	return &NullablePublicFetchAssociationsBatchRequest{value: val, isSet: true}
}

func (v NullablePublicFetchAssociationsBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicFetchAssociationsBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}