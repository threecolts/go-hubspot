/*
CRM cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_extensions

import (
	"encoding/json"
)

// ActionHookActionBody struct for ActionHookActionBody
type ActionHookActionBody struct {
	Type                  string                  `json:"type"`
	Confirmation          *ActionConfirmationBody `json:"confirmation,omitempty"`
	HttpMethod            string                  `json:"httpMethod"`
	Url                   string                  `json:"url"`
	Label                 *string                 `json:"label,omitempty"`
	PropertyNamesIncluded []string                `json:"propertyNamesIncluded"`
}

// NewActionHookActionBody instantiates a new ActionHookActionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionHookActionBody(type_ string, httpMethod string, url string, propertyNamesIncluded []string) *ActionHookActionBody {
	this := ActionHookActionBody{}
	this.Type = type_
	this.HttpMethod = httpMethod
	this.Url = url
	this.PropertyNamesIncluded = propertyNamesIncluded
	return &this
}

// NewActionHookActionBodyWithDefaults instantiates a new ActionHookActionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionHookActionBodyWithDefaults() *ActionHookActionBody {
	this := ActionHookActionBody{}
	var type_ string = "ACTION_HOOK"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ActionHookActionBody) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActionHookActionBody) SetType(v string) {
	o.Type = v
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise.
func (o *ActionHookActionBody) GetConfirmation() ActionConfirmationBody {
	if o == nil || o.Confirmation == nil {
		var ret ActionConfirmationBody
		return ret
	}
	return *o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetConfirmationOk() (*ActionConfirmationBody, bool) {
	if o == nil || o.Confirmation == nil {
		return nil, false
	}
	return o.Confirmation, true
}

// HasConfirmation returns a boolean if a field has been set.
func (o *ActionHookActionBody) HasConfirmation() bool {
	if o != nil && o.Confirmation != nil {
		return true
	}

	return false
}

// SetConfirmation gets a reference to the given ActionConfirmationBody and assigns it to the Confirmation field.
func (o *ActionHookActionBody) SetConfirmation(v ActionConfirmationBody) {
	o.Confirmation = &v
}

// GetHttpMethod returns the HttpMethod field value
func (o *ActionHookActionBody) GetHttpMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetHttpMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpMethod, true
}

// SetHttpMethod sets field value
func (o *ActionHookActionBody) SetHttpMethod(v string) {
	o.HttpMethod = v
}

// GetUrl returns the Url field value
func (o *ActionHookActionBody) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ActionHookActionBody) SetUrl(v string) {
	o.Url = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ActionHookActionBody) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ActionHookActionBody) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ActionHookActionBody) SetLabel(v string) {
	o.Label = &v
}

// GetPropertyNamesIncluded returns the PropertyNamesIncluded field value
func (o *ActionHookActionBody) GetPropertyNamesIncluded() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyNamesIncluded
}

// GetPropertyNamesIncludedOk returns a tuple with the PropertyNamesIncluded field value
// and a boolean to check if the value has been set.
func (o *ActionHookActionBody) GetPropertyNamesIncludedOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyNamesIncluded, true
}

// SetPropertyNamesIncluded sets field value
func (o *ActionHookActionBody) SetPropertyNamesIncluded(v []string) {
	o.PropertyNamesIncluded = v
}

func (o ActionHookActionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Confirmation != nil {
		toSerialize["confirmation"] = o.Confirmation
	}
	if true {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["propertyNamesIncluded"] = o.PropertyNamesIncluded
	}
	return json.Marshal(toSerialize)
}

type NullableActionHookActionBody struct {
	value *ActionHookActionBody
	isSet bool
}

func (v NullableActionHookActionBody) Get() *ActionHookActionBody {
	return v.value
}

func (v *NullableActionHookActionBody) Set(val *ActionHookActionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableActionHookActionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableActionHookActionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionHookActionBody(val *ActionHookActionBody) *NullableActionHookActionBody {
	return &NullableActionHookActionBody{value: val, isSet: true}
}

func (v NullableActionHookActionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionHookActionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
