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

func (gp *gatewayProxy) MatrixList(ctx context.Context, request *emptypb.Empty) (*proto.MatrixListResponse, error) {
	matricies, err := repository.Matricies.List(ctx)
	if err != nil {
		return nil, err
	}

	response := &proto.MatrixListResponse{
		GradeProgresses: make([]*proto.Matrix, 0, len(matricies)),
	}

	for i := range matricies {
		response.GradeProgresses = append(response.GradeProgresses, &proto.Matrix{
			Id:        matricies[i].Id,
			VacancyId: matricies[i].VacancyId,
			GradesIds: matricies[i].GradesIds,
		})
	}

	return response, nil
}

func (gp *gatewayProxy) MatrixAdd(ctx context.Context, request *proto.MatrixAddRequest) (*proto.Matrix, error) {
	matrix := &repository.Matrix{
		VacancyId: request.Params.GetVacancyId(),
		GradesIds: repository.Uint32Array(request.Params.GetGradesIds()),
	}

	matrix, err := repository.Matricies.Add(ctx, matrix)
	if err != nil {
		return nil, err
	}

	return &proto.Matrix{
		Id:        matrix.Id,
		VacancyId: matrix.VacancyId,
		GradesIds: matrix.GradesIds,
	}, nil
}

func (gp *gatewayProxy) MatrixEdit(ctx context.Context, request *proto.MatrixEditRequest) (*proto.Matrix, error) {
	log.Print(request.GetId())
	matrix, err := repository.Matricies.Get(ctx, request.GetId())
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
		case "vacancyId":
			matrix.VacancyId = request.Params.GetVacancyId()
		case "gradesIds":
			matrix.GradesIds = repository.Uint32Array(request.Params.GetGradesIds())
		}
	}

	matrix, err = repository.Matricies.Edit(ctx, matrix)
	if err != nil {
		return nil, err
	}

	return &proto.Matrix{
		Id:        matrix.Id,
		VacancyId: matrix.VacancyId,
		GradesIds: matrix.GradesIds,
	}, nil
}

func (gp *gatewayProxy) MatrixRemove(ctx context.Context, request *proto.MatrixRemoveRequest) (*proto.MatrixRemoveResponse, error) {
	err := repository.Matricies.Remove(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.MatrixRemoveResponse{Id: request.GetId()}, nil
}
