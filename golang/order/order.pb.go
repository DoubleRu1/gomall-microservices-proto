// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: order/order.proto

package order

import (
	cart "github.com/DoubleRu1/gomall-microservices-proto/golang/cart"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StreetAddress string                 `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string                 `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       int32                  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency  string                 `protobuf:"bytes,2,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address       *Address               `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	OrderItems    []*OrderItem           `protobuf:"bytes,5,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	mi := &file_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *PlaceOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *PlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceOrderReq) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *cart.CartItem         `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cost          float32                `protobuf:"fixed32,2,opt,name=cost,proto3" json:"cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetItem() *cart.CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *OrderItem) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type OrderResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderResult) Reset() {
	*x = OrderResult{}
	mi := &file_order_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResult) ProtoMessage() {}

func (x *OrderResult) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResult.ProtoReflect.Descriptor instead.
func (*OrderResult) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderResult) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type PlaceOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *OrderResult           `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderResp) Reset() {
	*x = PlaceOrderResp{}
	mi := &file_order_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResp) ProtoMessage() {}

func (x *PlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResp.ProtoReflect.Descriptor instead.
func (*PlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderResp) GetOrder() *OrderResult {
	if x != nil {
		return x.Order
	}
	return nil
}

type ListOrderReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderReq) Reset() {
	*x = ListOrderReq{}
	mi := &file_order_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderReq) ProtoMessage() {}

func (x *ListOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderReq.ProtoReflect.Descriptor instead.
func (*ListOrderReq) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderItems    []*OrderItem           `protobuf:"bytes,1,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	OrderId       string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        uint32                 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserCurrency  string                 `protobuf:"bytes,4,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty"`
	Address       *Address               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     int32                  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Order) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ListOrderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrderResp) Reset() {
	*x = ListOrderResp{}
	mi := &file_order_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResp) ProtoMessage() {}

func (x *ListOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResp.ProtoReflect.Descriptor instead.
func (*ListOrderResp) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrderResp) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type MarkOrderPaidReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId       string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkOrderPaidReq) Reset() {
	*x = MarkOrderPaidReq{}
	mi := &file_order_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkOrderPaidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderPaidReq) ProtoMessage() {}

func (x *MarkOrderPaidReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderPaidReq.ProtoReflect.Descriptor instead.
func (*MarkOrderPaidReq) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{8}
}

func (x *MarkOrderPaidReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MarkOrderPaidReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type MarkOrderPaidResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkOrderPaidResp) Reset() {
	*x = MarkOrderPaidResp{}
	mi := &file_order_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkOrderPaidResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderPaidResp) ProtoMessage() {}

func (x *MarkOrderPaidResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderPaidResp.ProtoReflect.Descriptor instead.
func (*MarkOrderPaidResp) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{9}
}

var File_order_order_proto protoreflect.FileDescriptor

var file_order_order_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0f, 0x63, 0x61, 0x72, 0x74,
	0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc0, 0x01,
	0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x0a,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x43, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3a, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xf2, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x22, 0x46, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x32, 0xcb, 0x01,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x52, 0x75, 0x31, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_order_order_proto_rawDescOnce sync.Once
	file_order_order_proto_rawDescData []byte
)

func file_order_order_proto_rawDescGZIP() []byte {
	file_order_order_proto_rawDescOnce.Do(func() {
		file_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)))
	})
	return file_order_order_proto_rawDescData
}

var file_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_order_order_proto_goTypes = []any{
	(*Address)(nil),           // 0: order.Address
	(*PlaceOrderReq)(nil),     // 1: order.PlaceOrderReq
	(*OrderItem)(nil),         // 2: order.OrderItem
	(*OrderResult)(nil),       // 3: order.OrderResult
	(*PlaceOrderResp)(nil),    // 4: order.PlaceOrderResp
	(*ListOrderReq)(nil),      // 5: order.ListOrderReq
	(*Order)(nil),             // 6: order.Order
	(*ListOrderResp)(nil),     // 7: order.ListOrderResp
	(*MarkOrderPaidReq)(nil),  // 8: order.MarkOrderPaidReq
	(*MarkOrderPaidResp)(nil), // 9: order.MarkOrderPaidResp
	(*cart.CartItem)(nil),     // 10: cart.CartItem
}
var file_order_order_proto_depIdxs = []int32{
	0,  // 0: order.PlaceOrderReq.address:type_name -> order.Address
	2,  // 1: order.PlaceOrderReq.order_items:type_name -> order.OrderItem
	10, // 2: order.OrderItem.item:type_name -> cart.CartItem
	3,  // 3: order.PlaceOrderResp.order:type_name -> order.OrderResult
	2,  // 4: order.Order.order_items:type_name -> order.OrderItem
	0,  // 5: order.Order.address:type_name -> order.Address
	6,  // 6: order.ListOrderResp.orders:type_name -> order.Order
	1,  // 7: order.OrderService.PlaceOrder:input_type -> order.PlaceOrderReq
	5,  // 8: order.OrderService.ListOrder:input_type -> order.ListOrderReq
	8,  // 9: order.OrderService.MarkOrderPaid:input_type -> order.MarkOrderPaidReq
	4,  // 10: order.OrderService.PlaceOrder:output_type -> order.PlaceOrderResp
	7,  // 11: order.OrderService.ListOrder:output_type -> order.ListOrderResp
	9,  // 12: order.OrderService.MarkOrderPaid:output_type -> order.MarkOrderPaidResp
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_order_order_proto_init() }
func file_order_order_proto_init() {
	if File_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_order_proto_goTypes,
		DependencyIndexes: file_order_order_proto_depIdxs,
		MessageInfos:      file_order_order_proto_msgTypes,
	}.Build()
	File_order_order_proto = out.File
	file_order_order_proto_goTypes = nil
	file_order_order_proto_depIdxs = nil
}
