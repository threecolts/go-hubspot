/*
Transactional Email

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"encoding/json"
)

// checks if the SmtpApiTokenRequestEgg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpApiTokenRequestEgg{}

// SmtpApiTokenRequestEgg A request object to create a SMTP API token
type SmtpApiTokenRequestEgg struct {
	// Indicates whether a contact should be created for email recipients.
	CreateContact bool `json:"createContact"`
	// A name for the campaign tied to the SMTP API token.
	CampaignName string `json:"campaignName"`
}

// NewSmtpApiTokenRequestEgg instantiates a new SmtpApiTokenRequestEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpApiTokenRequestEgg(createContact bool, campaignName string) *SmtpApiTokenRequestEgg {
	this := SmtpApiTokenRequestEgg{}
	this.CreateContact = createContact
	this.CampaignName = campaignName
	return &this
}

// NewSmtpApiTokenRequestEggWithDefaults instantiates a new SmtpApiTokenRequestEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpApiTokenRequestEggWithDefaults() *SmtpApiTokenRequestEgg {
	this := SmtpApiTokenRequestEgg{}
	return &this
}

// GetCreateContact returns the CreateContact field value
func (o *SmtpApiTokenRequestEgg) GetCreateContact() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateContact
}

// GetCreateContactOk returns a tuple with the CreateContact field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenRequestEgg) GetCreateContactOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateContact, true
}

// SetCreateContact sets field value
func (o *SmtpApiTokenRequestEgg) SetCreateContact(v bool) {
	o.CreateContact = v
}

// GetCampaignName returns the CampaignName field value
func (o *SmtpApiTokenRequestEgg) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenRequestEgg) GetCampaignNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignName, true
}

// SetCampaignName sets field value
func (o *SmtpApiTokenRequestEgg) SetCampaignName(v string) {
	o.CampaignName = v
}

func (o SmtpApiTokenRequestEgg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpApiTokenRequestEgg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createContact"] = o.CreateContact
	toSerialize["campaignName"] = o.CampaignName
	return toSerialize, nil
}

type NullableSmtpApiTokenRequestEgg struct {
	value *SmtpApiTokenRequestEgg
	isSet bool
}

func (v NullableSmtpApiTokenRequestEgg) Get() *SmtpApiTokenRequestEgg {
	return v.value
}

func (v *NullableSmtpApiTokenRequestEgg) Set(val *SmtpApiTokenRequestEgg) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpApiTokenRequestEgg) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpApiTokenRequestEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpApiTokenRequestEgg(val *SmtpApiTokenRequestEgg) *NullableSmtpApiTokenRequestEgg {
	return &NullableSmtpApiTokenRequestEgg{value: val, isSet: true}
}

func (v NullableSmtpApiTokenRequestEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpApiTokenRequestEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
