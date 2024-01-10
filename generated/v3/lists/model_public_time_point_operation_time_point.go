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

// PublicTimePointOperationTimePoint - struct for PublicTimePointOperationTimePoint
type PublicTimePointOperationTimePoint struct {
	PublicDatePoint              *PublicDatePoint
	PublicIndexedTimePoint       *PublicIndexedTimePoint
	PublicPropertyReferencedTime *PublicPropertyReferencedTime
}

// PublicDatePointAsPublicTimePointOperationTimePoint is a convenience function that returns PublicDatePoint wrapped in PublicTimePointOperationTimePoint
func PublicDatePointAsPublicTimePointOperationTimePoint(v *PublicDatePoint) PublicTimePointOperationTimePoint {
	return PublicTimePointOperationTimePoint{
		PublicDatePoint: v,
	}
}

// PublicIndexedTimePointAsPublicTimePointOperationTimePoint is a convenience function that returns PublicIndexedTimePoint wrapped in PublicTimePointOperationTimePoint
func PublicIndexedTimePointAsPublicTimePointOperationTimePoint(v *PublicIndexedTimePoint) PublicTimePointOperationTimePoint {
	return PublicTimePointOperationTimePoint{
		PublicIndexedTimePoint: v,
	}
}

// PublicPropertyReferencedTimeAsPublicTimePointOperationTimePoint is a convenience function that returns PublicPropertyReferencedTime wrapped in PublicTimePointOperationTimePoint
func PublicPropertyReferencedTimeAsPublicTimePointOperationTimePoint(v *PublicPropertyReferencedTime) PublicTimePointOperationTimePoint {
	return PublicTimePointOperationTimePoint{
		PublicPropertyReferencedTime: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PublicTimePointOperationTimePoint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PublicDatePoint
	err = json.Unmarshal(data, &dst.PublicDatePoint)
	if err == nil {
		jsonPublicDatePoint, _ := json.Marshal(dst.PublicDatePoint)
		if string(jsonPublicDatePoint) == "{}" { // empty struct
			dst.PublicDatePoint = nil
		} else {
			match++
		}
	} else {
		dst.PublicDatePoint = nil
	}

	// try to unmarshal data into PublicIndexedTimePoint
	err = json.Unmarshal(data, &dst.PublicIndexedTimePoint)
	if err == nil {
		jsonPublicIndexedTimePoint, _ := json.Marshal(dst.PublicIndexedTimePoint)
		if string(jsonPublicIndexedTimePoint) == "{}" { // empty struct
			dst.PublicIndexedTimePoint = nil
		} else {
			match++
		}
	} else {
		dst.PublicIndexedTimePoint = nil
	}

	// try to unmarshal data into PublicPropertyReferencedTime
	err = json.Unmarshal(data, &dst.PublicPropertyReferencedTime)
	if err == nil {
		jsonPublicPropertyReferencedTime, _ := json.Marshal(dst.PublicPropertyReferencedTime)
		if string(jsonPublicPropertyReferencedTime) == "{}" { // empty struct
			dst.PublicPropertyReferencedTime = nil
		} else {
			match++
		}
	} else {
		dst.PublicPropertyReferencedTime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PublicDatePoint = nil
		dst.PublicIndexedTimePoint = nil
		dst.PublicPropertyReferencedTime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PublicTimePointOperationTimePoint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PublicTimePointOperationTimePoint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PublicTimePointOperationTimePoint) MarshalJSON() ([]byte, error) {
	if src.PublicDatePoint != nil {
		return json.Marshal(&src.PublicDatePoint)
	}

	if src.PublicIndexedTimePoint != nil {
		return json.Marshal(&src.PublicIndexedTimePoint)
	}

	if src.PublicPropertyReferencedTime != nil {
		return json.Marshal(&src.PublicPropertyReferencedTime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PublicTimePointOperationTimePoint) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PublicDatePoint != nil {
		return obj.PublicDatePoint
	}

	if obj.PublicIndexedTimePoint != nil {
		return obj.PublicIndexedTimePoint
	}

	if obj.PublicPropertyReferencedTime != nil {
		return obj.PublicPropertyReferencedTime
	}

	// all schemas are nil
	return nil
}

type NullablePublicTimePointOperationTimePoint struct {
	value *PublicTimePointOperationTimePoint
	isSet bool
}

func (v NullablePublicTimePointOperationTimePoint) Get() *PublicTimePointOperationTimePoint {
	return v.value
}

func (v *NullablePublicTimePointOperationTimePoint) Set(val *PublicTimePointOperationTimePoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicTimePointOperationTimePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicTimePointOperationTimePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicTimePointOperationTimePoint(val *PublicTimePointOperationTimePoint) *NullablePublicTimePointOperationTimePoint {
	return &NullablePublicTimePointOperationTimePoint{value: val, isSet: true}
}

func (v NullablePublicTimePointOperationTimePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicTimePointOperationTimePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
