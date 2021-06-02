// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Contact_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Contact

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUser(ctx, &lark.CreateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteUser(ctx, &lark.DeleteUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUser(ctx, &lark.GetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUser(ctx, &lark.BatchGetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUserByID(ctx, &lark.BatchGetUserByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserList(ctx, &lark.GetUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUser(ctx, &lark.UpdateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDepartment(ctx, &lark.GetDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Contact

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactCreateUser(func(ctx context.Context, request *lark.CreateUserReq, options ...lark.MethodOptionFunc) (*lark.CreateUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactCreateUser()

			_, _, err := moduleCli.CreateUser(ctx, &lark.CreateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactDeleteUser(func(ctx context.Context, request *lark.DeleteUserReq, options ...lark.MethodOptionFunc) (*lark.DeleteUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactDeleteUser()

			_, _, err := moduleCli.DeleteUser(ctx, &lark.DeleteUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactGetUser(func(ctx context.Context, request *lark.GetUserReq, options ...lark.MethodOptionFunc) (*lark.GetUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactGetUser()

			_, _, err := moduleCli.GetUser(ctx, &lark.GetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactBatchGetUser(func(ctx context.Context, request *lark.BatchGetUserReq, options ...lark.MethodOptionFunc) (*lark.BatchGetUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactBatchGetUser()

			_, _, err := moduleCli.BatchGetUser(ctx, &lark.BatchGetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactBatchGetUserByID(func(ctx context.Context, request *lark.BatchGetUserByIDReq, options ...lark.MethodOptionFunc) (*lark.BatchGetUserByIDResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactBatchGetUserByID()

			_, _, err := moduleCli.BatchGetUserByID(ctx, &lark.BatchGetUserByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactGetUserList(func(ctx context.Context, request *lark.GetUserListReq, options ...lark.MethodOptionFunc) (*lark.GetUserListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactGetUserList()

			_, _, err := moduleCli.GetUserList(ctx, &lark.GetUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactUpdateUserPatch(func(ctx context.Context, request *lark.UpdateUserPatchReq, options ...lark.MethodOptionFunc) (*lark.UpdateUserPatchResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactUpdateUserPatch()

			_, _, err := moduleCli.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactUpdateUser(func(ctx context.Context, request *lark.UpdateUserReq, options ...lark.MethodOptionFunc) (*lark.UpdateUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactUpdateUser()

			_, _, err := moduleCli.UpdateUser(ctx, &lark.UpdateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactCreateDepartment(func(ctx context.Context, request *lark.CreateDepartmentReq, options ...lark.MethodOptionFunc) (*lark.CreateDepartmentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactCreateDepartment()

			_, _, err := moduleCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactGetDepartment(func(ctx context.Context, request *lark.GetDepartmentReq, options ...lark.MethodOptionFunc) (*lark.GetDepartmentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactGetDepartment()

			_, _, err := moduleCli.GetDepartment(ctx, &lark.GetDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactGetDepartmentList(func(ctx context.Context, request *lark.GetDepartmentListReq, options ...lark.MethodOptionFunc) (*lark.GetDepartmentListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactGetDepartmentList()

			_, _, err := moduleCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactGetParentDepartment(func(ctx context.Context, request *lark.GetParentDepartmentReq, options ...lark.MethodOptionFunc) (*lark.GetParentDepartmentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactGetParentDepartment()

			_, _, err := moduleCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactUpdateDepartmentPatch(func(ctx context.Context, request *lark.UpdateDepartmentPatchReq, options ...lark.MethodOptionFunc) (*lark.UpdateDepartmentPatchResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactUpdateDepartmentPatch()

			_, _, err := moduleCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactUpdateDepartment(func(ctx context.Context, request *lark.UpdateDepartmentReq, options ...lark.MethodOptionFunc) (*lark.UpdateDepartmentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactUpdateDepartment()

			_, _, err := moduleCli.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockContactDeleteDepartment(func(ctx context.Context, request *lark.DeleteDepartmentReq, options ...lark.MethodOptionFunc) (*lark.DeleteDepartmentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockContactDeleteDepartment()

			_, _, err := moduleCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Contact

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUser(ctx, &lark.CreateUserReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteUser(ctx, &lark.DeleteUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUser(ctx, &lark.GetUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUser(ctx, &lark.BatchGetUserReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUserByID(ctx, &lark.BatchGetUserByIDReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserList(ctx, &lark.GetUserListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUser(ctx, &lark.UpdateUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDepartment(ctx, &lark.GetDepartmentReq{
				DepartmentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{
				DepartmentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{
				DepartmentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{
				DepartmentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
