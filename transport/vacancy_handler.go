package transport

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (gp *gatewayProxy) VacanciesList(ctx context.Context, request *emptypb.Empty) (*proto.VacanciesListResponse, error) {
	vacancies, err := repository.Vacancies.List(ctx)
	if err != nil {
		return nil, err
	}

	response := &proto.VacanciesListResponse{
		Vacancies: make([]*proto.Vacancy, 0, len(vacancies)),
	}

	for i := range vacancies {
		response.Vacancies = append(response.Vacancies, &proto.Vacancy{
			Id:    uint32(vacancies[i].Id),
			Label: vacancies[i].Label,
		})
	}

	return response, nil
}

func (gp *gatewayProxy) VacancyAdd(ctx context.Context, request *proto.VacancyAddRequest) (*proto.Vacancy, error) {
	vacancy, err := repository.Vacancies.Add(ctx, request.GetParams().GetLabel())
	if err != nil {
		return nil, err
	}

	return &proto.Vacancy{Id: uint32(vacancy.Id), Label: vacancy.Label}, nil
}

func (gp *gatewayProxy) VacancyEdit(ctx context.Context, request *proto.VacancyEditRequest) (*proto.Vacancy, error) {
	response := &proto.Vacancy{}

	return response, nil
}

func (gp *gatewayProxy) VacancyRemove(ctx context.Context, request *proto.VacancyRemoveRequest) (*proto.VacancyRemoveResponse, error) {
	log.Print(request.Id)
	err := repository.Vacancies.Remove(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.VacancyRemoveResponse{Id: request.Id}, nil
}
