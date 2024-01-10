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

// checks if the PublicAssociationInListFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationInListFilter{}

// PublicAssociationInListFilter struct for PublicAssociationInListFilter
type PublicAssociationInListFilter struct {
	FilterType           string                                       `json:"filterType"`
	CoalescingRefineBy   PublicEventAnalyticsFilterCoalescingRefineBy `json:"coalescingRefineBy"`
	Operator             string                                       `json:"operator"`
	ListId               int32                                        `json:"listId"`
	ToObjectTypeId       *string                                      `json:"toObjectTypeId,omitempty"`
	AssociationTypeId    int32                                        `json:"associationTypeId"`
	AssociationCategory  string                                       `json:"associationCategory"`
	ToObjectType         *string                                      `json:"toObjectType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicAssociationInListFilter PublicAssociationInListFilter

// NewPublicAssociationInListFilter instantiates a new PublicAssociationInListFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationInListFilter(filterType string, coalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, operator string, listId int32, associationTypeId int32, associationCategory string) *PublicAssociationInListFilter {
	this := PublicAssociationInListFilter{}
	this.FilterType = filterType
	this.CoalescingRefineBy = coalescingRefineBy
	this.Operator = operator
	this.ListId = listId
	this.AssociationTypeId = associationTypeId
	this.AssociationCategory = associationCategory
	return &this
}

// NewPublicAssociationInListFilterWithDefaults instantiates a new PublicAssociationInListFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationInListFilterWithDefaults() *PublicAssociationInListFilter {
	this := PublicAssociationInListFilter{}
	var filterType string = "ASSOCIATION"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicAssociationInListFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicAssociationInListFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetCoalescingRefineBy returns the CoalescingRefineBy field value
func (o *PublicAssociationInListFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy {
	if o == nil {
		var ret PublicEventAnalyticsFilterCoalescingRefineBy
		return ret
	}

	return o.CoalescingRefineBy
}

// GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoalescingRefineBy, true
}

// SetCoalescingRefineBy sets field value
func (o *PublicAssociationInListFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy) {
	o.CoalescingRefineBy = v
}

// GetOperator returns the Operator field value
func (o *PublicAssociationInListFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicAssociationInListFilter) SetOperator(v string) {
	o.Operator = v
}

// GetListId returns the ListId field value
func (o *PublicAssociationInListFilter) GetListId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetListIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicAssociationInListFilter) SetListId(v int32) {
	o.ListId = v
}

// GetToObjectTypeId returns the ToObjectTypeId field value if set, zero value otherwise.
func (o *PublicAssociationInListFilter) GetToObjectTypeId() string {
	if o == nil || IsNil(o.ToObjectTypeId) {
		var ret string
		return ret
	}
	return *o.ToObjectTypeId
}

// GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetToObjectTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ToObjectTypeId) {
		return nil, false
	}
	return o.ToObjectTypeId, true
}

// HasToObjectTypeId returns a boolean if a field has been set.
func (o *PublicAssociationInListFilter) HasToObjectTypeId() bool {
	if o != nil && !IsNil(o.ToObjectTypeId) {
		return true
	}

	return false
}

// SetToObjectTypeId gets a reference to the given string and assigns it to the ToObjectTypeId field.
func (o *PublicAssociationInListFilter) SetToObjectTypeId(v string) {
	o.ToObjectTypeId = &v
}

// GetAssociationTypeId returns the AssociationTypeId field value
func (o *PublicAssociationInListFilter) GetAssociationTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssociationTypeId
}

// GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetAssociationTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationTypeId, true
}

// SetAssociationTypeId sets field value
func (o *PublicAssociationInListFilter) SetAssociationTypeId(v int32) {
	o.AssociationTypeId = v
}

// GetAssociationCategory returns the AssociationCategory field value
func (o *PublicAssociationInListFilter) GetAssociationCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociationCategory
}

// GetAssociationCategoryOk returns a tuple with the AssociationCategory field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetAssociationCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationCategory, true
}

// SetAssociationCategory sets field value
func (o *PublicAssociationInListFilter) SetAssociationCategory(v string) {
	o.AssociationCategory = v
}

// GetToObjectType returns the ToObjectType field value if set, zero value otherwise.
func (o *PublicAssociationInListFilter) GetToObjectType() string {
	if o == nil || IsNil(o.ToObjectType) {
		var ret string
		return ret
	}
	return *o.ToObjectType
}

// GetToObjectTypeOk returns a tuple with the ToObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAssociationInListFilter) GetToObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ToObjectType) {
		return nil, false
	}
	return o.ToObjectType, true
}

// HasToObjectType returns a boolean if a field has been set.
func (o *PublicAssociationInListFilter) HasToObjectType() bool {
	if o != nil && !IsNil(o.ToObjectType) {
		return true
	}

	return false
}

// SetToObjectType gets a reference to the given string and assigns it to the ToObjectType field.
func (o *PublicAssociationInListFilter) SetToObjectType(v string) {
	o.ToObjectType = &v
}

func (o PublicAssociationInListFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationInListFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	toSerialize["coalescingRefineBy"] = o.CoalescingRefineBy
	toSerialize["operator"] = o.Operator
	toSerialize["listId"] = o.ListId
	if !IsNil(o.ToObjectTypeId) {
		toSerialize["toObjectTypeId"] = o.ToObjectTypeId
	}
	toSerialize["associationTypeId"] = o.AssociationTypeId
	toSerialize["associationCategory"] = o.AssociationCategory
	if !IsNil(o.ToObjectType) {
		toSerialize["toObjectType"] = o.ToObjectType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicAssociationInListFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"coalescingRefineBy",
		"operator",
		"listId",
		"associationTypeId",
		"associationCategory",
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

	varPublicAssociationInListFilter := _PublicAssociationInListFilter{}

	err = json.Unmarshal(data, &varPublicAssociationInListFilter)

	if err != nil {
		return err
	}

	*o = PublicAssociationInListFilter(varPublicAssociationInListFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "coalescingRefineBy")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "listId")
		delete(additionalProperties, "toObjectTypeId")
		delete(additionalProperties, "associationTypeId")
		delete(additionalProperties, "associationCategory")
		delete(additionalProperties, "toObjectType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicAssociationInListFilter struct {
	value *PublicAssociationInListFilter
	isSet bool
}

func (v NullablePublicAssociationInListFilter) Get() *PublicAssociationInListFilter {
	return v.value
}

func (v *NullablePublicAssociationInListFilter) Set(val *PublicAssociationInListFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationInListFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationInListFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationInListFilter(val *PublicAssociationInListFilter) *NullablePublicAssociationInListFilter {
	return &NullablePublicAssociationInListFilter{value: val, isSet: true}
}

func (v NullablePublicAssociationInListFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationInListFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
