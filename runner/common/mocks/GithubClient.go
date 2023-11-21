// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v55/github"
	mock "github.com/stretchr/testify/mock"
)

// GithubClient is an autogenerated mock type for the GithubClient type
type GithubClient struct {
	mock.Mock
}

// CreateOrgHook provides a mock function with given fields: ctx, org, hook
func (_m *GithubClient) CreateOrgHook(ctx context.Context, org string, hook *github.Hook) (*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, org, hook)

	var r0 *github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.Hook) (*github.Hook, *github.Response, error)); ok {
		return rf(ctx, org, hook)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.Hook) *github.Hook); ok {
		r0 = rf(ctx, org, hook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.Hook) *github.Response); ok {
		r1 = rf(ctx, org, hook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.Hook) error); ok {
		r2 = rf(ctx, org, hook)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateOrganizationRegistrationToken provides a mock function with given fields: ctx, owner
func (_m *GithubClient) CreateOrganizationRegistrationToken(ctx context.Context, owner string) (*github.RegistrationToken, *github.Response, error) {
	ret := _m.Called(ctx, owner)

	var r0 *github.RegistrationToken
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*github.RegistrationToken, *github.Response, error)); ok {
		return rf(ctx, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *github.RegistrationToken); ok {
		r0 = rf(ctx, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RegistrationToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *github.Response); ok {
		r1 = rf(ctx, owner)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, owner)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateRegistrationToken provides a mock function with given fields: ctx, owner, repo
func (_m *GithubClient) CreateRegistrationToken(ctx context.Context, owner string, repo string) (*github.RegistrationToken, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 *github.RegistrationToken
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*github.RegistrationToken, *github.Response, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *github.RegistrationToken); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RegistrationToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateRepoHook provides a mock function with given fields: ctx, owner, repo, hook
func (_m *GithubClient) CreateRepoHook(ctx context.Context, owner string, repo string, hook *github.Hook) (*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, hook)

	var r0 *github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.Hook) (*github.Hook, *github.Response, error)); ok {
		return rf(ctx, owner, repo, hook)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.Hook) *github.Hook); ok {
		r0 = rf(ctx, owner, repo, hook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.Hook) *github.Response); ok {
		r1 = rf(ctx, owner, repo, hook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.Hook) error); ok {
		r2 = rf(ctx, owner, repo, hook)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteOrgHook provides a mock function with given fields: ctx, org, id
func (_m *GithubClient) DeleteOrgHook(ctx context.Context, org string, id int64) (*github.Response, error) {
	ret := _m.Called(ctx, org, id)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*github.Response, error)); ok {
		return rf(ctx, org, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *github.Response); ok {
		r0 = rf(ctx, org, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, org, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepoHook provides a mock function with given fields: ctx, owner, repo, id
func (_m *GithubClient) DeleteRepoHook(ctx context.Context, owner string, repo string, id int64) (*github.Response, error) {
	ret := _m.Called(ctx, owner, repo, id)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*github.Response, error)); ok {
		return rf(ctx, owner, repo, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Response); ok {
		r0 = rf(ctx, owner, repo, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, owner, repo, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateOrgJITConfig provides a mock function with given fields: ctx, owner, request
func (_m *GithubClient) GenerateOrgJITConfig(ctx context.Context, owner string, request *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error) {
	ret := _m.Called(ctx, owner, request)

	var r0 *github.JITRunnerConfig
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error)); ok {
		return rf(ctx, owner, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.GenerateJITConfigRequest) *github.JITRunnerConfig); ok {
		r0 = rf(ctx, owner, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.JITRunnerConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.GenerateJITConfigRequest) *github.Response); ok {
		r1 = rf(ctx, owner, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.GenerateJITConfigRequest) error); ok {
		r2 = rf(ctx, owner, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GenerateRepoJITConfig provides a mock function with given fields: ctx, owner, repo, request
func (_m *GithubClient) GenerateRepoJITConfig(ctx context.Context, owner string, repo string, request *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, request)

	var r0 *github.JITRunnerConfig
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.GenerateJITConfigRequest) (*github.JITRunnerConfig, *github.Response, error)); ok {
		return rf(ctx, owner, repo, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.GenerateJITConfigRequest) *github.JITRunnerConfig); ok {
		r0 = rf(ctx, owner, repo, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.JITRunnerConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.GenerateJITConfigRequest) *github.Response); ok {
		r1 = rf(ctx, owner, repo, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.GenerateJITConfigRequest) error); ok {
		r2 = rf(ctx, owner, repo, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrgHook provides a mock function with given fields: ctx, org, id
func (_m *GithubClient) GetOrgHook(ctx context.Context, org string, id int64) (*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, org, id)

	var r0 *github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*github.Hook, *github.Response, error)); ok {
		return rf(ctx, org, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *github.Hook); ok {
		r0 = rf(ctx, org, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) *github.Response); ok {
		r1 = rf(ctx, org, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, org, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRepoHook provides a mock function with given fields: ctx, owner, repo, id
func (_m *GithubClient) GetRepoHook(ctx context.Context, owner string, repo string, id int64) (*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, id)

	var r0 *github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*github.Hook, *github.Response, error)); ok {
		return rf(ctx, owner, repo, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Hook); ok {
		r0 = rf(ctx, owner, repo, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) *github.Response); ok {
		r1 = rf(ctx, owner, repo, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int64) error); ok {
		r2 = rf(ctx, owner, repo, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWorkflowJobByID provides a mock function with given fields: ctx, owner, repo, jobID
func (_m *GithubClient) GetWorkflowJobByID(ctx context.Context, owner string, repo string, jobID int64) (*github.WorkflowJob, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, jobID)

	var r0 *github.WorkflowJob
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*github.WorkflowJob, *github.Response, error)); ok {
		return rf(ctx, owner, repo, jobID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.WorkflowJob); ok {
		r0 = rf(ctx, owner, repo, jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.WorkflowJob)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) *github.Response); ok {
		r1 = rf(ctx, owner, repo, jobID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int64) error); ok {
		r2 = rf(ctx, owner, repo, jobID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListOrgHooks provides a mock function with given fields: ctx, org, opts
func (_m *GithubClient) ListOrgHooks(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, org, opts)

	var r0 []*github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) ([]*github.Hook, *github.Response, error)); ok {
		return rf(ctx, org, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) []*github.Hook); ok {
		r0 = rf(ctx, org, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, org, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, org, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListOrganizationRunnerApplicationDownloads provides a mock function with given fields: ctx, owner
func (_m *GithubClient) ListOrganizationRunnerApplicationDownloads(ctx context.Context, owner string) ([]*github.RunnerApplicationDownload, *github.Response, error) {
	ret := _m.Called(ctx, owner)

	var r0 []*github.RunnerApplicationDownload
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*github.RunnerApplicationDownload, *github.Response, error)); ok {
		return rf(ctx, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*github.RunnerApplicationDownload); ok {
		r0 = rf(ctx, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RunnerApplicationDownload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *github.Response); ok {
		r1 = rf(ctx, owner)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, owner)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListOrganizationRunnerGroups provides a mock function with given fields: ctx, org, opts
func (_m *GithubClient) ListOrganizationRunnerGroups(ctx context.Context, org string, opts *github.ListOrgRunnerGroupOptions) (*github.RunnerGroups, *github.Response, error) {
	ret := _m.Called(ctx, org, opts)

	var r0 *github.RunnerGroups
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOrgRunnerGroupOptions) (*github.RunnerGroups, *github.Response, error)); ok {
		return rf(ctx, org, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOrgRunnerGroupOptions) *github.RunnerGroups); ok {
		r0 = rf(ctx, org, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RunnerGroups)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.ListOrgRunnerGroupOptions) *github.Response); ok {
		r1 = rf(ctx, org, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.ListOrgRunnerGroupOptions) error); ok {
		r2 = rf(ctx, org, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListOrganizationRunners provides a mock function with given fields: ctx, owner, opts
func (_m *GithubClient) ListOrganizationRunners(ctx context.Context, owner string, opts *github.ListOptions) (*github.Runners, *github.Response, error) {
	ret := _m.Called(ctx, owner, opts)

	var r0 *github.Runners
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) (*github.Runners, *github.Response, error)); ok {
		return rf(ctx, owner, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.ListOptions) *github.Runners); ok {
		r0 = rf(ctx, owner, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Runners)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRepoHooks provides a mock function with given fields: ctx, owner, repo, opts
func (_m *GithubClient) ListRepoHooks(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opts)

	var r0 []*github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) ([]*github.Hook, *github.Response, error)); ok {
		return rf(ctx, owner, repo, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) []*github.Hook); ok {
		r0 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, repo, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRunnerApplicationDownloads provides a mock function with given fields: ctx, owner, repo
func (_m *GithubClient) ListRunnerApplicationDownloads(ctx context.Context, owner string, repo string) ([]*github.RunnerApplicationDownload, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 []*github.RunnerApplicationDownload
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]*github.RunnerApplicationDownload, *github.Response, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*github.RunnerApplicationDownload); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RunnerApplicationDownload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRunners provides a mock function with given fields: ctx, owner, repo, opts
func (_m *GithubClient) ListRunners(ctx context.Context, owner string, repo string, opts *github.ListOptions) (*github.Runners, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opts)

	var r0 *github.Runners
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) (*github.Runners, *github.Response, error)); ok {
		return rf(ctx, owner, repo, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) *github.Runners); ok {
		r0 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Runners)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, repo, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PingOrgHook provides a mock function with given fields: ctx, org, id
func (_m *GithubClient) PingOrgHook(ctx context.Context, org string, id int64) (*github.Response, error) {
	ret := _m.Called(ctx, org, id)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*github.Response, error)); ok {
		return rf(ctx, org, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *github.Response); ok {
		r0 = rf(ctx, org, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, org, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingRepoHook provides a mock function with given fields: ctx, owner, repo, id
func (_m *GithubClient) PingRepoHook(ctx context.Context, owner string, repo string, id int64) (*github.Response, error) {
	ret := _m.Called(ctx, owner, repo, id)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*github.Response, error)); ok {
		return rf(ctx, owner, repo, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Response); ok {
		r0 = rf(ctx, owner, repo, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, owner, repo, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveOrganizationRunner provides a mock function with given fields: ctx, owner, runnerID
func (_m *GithubClient) RemoveOrganizationRunner(ctx context.Context, owner string, runnerID int64) (*github.Response, error) {
	ret := _m.Called(ctx, owner, runnerID)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*github.Response, error)); ok {
		return rf(ctx, owner, runnerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *github.Response); ok {
		r0 = rf(ctx, owner, runnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, owner, runnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRunner provides a mock function with given fields: ctx, owner, repo, runnerID
func (_m *GithubClient) RemoveRunner(ctx context.Context, owner string, repo string, runnerID int64) (*github.Response, error) {
	ret := _m.Called(ctx, owner, repo, runnerID)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*github.Response, error)); ok {
		return rf(ctx, owner, repo, runnerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Response); ok {
		r0 = rf(ctx, owner, repo, runnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, owner, repo, runnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGithubClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGithubClient creates a new instance of GithubClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGithubClient(t mockConstructorTestingTNewGithubClient) *GithubClient {
	mock := &GithubClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
