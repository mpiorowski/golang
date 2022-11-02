// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: food.proto

package homeit

import (
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

type SupplyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyId string `protobuf:"bytes,1,opt,name=supplyId,proto3" json:"supplyId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *SupplyId) Reset() {
	*x = SupplyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyId) ProtoMessage() {}

func (x *SupplyId) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyId.ProtoReflect.Descriptor instead.
func (*SupplyId) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *SupplyId) GetSupplyId() string {
	if x != nil {
		return x.SupplyId
	}
	return ""
}

func (x *SupplyId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Supply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created  string `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  bool   `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	UserId   string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	Name     string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	Amount   int32  `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Priority int32  `protobuf:"varint,9,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *Supply) Reset() {
	*x = Supply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supply) ProtoMessage() {}

func (x *Supply) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supply.ProtoReflect.Descriptor instead.
func (*Supply) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *Supply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supply) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Supply) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Supply) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Supply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Supply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supply) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Supply) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Supply) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type MenuId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId string `protobuf:"bytes,1,opt,name=menuId,proto3" json:"menuId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MenuId) Reset() {
	*x = MenuId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuId) ProtoMessage() {}

func (x *MenuId) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuId.ProtoReflect.Descriptor instead.
func (*MenuId) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *MenuId) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

func (x *MenuId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created  string  `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string  `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  *string `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	UserId   string  `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	RecipeId string  `protobuf:"bytes,6,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	Date     string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Type     string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Recipe   *Recipe `protobuf:"bytes,9,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *Menu) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Menu) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Menu) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Menu) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Menu) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Menu) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *Menu) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Menu) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Menu) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

type RecipeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId string `protobuf:"bytes,1,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RecipeId) Reset() {
	*x = RecipeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeId) ProtoMessage() {}

func (x *RecipeId) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeId.ProtoReflect.Descriptor instead.
func (*RecipeId) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *RecipeId) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *RecipeId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created     string        `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated     string        `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted     *string       `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	UserId      string        `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	Name        string        `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description string        `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Category    string        `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Ingredients []*Ingredient `protobuf:"bytes,9,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *Recipe) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Recipe) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Recipe) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Recipe) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Recipe) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Recipe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipe) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Recipe) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Recipe) GetIngredients() []*Ingredient {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

type IngredientId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngredientId string `protobuf:"bytes,1,opt,name=ingredientId,proto3" json:"ingredientId,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *IngredientId) Reset() {
	*x = IngredientId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngredientId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientId) ProtoMessage() {}

func (x *IngredientId) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientId.ProtoReflect.Descriptor instead.
func (*IngredientId) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{6}
}

func (x *IngredientId) GetIngredientId() string {
	if x != nil {
		return x.IngredientId
	}
	return ""
}

func (x *IngredientId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Ingredient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created  string  `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string  `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  *string `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	UserId   string  `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	RecipeId string  `protobuf:"bytes,6,opt,name=recipeId,proto3" json:"recipeId,omitempty"`
	Name     string  `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Amount   float64 `protobuf:"fixed64,8,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Ingredient) Reset() {
	*x = Ingredient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingredient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredient) ProtoMessage() {}

func (x *Ingredient) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredient.ProtoReflect.Descriptor instead.
func (*Ingredient) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{7}
}

func (x *Ingredient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ingredient) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Ingredient) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Ingredient) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Ingredient) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ingredient) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *Ingredient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ingredient) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x22, 0x3e, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x06, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xf7, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x08, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x95, 0x02, 0x0a, 0x06,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xdb, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x4a, 0x0a,
	0x14, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x70, 0x69, 0x6f, 0x72, 0x6f, 0x77, 0x73, 0x6b, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_food_proto_goTypes = []interface{}{
	(*SupplyId)(nil),     // 0: food.SupplyId
	(*Supply)(nil),       // 1: food.Supply
	(*MenuId)(nil),       // 2: food.MenuId
	(*Menu)(nil),         // 3: food.Menu
	(*RecipeId)(nil),     // 4: food.RecipeId
	(*Recipe)(nil),       // 5: food.Recipe
	(*IngredientId)(nil), // 6: food.IngredientId
	(*Ingredient)(nil),   // 7: food.Ingredient
}
var file_food_proto_depIdxs = []int32{
	5, // 0: food.Menu.recipe:type_name -> food.Recipe
	7, // 1: food.Recipe.ingredients:type_name -> food.Ingredient
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyId); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supply); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuId); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeId); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngredientId); i {
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
		file_food_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingredient); i {
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
	file_food_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_food_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_food_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}