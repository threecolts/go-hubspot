/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"fmt"
)

// PublicPropertyAssociationFilterBranchFilterBranchesInner - struct for PublicPropertyAssociationFilterBranchFilterBranchesInner
type PublicPropertyAssociationFilterBranchFilterBranchesInner struct {
	PublicAndFilterBranch                 *PublicAndFilterBranch
	PublicAssociationFilterBranch         *PublicAssociationFilterBranch
	PublicNotAllFilterBranch              *PublicNotAllFilterBranch
	PublicNotAnyFilterBranch              *PublicNotAnyFilterBranch
	PublicOrFilterBranch                  *PublicOrFilterBranch
	PublicPropertyAssociationFilterBranch *PublicPropertyAssociationFilterBranch
	PublicRestrictedFilterBranch          *PublicRestrictedFilterBranch
	PublicUnifiedEventsFilterBranch       *PublicUnifiedEventsFilterBranch
}

// PublicAndFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicAndFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicAndFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicAndFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicAndFilterBranch: v,
	}
}

// PublicAssociationFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicAssociationFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicAssociationFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicAssociationFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicAssociationFilterBranch: v,
	}
}

// PublicNotAllFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicNotAllFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicNotAllFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicNotAllFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicNotAllFilterBranch: v,
	}
}

// PublicNotAnyFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicNotAnyFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicNotAnyFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicNotAnyFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicNotAnyFilterBranch: v,
	}
}

// PublicOrFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicOrFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicOrFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicOrFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicOrFilterBranch: v,
	}
}

// PublicPropertyAssociationFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicPropertyAssociationFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicPropertyAssociationFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicPropertyAssociationFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicPropertyAssociationFilterBranch: v,
	}
}

// PublicRestrictedFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicRestrictedFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicRestrictedFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicRestrictedFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicRestrictedFilterBranch: v,
	}
}

// PublicUnifiedEventsFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner is a convenience function that returns PublicUnifiedEventsFilterBranch wrapped in PublicPropertyAssociationFilterBranchFilterBranchesInner
func PublicUnifiedEventsFilterBranchAsPublicPropertyAssociationFilterBranchFilterBranchesInner(v *PublicUnifiedEventsFilterBranch) PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return PublicPropertyAssociationFilterBranchFilterBranchesInner{
		PublicUnifiedEventsFilterBranch: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PublicPropertyAssociationFilterBranchFilterBranchesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PublicAndFilterBranch
	err = json.Unmarshal(data, &dst.PublicAndFilterBranch)
	if err == nil {
		jsonPublicAndFilterBranch, _ := json.Marshal(dst.PublicAndFilterBranch)
		if string(jsonPublicAndFilterBranch) == "{}" { // empty struct
			dst.PublicAndFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicAndFilterBranch = nil
	}

	// try to unmarshal data into PublicAssociationFilterBranch
	err = json.Unmarshal(data, &dst.PublicAssociationFilterBranch)
	if err == nil {
		jsonPublicAssociationFilterBranch, _ := json.Marshal(dst.PublicAssociationFilterBranch)
		if string(jsonPublicAssociationFilterBranch) == "{}" { // empty struct
			dst.PublicAssociationFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicAssociationFilterBranch = nil
	}

	// try to unmarshal data into PublicNotAllFilterBranch
	err = json.Unmarshal(data, &dst.PublicNotAllFilterBranch)
	if err == nil {
		jsonPublicNotAllFilterBranch, _ := json.Marshal(dst.PublicNotAllFilterBranch)
		if string(jsonPublicNotAllFilterBranch) == "{}" { // empty struct
			dst.PublicNotAllFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicNotAllFilterBranch = nil
	}

	// try to unmarshal data into PublicNotAnyFilterBranch
	err = json.Unmarshal(data, &dst.PublicNotAnyFilterBranch)
	if err == nil {
		jsonPublicNotAnyFilterBranch, _ := json.Marshal(dst.PublicNotAnyFilterBranch)
		if string(jsonPublicNotAnyFilterBranch) == "{}" { // empty struct
			dst.PublicNotAnyFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicNotAnyFilterBranch = nil
	}

	// try to unmarshal data into PublicOrFilterBranch
	err = json.Unmarshal(data, &dst.PublicOrFilterBranch)
	if err == nil {
		jsonPublicOrFilterBranch, _ := json.Marshal(dst.PublicOrFilterBranch)
		if string(jsonPublicOrFilterBranch) == "{}" { // empty struct
			dst.PublicOrFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicOrFilterBranch = nil
	}

	// try to unmarshal data into PublicPropertyAssociationFilterBranch
	err = json.Unmarshal(data, &dst.PublicPropertyAssociationFilterBranch)
	if err == nil {
		jsonPublicPropertyAssociationFilterBranch, _ := json.Marshal(dst.PublicPropertyAssociationFilterBranch)
		if string(jsonPublicPropertyAssociationFilterBranch) == "{}" { // empty struct
			dst.PublicPropertyAssociationFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicPropertyAssociationFilterBranch = nil
	}

	// try to unmarshal data into PublicRestrictedFilterBranch
	err = json.Unmarshal(data, &dst.PublicRestrictedFilterBranch)
	if err == nil {
		jsonPublicRestrictedFilterBranch, _ := json.Marshal(dst.PublicRestrictedFilterBranch)
		if string(jsonPublicRestrictedFilterBranch) == "{}" { // empty struct
			dst.PublicRestrictedFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicRestrictedFilterBranch = nil
	}

	// try to unmarshal data into PublicUnifiedEventsFilterBranch
	err = json.Unmarshal(data, &dst.PublicUnifiedEventsFilterBranch)
	if err == nil {
		jsonPublicUnifiedEventsFilterBranch, _ := json.Marshal(dst.PublicUnifiedEventsFilterBranch)
		if string(jsonPublicUnifiedEventsFilterBranch) == "{}" { // empty struct
			dst.PublicUnifiedEventsFilterBranch = nil
		} else {
			match++
		}
	} else {
		dst.PublicUnifiedEventsFilterBranch = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PublicAndFilterBranch = nil
		dst.PublicAssociationFilterBranch = nil
		dst.PublicNotAllFilterBranch = nil
		dst.PublicNotAnyFilterBranch = nil
		dst.PublicOrFilterBranch = nil
		dst.PublicPropertyAssociationFilterBranch = nil
		dst.PublicRestrictedFilterBranch = nil
		dst.PublicUnifiedEventsFilterBranch = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PublicPropertyAssociationFilterBranchFilterBranchesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PublicPropertyAssociationFilterBranchFilterBranchesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PublicPropertyAssociationFilterBranchFilterBranchesInner) MarshalJSON() ([]byte, error) {
	if src.PublicAndFilterBranch != nil {
		return json.Marshal(&src.PublicAndFilterBranch)
	}

	if src.PublicAssociationFilterBranch != nil {
		return json.Marshal(&src.PublicAssociationFilterBranch)
	}

	if src.PublicNotAllFilterBranch != nil {
		return json.Marshal(&src.PublicNotAllFilterBranch)
	}

	if src.PublicNotAnyFilterBranch != nil {
		return json.Marshal(&src.PublicNotAnyFilterBranch)
	}

	if src.PublicOrFilterBranch != nil {
		return json.Marshal(&src.PublicOrFilterBranch)
	}

	if src.PublicPropertyAssociationFilterBranch != nil {
		return json.Marshal(&src.PublicPropertyAssociationFilterBranch)
	}

	if src.PublicRestrictedFilterBranch != nil {
		return json.Marshal(&src.PublicRestrictedFilterBranch)
	}

	if src.PublicUnifiedEventsFilterBranch != nil {
		return json.Marshal(&src.PublicUnifiedEventsFilterBranch)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PublicAndFilterBranch != nil {
		return obj.PublicAndFilterBranch
	}

	if obj.PublicAssociationFilterBranch != nil {
		return obj.PublicAssociationFilterBranch
	}

	if obj.PublicNotAllFilterBranch != nil {
		return obj.PublicNotAllFilterBranch
	}

	if obj.PublicNotAnyFilterBranch != nil {
		return obj.PublicNotAnyFilterBranch
	}

	if obj.PublicOrFilterBranch != nil {
		return obj.PublicOrFilterBranch
	}

	if obj.PublicPropertyAssociationFilterBranch != nil {
		return obj.PublicPropertyAssociationFilterBranch
	}

	if obj.PublicRestrictedFilterBranch != nil {
		return obj.PublicRestrictedFilterBranch
	}

	if obj.PublicUnifiedEventsFilterBranch != nil {
		return obj.PublicUnifiedEventsFilterBranch
	}

	// all schemas are nil
	return nil
}

type NullablePublicPropertyAssociationFilterBranchFilterBranchesInner struct {
	value *PublicPropertyAssociationFilterBranchFilterBranchesInner
	isSet bool
}

func (v NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) Get() *PublicPropertyAssociationFilterBranchFilterBranchesInner {
	return v.value
}

func (v *NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) Set(val *PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPropertyAssociationFilterBranchFilterBranchesInner(val *PublicPropertyAssociationFilterBranchFilterBranchesInner) *NullablePublicPropertyAssociationFilterBranchFilterBranchesInner {
	return &NullablePublicPropertyAssociationFilterBranchFilterBranchesInner{value: val, isSet: true}
}

func (v NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPropertyAssociationFilterBranchFilterBranchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}