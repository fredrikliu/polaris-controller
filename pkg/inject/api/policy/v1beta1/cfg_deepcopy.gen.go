// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/cfg.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/polarismesh/polaris-controller/pkg/inject/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using AttributeManifest within kubernetes types, where deepcopy-gen is used.
func (in *AttributeManifest) DeepCopyInto(out *AttributeManifest) {
	p := proto.Clone(in).(*AttributeManifest)
	*out = *p
}

// DeepCopyInto supports using Rule within kubernetes types, where deepcopy-gen is used.
func (in *Rule) DeepCopyInto(out *Rule) {
	p := proto.Clone(in).(*Rule)
	*out = *p
}

// DeepCopyInto supports using Instance within kubernetes types, where deepcopy-gen is used.
func (in *Instance) DeepCopyInto(out *Instance) {
	p := proto.Clone(in).(*Instance)
	*out = *p
}

// DeepCopyInto supports using Handler within kubernetes types, where deepcopy-gen is used.
func (in *Handler) DeepCopyInto(out *Handler) {
	p := proto.Clone(in).(*Handler)
	*out = *p
}
