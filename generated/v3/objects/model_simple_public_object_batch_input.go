/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// checks if the SimplePublicObjectBatchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectBatchInput{}

// SimplePublicObjectBatchInput struct for SimplePublicObjectBatchInput
type SimplePublicObjectBatchInput struct {
	Properties map[string]string `json:"properties"`
	Id         string            `json:"id"`
}

// NewSimplePublicObjectBatchInput instantiates a new SimplePublicObjectBatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectBatchInput(properties map[string]string, id string) *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	this.Properties = properties
	this.Id = id
	return &this
}

// NewSimplePublicObjectBatchInputWithDefaults instantiates a new SimplePublicObjectBatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectBatchInputWithDefaults() *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	return &this
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectBatchInput) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectBatchInput) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetId returns the Id field value
func (o *SimplePublicObjectBatchInput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectBatchInput) SetId(v string) {
	o.Id = v
}

func (o SimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectBatchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSimplePublicObjectBatchInput struct {
	value *SimplePublicObjectBatchInput
	isSet bool
}

func (v NullableSimplePublicObjectBatchInput) Get() *SimplePublicObjectBatchInput {
	return v.value
}

func (v *NullableSimplePublicObjectBatchInput) Set(val *SimplePublicObjectBatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectBatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectBatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectBatchInput(val *SimplePublicObjectBatchInput) *NullableSimplePublicObjectBatchInput {
	return &NullableSimplePublicObjectBatchInput{value: val, isSet: true}
}

func (v NullableSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectBatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
