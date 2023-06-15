/*
Webhooks API

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
)

// checks if the SubscriptionListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionListResponse{}

// SubscriptionListResponse List of event subscriptions for your app
type SubscriptionListResponse struct {
	// List of event subscriptions for your app
	Results []SubscriptionResponse `json:"results"`
}

// NewSubscriptionListResponse instantiates a new SubscriptionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionListResponse(results []SubscriptionResponse) *SubscriptionListResponse {
	this := SubscriptionListResponse{}
	this.Results = results
	return &this
}

// NewSubscriptionListResponseWithDefaults instantiates a new SubscriptionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionListResponseWithDefaults() *SubscriptionListResponse {
	this := SubscriptionListResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *SubscriptionListResponse) GetResults() []SubscriptionResponse {
	if o == nil {
		var ret []SubscriptionResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SubscriptionListResponse) GetResultsOk() ([]SubscriptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *SubscriptionListResponse) SetResults(v []SubscriptionResponse) {
	o.Results = v
}

func (o SubscriptionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableSubscriptionListResponse struct {
	value *SubscriptionListResponse
	isSet bool
}

func (v NullableSubscriptionListResponse) Get() *SubscriptionListResponse {
	return v.value
}

func (v *NullableSubscriptionListResponse) Set(val *SubscriptionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionListResponse(val *SubscriptionListResponse) *NullableSubscriptionListResponse {
	return &NullableSubscriptionListResponse{value: val, isSet: true}
}

func (v NullableSubscriptionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
