/*
Subscriptions

Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package communications_status

import (
	"encoding/json"
	"time"
)

// checks if the SubscriptionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDefinition{}

// SubscriptionDefinition struct for SubscriptionDefinition
type SubscriptionDefinition struct {
	// The ID of the definition.
	Id string `json:"id"`
	// The name of the subscription.
	Name string `json:"name"`
	// A description of the subscription.
	Description string `json:"description"`
	// The purpose of this subscription or the department in your organization that uses it.
	Purpose *string `json:"purpose,omitempty"`
	// The method or technology used to contact.
	CommunicationMethod *string `json:"communicationMethod,omitempty"`
	// Whether the definition is active or archived.
	IsActive bool `json:"isActive"`
	// A subscription definition created by HubSpot.
	IsDefault bool `json:"isDefault"`
	// A default description that is used by some HubSpot tools and cannot be edited.
	IsInternal bool `json:"isInternal"`
	// Time at which the definition was created.
	CreatedAt time.Time `json:"createdAt"`
	// Time at which the definition was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewSubscriptionDefinition instantiates a new SubscriptionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDefinition(id string, name string, description string, isActive bool, isDefault bool, isInternal bool, createdAt time.Time, updatedAt time.Time) *SubscriptionDefinition {
	this := SubscriptionDefinition{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.IsActive = isActive
	this.IsDefault = isDefault
	this.IsInternal = isInternal
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionDefinitionWithDefaults instantiates a new SubscriptionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDefinitionWithDefaults() *SubscriptionDefinition {
	this := SubscriptionDefinition{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionDefinition) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SubscriptionDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SubscriptionDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SubscriptionDefinition) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *SubscriptionDefinition) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *SubscriptionDefinition) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *SubscriptionDefinition) SetPurpose(v string) {
	o.Purpose = &v
}

// GetCommunicationMethod returns the CommunicationMethod field value if set, zero value otherwise.
func (o *SubscriptionDefinition) GetCommunicationMethod() string {
	if o == nil || IsNil(o.CommunicationMethod) {
		var ret string
		return ret
	}
	return *o.CommunicationMethod
}

// GetCommunicationMethodOk returns a tuple with the CommunicationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetCommunicationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CommunicationMethod) {
		return nil, false
	}
	return o.CommunicationMethod, true
}

// HasCommunicationMethod returns a boolean if a field has been set.
func (o *SubscriptionDefinition) HasCommunicationMethod() bool {
	if o != nil && !IsNil(o.CommunicationMethod) {
		return true
	}

	return false
}

// SetCommunicationMethod gets a reference to the given string and assigns it to the CommunicationMethod field.
func (o *SubscriptionDefinition) SetCommunicationMethod(v string) {
	o.CommunicationMethod = &v
}

// GetIsActive returns the IsActive field value
func (o *SubscriptionDefinition) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *SubscriptionDefinition) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsDefault returns the IsDefault field value
func (o *SubscriptionDefinition) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *SubscriptionDefinition) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIsInternal returns the IsInternal field value
func (o *SubscriptionDefinition) GetIsInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetIsInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInternal, true
}

// SetIsInternal sets field value
func (o *SubscriptionDefinition) SetIsInternal(v bool) {
	o.IsInternal = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SubscriptionDefinition) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SubscriptionDefinition) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SubscriptionDefinition) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinition) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SubscriptionDefinition) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SubscriptionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.CommunicationMethod) {
		toSerialize["communicationMethod"] = o.CommunicationMethod
	}
	toSerialize["isActive"] = o.IsActive
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["isInternal"] = o.IsInternal
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableSubscriptionDefinition struct {
	value *SubscriptionDefinition
	isSet bool
}

func (v NullableSubscriptionDefinition) Get() *SubscriptionDefinition {
	return v.value
}

func (v *NullableSubscriptionDefinition) Set(val *SubscriptionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDefinition(val *SubscriptionDefinition) *NullableSubscriptionDefinition {
	return &NullableSubscriptionDefinition{value: val, isSet: true}
}

func (v NullableSubscriptionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
