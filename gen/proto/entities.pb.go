// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: entities.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Skill) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Skill) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Vacancy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Vacancy) Reset() {
	*x = Vacancy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vacancy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vacancy) ProtoMessage() {}

func (x *Vacancy) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vacancy.ProtoReflect.Descriptor instead.
func (*Vacancy) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Vacancy) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vacancy) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type VacancyProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VacancyId uint32   `protobuf:"varint,2,opt,name=vacancyId,proto3" json:"vacancyId,omitempty"`
	GradeIds  []uint32 `protobuf:"varint,3,rep,packed,name=gradeIds,proto3" json:"gradeIds,omitempty"`
}

func (x *VacancyProgress) Reset() {
	*x = VacancyProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacancyProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacancyProgress) ProtoMessage() {}

func (x *VacancyProgress) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacancyProgress.ProtoReflect.Descriptor instead.
func (*VacancyProgress) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{2}
}

func (x *VacancyProgress) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VacancyProgress) GetVacancyId() uint32 {
	if x != nil {
		return x.VacancyId
	}
	return 0
}

func (x *VacancyProgress) GetGradeIds() []uint32 {
	if x != nil {
		return x.GradeIds
	}
	return nil
}

type Grade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label         string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Experience    uint32 `protobuf:"varint,4,opt,name=experience,proto3" json:"experience,omitempty"`
	NeedsApproval bool   `protobuf:"varint,5,opt,name=needsApproval,proto3" json:"needsApproval,omitempty"`
}

func (x *Grade) Reset() {
	*x = Grade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grade) ProtoMessage() {}

func (x *Grade) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grade.ProtoReflect.Descriptor instead.
func (*Grade) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{3}
}

func (x *Grade) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Grade) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Grade) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Grade) GetExperience() uint32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *Grade) GetNeedsApproval() bool {
	if x != nil {
		return x.NeedsApproval
	}
	return false
}

type Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VacancyId uint32   `protobuf:"varint,1,opt,name=vacancyId,proto3" json:"vacancyId,omitempty"`
	GradesIds []uint32 `protobuf:"varint,2,rep,packed,name=grades_ids,json=gradesIds,proto3" json:"grades_ids,omitempty"`
	Id        uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Matrix) Reset() {
	*x = Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matrix) ProtoMessage() {}

func (x *Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matrix.ProtoReflect.Descriptor instead.
func (*Matrix) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{4}
}

func (x *Matrix) GetVacancyId() uint32 {
	if x != nil {
		return x.VacancyId
	}
	return 0
}

func (x *Matrix) GetGradesIds() []uint32 {
	if x != nil {
		return x.GradesIds
	}
	return nil
}

func (x *Matrix) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AuthorId    uint32   `protobuf:"varint,4,opt,name=authorId,proto3" json:"authorId,omitempty"`
	TotalViewed uint32   `protobuf:"varint,5,opt,name=totalViewed,proto3" json:"totalViewed,omitempty"`
	SkillsIds   []uint32 `protobuf:"varint,6,rep,packed,name=skillsIds,proto3" json:"skillsIds,omitempty"`
	Experience  uint32   `protobuf:"varint,7,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{5}
}

func (x *Article) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Article) GetAuthorId() uint32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Article) GetTotalViewed() uint32 {
	if x != nil {
		return x.TotalViewed
	}
	return 0
}

func (x *Article) GetSkillsIds() []uint32 {
	if x != nil {
		return x.SkillsIds
	}
	return nil
}

func (x *Article) GetExperience() uint32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TeamleadId  uint32   `protobuf:"varint,4,opt,name=teamleadId,proto3" json:"teamleadId,omitempty"`
	SkillsIds   []uint32 `protobuf:"varint,5,rep,packed,name=skillsIds,proto3" json:"skillsIds,omitempty"`
	UserIds     []uint32 `protobuf:"varint,6,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	Experience  uint32   `protobuf:"varint,7,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{6}
}

func (x *Project) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetTeamleadId() uint32 {
	if x != nil {
		return x.TeamleadId
	}
	return 0
}

func (x *Project) GetSkillsIds() []uint32 {
	if x != nil {
		return x.SkillsIds
	}
	return nil
}

func (x *Project) GetUserIds() []uint32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *Project) GetExperience() uint32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login               string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	ViewedPostsIds      []uint32 `protobuf:"varint,3,rep,packed,name=viewedPostsIds,proto3" json:"viewedPostsIds,omitempty"`
	FirstName           string   `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName            string   `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName          string   `protobuf:"bytes,6,opt,name=middleName,proto3" json:"middleName,omitempty"`
	AvatarUri           string   `protobuf:"bytes,7,opt,name=avatarUri,proto3" json:"avatarUri,omitempty"`
	CurentProjectId     uint32   `protobuf:"varint,8,opt,name=curentProjectId,proto3" json:"curentProjectId,omitempty"`
	CompletedProjectIds []uint32 `protobuf:"varint,9,rep,packed,name=completedProjectIds,proto3" json:"completedProjectIds,omitempty"`
	SkillsIds           []uint32 `protobuf:"varint,10,rep,packed,name=skillsIds,proto3" json:"skillsIds,omitempty"`
	Role                string   `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
	TotalExperience     uint32   `protobuf:"varint,12,opt,name=totalExperience,proto3" json:"totalExperience,omitempty"`
	VacancyId           uint32   `protobuf:"varint,13,opt,name=vacancyId,proto3" json:"vacancyId,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetViewedPostsIds() []uint32 {
	if x != nil {
		return x.ViewedPostsIds
	}
	return nil
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetAvatarUri() string {
	if x != nil {
		return x.AvatarUri
	}
	return ""
}

