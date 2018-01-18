// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: main.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	PublishRequest
	RetResponse
	AddUserRequest
	DelUserRequest
	GetUserRequest
	GetUserResponse
	AddMembersRequest
	DelMembersRequest
	ListMembersRequest
	ListMembersResponse
*/
package main

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PublishRequest) Validate() error {
	if this.Topic == "" {
		return go_proto_validators.FieldError("Topic", fmt.Errorf(`value '%v' must not be an empty string`, this.Topic))
	}
	if this.Payload == "" {
		return go_proto_validators.FieldError("Payload", fmt.Errorf(`value '%v' must not be an empty string`, this.Payload))
	}
	if !(this.Qos > -1) {
		return go_proto_validators.FieldError("Qos", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Qos))
	}
	if !(this.Qos < 3) {
		return go_proto_validators.FieldError("Qos", fmt.Errorf(`value '%v' must be less than '3'`, this.Qos))
	}
	return nil
}
func (this *RetResponse) Validate() error {
	return nil
}
func (this *AddUserRequest) Validate() error {
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *DelUserRequest) Validate() error {
	return nil
}
func (this *GetUserRequest) Validate() error {
	return nil
}
func (this *GetUserResponse) Validate() error {
	return nil
}
func (this *AddMembersRequest) Validate() error {
	if this.GroupId == "" {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf(`value '%v' must not be an empty string`, this.GroupId))
	}
	if len(this.Members) < 1 {
		return go_proto_validators.FieldError("Members", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Members))
	}
	return nil
}
func (this *DelMembersRequest) Validate() error {
	if this.GroupId == "" {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf(`value '%v' must not be an empty string`, this.GroupId))
	}
	if len(this.Members) < 1 {
		return go_proto_validators.FieldError("Members", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Members))
	}
	return nil
}
func (this *ListMembersRequest) Validate() error {
	if this.GroupId == "" {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf(`value '%v' must not be an empty string`, this.GroupId))
	}
	return nil
}
func (this *ListMembersResponse) Validate() error {
	return nil
}