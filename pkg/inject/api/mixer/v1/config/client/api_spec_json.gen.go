// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/config/client/api_spec.proto

package client

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/polarismesh/polaris-controller/pkg/inject/api/mixer/v1"
	_ "github.com/polarismesh/polaris-controller/pkg/inject/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for HTTPAPISpec
func (this *HTTPAPISpec) MarshalJSON() ([]byte, error) {
	str, err := ApiSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPAPISpec
func (this *HTTPAPISpec) UnmarshalJSON(b []byte) error {
	return ApiSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPAPISpecPattern
func (this *HTTPAPISpecPattern) MarshalJSON() ([]byte, error) {
	str, err := ApiSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPAPISpecPattern
func (this *HTTPAPISpecPattern) UnmarshalJSON(b []byte) error {
	return ApiSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for APIKey
func (this *APIKey) MarshalJSON() ([]byte, error) {
	str, err := ApiSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for APIKey
func (this *APIKey) UnmarshalJSON(b []byte) error {
	return ApiSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPAPISpecReference
func (this *HTTPAPISpecReference) MarshalJSON() ([]byte, error) {
	str, err := ApiSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPAPISpecReference
func (this *HTTPAPISpecReference) UnmarshalJSON(b []byte) error {
	return ApiSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPAPISpecBinding
func (this *HTTPAPISpecBinding) MarshalJSON() ([]byte, error) {
	str, err := ApiSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPAPISpecBinding
func (this *HTTPAPISpecBinding) UnmarshalJSON(b []byte) error {
	return ApiSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ApiSpecMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ApiSpecUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
