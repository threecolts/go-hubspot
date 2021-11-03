/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"encoding/json"
	"time"
)

// SimplePublicObject struct for SimplePublicObject
type SimplePublicObject struct {
	Id         string            `json:"id"`
	Properties map[string]string `json:"properties"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	Archived   *bool             `json:"archived,omitempty"`
	ArchivedAt *time.Time        `json:"archivedAt,omitempty"`
}

// NewSimplePublicObject instantiates a new SimplePublicObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObject(id string, properties map[string]string, createdAt time.Time, updatedAt time.Time) *SimplePublicObject {
	this := SimplePublicObject{}
	this.Id = id
	this.Properties = properties
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSimplePublicObjectWithDefaults instantiates a new SimplePublicObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectWithDefaults() *SimplePublicObject {
	this := SimplePublicObject{}
	return &this
}

// GetId returns the Id field value
func (o *SimplePublicObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObject) SetId(v string) {
	o.Id = v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObject) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObject) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SimplePublicObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SimplePublicObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SimplePublicObject) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SimplePublicObject) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SimplePublicObject) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SimplePublicObject) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *SimplePublicObject) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *SimplePublicObject) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *SimplePublicObject) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *SimplePublicObject) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

func (o SimplePublicObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSimplePublicObject struct {
	value *SimplePublicObject
	isSet bool
}

func (v NullableSimplePublicObject) Get() *SimplePublicObject {
	return v.value
}

func (v *NullableSimplePublicObject) Set(val *SimplePublicObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObject(val *SimplePublicObject) *NullableSimplePublicObject {
	return &NullableSimplePublicObject{value: val, isSet: true}
}

func (v NullableSimplePublicObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
