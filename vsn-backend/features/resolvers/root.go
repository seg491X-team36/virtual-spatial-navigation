package resolvers

import "github.com/seg491X-team36/vsn-backend/codegen/graph"

type Root struct {
	graph.ExperimentResolver
	graph.ExperimentResultResolver
	graph.ExperimentConfigResolver
	graph.InviteResolver
	graph.MutationResolver
	graph.QueryResolver
	graph.UserResolver
}

func (r *Root) Experiment() graph.ExperimentResolver {
	return r.ExperimentResolver
}

func (r *Root) ExperimentResult() graph.ExperimentResultResolver {
	return r.ExperimentResultResolver
}

func (r *Root) ExperimentConfig() graph.ExperimentConfigResolver {
	return r.ExperimentConfigResolver
}

func (r *Root) Invite() graph.InviteResolver {
	return r.InviteResolver
}

func (r *Root) Mutation() graph.MutationResolver {
	return r.MutationResolver
}

func (r *Root) Query() graph.QueryResolver {
	return r.QueryResolver
}

func (r *Root) User() graph.UserResolver {
	return r.UserResolver
}
