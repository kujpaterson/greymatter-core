// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.5
// source: source/filters/network/proto/observables.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type ObservablesTCPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to emit response (or otherwise just request)
	EmitFullResponse bool `protobuf:"varint,1,opt,name=emitFullResponse,proto3" json:"emitFullResponse,omitempty"`
	// Whether to use Kafka
	UseKafka bool `protobuf:"varint,2,opt,name=useKafka,proto3" json:"useKafka,omitempty"`
	// Kafka TLS configuration
	// -----------------------
	// Whether to use TLS when connecting to Kafka
	UseKafkaTLS bool `protobuf:"varint,3,opt,name=useKafkaTLS,proto3" json:"useKafkaTLS,omitempty"`
	// Kafka Certificate Authorities
	KafkaCAs string `protobuf:"bytes,4,opt,name=kafkaCAs,proto3" json:"kafkaCAs,omitempty"`
	// Kafka TLC cert key
	KafkaCertificate string `protobuf:"bytes,5,opt,name=kafkaCertificate,proto3" json:"kafkaCertificate,omitempty"`
	// Kafka TLC cert key
	KafkaCertificateKey string `protobuf:"bytes,6,opt,name=kafkaCertificateKey,proto3" json:"kafkaCertificateKey,omitempty"`
	// Name of Kafka server
	KafkaServerName string `protobuf:"bytes,7,opt,name=kafkaServerName,proto3" json:"kafkaServerName,omitempty"`
	// The topic name to embed in the event.
	Topic string `protobuf:"bytes,8,opt,name=topic,proto3" json:"topic,omitempty"`
	// Kafka topic to publish to.
	EventTopic string `protobuf:"bytes,9,opt,name=eventTopic,proto3" json:"eventTopic,omitempty"`
	// Whether to use Zookeeper for Kafka discovery (if not using file storage)
	KafkaZKDiscover bool `protobuf:"varint,10,opt,name=kafkaZKDiscover,proto3" json:"kafkaZKDiscover,omitempty"`
	// Kafka connection string (if not using file storage)
	KafkaServerConnection string `protobuf:"bytes,11,opt,name=kafkaServerConnection,proto3" json:"kafkaServerConnection,omitempty"`
	// File to store event to use (if not using Kafka)
	FileName string `protobuf:"bytes,12,opt,name=fileName,proto3" json:"fileName,omitempty"`
	// Log level to use ("warn", "debug" or "info")
	LogLevel string `protobuf:"bytes,13,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	// Algorithm used to encrypt
	EncryptionAlgorithm string `protobuf:"bytes,14,opt,name=encryptionAlgorithm,proto3" json:"encryptionAlgorithm,omitempty"`
	// Key to encrypt event
	EncryptionKey   string `protobuf:"bytes,15,opt,name=encryptionKey,proto3" json:"encryptionKey,omitempty"`
	EncryptionKeyID uint32 `protobuf:"varint,16,opt,name=encryptionKeyID,proto3" json:"encryptionKeyID,omitempty"`
	// Kafka timeout
	TimeoutMs    int32 `protobuf:"varint,17,opt,name=timeoutMs,proto3" json:"timeoutMs,omitempty"`
	EnforceAudit bool  `protobuf:"varint,18,opt,name=enforceAudit,proto3" json:"enforceAudit,omitempty"`
	// Decode
	DecodeToProtocol string `protobuf:"bytes,19,opt,name=decodeToProtocol,proto3" json:"decodeToProtocol,omitempty"`
	DecodeSkipFail   bool   `protobuf:"varint,20,opt,name=decodeSkipFail,proto3" json:"decodeSkipFail,omitempty"`
}

func (x *ObservablesTCPConfig) Reset() {
	*x = ObservablesTCPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_network_proto_observables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservablesTCPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservablesTCPConfig) ProtoMessage() {}

func (x *ObservablesTCPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_network_proto_observables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservablesTCPConfig.ProtoReflect.Descriptor instead.
func (*ObservablesTCPConfig) Descriptor() ([]byte, []int) {
	return file_source_filters_network_proto_observables_proto_rawDescGZIP(), []int{0}
}

func (x *ObservablesTCPConfig) GetEmitFullResponse() bool {
	if x != nil {
		return x.EmitFullResponse
	}
	return false
}

func (x *ObservablesTCPConfig) GetUseKafka() bool {
	if x != nil {
		return x.UseKafka
	}
	return false
}

func (x *ObservablesTCPConfig) GetUseKafkaTLS() bool {
	if x != nil {
		return x.UseKafkaTLS
	}
	return false
}

func (x *ObservablesTCPConfig) GetKafkaCAs() string {
	if x != nil {
		return x.KafkaCAs
	}
	return ""
}

func (x *ObservablesTCPConfig) GetKafkaCertificate() string {
	if x != nil {
		return x.KafkaCertificate
	}
	return ""
}

func (x *ObservablesTCPConfig) GetKafkaCertificateKey() string {
	if x != nil {
		return x.KafkaCertificateKey
	}
	return ""
}

func (x *ObservablesTCPConfig) GetKafkaServerName() string {
	if x != nil {
		return x.KafkaServerName
	}
	return ""
}

func (x *ObservablesTCPConfig) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ObservablesTCPConfig) GetEventTopic() string {
	if x != nil {
		return x.EventTopic
	}
	return ""
}

func (x *ObservablesTCPConfig) GetKafkaZKDiscover() bool {
	if x != nil {
		return x.KafkaZKDiscover
	}
	return false
}

func (x *ObservablesTCPConfig) GetKafkaServerConnection() string {
	if x != nil {
		return x.KafkaServerConnection
	}
	return ""
}

func (x *ObservablesTCPConfig) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ObservablesTCPConfig) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *ObservablesTCPConfig) GetEncryptionAlgorithm() string {
	if x != nil {
		return x.EncryptionAlgorithm
	}
	return ""
}

func (x *ObservablesTCPConfig) GetEncryptionKey() string {
	if x != nil {
		return x.EncryptionKey
	}
	return ""
}

func (x *ObservablesTCPConfig) GetEncryptionKeyID() uint32 {
	if x != nil {
		return x.EncryptionKeyID
	}
	return 0
}

func (x *ObservablesTCPConfig) GetTimeoutMs() int32 {
	if x != nil {
		return x.TimeoutMs
	}
	return 0
}

func (x *ObservablesTCPConfig) GetEnforceAudit() bool {
	if x != nil {
		return x.EnforceAudit
	}
	return false
}

func (x *ObservablesTCPConfig) GetDecodeToProtocol() string {
	if x != nil {
		return x.DecodeToProtocol
	}
	return ""
}

func (x *ObservablesTCPConfig) GetDecodeSkipFail() bool {
	if x != nil {
		return x.DecodeSkipFail
	}
	return false
}

var File_source_filters_network_proto_observables_proto protoreflect.FileDescriptor

var file_source_filters_network_proto_observables_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2d, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e,
	0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22,
	0x8a, 0x06, 0x0a, 0x14, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x54,
	0x43, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x6d, 0x69, 0x74,
	0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x65, 0x6d, 0x69, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x73, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x4c, 0x53, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54,
	0x4c, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x41, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x41, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x28, 0x0a, 0x0f,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5a, 0x4b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5a, 0x4b, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x4d, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x4d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x6b,
	0x69, 0x70, 0x46, 0x61, 0x69, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x46, 0x61, 0x69, 0x6c, 0x42, 0x40, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x79, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_filters_network_proto_observables_proto_rawDescOnce sync.Once
	file_source_filters_network_proto_observables_proto_rawDescData = file_source_filters_network_proto_observables_proto_rawDesc
)

func file_source_filters_network_proto_observables_proto_rawDescGZIP() []byte {
	file_source_filters_network_proto_observables_proto_rawDescOnce.Do(func() {
		file_source_filters_network_proto_observables_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_filters_network_proto_observables_proto_rawDescData)
	})
	return file_source_filters_network_proto_observables_proto_rawDescData
}

var file_source_filters_network_proto_observables_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_source_filters_network_proto_observables_proto_goTypes = []interface{}{
	(*ObservablesTCPConfig)(nil), // 0: greymatter_io.gm_proxy.source.filters.network.ObservablesTCPConfig
}
var file_source_filters_network_proto_observables_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_source_filters_network_proto_observables_proto_init() }
func file_source_filters_network_proto_observables_proto_init() {
	if File_source_filters_network_proto_observables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_filters_network_proto_observables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservablesTCPConfig); i {
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
			RawDescriptor: file_source_filters_network_proto_observables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_filters_network_proto_observables_proto_goTypes,
		DependencyIndexes: file_source_filters_network_proto_observables_proto_depIdxs,
		MessageInfos:      file_source_filters_network_proto_observables_proto_msgTypes,
	}.Build()
	File_source_filters_network_proto_observables_proto = out.File
	file_source_filters_network_proto_observables_proto_rawDesc = nil
	file_source_filters_network_proto_observables_proto_goTypes = nil
	file_source_filters_network_proto_observables_proto_depIdxs = nil
}