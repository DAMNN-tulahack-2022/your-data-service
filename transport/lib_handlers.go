package transport

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/repository"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (gp *gatewayProxy) DataLib(ctx context.Context, request *emptypb.Empty) (*proto.DataLibResponse, error) {
	response := new(proto.DataLibResponse)
	g := new(errgroup.Group)

	g.Go(func() error {
		articles, err := repository.Articles.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.Articles = make([]*proto.Article, 0, len(articles))
		for i := range articles {
			response.Articles = append(response.Articles, &proto.Article{
				Id:          articles[i].Id,
				Title:       articles[i].Title,
				Description: articles[i].Description,
				AuthorId:    articles[i].AuthorId,
				TotalViewed: articles[i].TotalViewed,
				SkillsIds:   articles[i].SkillsIds,
				Experience:  articles[i].Exp,
			})
		}
		return nil
	})

	g.Go(func() error {
		projects, err := repository.Projects.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.Projects = make([]*proto.Project, 0, len(projects))
		for i := range projects {
			response.Projects = append(response.Projects, &proto.Project{
				Id:          projects[i].Id,
				Title:       projects[i].Title,
				Description: projects[i].Description,
				TeamleadId:  projects[i].TeamLeadId,
				SkillsIds:   projects[i].SkillsIds,
				UserIds:     projects[i].UserIds,
				Experience:  projects[i].Exp,
			})
		}

		return nil
	})

	g.Go(func() error {
		users, err := repository.Users.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.Users = make([]*proto.User, 0, len(users))
		for i := range users {
			response.Users = append(response.Users, &proto.User{
				Id:                  users[i].Id,
				Login:               users[i].Login,
				ViewedPostsIds:      users[i].ViewedPostIds,
				FirstName:           users[i].FirstName,
				LastName:            users[i].LastName,
				MiddleName:          users[i].MiddleName,
				AvatarUrl:           users[i].AvaterURL,
				CurrentProjectId:    users[i].CurrentProjectId,
				CompletedProjectIds: users[i].CompletedProjectsIds,
				SkillsIds:           users[i].SkillsIds,
				Role:                users[i].Role,
				TotalExperience:     users[i].TotalExp,
			})
		}
		return nil
	})

	g.Go(func() error {
		matrixes, err := repository.Matricies.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.VacanciesProgresses = make([]*proto.VacancyProgress, 0, len(matrixes))
		for i := range matrixes {
			response.VacanciesProgresses = append(response.VacanciesProgresses, &proto.VacancyProgress{
				Id:        matrixes[i].Id,
				VacancyId: matrixes[i].VacancyId,
				GradeIds:  matrixes[i].GradesIds,
			})
		}
		return nil
	})

	g.Go(func() error {
		grades, err := repository.Grades.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.Grades = make([]*proto.Grade, 0, len(grades))
		for i := range grades {
			response.Grades = append(response.Grades, &proto.Grade{
				Id:            grades[i].Id,
				Label:         grades[i].Label,
				Description:   grades[i].Description,
				Experience:    grades[i].Exp,
				NeedsApproval: grades[i].NeedsApproval,
			})
		}
		return nil
	})

	g.Go(func() error {
		vacancies, err := repository.Vacancies.List(ctx)
		if err != nil {
			return status.Error(codes.Internal, "errors.common.internal: "+err.Error())
		}

		response.Vacancies = make([]*proto.Vacancy, 0, len(vacancies))
		for i := range vacancies {
			response.Vacancies = append(response.Vacancies, &proto.Vacancy{
				Id:    vacancies[i].Id,
				Label: vacancies[i].Label,
			})
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return response, nil
}
