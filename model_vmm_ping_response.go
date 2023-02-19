/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VmmPingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmmPingResponse{}

// VmmPingResponse Virtual Machine Monitor information
type VmmPingResponse struct {
	Version string `json:"version"`
}

// NewVmmPingResponse instantiates a new VmmPingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmmPingResponse(version string) *VmmPingResponse {
	this := VmmPingResponse{}
	this.Version = version
	return &this
}

// NewVmmPingResponseWithDefaults instantiates a new VmmPingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmmPingResponseWithDefaults() *VmmPingResponse {
	this := VmmPingResponse{}
	return &this
}

// GetVersion returns the Version field value
func (o *VmmPingResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VmmPingResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VmmPingResponse) SetVersion(v string) {
	o.Version = v
}

func (o VmmPingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmmPingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableVmmPingResponse struct {
	value *VmmPingResponse
	isSet bool
}

func (v NullableVmmPingResponse) Get() *VmmPingResponse {
	return v.value
}

func (v *NullableVmmPingResponse) Set(val *VmmPingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVmmPingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVmmPingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmmPingResponse(val *VmmPingResponse) *NullableVmmPingResponse {
	return &NullableVmmPingResponse{value: val, isSet: true}
}

func (v NullableVmmPingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmmPingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


