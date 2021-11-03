/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
)

// CreateInvoiceFeature Indicates what elements of creating invoices in HubSpot is supported for the integration.
type CreateInvoiceFeature struct {
	// Indicates if creating invoices in HubSpot is supported for the integration.
	Enabled     bool                     `json:"enabled"`
	SubFeatures CreateInvoiceSubFeatures `json:"subFeatures"`
}

// NewCreateInvoiceFeature instantiates a new CreateInvoiceFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvoiceFeature(enabled bool, subFeatures CreateInvoiceSubFeatures) *CreateInvoiceFeature {
	this := CreateInvoiceFeature{}
	this.Enabled = enabled
	this.SubFeatures = subFeatures
	return &this
}

// NewCreateInvoiceFeatureWithDefaults instantiates a new CreateInvoiceFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvoiceFeatureWithDefaults() *CreateInvoiceFeature {
	this := CreateInvoiceFeature{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *CreateInvoiceFeature) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceFeature) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateInvoiceFeature) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSubFeatures returns the SubFeatures field value
func (o *CreateInvoiceFeature) GetSubFeatures() CreateInvoiceSubFeatures {
	if o == nil {
		var ret CreateInvoiceSubFeatures
		return ret
	}

	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceFeature) GetSubFeaturesOk() (*CreateInvoiceSubFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubFeatures, true
}

// SetSubFeatures sets field value
func (o *CreateInvoiceFeature) SetSubFeatures(v CreateInvoiceSubFeatures) {
	o.SubFeatures = v
}

func (o CreateInvoiceFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInvoiceFeature struct {
	value *CreateInvoiceFeature
	isSet bool
}

func (v NullableCreateInvoiceFeature) Get() *CreateInvoiceFeature {
	return v.value
}

func (v *NullableCreateInvoiceFeature) Set(val *CreateInvoiceFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceFeature(val *CreateInvoiceFeature) *NullableCreateInvoiceFeature {
	return &NullableCreateInvoiceFeature{value: val, isSet: true}
}

func (v NullableCreateInvoiceFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
