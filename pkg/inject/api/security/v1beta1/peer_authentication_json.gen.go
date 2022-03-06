// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/peer_authentication.proto

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/polarismesh/polaris-controller/pkg/inject/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for PeerAuthentication
func (this *PeerAuthentication) MarshalJSON() ([]byte, error) {
	str, err := PeerAuthenticationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PeerAuthentication
func (this *PeerAuthentication) UnmarshalJSON(b []byte) error {
	return PeerAuthenticationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PeerAuthentication_MutualTLS
func (this *PeerAuthentication_MutualTLS) MarshalJSON() ([]byte, error) {
	str, err := PeerAuthenticationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PeerAuthentication_MutualTLS
func (this *PeerAuthentication_MutualTLS) UnmarshalJSON(b []byte) error {
	return PeerAuthenticationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	PeerAuthenticationMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	PeerAuthenticationUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
