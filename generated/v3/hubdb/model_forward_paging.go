/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
)

// checks if the ForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardPaging{}

// ForwardPaging struct for ForwardPaging
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NewForwardPaging instantiates a new ForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardPaging() *ForwardPaging {
	this := ForwardPaging{}
	return &this
}

// NewForwardPagingWithDefaults instantiates a new ForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardPagingWithDefaults() *ForwardPaging {
	this := ForwardPaging{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ForwardPaging) GetNext() NextPage {
	if o == nil || IsNil(o.Next) {
		var ret NextPage
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardPaging) GetNextOk() (*NextPage, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ForwardPaging) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextPage and assigns it to the Next field.
func (o *ForwardPaging) SetNext(v NextPage) {
	o.Next = &v
}

func (o ForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableForwardPaging struct {
	value *ForwardPaging
	isSet bool
}

func (v NullableForwardPaging) Get() *ForwardPaging {
	return v.value
}

func (v *NullableForwardPaging) Set(val *ForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardPaging(val *ForwardPaging) *NullableForwardPaging {
	return &NullableForwardPaging{value: val, isSet: true}
}

func (v NullableForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
