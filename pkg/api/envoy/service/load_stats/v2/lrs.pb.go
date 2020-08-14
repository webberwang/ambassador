// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/load_stats/v2/lrs.proto

package envoy_service_load_stats_v2

import (
	context "context"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	endpoint "github.com/datawire/ambassador/pkg/api/envoy/api/v2/endpoint"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoadStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node         *core.Node               `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
}

func (x *LoadStatsRequest) Reset() {
	*x = LoadStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_load_stats_v2_lrs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadStatsRequest) ProtoMessage() {}

func (x *LoadStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_load_stats_v2_lrs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadStatsRequest.ProtoReflect.Descriptor instead.
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_load_stats_v2_lrs_proto_rawDescGZIP(), []int{0}
}

func (x *LoadStatsRequest) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if x != nil {
		return x.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	SendAllClusters           bool               `protobuf:"varint,4,opt,name=send_all_clusters,json=sendAllClusters,proto3" json:"send_all_clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
}

func (x *LoadStatsResponse) Reset() {
	*x = LoadStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_load_stats_v2_lrs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadStatsResponse) ProtoMessage() {}

func (x *LoadStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_load_stats_v2_lrs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadStatsResponse.ProtoReflect.Descriptor instead.
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_load_stats_v2_lrs_proto_rawDescGZIP(), []int{1}
}

func (x *LoadStatsResponse) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *LoadStatsResponse) GetSendAllClusters() bool {
	if x != nil {
		return x.SendAllClusters
	}
	return false
}

func (x *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if x != nil {
		return x.LoadReportingInterval
	}
	return nil
}

func (x *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if x != nil {
		return x.ReportEndpointGranularity
	}
	return false
}

var File_envoy_service_load_stats_v2_lrs_proto protoreflect.FileDescriptor

var file_envoy_service_load_stats_v2_lrs_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22,
	0xee, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x65,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a,
	0x17, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x3e, 0x0a, 0x1b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x32, 0x8e, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x42, 0x0a, 0x29, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x08,
	0x4c, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_load_stats_v2_lrs_proto_rawDescOnce sync.Once
	file_envoy_service_load_stats_v2_lrs_proto_rawDescData = file_envoy_service_load_stats_v2_lrs_proto_rawDesc
)

func file_envoy_service_load_stats_v2_lrs_proto_rawDescGZIP() []byte {
	file_envoy_service_load_stats_v2_lrs_proto_rawDescOnce.Do(func() {
		file_envoy_service_load_stats_v2_lrs_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_load_stats_v2_lrs_proto_rawDescData)
	})
	return file_envoy_service_load_stats_v2_lrs_proto_rawDescData
}

var file_envoy_service_load_stats_v2_lrs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_service_load_stats_v2_lrs_proto_goTypes = []interface{}{
	(*LoadStatsRequest)(nil),      // 0: envoy.service.load_stats.v2.LoadStatsRequest
	(*LoadStatsResponse)(nil),     // 1: envoy.service.load_stats.v2.LoadStatsResponse
	(*core.Node)(nil),             // 2: envoy.api.v2.core.Node
	(*endpoint.ClusterStats)(nil), // 3: envoy.api.v2.endpoint.ClusterStats
	(*duration.Duration)(nil),     // 4: google.protobuf.Duration
}
var file_envoy_service_load_stats_v2_lrs_proto_depIdxs = []int32{
	2, // 0: envoy.service.load_stats.v2.LoadStatsRequest.node:type_name -> envoy.api.v2.core.Node
	3, // 1: envoy.service.load_stats.v2.LoadStatsRequest.cluster_stats:type_name -> envoy.api.v2.endpoint.ClusterStats
	4, // 2: envoy.service.load_stats.v2.LoadStatsResponse.load_reporting_interval:type_name -> google.protobuf.Duration
	0, // 3: envoy.service.load_stats.v2.LoadReportingService.StreamLoadStats:input_type -> envoy.service.load_stats.v2.LoadStatsRequest
	1, // 4: envoy.service.load_stats.v2.LoadReportingService.StreamLoadStats:output_type -> envoy.service.load_stats.v2.LoadStatsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_service_load_stats_v2_lrs_proto_init() }
func file_envoy_service_load_stats_v2_lrs_proto_init() {
	if File_envoy_service_load_stats_v2_lrs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_load_stats_v2_lrs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_load_stats_v2_lrs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_load_stats_v2_lrs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_load_stats_v2_lrs_proto_goTypes,
		DependencyIndexes: file_envoy_service_load_stats_v2_lrs_proto_depIdxs,
		MessageInfos:      file_envoy_service_load_stats_v2_lrs_proto_msgTypes,
	}.Build()
	File_envoy_service_load_stats_v2_lrs_proto = out.File
	file_envoy_service_load_stats_v2_lrs_proto_rawDesc = nil
	file_envoy_service_load_stats_v2_lrs_proto_goTypes = nil
	file_envoy_service_load_stats_v2_lrs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadReportingServiceClient(cc grpc.ClientConnInterface) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}
