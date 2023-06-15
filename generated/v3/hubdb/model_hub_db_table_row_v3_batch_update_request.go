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

// checks if the HubDbTableRowV3BatchUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubDbTableRowV3BatchUpdateRequest{}

// HubDbTableRowV3BatchUpdateRequest struct for HubDbTableRowV3BatchUpdateRequest
type HubDbTableRowV3BatchUpdateRequest struct {
	// Specifies the value for `hs_path` column, which will be used as slug in the dynamic pages
	Path *string `json:"path,omitempty"`
	// Specifies the value for `hs_name` column, which will be used as title in the dynamic pages
	Name *string `json:"name,omitempty"`
	// Specifies the value for the column child table id
	ChildTableId *int32 `json:"childTableId,omitempty"`
	// List of key value pairs with the column name and column value
	Values       map[string]map[string]interface{} `json:"values"`
	DisplayIndex *int32                            `json:"displayIndex,omitempty"`
	// The id of the table row
	Id int64 `json:"id"`
}

// NewHubDbTableRowV3BatchUpdateRequest instantiates a new HubDbTableRowV3BatchUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubDbTableRowV3BatchUpdateRequest(values map[string]map[string]interface{}, id int64) *HubDbTableRowV3BatchUpdateRequest {
	this := HubDbTableRowV3BatchUpdateRequest{}
	this.Values = values
	this.Id = id
	return &this
}

// NewHubDbTableRowV3BatchUpdateRequestWithDefaults instantiates a new HubDbTableRowV3BatchUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubDbTableRowV3BatchUpdateRequestWithDefaults() *HubDbTableRowV3BatchUpdateRequest {
	this := HubDbTableRowV3BatchUpdateRequest{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HubDbTableRowV3BatchUpdateRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HubDbTableRowV3BatchUpdateRequest) SetPath(v string) {
	o.Path = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HubDbTableRowV3BatchUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HubDbTableRowV3BatchUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetChildTableId returns the ChildTableId field value if set, zero value otherwise.
func (o *HubDbTableRowV3BatchUpdateRequest) GetChildTableId() int32 {
	if o == nil || IsNil(o.ChildTableId) {
		var ret int32
		return ret
	}
	return *o.ChildTableId
}

// GetChildTableIdOk returns a tuple with the ChildTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetChildTableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChildTableId) {
		return nil, false
	}
	return o.ChildTableId, true
}

// HasChildTableId returns a boolean if a field has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) HasChildTableId() bool {
	if o != nil && !IsNil(o.ChildTableId) {
		return true
	}

	return false
}

// SetChildTableId gets a reference to the given int32 and assigns it to the ChildTableId field.
func (o *HubDbTableRowV3BatchUpdateRequest) SetChildTableId(v int32) {
	o.ChildTableId = &v
}

// GetValues returns the Values field value
func (o *HubDbTableRowV3BatchUpdateRequest) GetValues() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetValuesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *HubDbTableRowV3BatchUpdateRequest) SetValues(v map[string]map[string]interface{}) {
	o.Values = v
}

// GetDisplayIndex returns the DisplayIndex field value if set, zero value otherwise.
func (o *HubDbTableRowV3BatchUpdateRequest) GetDisplayIndex() int32 {
	if o == nil || IsNil(o.DisplayIndex) {
		var ret int32
		return ret
	}
	return *o.DisplayIndex
}

// GetDisplayIndexOk returns a tuple with the DisplayIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetDisplayIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayIndex) {
		return nil, false
	}
	return o.DisplayIndex, true
}

// HasDisplayIndex returns a boolean if a field has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) HasDisplayIndex() bool {
	if o != nil && !IsNil(o.DisplayIndex) {
		return true
	}

	return false
}

// SetDisplayIndex gets a reference to the given int32 and assigns it to the DisplayIndex field.
func (o *HubDbTableRowV3BatchUpdateRequest) SetDisplayIndex(v int32) {
	o.DisplayIndex = &v
}

// GetId returns the Id field value
func (o *HubDbTableRowV3BatchUpdateRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3BatchUpdateRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HubDbTableRowV3BatchUpdateRequest) SetId(v int64) {
	o.Id = v
}

func (o HubDbTableRowV3BatchUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubDbTableRowV3BatchUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ChildTableId) {
		toSerialize["childTableId"] = o.ChildTableId
	}
	toSerialize["values"] = o.Values
	if !IsNil(o.DisplayIndex) {
		toSerialize["displayIndex"] = o.DisplayIndex
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableHubDbTableRowV3BatchUpdateRequest struct {
	value *HubDbTableRowV3BatchUpdateRequest
	isSet bool
}

func (v NullableHubDbTableRowV3BatchUpdateRequest) Get() *HubDbTableRowV3BatchUpdateRequest {
	return v.value
}

func (v *NullableHubDbTableRowV3BatchUpdateRequest) Set(val *HubDbTableRowV3BatchUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHubDbTableRowV3BatchUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHubDbTableRowV3BatchUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubDbTableRowV3BatchUpdateRequest(val *HubDbTableRowV3BatchUpdateRequest) *NullableHubDbTableRowV3BatchUpdateRequest {
	return &NullableHubDbTableRowV3BatchUpdateRequest{value: val, isSet: true}
}

func (v NullableHubDbTableRowV3BatchUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubDbTableRowV3BatchUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
