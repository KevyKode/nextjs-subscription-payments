// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: commands.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_commands_proto protoreflect.FileDescriptor

var file_commands_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x66, 0x69, 0x78, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6c, 0x6f, 0x67, 0x73, 0x5f,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa5, 0x07, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x43, 0x4c, 0x49, 0x12, 0x43, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x46, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69,
	0x78, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x73, 0x54, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x14, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c,
	0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_commands_proto_goTypes = []interface{}{
	(*EventsResendRequest)(nil),           // 0: rpc.EventsResendRequest
	(*FixtureRequest)(nil),                // 1: rpc.FixtureRequest
	(*ListenRequest)(nil),                 // 2: rpc.ListenRequest
	(*LoginRequest)(nil),                  // 3: rpc.LoginRequest
	(*LoginStatusRequest)(nil),            // 4: rpc.LoginStatusRequest
	(*LogsTailRequest)(nil),               // 5: rpc.LogsTailRequest
	(*SampleConfigsRequest)(nil),          // 6: rpc.SampleConfigsRequest
	(*SampleCreateRequest)(nil),           // 7: rpc.SampleCreateRequest
	(*SamplesListRequest)(nil),            // 8: rpc.SamplesListRequest
	(*TriggerRequest)(nil),                // 9: rpc.TriggerRequest
	(*TriggersListRequest)(nil),           // 10: rpc.TriggersListRequest
	(*VersionRequest)(nil),                // 11: rpc.VersionRequest
	(*WebhookEndpointCreateRequest)(nil),  // 12: rpc.WebhookEndpointCreateRequest
	(*WebhookEndpointsListRequest)(nil),   // 13: rpc.WebhookEndpointsListRequest
	(*EventsResendResponse)(nil),          // 14: rpc.EventsResendResponse
	(*FixtureResponse)(nil),               // 15: rpc.FixtureResponse
	(*ListenResponse)(nil),                // 16: rpc.ListenResponse
	(*LoginResponse)(nil),                 // 17: rpc.LoginResponse
	(*LoginStatusResponse)(nil),           // 18: rpc.LoginStatusResponse
	(*LogsTailResponse)(nil),              // 19: rpc.LogsTailResponse
	(*SampleConfigsResponse)(nil),         // 20: rpc.SampleConfigsResponse
	(*SampleCreateResponse)(nil),          // 21: rpc.SampleCreateResponse
	(*SamplesListResponse)(nil),           // 22: rpc.SamplesListResponse
	(*TriggerResponse)(nil),               // 23: rpc.TriggerResponse
	(*TriggersListResponse)(nil),          // 24: rpc.TriggersListResponse
	(*VersionResponse)(nil),               // 25: rpc.VersionResponse
	(*WebhookEndpointCreateResponse)(nil), // 26: rpc.WebhookEndpointCreateResponse
	(*WebhookEndpointsListResponse)(nil),  // 27: rpc.WebhookEndpointsListResponse
}
var file_commands_proto_depIdxs = []int32{
	0,  // 0: rpc.StripeCLI.EventsResend:input_type -> rpc.EventsResendRequest
	1,  // 1: rpc.StripeCLI.Fixture:input_type -> rpc.FixtureRequest
	2,  // 2: rpc.StripeCLI.Listen:input_type -> rpc.ListenRequest
	3,  // 3: rpc.StripeCLI.Login:input_type -> rpc.LoginRequest
	4,  // 4: rpc.StripeCLI.LoginStatus:input_type -> rpc.LoginStatusRequest
	5,  // 5: rpc.StripeCLI.LogsTail:input_type -> rpc.LogsTailRequest
	6,  // 6: rpc.StripeCLI.SampleConfigs:input_type -> rpc.SampleConfigsRequest
	7,  // 7: rpc.StripeCLI.SampleCreate:input_type -> rpc.SampleCreateRequest
	8,  // 8: rpc.StripeCLI.SamplesList:input_type -> rpc.SamplesListRequest
	9,  // 9: rpc.StripeCLI.Trigger:input_type -> rpc.TriggerRequest
	10, // 10: rpc.StripeCLI.TriggersList:input_type -> rpc.TriggersListRequest
	11, // 11: rpc.StripeCLI.Version:input_type -> rpc.VersionRequest
	12, // 12: rpc.StripeCLI.WebhookEndpointCreate:input_type -> rpc.WebhookEndpointCreateRequest
	13, // 13: rpc.StripeCLI.WebhookEndpointsList:input_type -> rpc.WebhookEndpointsListRequest
	14, // 14: rpc.StripeCLI.EventsResend:output_type -> rpc.EventsResendResponse
	15, // 15: rpc.StripeCLI.Fixture:output_type -> rpc.FixtureResponse
	16, // 16: rpc.StripeCLI.Listen:output_type -> rpc.ListenResponse
	17, // 17: rpc.StripeCLI.Login:output_type -> rpc.LoginResponse
	18, // 18: rpc.StripeCLI.LoginStatus:output_type -> rpc.LoginStatusResponse
	19, // 19: rpc.StripeCLI.LogsTail:output_type -> rpc.LogsTailResponse
	20, // 20: rpc.StripeCLI.SampleConfigs:output_type -> rpc.SampleConfigsResponse
	21, // 21: rpc.StripeCLI.SampleCreate:output_type -> rpc.SampleCreateResponse
	22, // 22: rpc.StripeCLI.SamplesList:output_type -> rpc.SamplesListResponse
	23, // 23: rpc.StripeCLI.Trigger:output_type -> rpc.TriggerResponse
	24, // 24: rpc.StripeCLI.TriggersList:output_type -> rpc.TriggersListResponse
	25, // 25: rpc.StripeCLI.Version:output_type -> rpc.VersionResponse
	26, // 26: rpc.StripeCLI.WebhookEndpointCreate:output_type -> rpc.WebhookEndpointCreateResponse
	27, // 27: rpc.StripeCLI.WebhookEndpointsList:output_type -> rpc.WebhookEndpointsListResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_commands_proto_init() }
func file_commands_proto_init() {
	if File_commands_proto != nil {
		return
	}
	file_events_resend_proto_init()
	file_fixtures_proto_init()
	file_listen_proto_init()
	file_login_proto_init()
	file_login_status_proto_init()
	file_logs_tail_proto_init()
	file_sample_configs_proto_init()
	file_sample_create_proto_init()
	file_samples_list_proto_init()
	file_trigger_proto_init()
	file_triggers_list_proto_init()
	file_version_proto_init()
	file_webhook_endpoint_create_proto_init()
	file_webhook_endpoints_list_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commands_proto_goTypes,
		DependencyIndexes: file_commands_proto_depIdxs,
	}.Build()
	File_commands_proto = out.File
	file_commands_proto_rawDesc = nil
	file_commands_proto_goTypes = nil
	file_commands_proto_depIdxs = nil
}
