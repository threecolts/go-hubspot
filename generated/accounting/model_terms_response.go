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

// TermsResponse A response to a search for payment terms.
type TermsResponse struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result *string `json:"@result,omitempty"`
	// The list of payment terms that matched the search criteria.
	Terms []AccountingExtensionTerm `json:"terms"`
}

// NewTermsResponse instantiates a new TermsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermsResponse(terms []AccountingExtensionTerm) *TermsResponse {
	this := TermsResponse{}
	this.Terms = terms
	return &this
}

// NewTermsResponseWithDefaults instantiates a new TermsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsResponseWithDefaults() *TermsResponse {
	this := TermsResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TermsResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TermsResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *TermsResponse) SetResult(v string) {
	o.Result = &v
}

// GetTerms returns the Terms field value
func (o *TermsResponse) GetTerms() []AccountingExtensionTerm {
	if o == nil {
		var ret []AccountingExtensionTerm
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *TermsResponse) GetTermsOk() (*[]AccountingExtensionTerm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terms, true
}

// SetTerms sets field value
func (o *TermsResponse) SetTerms(v []AccountingExtensionTerm) {
	o.Terms = v
}

func (o TermsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["@result"] = o.Result
	}
	if true {
		toSerialize["terms"] = o.Terms
	}
	return json.Marshal(toSerialize)
}

type NullableTermsResponse struct {
	value *TermsResponse
	isSet bool
}

func (v NullableTermsResponse) Get() *TermsResponse {
	return v.value
}

func (v *NullableTermsResponse) Set(val *TermsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTermsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTermsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermsResponse(val *TermsResponse) *NullableTermsResponse {
	return &NullableTermsResponse{value: val, isSet: true}
}

func (v NullableTermsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
