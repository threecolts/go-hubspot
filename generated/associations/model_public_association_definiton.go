/*
Associations

Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"encoding/json"
)

// PublicAssociationDefiniton struct for PublicAssociationDefiniton
type PublicAssociationDefiniton struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// NewPublicAssociationDefiniton instantiates a new PublicAssociationDefiniton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationDefiniton(id string, name string) *PublicAssociationDefiniton {
	this := PublicAssociationDefiniton{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPublicAssociationDefinitonWithDefaults instantiates a new PublicAssociationDefiniton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationDefinitonWithDefaults() *PublicAssociationDefiniton {
	this := PublicAssociationDefiniton{}
	return &this
}

// GetId returns the Id field value
func (o *PublicAssociationDefiniton) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefiniton) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicAssociationDefiniton) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PublicAssociationDefiniton) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefiniton) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicAssociationDefiniton) SetName(v string) {
	o.Name = v
}

func (o PublicAssociationDefiniton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAssociationDefiniton struct {
	value *PublicAssociationDefiniton
	isSet bool
}

func (v NullablePublicAssociationDefiniton) Get() *PublicAssociationDefiniton {
	return v.value
}

func (v *NullablePublicAssociationDefiniton) Set(val *PublicAssociationDefiniton) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationDefiniton) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationDefiniton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationDefiniton(val *PublicAssociationDefiniton) *NullablePublicAssociationDefiniton {
	return &NullablePublicAssociationDefiniton{value: val, isSet: true}
}

func (v NullablePublicAssociationDefiniton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationDefiniton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
