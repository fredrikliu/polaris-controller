// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/request_authentication.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/polarismesh/polaris-controller/pkg/inject/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using RequestAuthentication within kubernetes types, where deepcopy-gen is used.
func (in *RequestAuthentication) DeepCopyInto(out *RequestAuthentication) {
	p := proto.Clone(in).(*RequestAuthentication)
	*out = *p
}
