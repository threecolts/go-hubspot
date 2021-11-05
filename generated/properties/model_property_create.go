/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
)

// PropertyCreate struct for PropertyCreate
type PropertyCreate struct {
	// The internal property name, which must be used when referencing the property via the API.
	Name string `json:"name"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`
	// The data type of the property.
	Type string `json:"type"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`
	// The name of the property group the property belongs to.
	GroupName string `json:"groupName"`
	// A description of the property that will be shown as help text in HubSpot.
	Description *string `json:"description,omitempty"`
	// A list of valid options for the property. This field is required for enumerated properties.
	Options *[]OptionInput `json:"options,omitempty"`
	// Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	// If true, the property won't be visible and can't be used in HubSpot.
	Hidden *bool `json:"hidden,omitempty"`
	// Whether or not the property can be used in a HubSpot form.
	FormField *bool `json:"formField,omitempty"`
}

// NewPropertyCreate instantiates a new PropertyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyCreate(name string, label string, type_ string, fieldType string, groupName string) *PropertyCreate {
	this := PropertyCreate{}
	this.Name = name
	this.Label = label
	this.Type = type_
	this.FieldType = fieldType
	this.GroupName = groupName
	return &this
}

// NewPropertyCreateWithDefaults instantiates a new PropertyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyCreateWithDefaults() *PropertyCreate {
	this := PropertyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *PropertyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropertyCreate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *PropertyCreate) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PropertyCreate) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *PropertyCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PropertyCreate) SetType(v string) {
	o.Type = v
}

// GetFieldType returns the FieldType field value
func (o *PropertyCreate) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *PropertyCreate) SetFieldType(v string) {
	o.FieldType = v
}

// GetGroupName returns the GroupName field value
func (o *PropertyCreate) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *PropertyCreate) SetGroupName(v string) {
	o.GroupName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyCreate) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PropertyCreate) GetOptions() []OptionInput {
	if o == nil || o.Options == nil {
		var ret []OptionInput
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetOptionsOk() (*[]OptionInput, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PropertyCreate) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []OptionInput and assigns it to the Options field.
func (o *PropertyCreate) SetOptions(v []OptionInput) {
	o.Options = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *PropertyCreate) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder == nil {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *PropertyCreate) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *PropertyCreate) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetHasUniqueValue returns the HasUniqueValue field value if set, zero value otherwise.
func (o *PropertyCreate) GetHasUniqueValue() bool {
	if o == nil || o.HasUniqueValue == nil {
		var ret bool
		return ret
	}
	return *o.HasUniqueValue
}

// GetHasUniqueValueOk returns a tuple with the HasUniqueValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetHasUniqueValueOk() (*bool, bool) {
	if o == nil || o.HasUniqueValue == nil {
		return nil, false
	}
	return o.HasUniqueValue, true
}

// HasHasUniqueValue returns a boolean if a field has been set.
func (o *PropertyCreate) HasHasUniqueValue() bool {
	if o != nil && o.HasUniqueValue != nil {
		return true
	}

	return false
}

// SetHasUniqueValue gets a reference to the given bool and assigns it to the HasUniqueValue field.
func (o *PropertyCreate) SetHasUniqueValue(v bool) {
	o.HasUniqueValue = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *PropertyCreate) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *PropertyCreate) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *PropertyCreate) SetHidden(v bool) {
	o.Hidden = &v
}

// GetFormField returns the FormField field value if set, zero value otherwise.
func (o *PropertyCreate) GetFormField() bool {
	if o == nil || o.FormField == nil {
		var ret bool
		return ret
	}
	return *o.FormField
}

// GetFormFieldOk returns a tuple with the FormField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetFormFieldOk() (*bool, bool) {
	if o == nil || o.FormField == nil {
		return nil, false
	}
	return o.FormField, true
}

// HasFormField returns a boolean if a field has been set.
func (o *PropertyCreate) HasFormField() bool {
	if o != nil && o.FormField != nil {
		return true
	}

	return false
}

// SetFormField gets a reference to the given bool and assigns it to the FormField field.
func (o *PropertyCreate) SetFormField(v bool) {
	o.FormField = &v
}

func (o PropertyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["fieldType"] = o.FieldType
	}
	if true {
		toSerialize["groupName"] = o.GroupName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.HasUniqueValue != nil {
		toSerialize["hasUniqueValue"] = o.HasUniqueValue
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.FormField != nil {
		toSerialize["formField"] = o.FormField
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyCreate struct {
	value *PropertyCreate
	isSet bool
}

func (v NullablePropertyCreate) Get() *PropertyCreate {
	return v.value
}

func (v *NullablePropertyCreate) Set(val *PropertyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyCreate(val *PropertyCreate) *NullablePropertyCreate {
	return &NullablePropertyCreate{value: val, isSet: true}
}

func (v NullablePropertyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}