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

// checks if the PublicFormSubmissionOnPageFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicFormSubmissionOnPageFilter{}

// PublicFormSubmissionOnPageFilter struct for PublicFormSubmissionOnPageFilter
type PublicFormSubmissionOnPageFilter struct {
	FilterType           string                                        `json:"filterType"`
	FormId               *string                                       `json:"formId,omitempty"`
	PruningRefineBy      *PublicEventAnalyticsFilterCoalescingRefineBy `json:"pruningRefineBy,omitempty"`
	CoalescingRefineBy   *PublicEventAnalyticsFilterCoalescingRefineBy `json:"coalescingRefineBy,omitempty"`
	Operator             string                                        `json:"operator"`
	PageId               string                                        `json:"pageId"`
	AdditionalProperties map[string]interface{}
}

type _PublicFormSubmissionOnPageFilter PublicFormSubmissionOnPageFilter

// NewPublicFormSubmissionOnPageFilter instantiates a new PublicFormSubmissionOnPageFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicFormSubmissionOnPageFilter(filterType string, operator string, pageId string) *PublicFormSubmissionOnPageFilter {
	this := PublicFormSubmissionOnPageFilter{}
	this.FilterType = filterType
	this.Operator = operator
	this.PageId = pageId
	return &this
}

// NewPublicFormSubmissionOnPageFilterWithDefaults instantiates a new PublicFormSubmissionOnPageFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicFormSubmissionOnPageFilterWithDefaults() *PublicFormSubmissionOnPageFilter {
	this := PublicFormSubmissionOnPageFilter{}
	var filterType string = "FORM_SUBMISSION_ON_PAGE"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicFormSubmissionOnPageFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicFormSubmissionOnPageFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetFormId returns the FormId field value if set, zero value otherwise.
func (o *PublicFormSubmissionOnPageFilter) GetFormId() string {
	if o == nil || IsNil(o.FormId) {
		var ret string
		return ret
	}
	return *o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetFormIdOk() (*string, bool) {
	if o == nil || IsNil(o.FormId) {
		return nil, false
	}
	return o.FormId, true
}

// HasFormId returns a boolean if a field has been set.
func (o *PublicFormSubmissionOnPageFilter) HasFormId() bool {
	if o != nil && !IsNil(o.FormId) {
		return true
	}

	return false
}

// SetFormId gets a reference to the given string and assigns it to the FormId field.
func (o *PublicFormSubmissionOnPageFilter) SetFormId(v string) {
	o.FormId = &v
}

// GetPruningRefineBy returns the PruningRefineBy field value if set, zero value otherwise.
func (o *PublicFormSubmissionOnPageFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy {
	if o == nil || IsNil(o.PruningRefineBy) {
		var ret PublicEventAnalyticsFilterCoalescingRefineBy
		return ret
	}
	return *o.PruningRefineBy
}

// GetPruningRefineByOk returns a tuple with the PruningRefineBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool) {
	if o == nil || IsNil(o.PruningRefineBy) {
		return nil, false
	}
	return o.PruningRefineBy, true
}

// HasPruningRefineBy returns a boolean if a field has been set.
func (o *PublicFormSubmissionOnPageFilter) HasPruningRefineBy() bool {
	if o != nil && !IsNil(o.PruningRefineBy) {
		return true
	}

	return false
}

// SetPruningRefineBy gets a reference to the given PublicEventAnalyticsFilterCoalescingRefineBy and assigns it to the PruningRefineBy field.
func (o *PublicFormSubmissionOnPageFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy) {
	o.PruningRefineBy = &v
}

// GetCoalescingRefineBy returns the CoalescingRefineBy field value if set, zero value otherwise.
func (o *PublicFormSubmissionOnPageFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy {
	if o == nil || IsNil(o.CoalescingRefineBy) {
		var ret PublicEventAnalyticsFilterCoalescingRefineBy
		return ret
	}
	return *o.CoalescingRefineBy
}

// GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool) {
	if o == nil || IsNil(o.CoalescingRefineBy) {
		return nil, false
	}
	return o.CoalescingRefineBy, true
}

// HasCoalescingRefineBy returns a boolean if a field has been set.
func (o *PublicFormSubmissionOnPageFilter) HasCoalescingRefineBy() bool {
	if o != nil && !IsNil(o.CoalescingRefineBy) {
		return true
	}

	return false
}

// SetCoalescingRefineBy gets a reference to the given PublicEventAnalyticsFilterCoalescingRefineBy and assigns it to the CoalescingRefineBy field.
func (o *PublicFormSubmissionOnPageFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy) {
	o.CoalescingRefineBy = &v
}

// GetOperator returns the Operator field value
func (o *PublicFormSubmissionOnPageFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicFormSubmissionOnPageFilter) SetOperator(v string) {
	o.Operator = v
}

// GetPageId returns the PageId field value
func (o *PublicFormSubmissionOnPageFilter) GetPageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value
// and a boolean to check if the value has been set.
func (o *PublicFormSubmissionOnPageFilter) GetPageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageId, true
}

// SetPageId sets field value
func (o *PublicFormSubmissionOnPageFilter) SetPageId(v string) {
	o.PageId = v
}

func (o PublicFormSubmissionOnPageFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicFormSubmissionOnPageFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	if !IsNil(o.FormId) {
		toSerialize["formId"] = o.FormId
	}
	if !IsNil(o.PruningRefineBy) {
		toSerialize["pruningRefineBy"] = o.PruningRefineBy
	}
	if !IsNil(o.CoalescingRefineBy) {
		toSerialize["coalescingRefineBy"] = o.CoalescingRefineBy
	}
	toSerialize["operator"] = o.Operator
	toSerialize["pageId"] = o.PageId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicFormSubmissionOnPageFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"operator",
		"pageId",
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

	varPublicFormSubmissionOnPageFilter := _PublicFormSubmissionOnPageFilter{}

	err = json.Unmarshal(data, &varPublicFormSubmissionOnPageFilter)

	if err != nil {
		return err
	}

	*o = PublicFormSubmissionOnPageFilter(varPublicFormSubmissionOnPageFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "formId")
		delete(additionalProperties, "pruningRefineBy")
		delete(additionalProperties, "coalescingRefineBy")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "pageId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicFormSubmissionOnPageFilter struct {
	value *PublicFormSubmissionOnPageFilter
	isSet bool
}

func (v NullablePublicFormSubmissionOnPageFilter) Get() *PublicFormSubmissionOnPageFilter {
	return v.value
}

func (v *NullablePublicFormSubmissionOnPageFilter) Set(val *PublicFormSubmissionOnPageFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicFormSubmissionOnPageFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicFormSubmissionOnPageFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicFormSubmissionOnPageFilter(val *PublicFormSubmissionOnPageFilter) *NullablePublicFormSubmissionOnPageFilter {
	return &NullablePublicFormSubmissionOnPageFilter{value: val, isSet: true}
}

func (v NullablePublicFormSubmissionOnPageFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicFormSubmissionOnPageFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
