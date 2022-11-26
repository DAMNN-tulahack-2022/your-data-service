package transport

import (
	"context"
	"log"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (gp *gatewayProxy) GradesList(ctx context.Context, request *emptypb.Empty) (*proto.GradesListResponse, error) {
	grades, err := repository.Grades.List(ctx)
	if err != nil {
		return nil, err
	}

	response := &proto.GradesListResponse{
		Grades: make([]*proto.Grade, 0, len(grades)),
	}

	for i := range grades {
		response.Grades = append(response.Grades, &proto.Grade{
			Id:    grades[i].Id,
			Label: grades[i].Label,
		})
	}

	return response, nil
}

func (gp *gatewayProxy) GradeAdd(ctx context.Context, request *proto.GradeAddRequest) (*proto.Grade, error) {
	grade := &repository.Grade{
		Label:         request.Params.GetLabel(),
		Description:   request.Params.GetDescription(),
		Exp:           request.Params.GetExperience(),
		NeedsApproval: request.Params.GetNeedsApproval(),
	}

	grade, err := repository.Grades.Add(ctx, grade)
	if err != nil {
		return nil, err
	}

	return &proto.Grade{
		Id:            grade.Id,
		Label:         grade.Label,
		Description:   grade.Description,
		Experience:    grade.Exp,
		NeedsApproval: grade.NeedsApproval,
	}, nil
}

func (gp *gatewayProxy) GradeEdit(ctx context.Context, request *proto.GradeEditRequest) (*proto.Grade, error) {
	log.Print(request.GetId())
	grade, err := repository.Grades.Get(ctx, request.GetId())
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	mask := request.UpdateMask
	if mask == nil {
		return nil, status.Error(codes.Unknown, "errors.common.internalError")
	}

	for _, v := range mask.Paths {
		switch v {
		case "label":
			grade.Label = request.Params.GetLabel()
		case "description":
			grade.Description = request.Params.GetDescription()
		case "experience":
			grade.Exp = request.Params.GetExperience()
		case "needsApproval":
			grade.NeedsApproval = request.Params.GetNeedsApproval()
		}
	}

	grade, err = repository.Grades.Edit(ctx, grade)
	if err != nil {
		return nil, err
	}

	return &proto.Grade{
		Id:            grade.Id,
		Label:         grade.Label,
		Description:   grade.Description,
		Experience:    grade.Exp,
		NeedsApproval: grade.NeedsApproval,
	}, nil
}

func (gp *gatewayProxy) GradeRemove(ctx context.Context, request *proto.GradeRemoveRequest) (*proto.GradeRemoveResponse, error) {
	err := repository.Grades.Remove(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.GradeRemoveResponse{Id: request.GetId()}, nil
}
