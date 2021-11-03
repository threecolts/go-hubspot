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

// PublicSingleSendEmail A JSON object containing anything you want to override.
type PublicSingleSendEmail struct {
	// The From header for the email.
	From *string `json:"from,omitempty"`
	// The recipient of the email.
	To *string `json:"to,omitempty"`
	// ID for a particular send. No more than one email will be sent per sendId.
	SendId *string `json:"sendId,omitempty"`
	// List of Reply-To header values for the email.
	ReplyTo []string `json:"replyTo"`
	// List of email addresses to send as Cc.
	Cc []string `json:"cc"`
	// List of email addresses to send as Bcc.
	Bcc []string `json:"bcc"`
}

// NewPublicSingleSendEmail instantiates a new PublicSingleSendEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSingleSendEmail(replyTo []string, cc []string, bcc []string) *PublicSingleSendEmail {
	this := PublicSingleSendEmail{}
	this.ReplyTo = replyTo
	this.Cc = cc
	this.Bcc = bcc
	return &this
}

// NewPublicSingleSendEmailWithDefaults instantiates a new PublicSingleSendEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSingleSendEmailWithDefaults() *PublicSingleSendEmail {
	this := PublicSingleSendEmail{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PublicSingleSendEmail) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *PublicSingleSendEmail) SetTo(v string) {
	o.To = &v
}

// GetSendId returns the SendId field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetSendId() string {
	if o == nil || o.SendId == nil {
		var ret string
		return ret
	}
	return *o.SendId
}

// GetSendIdOk returns a tuple with the SendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetSendIdOk() (*string, bool) {
	if o == nil || o.SendId == nil {
		return nil, false
	}
	return o.SendId, true
}

// HasSendId returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasSendId() bool {
	if o != nil && o.SendId != nil {
		return true
	}

	return false
}

// SetSendId gets a reference to the given string and assigns it to the SendId field.
func (o *PublicSingleSendEmail) SetSendId(v string) {
	o.SendId = &v
}

// GetReplyTo returns the ReplyTo field value
func (o *PublicSingleSendEmail) GetReplyTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetReplyToOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplyTo, true
}

// SetReplyTo sets field value
func (o *PublicSingleSendEmail) SetReplyTo(v []string) {
	o.ReplyTo = v
}

// GetCc returns the Cc field value
func (o *PublicSingleSendEmail) GetCc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cc
}

// GetCcOk returns a tuple with the Cc field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetCcOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cc, true
}

// SetCc sets field value
func (o *PublicSingleSendEmail) SetCc(v []string) {
	o.Cc = v
}

// GetBcc returns the Bcc field value
func (o *PublicSingleSendEmail) GetBcc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetBccOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bcc, true
}

// SetBcc sets field value
func (o *PublicSingleSendEmail) SetBcc(v []string) {
	o.Bcc = v
}

func (o PublicSingleSendEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.SendId != nil {
		toSerialize["sendId"] = o.SendId
	}
	if true {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if true {
		toSerialize["cc"] = o.Cc
	}
	if true {
		toSerialize["bcc"] = o.Bcc
	}
	return json.Marshal(toSerialize)
}

type NullablePublicSingleSendEmail struct {
	value *PublicSingleSendEmail
	isSet bool
}

func (v NullablePublicSingleSendEmail) Get() *PublicSingleSendEmail {
	return v.value
}

func (v *NullablePublicSingleSendEmail) Set(val *PublicSingleSendEmail) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSingleSendEmail) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSingleSendEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSingleSendEmail(val *PublicSingleSendEmail) *NullablePublicSingleSendEmail {
	return &NullablePublicSingleSendEmail{value: val, isSet: true}
}

func (v NullablePublicSingleSendEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSingleSendEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
