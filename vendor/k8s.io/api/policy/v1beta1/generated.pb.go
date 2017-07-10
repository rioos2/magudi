/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: gitlab.com/sankish/magudi/vendor/k8s.io/api/policy/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		gitlab.com/sankish/magudi/vendor/k8s.io/api/policy/v1beta1/generated.proto

	It has these top-level messages:
		Eviction
		PodDisruptionBudget
		PodDisruptionBudgetList
		PodDisruptionBudgetSpec
		PodDisruptionBudgetStatus
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apimachinery_pkg_util_intstr "k8s.io/apimachinery/pkg/util/intstr"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Eviction) Reset()                    { *m = Eviction{} }
func (*Eviction) ProtoMessage()               {}
func (*Eviction) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PodDisruptionBudget) Reset()                    { *m = PodDisruptionBudget{} }
func (*PodDisruptionBudget) ProtoMessage()               {}
func (*PodDisruptionBudget) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodDisruptionBudgetList) Reset()                    { *m = PodDisruptionBudgetList{} }
func (*PodDisruptionBudgetList) ProtoMessage()               {}
func (*PodDisruptionBudgetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodDisruptionBudgetSpec) Reset()                    { *m = PodDisruptionBudgetSpec{} }
func (*PodDisruptionBudgetSpec) ProtoMessage()               {}
func (*PodDisruptionBudgetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *PodDisruptionBudgetStatus) Reset()      { *m = PodDisruptionBudgetStatus{} }
func (*PodDisruptionBudgetStatus) ProtoMessage() {}
func (*PodDisruptionBudgetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{4}
}

