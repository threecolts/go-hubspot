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

// checks if the BatchInputSimplePublicObjectInputForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputSimplePublicObjectInputForCreate{}

// BatchInputSimplePublicObjectInputForCreate struct for BatchInputSimplePublicObjectInputForCreate
type BatchInputSimplePublicObjectInputForCreate struct {
	Inputs []SimplePublicObjectInputForCreate `json:"inputs"`
}

// NewBatchInputSimplePublicObjectInputForCreate instantiates a new BatchInputSimplePublicObjectInputForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectInputForCreate(inputs []SimplePublicObjectInputForCreate) *BatchInputSimplePublicObjectInputForCreate {
	this := BatchInputSimplePublicObjectInputForCreate{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectInputForCreateWithDefaults instantiates a new BatchInputSimplePublicObjectInputForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectInputForCreateWithDefaults() *BatchInputSimplePublicObjectInputForCreate {
	this := BatchInputSimplePublicObjectInputForCreate{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectInputForCreate) GetInputs() []SimplePublicObjectInputForCreate {
	if o == nil {
		var ret []SimplePublicObjectInputForCreate
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectInputForCreate) GetInputsOk() ([]SimplePublicObjectInputForCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectInputForCreate) SetInputs(v []SimplePublicObjectInputForCreate) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputSimplePublicObjectInputForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputSimplePublicObjectInputForCreate struct {
	value *BatchInputSimplePublicObjectInputForCreate
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) Get() *BatchInputSimplePublicObjectInputForCreate {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) Set(val *BatchInputSimplePublicObjectInputForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectInputForCreate(val *BatchInputSimplePublicObjectInputForCreate) *NullableBatchInputSimplePublicObjectInputForCreate {
	return &NullableBatchInputSimplePublicObjectInputForCreate{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
