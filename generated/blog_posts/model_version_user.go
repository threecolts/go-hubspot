/*
Blog Post endpoints

\"Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags\"

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
)

// VersionUser Model definition for a VersionUser. Contains addition information about the user who created a version.
type VersionUser struct {
	// The unique ID of the User.
	Id string `json:"id"`
	// The email address of the user.
	Email string `json:"email"`
	// The first and last name of the User.
	FullName string `json:"fullName"`
}

// NewVersionUser instantiates a new VersionUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionUser(id string, email string, fullName string) *VersionUser {
	this := VersionUser{}
	this.Id = id
	this.Email = email
	this.FullName = fullName
	return &this
}

// NewVersionUserWithDefaults instantiates a new VersionUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionUserWithDefaults() *VersionUser {
	this := VersionUser{}
	return &this
}

// GetId returns the Id field value
func (o *VersionUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VersionUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VersionUser) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *VersionUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *VersionUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *VersionUser) SetEmail(v string) {
	o.Email = v
}

// GetFullName returns the FullName field value
func (o *VersionUser) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *VersionUser) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *VersionUser) SetFullName(v string) {
	o.FullName = v
}

func (o VersionUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["fullName"] = o.FullName
	}
	return json.Marshal(toSerialize)
}

type NullableVersionUser struct {
	value *VersionUser
	isSet bool
}

func (v NullableVersionUser) Get() *VersionUser {
	return v.value
}

func (v *NullableVersionUser) Set(val *VersionUser) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionUser) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionUser(val *VersionUser) *NullableVersionUser {
	return &NullableVersionUser{value: val, isSet: true}
}

func (v NullableVersionUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}