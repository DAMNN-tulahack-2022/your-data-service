package transport

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (gp *gatewayProxy) SkillsList(ctx context.Context, request *emptypb.Empty) (*proto.SkillsListResponse, error) {
	skills, err := repository.Skills.List(ctx)
	if err != nil {
		return nil, err
	}

	response := &proto.SkillsListResponse{
		Skills: make([]*proto.Skill, 0, len(skills)),
	}

	for i := range skills {
		response.Skills = append(response.Skills, &proto.Skill{
			Id:    uint32(skills[i].Id),
			Label: skills[i].Label,
		})
	}

	return response, nil
}

func (gp *gatewayProxy) SkillAdd(ctx context.Context, request *proto.SkillAddRequest) (*proto.Skill, error) {
	skill, err := repository.Skills.Add(ctx, request.GetParams().GetLabel())
	if err != nil {
		return nil, err
	}

	return &proto.Skill{Id: uint32(skill.Id), Label: skill.Label}, nil
}

func (gp *gatewayProxy) SkillEdit(ctx context.Context, request *proto.SkillEditRequest) (*proto.Skill, error) {
	response := &proto.Skill{}

	return response, nil
}

func (gp *gatewayProxy) SkillRemove(ctx context.Context, request *proto.SkillRemoveRequest) (*proto.SkillRemoveResponse, error) {
	response := &proto.SkillRemoveResponse{}

	return response, nil
}
