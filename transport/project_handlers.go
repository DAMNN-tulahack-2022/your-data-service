package transport

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
)

func (gp *gatewayProxy) ProjectAdd(ctx context.Context, request *proto.ProjectAddRequest) (*proto.Project, error) {
	project := &repository.Project{
		Id:          request.Params.GetId(),
		Title:       request.Params.GetTitle(),
		Description: request.Params.GetDescription(),
		SkillsIds:   repository.Uint32Array(request.Params.GetSkillsIds()),
		UserIds:     repository.Uint32Array(request.Params.GetUserIds()),
		TeamLeadId:  request.Params.GetTeamleadId(),
		Exp:         request.Params.GetExperience(),
	}

	project, err := repository.Projects.Add(ctx, project)
	if err != nil {
		return nil, err
	}

	return &proto.Project{
		Id:          project.Id,
		Title:       project.Title,
		Description: project.Description,
		SkillsIds:   project.SkillsIds,
		UserIds:     project.UserIds,
		TeamleadId:  project.TeamLeadId,
		Experience:  project.Exp,
	}, nil
}
