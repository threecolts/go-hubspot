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

// IFrameActionBody struct for IFrameActionBody
type IFrameActionBody struct {
	Type                  string   `json:"type"`
	Width                 int32    `json:"width"`
	Height                int32    `json:"height"`
	Url                   string   `json:"url"`
	Label                 *string  `json:"label,omitempty"`
	PropertyNamesIncluded []string `json:"propertyNamesIncluded"`
}

// NewIFrameActionBody instantiates a new IFrameActionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIFrameActionBody(type_ string, width int32, height int32, url string, propertyNamesIncluded []string) *IFrameActionBody {
	this := IFrameActionBody{}
	this.Type = type_
	this.Width = width
	this.Height = height
	this.Url = url
	this.PropertyNamesIncluded = propertyNamesIncluded
	return &this
}

// NewIFrameActionBodyWithDefaults instantiates a new IFrameActionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIFrameActionBodyWithDefaults() *IFrameActionBody {
	this := IFrameActionBody{}
	var type_ string = "IFRAME"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *IFrameActionBody) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IFrameActionBody) SetType(v string) {
	o.Type = v
}

// GetWidth returns the Width field value
func (o *IFrameActionBody) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *IFrameActionBody) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *IFrameActionBody) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *IFrameActionBody) SetHeight(v int32) {
	o.Height = v
}

// GetUrl returns the Url field value
func (o *IFrameActionBody) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IFrameActionBody) SetUrl(v string) {
	o.Url = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IFrameActionBody) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IFrameActionBody) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IFrameActionBody) SetLabel(v string) {
	o.Label = &v
}

// GetPropertyNamesIncluded returns the PropertyNamesIncluded field value
func (o *IFrameActionBody) GetPropertyNamesIncluded() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyNamesIncluded
}

// GetPropertyNamesIncludedOk returns a tuple with the PropertyNamesIncluded field value
// and a boolean to check if the value has been set.
func (o *IFrameActionBody) GetPropertyNamesIncludedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyNamesIncluded, true
}

// SetPropertyNamesIncluded sets field value
func (o *IFrameActionBody) SetPropertyNamesIncluded(v []string) {
	o.PropertyNamesIncluded = v
}

func (o IFrameActionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if true {
		toSerialize["height"] = o.Height
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

type NullableIFrameActionBody struct {
	value *IFrameActionBody
	isSet bool
}

func (v NullableIFrameActionBody) Get() *IFrameActionBody {
	return v.value
}

func (v *NullableIFrameActionBody) Set(val *IFrameActionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIFrameActionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIFrameActionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIFrameActionBody(val *IFrameActionBody) *NullableIFrameActionBody {
	return &NullableIFrameActionBody{value: val, isSet: true}
}

func (v NullableIFrameActionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIFrameActionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}