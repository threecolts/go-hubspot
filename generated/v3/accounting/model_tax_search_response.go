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

// checks if the TaxSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxSearchResponse{}

// TaxSearchResponse A response to a search for taxes.
type TaxSearchResponse struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result *string `json:"@result,omitempty"`
	// The list of taxes that matched the search criteria
	Taxes []Tax `json:"taxes"`
}

// NewTaxSearchResponse instantiates a new TaxSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxSearchResponse(taxes []Tax) *TaxSearchResponse {
	this := TaxSearchResponse{}
	this.Taxes = taxes
	return &this
}

// NewTaxSearchResponseWithDefaults instantiates a new TaxSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxSearchResponseWithDefaults() *TaxSearchResponse {
	this := TaxSearchResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TaxSearchResponse) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxSearchResponse) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TaxSearchResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *TaxSearchResponse) SetResult(v string) {
	o.Result = &v
}

// GetTaxes returns the Taxes field value
func (o *TaxSearchResponse) GetTaxes() []Tax {
	if o == nil {
		var ret []Tax
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
func (o *TaxSearchResponse) GetTaxesOk() ([]Tax, bool) {
	if o == nil {
		return nil, false
	}
	return o.Taxes, true
}

// SetTaxes sets field value
func (o *TaxSearchResponse) SetTaxes(v []Tax) {
	o.Taxes = v
}

func (o TaxSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["@result"] = o.Result
	}
	toSerialize["taxes"] = o.Taxes
	return toSerialize, nil
}

type NullableTaxSearchResponse struct {
	value *TaxSearchResponse
	isSet bool
}

func (v NullableTaxSearchResponse) Get() *TaxSearchResponse {
	return v.value
}

func (v *NullableTaxSearchResponse) Set(val *TaxSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxSearchResponse(val *TaxSearchResponse) *NullableTaxSearchResponse {
	return &NullableTaxSearchResponse{value: val, isSet: true}
}

func (v NullableTaxSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
