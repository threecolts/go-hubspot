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

// checks if the PublicAssociationMultiWithLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationMultiWithLabel{}

// PublicAssociationMultiWithLabel struct for PublicAssociationMultiWithLabel
type PublicAssociationMultiWithLabel struct {
	From   PublicObjectId                   `json:"from"`
	To     []MultiAssociatedObjectWithLabel `json:"to"`
	Paging *Paging                          `json:"paging,omitempty"`
}

// NewPublicAssociationMultiWithLabel instantiates a new PublicAssociationMultiWithLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationMultiWithLabel(from PublicObjectId, to []MultiAssociatedObjectWithLabel) *PublicAssociationMultiWithLabel {
	this := PublicAssociationMultiWithLabel{}
	this.From = from
	this.To = to
	return &this
}

// NewPublicAssociationMultiWithLabelWithDefaults instantiates a new PublicAssociationMultiWithLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationMultiWithLabelWithDefaults() *PublicAssociationMultiWithLabel {
	this := PublicAssociationMultiWithLabel{}
	return &this
}

// GetFrom returns the From field value
func (o *PublicAssociationMultiWithLabel) GetFrom() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationMultiWithLabel) GetFromOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PublicAssociationMultiWithLabel) SetFrom(v PublicObjectId) {
	o.From = v
}

// GetTo returns the To field value
func (o *PublicAssociationMultiWithLabel) GetTo() []MultiAssociatedObjectWithLabel {
	if o == nil {
		var ret []MultiAssociatedObjectWithLabel
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationMultiWithLabel) GetToOk() ([]MultiAssociatedObjectWithLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *PublicAssociationMultiWithLabel) SetTo(v []MultiAssociatedObjectWithLabel) {
	o.To = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *PublicAssociationMultiWithLabel) GetPaging() Paging {
	if o == nil || IsNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAssociationMultiWithLabel) GetPagingOk() (*Paging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *PublicAssociationMultiWithLabel) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *PublicAssociationMultiWithLabel) SetPaging(v Paging) {
	o.Paging = &v
}

func (o PublicAssociationMultiWithLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationMultiWithLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullablePublicAssociationMultiWithLabel struct {
	value *PublicAssociationMultiWithLabel
	isSet bool
}

func (v NullablePublicAssociationMultiWithLabel) Get() *PublicAssociationMultiWithLabel {
	return v.value
}

func (v *NullablePublicAssociationMultiWithLabel) Set(val *PublicAssociationMultiWithLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationMultiWithLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationMultiWithLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationMultiWithLabel(val *PublicAssociationMultiWithLabel) *NullablePublicAssociationMultiWithLabel {
	return &NullablePublicAssociationMultiWithLabel{value: val, isSet: true}
}

func (v NullablePublicAssociationMultiWithLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationMultiWithLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
