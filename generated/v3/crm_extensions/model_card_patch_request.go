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

// checks if the CardPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardPatchRequest{}

// CardPatchRequest Body for a patch with optional fields
type CardPatchRequest struct {
	// The top-level title for this card. Displayed to users in the CRM UI.
	Title   *string             `json:"title,omitempty"`
	Fetch   *CardFetchBodyPatch `json:"fetch,omitempty"`
	Display *CardDisplayBody    `json:"display,omitempty"`
	Actions *CardActions        `json:"actions,omitempty"`
}

// NewCardPatchRequest instantiates a new CardPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardPatchRequest() *CardPatchRequest {
	this := CardPatchRequest{}
	return &this
}

// NewCardPatchRequestWithDefaults instantiates a new CardPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardPatchRequestWithDefaults() *CardPatchRequest {
	this := CardPatchRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CardPatchRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardPatchRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CardPatchRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CardPatchRequest) SetTitle(v string) {
	o.Title = &v
}

// GetFetch returns the Fetch field value if set, zero value otherwise.
func (o *CardPatchRequest) GetFetch() CardFetchBodyPatch {
	if o == nil || IsNil(o.Fetch) {
		var ret CardFetchBodyPatch
		return ret
	}
	return *o.Fetch
}

// GetFetchOk returns a tuple with the Fetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardPatchRequest) GetFetchOk() (*CardFetchBodyPatch, bool) {
	if o == nil || IsNil(o.Fetch) {
		return nil, false
	}
	return o.Fetch, true
}

// HasFetch returns a boolean if a field has been set.
func (o *CardPatchRequest) HasFetch() bool {
	if o != nil && !IsNil(o.Fetch) {
		return true
	}

	return false
}

// SetFetch gets a reference to the given CardFetchBodyPatch and assigns it to the Fetch field.
func (o *CardPatchRequest) SetFetch(v CardFetchBodyPatch) {
	o.Fetch = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *CardPatchRequest) GetDisplay() CardDisplayBody {
	if o == nil || IsNil(o.Display) {
		var ret CardDisplayBody
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardPatchRequest) GetDisplayOk() (*CardDisplayBody, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *CardPatchRequest) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given CardDisplayBody and assigns it to the Display field.
func (o *CardPatchRequest) SetDisplay(v CardDisplayBody) {
	o.Display = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CardPatchRequest) GetActions() CardActions {
	if o == nil || IsNil(o.Actions) {
		var ret CardActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardPatchRequest) GetActionsOk() (*CardActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CardPatchRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given CardActions and assigns it to the Actions field.
func (o *CardPatchRequest) SetActions(v CardActions) {
	o.Actions = &v
}

func (o CardPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Fetch) {
		toSerialize["fetch"] = o.Fetch
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	return toSerialize, nil
}

type NullableCardPatchRequest struct {
	value *CardPatchRequest
	isSet bool
}

func (v NullableCardPatchRequest) Get() *CardPatchRequest {
	return v.value
}

func (v *NullableCardPatchRequest) Set(val *CardPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardPatchRequest(val *CardPatchRequest) *NullableCardPatchRequest {
	return &NullableCardPatchRequest{value: val, isSet: true}
}

func (v NullableCardPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
