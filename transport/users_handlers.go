package transport

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gp *gatewayProxy) UpUserExp(ctx context.Context, request *proto.UpUserExpRequest) (*proto.User, error) {
	user, err := repository.Users.UpExp(ctx, request.GetUserId(), request.GetExperience())
	if err != nil {
		return nil, err
	}

	return &proto.User{
		Id:                  user.Id,
		Login:               user.Login,
		ViewedPostsIds:      user.ViewedPostIds,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		MiddleName:          user.MiddleName,
		AvatarUrl:           user.AvaterURL,
		CurrentProjectId:    user.CurrentProjectId,
		CompletedProjectIds: user.CompletedProjectsIds,
		SkillsIds:           user.SkillsIds,
		Role:                user.Role,
		TotalExperience:     user.TotalExp,
	}, nil
}

func (gp *gatewayProxy) PublicPost(ctx context.Context, request *proto.PublicPostRequest) (*proto.PublicPostResponse, error) {
	article := &repository.Article{
		AuthorId:    request.GetUserId(),
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		SkillsIds:   repository.Uint32Array(request.GetSkillsIds()),
	}

	article, err := repository.Articles.Add(ctx, article)
	if err != nil {
		return nil, err
	}

	return &proto.PublicPostResponse{
		Post: &proto.Article{
			Id:          article.Id,
			AuthorId:    article.AuthorId,
			Title:       article.Title,
			Description: article.Description,
			TotalViewed: article.TotalViewed,
			SkillsIds:   article.SkillsIds,
			Experience:  article.Exp,
		},
	}, nil
}

func (gp *gatewayProxy) ReadPost(ctx context.Context, request *proto.ReadPostRequest) (*proto.ReadPostResponse, error) {
	user, err := repository.Users.Get(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "errors.validation.userNotFound")
	}

	for _, id := range user.ViewedPostIds {
		if id == request.PostId {
			return nil, status.Error(codes.InvalidArgument, "errors.validation.postAlreadyRead")
		}
	}

	tx, err := repository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	article, err := repository.Articles.View(ctx, tx, request.GetPostId())
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.NotFound, "errors.validation.postNotFound")
	}

	user.TotalExp += article.Exp
	user.ViewedPostIds = append(user.ViewedPostIds, request.PostId)
	_, err = repository.Users.ReadPost(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &proto.ReadPostResponse{
		PostId: article.Id,
		UserId: request.UserId,
	}, nil
}