func (x *User) GetCurentProjectId() uint32 {
	if x != nil {
		return x.CurentProjectId
	}
	return 0
}

func (x *User) GetCompletedProjectIds() []uint32 {
	if x != nil {
		return x.CompletedProjectIds
	}
	return nil
}

func (x *User) GetSkillsIds() []uint32 {
	if x != nil {
		return x.SkillsIds
	}
	return nil
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetTotalExperience() uint32 {
	if x != nil {
		return x.TotalExperience
	}
	return 0
}

func (x *User) GetVacancyId() uint32 {
	if x != nil {
		return x.VacancyId
	}
	return 0
}

type DataLibResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users               []*User            `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Projects            []*Project         `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	Skills              []*Skill           `protobuf:"bytes,3,rep,name=skills,proto3" json:"skills,omitempty"`
	VacanciesProgresses []*VacancyProgress `protobuf:"bytes,4,rep,name=vacanciesProgresses,proto3" json:"vacanciesProgresses,omitempty"`
	Vacancies           []*Vacancy         `protobuf:"bytes,5,rep,name=vacancies,proto3" json:"vacancies,omitempty"`
	Articles            []*Article         `protobuf:"bytes,6,rep,name=articles,proto3" json:"articles,omitempty"`
	Grades              []*Grade           `protobuf:"bytes,7,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *DataLibResponse) Reset() {
	*x = DataLibResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataLibResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataLibResponse) ProtoMessage() {}

func (x *DataLibResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entities_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataLibResponse.ProtoReflect.Descriptor instead.
func (*DataLibResponse) Descriptor() ([]byte, []int) {
	return file_entities_proto_rawDescGZIP(), []int{8}
}

func (x *DataLibResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *DataLibResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *DataLibResponse) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *DataLibResponse) GetVacanciesProgresses() []*VacancyProgress {
	if x != nil {
		return x.VacanciesProgresses
	}
	return nil
}

func (x *DataLibResponse) GetVacancies() []*Vacancy {
	if x != nil {
		return x.Vacancies
	}
	return nil
}

func (x *DataLibResponse) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *DataLibResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

var File_entities_proto protoreflect.FileDescriptor

var file_entities_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x2f, 0x0a, 0x07, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0x5b, 0x0a, 0x0f, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e,
	0x63, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61,
	0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6e, 0x65, 0x65, 0x64,
	0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x22, 0x55, 0x0a, 0x06, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x49, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xcd, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x64, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0xc9, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x6c, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x6c, 0x65,
	0x61, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49,
	0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xa2, 0x03, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0e, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x69, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x75, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x49, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x49,
	0x64, 0x22, 0xd0, 0x02, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x76, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x13, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x42, 0x41, 0x5a, 0x21, 0x64, 0x61, 0x6d, 0x6d, 0x6e, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x70, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x1b, 0x1a, 0x19, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x6d, 0x6e, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_proto_rawDescOnce sync.Once
	file_entities_proto_rawDescData = file_entities_proto_rawDesc
)

func file_entities_proto_rawDescGZIP() []byte {
	file_entities_proto_rawDescOnce.Do(func() {
		file_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_proto_rawDescData)
	})
	return file_entities_proto_rawDescData
}

var file_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_entities_proto_goTypes = []interface{}{
	(*Skill)(nil),           // 0: proto.Skill
	(*Vacancy)(nil),         // 1: proto.Vacancy
	(*VacancyProgress)(nil), // 2: proto.VacancyProgress
	(*Grade)(nil),           // 3: proto.Grade
	(*Matrix)(nil),          // 4: proto.Matrix
	(*Article)(nil),         // 5: proto.Article
	(*Project)(nil),         // 6: proto.Project
	(*User)(nil),            // 7: proto.User
	(*DataLibResponse)(nil), // 8: proto.DataLibResponse
}
var file_entities_proto_depIdxs = []int32{
	7, // 0: proto.DataLibResponse.users:type_name -> proto.User
	6, // 1: proto.DataLibResponse.projects:type_name -> proto.Project
	0, // 2: proto.DataLibResponse.skills:type_name -> proto.Skill
	2, // 3: proto.DataLibResponse.vacanciesProgresses:type_name -> proto.VacancyProgress
	1, // 4: proto.DataLibResponse.vacancies:type_name -> proto.Vacancy
	5, // 5: proto.DataLibResponse.articles:type_name -> proto.Article
	3, // 6: proto.DataLibResponse.grades:type_name -> proto.Grade
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_entities_proto_init() }
func file_entities_proto_init() {
	if File_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vacancy); i {
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
		file_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacancyProgress); i {
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
		file_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grade); i {
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
		file_entities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matrix); i {
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
		file_entities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_entities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_entities_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_entities_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataLibResponse); i {
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
			RawDescriptor: file_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_proto_goTypes,
		DependencyIndexes: file_entities_proto_depIdxs,
		MessageInfos:      file_entities_proto_msgTypes,
	}.Build()
	File_entities_proto = out.File
	file_entities_proto_rawDesc = nil
	file_entities_proto_goTypes = nil
	file_entities_proto_depIdxs = nil
}