func init() {
	proto.RegisterType((*Eviction)(nil), "k8s.io.api.policy.v1beta1.Eviction")
	proto.RegisterType((*PodDisruptionBudget)(nil), "k8s.io.api.policy.v1beta1.PodDisruptionBudget")
	proto.RegisterType((*PodDisruptionBudgetList)(nil), "k8s.io.api.policy.v1beta1.PodDisruptionBudgetList")
	proto.RegisterType((*PodDisruptionBudgetSpec)(nil), "k8s.io.api.policy.v1beta1.PodDisruptionBudgetSpec")
	proto.RegisterType((*PodDisruptionBudgetStatus)(nil), "k8s.io.api.policy.v1beta1.PodDisruptionBudgetStatus")
}
func (m *Eviction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Eviction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.DeleteOptions != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.DeleteOptions.Size()))
		n2, err := m.DeleteOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PodDisruptionBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudget) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n3, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n4, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n5, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *PodDisruptionBudgetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n6, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PodDisruptionBudgetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MinAvailable != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.MinAvailable.Size()))
		n7, err := m.MinAvailable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Selector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n8, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.MaxUnavailable != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.MaxUnavailable.Size()))
		n9, err := m.MaxUnavailable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func (m *PodDisruptionBudgetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObservedGeneration))
	if len(m.DisruptedPods) > 0 {
		keysForDisruptedPods := make([]string, 0, len(m.DisruptedPods))
		for k := range m.DisruptedPods {
			keysForDisruptedPods = append(keysForDisruptedPods, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForDisruptedPods)
		for _, k := range keysForDisruptedPods {
			dAtA[i] = 0x12
			i++
			v := m.DisruptedPods[string(k)]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovGenerated(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + msgSize
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64((&v).Size()))
			n10, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n10
		}
	}
	dAtA[i] = 0x18
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.PodDisruptionsAllowed))
	dAtA[i] = 0x20
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.CurrentHealthy))
	dAtA[i] = 0x28
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.DesiredHealthy))
	dAtA[i] = 0x30
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ExpectedPods))
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Eviction) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.DeleteOptions != nil {
		l = m.DeleteOptions.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PodDisruptionBudget) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PodDisruptionBudgetList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PodDisruptionBudgetSpec) Size() (n int) {
	var l int
	_ = l
	if m.MinAvailable != nil {
		l = m.MinAvailable.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.MaxUnavailable != nil {
		l = m.MaxUnavailable.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PodDisruptionBudgetStatus) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.ObservedGeneration))
	if len(m.DisruptedPods) > 0 {
		for k, v := range m.DisruptedPods {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	n += 1 + sovGenerated(uint64(m.PodDisruptionsAllowed))
	n += 1 + sovGenerated(uint64(m.CurrentHealthy))
	n += 1 + sovGenerated(uint64(m.DesiredHealthy))
	n += 1 + sovGenerated(uint64(m.ExpectedPods))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Eviction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Eviction{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`DeleteOptions:` + strings.Replace(fmt.Sprintf("%v", this.DeleteOptions), "DeleteOptions", "k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudget) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudget{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "PodDisruptionBudgetSpec", "PodDisruptionBudgetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "PodDisruptionBudgetStatus", "PodDisruptionBudgetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudgetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PodDisruptionBudget", "PodDisruptionBudget", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodDisruptionBudgetSpec{`,
		`MinAvailable:` + strings.Replace(fmt.Sprintf("%v", this.MinAvailable), "IntOrString", "k8s_io_apimachinery_pkg_util_intstr.IntOrString", 1) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1) + `,`,
		`MaxUnavailable:` + strings.Replace(fmt.Sprintf("%v", this.MaxUnavailable), "IntOrString", "k8s_io_apimachinery_pkg_util_intstr.IntOrString", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodDisruptionBudgetStatus) String() string {
	if this == nil {
		return "nil"
	}
	keysForDisruptedPods := make([]string, 0, len(this.DisruptedPods))
	for k := range this.DisruptedPods {
		keysForDisruptedPods = append(keysForDisruptedPods, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDisruptedPods)
	mapStringForDisruptedPods := "map[string]k8s_io_apimachinery_pkg_apis_meta_v1.Time{"
	for _, k := range keysForDisruptedPods {
		mapStringForDisruptedPods += fmt.Sprintf("%v: %v,", k, this.DisruptedPods[k])
	}
	mapStringForDisruptedPods += "}"
	s := strings.Join([]string{`&PodDisruptionBudgetStatus{`,
		`ObservedGeneration:` + fmt.Sprintf("%v", this.ObservedGeneration) + `,`,
		`DisruptedPods:` + mapStringForDisruptedPods + `,`,
		`PodDisruptionsAllowed:` + fmt.Sprintf("%v", this.PodDisruptionsAllowed) + `,`,
		`CurrentHealthy:` + fmt.Sprintf("%v", this.CurrentHealthy) + `,`,
		`DesiredHealthy:` + fmt.Sprintf("%v", this.DesiredHealthy) + `,`,
		`ExpectedPods:` + fmt.Sprintf("%v", this.ExpectedPods) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Eviction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Eviction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Eviction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteOptions == nil {
				m.DeleteOptions = &k8s_io_apimachinery_pkg_apis_meta_v1.DeleteOptions{}
			}
			if err := m.DeleteOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PodDisruptionBudget{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAvailable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinAvailable == nil {
				m.MinAvailable = &k8s_io_apimachinery_pkg_util_intstr.IntOrString{}
			}
			if err := m.MinAvailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxUnavailable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxUnavailable == nil {
				m.MaxUnavailable = &k8s_io_apimachinery_pkg_util_intstr.IntOrString{}
			}
			if err := m.MaxUnavailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			m.ObservedGeneration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObservedGeneration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisruptedPods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.DisruptedPods == nil {
				m.DisruptedPods = make(map[string]k8s_io_apimachinery_pkg_apis_meta_v1.Time)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &k8s_io_apimachinery_pkg_apis_meta_v1.Time{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.DisruptedPods[mapkey] = *mapvalue
			} else {
				var mapvalue k8s_io_apimachinery_pkg_apis_meta_v1.Time
				m.DisruptedPods[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodDisruptionsAllowed", wireType)
			}
			m.PodDisruptionsAllowed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PodDisruptionsAllowed |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHealthy", wireType)
			}
			m.CurrentHealthy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentHealthy |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredHealthy", wireType)
			}
			m.DesiredHealthy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DesiredHealthy |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedPods", wireType)
			}
			m.ExpectedPods = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpectedPods |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gitlab.com/sankish/magudi/vendor/k8s.io/api/policy/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x6f, 0xdc, 0x44,
	0x14, 0xc7, 0xd7, 0xd9, 0x6c, 0x08, 0xc3, 0xee, 0x2a, 0x0c, 0x14, 0xd2, 0x95, 0xf0, 0xa2, 0x3d,
	0x21, 0xa4, 0x8e, 0x49, 0x5b, 0xa1, 0x88, 0x03, 0xa2, 0x26, 0x51, 0x29, 0x6a, 0x94, 0x6a, 0x52,
	0x2e, 0xa8, 0x48, 0x8c, 0xed, 0x57, 0x67, 0x58, 0xdb, 0x63, 0xcd, 0x8c, 0x4d, 0xf7, 0xc6, 0x81,
	0x0f, 0xc0, 0xf7, 0xe0, 0x8b, 0xe4, 0x80, 0x50, 0x8f, 0x15, 0x87, 0x15, 0x31, 0xe2, 0x7b, 0x20,
	0xdb, 0xb3, 0x9b, 0xf5, 0xee, 0x46, 0xdd, 0xe6, 0xc0, 0xcd, 0xf3, 0xde, 0xfb, 0xff, 0x9e, 0xdf,
	0x7f, 0x9e, 0x8d, 0xdc, 0xf1, 0xa1, 0x22, 0x5c, 0x38, 0xe3, 0xcc, 0x03, 0x99, 0x80, 0x06, 0xe5,
	0xe4, 0x90, 0x04, 0x42, 0x3a, 0x26, 0xc1, 0x52, 0xee, 0xa4, 0x22, 0xe2, 0xfe, 0xc4, 0xc9, 0x0f,
	0x3c, 0xd0, 0xec, 0xc0, 0x09, 0x21, 0x01, 0xc9, 0x34, 0x04, 0x24, 0x95, 0x42, 0x0b, 0x7c, 0xbb,
	0x2e, 0x25, 0x2c, 0xe5, 0xa4, 0x2e, 0x25, 0xa6, 0x74, 0x70, 0x27, 0xe4, 0xfa, 0x3c, 0xf3, 0x88,
	0x2f, 0x62, 0x27, 0x14, 0xa1, 0x70, 0x2a, 0x85, 0x97, 0x3d, 0xaf, 0x4e, 0xd5, 0xa1, 0x7a, 0xaa,
	0x49, 0x83, 0xd1, 0x42, 0x53, 0x5f, 0x48, 0x70, 0xf2, 0x95, 0x6e, 0x83, 0xfb, 0x57, 0x35, 0x31,
	0xf3, 0xcf, 0x79, 0x02, 0x72, 0xe2, 0xa4, 0xe3, 0xb0, 0x0c, 0x28, 0x27, 0x06, 0xcd, 0xd6, 0xa9,
	0x9c, 0xeb, 0x54, 0x32, 0x4b, 0x34, 0x8f, 0x61, 0x45, 0xf0, 0xf9, 0xeb, 0x04, 0xca, 0x3f, 0x87,
	0x98, 0xad, 0xe8, 0xee, 0x5d, 0xa7, 0xcb, 0x34, 0x8f, 0x1c, 0x9e, 0x68, 0xa5, 0xe5, 0xb2, 0x68,
	0xf4, 0x97, 0x85, 0x76, 0x8f, 0x73, 0xee, 0x6b, 0x2e, 0x12, 0xfc, 0x23, 0xda, 0x2d, 0xa7, 0x08,
	0x98, 0x66, 0xfb, 0xd6, 0xc7, 0xd6, 0x27, 0xef, 0xdc, 0xfd, 0x8c, 0x5c, 0x39, 0x3c, 0x87, 0x92,
	0x74, 0x1c, 0x96, 0x01, 0x45, 0xca, 0x6a, 0x92, 0x1f, 0x90, 0x53, 0xef, 0x27, 0xf0, 0xf5, 0x09,
	0x68, 0xe6, 0xe2, 0x8b, 0xe9, 0xb0, 0x55, 0x4c, 0x87, 0xe8, 0x2a, 0x46, 0xe7, 0x54, 0x1c, 0xa1,
	0x5e, 0x00, 0x11, 0x68, 0x38, 0x4d, 0xcb, 0x8e, 0x6a, 0x7f, 0xab, 0x6a, 0x73, 0x6f, 0xb3, 0x36,
	0x47, 0x8b, 0x52, 0xf7, 0xdd, 0x62, 0x3a, 0xec, 0x35, 0x42, 0xb4, 0x09, 0x1f, 0xfd, 0xbe, 0x85,
	0xde, 0x7b, 0x22, 0x82, 0x23, 0xae, 0x64, 0x56, 0x85, 0xdc, 0x2c, 0x08, 0x41, 0xff, 0x0f, 0x73,
	0x3e, 0x45, 0xdb, 0x2a, 0x05, 0xdf, 0x8c, 0x77, 0x97, 0x5c, 0xbb, 0xa7, 0x64, 0xcd, 0xfb, 0x9d,
	0xa5, 0xe0, 0xbb, 0x5d, 0xc3, 0xdf, 0x2e, 0x4f, 0xb4, 0xa2, 0xe1, 0x67, 0x68, 0x47, 0x69, 0xa6,
	0x33, 0xb5, 0xdf, 0xae, 0xb8, 0xf7, 0xdf, 0x90, 0x5b, 0x69, 0xdd, 0xbe, 0x21, 0xef, 0xd4, 0x67,
	0x6a, 0x98, 0xa3, 0x3f, 0x2c, 0xf4, 0xe1, 0x1a, 0xd5, 0x63, 0xae, 0x34, 0x7e, 0xb6, 0xe2, 0x18,
	0xd9, 0xcc, 0xb1, 0x52, 0x5d, 0xf9, 0xb5, 0x67, 0xba, 0xee, 0xce, 0x22, 0x0b, 0x6e, 0x9d, 0xa1,
	0x0e, 0xd7, 0x10, 0x97, 0xdb, 0xd0, 0x5e, 0x42, 0x6f, 0x30, 0x96, 0xdb, 0x33, 0xe8, 0xce, 0xa3,
	0x12, 0x42, 0x6b, 0xd6, 0xe8, 0xcf, 0xad, 0xb5, 0xe3, 0x94, 0x76, 0xe2, 0xe7, 0xa8, 0x1b, 0xf3,
	0xe4, 0x41, 0xce, 0x78, 0xc4, 0xbc, 0x08, 0x5e, 0xbb, 0x04, 0xe5, 0x17, 0x44, 0xea, 0x2f, 0x88,
	0x3c, 0x4a, 0xf4, 0xa9, 0x3c, 0xd3, 0x92, 0x27, 0xa1, 0xbb, 0x57, 0x4c, 0x87, 0xdd, 0x93, 0x05,
	0x12, 0x6d, 0x70, 0xf1, 0x0f, 0x68, 0x57, 0x41, 0x04, 0xbe, 0x16, 0xf2, 0xcd, 0x36, 0xfd, 0x31,
	0xf3, 0x20, 0x3a, 0x33, 0x52, 0xb7, 0x5b, 0xfa, 0x36, 0x3b, 0xd1, 0x39, 0x12, 0x47, 0xa8, 0x1f,
	0xb3, 0x17, 0xdf, 0x25, 0x6c, 0x3e, 0x48, 0xfb, 0x86, 0x83, 0xe0, 0x62, 0x3a, 0xec, 0x9f, 0x34,
	0x58, 0x74, 0x89, 0x3d, 0xfa, 0x77, 0x1b, 0xdd, 0xbe, 0x76, 0xab, 0xf0, 0xb7, 0x08, 0x0b, 0x4f,
	0x81, 0xcc, 0x21, 0x78, 0x58, 0xff, 0x63, 0xb8, 0x48, 0x2a, 0x63, 0xdb, 0xee, 0xc0, 0x5c, 0x10,
	0x3e, 0x5d, 0xa9, 0xa0, 0x6b, 0x54, 0xf8, 0x57, 0x0b, 0xf5, 0x82, 0xba, 0x0d, 0x04, 0x4f, 0x44,
	0x30, 0x5b, 0x8c, 0x87, 0x37, 0xd9, 0x77, 0x72, 0xb4, 0x48, 0x3a, 0x4e, 0xb4, 0x9c, 0xb8, 0xb7,
	0xcc, 0x0b, 0xf5, 0x1a, 0x39, 0xda, 0x6c, 0x8a, 0x4f, 0x10, 0x0e, 0xe6, 0x48, 0xf5, 0x20, 0x8a,
	0xc4, 0xcf, 0x10, 0x54, 0x16, 0x77, 0xdc, 0x8f, 0x0c, 0xe1, 0x56, 0xa3, 0xef, 0xac, 0x88, 0xae,
	0x11, 0xe2, 0x2f, 0x51, 0xdf, 0xcf, 0xa4, 0x84, 0x44, 0x7f, 0x03, 0x2c, 0xd2, 0xe7, 0x93, 0xfd,
	0xed, 0x0a, 0xf5, 0x81, 0x41, 0xf5, 0xbf, 0x6e, 0x64, 0xe9, 0x52, 0x75, 0xa9, 0x0f, 0x40, 0x71,
	0x09, 0xc1, 0x4c, 0xdf, 0x69, 0xea, 0x8f, 0x1a, 0x59, 0xba, 0x54, 0x8d, 0x0f, 0x51, 0x17, 0x5e,
	0xa4, 0xe0, 0xcf, 0x3c, 0xdd, 0xa9, 0xd4, 0xef, 0x1b, 0x75, 0xf7, 0x78, 0x21, 0x47, 0x1b, 0x95,
	0x83, 0x08, 0xe1, 0x55, 0x13, 0xf1, 0x1e, 0x6a, 0x8f, 0x61, 0x52, 0x5d, 0xf1, 0xdb, 0xb4, 0x7c,
	0xc4, 0x5f, 0xa1, 0x4e, 0xce, 0xa2, 0x0c, 0xcc, 0xae, 0x7f, 0xba, 0xd9, 0xae, 0x3f, 0xe5, 0x31,
	0xd0, 0x5a, 0xf8, 0xc5, 0xd6, 0xa1, 0xe5, 0xde, 0xb9, 0xb8, 0xb4, 0x5b, 0x2f, 0x2f, 0xed, 0xd6,
	0xab, 0x4b, 0xbb, 0xf5, 0x4b, 0x61, 0x5b, 0x17, 0x85, 0x6d, 0xbd, 0x2c, 0x6c, 0xeb, 0x55, 0x61,
	0x5b, 0x7f, 0x17, 0xb6, 0xf5, 0xdb, 0x3f, 0x76, 0xeb, 0xfb, 0xb7, 0xcc, 0xc5, 0xff, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x00, 0xc0, 0xac, 0xb5, 0x48, 0x08, 0x00, 0x00,
}
