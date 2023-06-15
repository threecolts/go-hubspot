/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// checks if the CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging{}

// CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging struct for CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging
type CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging struct {
	Results []MarketingEventExternalUniqueIdentifier `json:"results"`
}

// NewCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging instantiates a new CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging(results []MarketingEventExternalUniqueIdentifier) *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging {
	this := CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseMarketingEventExternalUniqueIdentifierNoPagingWithDefaults instantiates a new CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseMarketingEventExternalUniqueIdentifierNoPagingWithDefaults() *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging {
	this := CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) GetResults() []MarketingEventExternalUniqueIdentifier {
	if o == nil {
		var ret []MarketingEventExternalUniqueIdentifier
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) GetResultsOk() ([]MarketingEventExternalUniqueIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) SetResults(v []MarketingEventExternalUniqueIdentifier) {
	o.Results = v
}

func (o CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging struct {
	value *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging
	isSet bool
}

func (v NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) Get() *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging {
	return v.value
}

func (v *NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) Set(val *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging(val *CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) *NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging {
	return &NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseMarketingEventExternalUniqueIdentifierNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
