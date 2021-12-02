// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_MeetingRoom_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.MeetingRoom

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCustomization(ctx, &lark.GetMeetingRoomCustomizationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomSummary(ctx, &lark.BatchGetMeetingRoomSummaryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomBuildingList(ctx, &lark.GetMeetingRoomBuildingListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuilding(ctx, &lark.BatchGetMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomRoomList(ctx, &lark.GetMeetingRoomRoomListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoom(ctx, &lark.BatchGetMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomFreebusy(ctx, &lark.BatchGetMeetingRoomFreebusyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyMeetingRoomInstance(ctx, &lark.ReplyMeetingRoomInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomBuilding(ctx, &lark.CreateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomBuilding(ctx, &lark.UpdateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomBuilding(ctx, &lark.DeleteMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuildingID(ctx, &lark.BatchGetMeetingRoomBuildingIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomRoom(ctx, &lark.CreateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomRoom(ctx, &lark.UpdateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomRoom(ctx, &lark.DeleteMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoomID(ctx, &lark.BatchGetMeetingRoomRoomIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCountryList(ctx, &lark.GetMeetingRoomCountryListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomDistrictList(ctx, &lark.GetMeetingRoomDistrictListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.MeetingRoom

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomGetMeetingRoomCustomization(func(ctx context.Context, request *lark.GetMeetingRoomCustomizationReq, options ...lark.MethodOptionFunc) (*lark.GetMeetingRoomCustomizationResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomGetMeetingRoomCustomization()

			_, _, err := moduleCli.GetMeetingRoomCustomization(ctx, &lark.GetMeetingRoomCustomizationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomSummary(func(ctx context.Context, request *lark.BatchGetMeetingRoomSummaryReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomSummaryResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomSummary()

			_, _, err := moduleCli.BatchGetMeetingRoomSummary(ctx, &lark.BatchGetMeetingRoomSummaryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomGetMeetingRoomBuildingList(func(ctx context.Context, request *lark.GetMeetingRoomBuildingListReq, options ...lark.MethodOptionFunc) (*lark.GetMeetingRoomBuildingListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomGetMeetingRoomBuildingList()

			_, _, err := moduleCli.GetMeetingRoomBuildingList(ctx, &lark.GetMeetingRoomBuildingListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomBuilding(func(ctx context.Context, request *lark.BatchGetMeetingRoomBuildingReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomBuildingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomBuilding()

			_, _, err := moduleCli.BatchGetMeetingRoomBuilding(ctx, &lark.BatchGetMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomGetMeetingRoomRoomList(func(ctx context.Context, request *lark.GetMeetingRoomRoomListReq, options ...lark.MethodOptionFunc) (*lark.GetMeetingRoomRoomListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomGetMeetingRoomRoomList()

			_, _, err := moduleCli.GetMeetingRoomRoomList(ctx, &lark.GetMeetingRoomRoomListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomRoom(func(ctx context.Context, request *lark.BatchGetMeetingRoomRoomReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomRoomResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomRoom()

			_, _, err := moduleCli.BatchGetMeetingRoomRoom(ctx, &lark.BatchGetMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomFreebusy(func(ctx context.Context, request *lark.BatchGetMeetingRoomFreebusyReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomFreebusyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomFreebusy()

			_, _, err := moduleCli.BatchGetMeetingRoomFreebusy(ctx, &lark.BatchGetMeetingRoomFreebusyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomReplyMeetingRoomInstance(func(ctx context.Context, request *lark.ReplyMeetingRoomInstanceReq, options ...lark.MethodOptionFunc) (*lark.ReplyMeetingRoomInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomReplyMeetingRoomInstance()

			_, _, err := moduleCli.ReplyMeetingRoomInstance(ctx, &lark.ReplyMeetingRoomInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomCreateMeetingRoomBuilding(func(ctx context.Context, request *lark.CreateMeetingRoomBuildingReq, options ...lark.MethodOptionFunc) (*lark.CreateMeetingRoomBuildingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomCreateMeetingRoomBuilding()

			_, _, err := moduleCli.CreateMeetingRoomBuilding(ctx, &lark.CreateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomUpdateMeetingRoomBuilding(func(ctx context.Context, request *lark.UpdateMeetingRoomBuildingReq, options ...lark.MethodOptionFunc) (*lark.UpdateMeetingRoomBuildingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomUpdateMeetingRoomBuilding()

			_, _, err := moduleCli.UpdateMeetingRoomBuilding(ctx, &lark.UpdateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomDeleteMeetingRoomBuilding(func(ctx context.Context, request *lark.DeleteMeetingRoomBuildingReq, options ...lark.MethodOptionFunc) (*lark.DeleteMeetingRoomBuildingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomDeleteMeetingRoomBuilding()

			_, _, err := moduleCli.DeleteMeetingRoomBuilding(ctx, &lark.DeleteMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomBuildingID(func(ctx context.Context, request *lark.BatchGetMeetingRoomBuildingIDReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomBuildingIDResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomBuildingID()

			_, _, err := moduleCli.BatchGetMeetingRoomBuildingID(ctx, &lark.BatchGetMeetingRoomBuildingIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomCreateMeetingRoomRoom(func(ctx context.Context, request *lark.CreateMeetingRoomRoomReq, options ...lark.MethodOptionFunc) (*lark.CreateMeetingRoomRoomResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomCreateMeetingRoomRoom()

			_, _, err := moduleCli.CreateMeetingRoomRoom(ctx, &lark.CreateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomUpdateMeetingRoomRoom(func(ctx context.Context, request *lark.UpdateMeetingRoomRoomReq, options ...lark.MethodOptionFunc) (*lark.UpdateMeetingRoomRoomResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomUpdateMeetingRoomRoom()

			_, _, err := moduleCli.UpdateMeetingRoomRoom(ctx, &lark.UpdateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomDeleteMeetingRoomRoom(func(ctx context.Context, request *lark.DeleteMeetingRoomRoomReq, options ...lark.MethodOptionFunc) (*lark.DeleteMeetingRoomRoomResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomDeleteMeetingRoomRoom()

			_, _, err := moduleCli.DeleteMeetingRoomRoom(ctx, &lark.DeleteMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomBatchGetMeetingRoomRoomID(func(ctx context.Context, request *lark.BatchGetMeetingRoomRoomIDReq, options ...lark.MethodOptionFunc) (*lark.BatchGetMeetingRoomRoomIDResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomBatchGetMeetingRoomRoomID()

			_, _, err := moduleCli.BatchGetMeetingRoomRoomID(ctx, &lark.BatchGetMeetingRoomRoomIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomGetMeetingRoomCountryList(func(ctx context.Context, request *lark.GetMeetingRoomCountryListReq, options ...lark.MethodOptionFunc) (*lark.GetMeetingRoomCountryListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomGetMeetingRoomCountryList()

			_, _, err := moduleCli.GetMeetingRoomCountryList(ctx, &lark.GetMeetingRoomCountryListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockMeetingRoomGetMeetingRoomDistrictList(func(ctx context.Context, request *lark.GetMeetingRoomDistrictListReq, options ...lark.MethodOptionFunc) (*lark.GetMeetingRoomDistrictListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMeetingRoomGetMeetingRoomDistrictList()

			_, _, err := moduleCli.GetMeetingRoomDistrictList(ctx, &lark.GetMeetingRoomDistrictListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.MeetingRoom

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCustomization(ctx, &lark.GetMeetingRoomCustomizationReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomSummary(ctx, &lark.BatchGetMeetingRoomSummaryReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomBuildingList(ctx, &lark.GetMeetingRoomBuildingListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuilding(ctx, &lark.BatchGetMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomRoomList(ctx, &lark.GetMeetingRoomRoomListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoom(ctx, &lark.BatchGetMeetingRoomRoomReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomFreebusy(ctx, &lark.BatchGetMeetingRoomFreebusyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyMeetingRoomInstance(ctx, &lark.ReplyMeetingRoomInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomBuilding(ctx, &lark.CreateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomBuilding(ctx, &lark.UpdateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomBuilding(ctx, &lark.DeleteMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuildingID(ctx, &lark.BatchGetMeetingRoomBuildingIDReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomRoom(ctx, &lark.CreateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomRoom(ctx, &lark.UpdateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomRoom(ctx, &lark.DeleteMeetingRoomRoomReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoomID(ctx, &lark.BatchGetMeetingRoomRoomIDReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCountryList(ctx, &lark.GetMeetingRoomCountryListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomDistrictList(ctx, &lark.GetMeetingRoomDistrictListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.MeetingRoom
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCustomization(ctx, &lark.GetMeetingRoomCustomizationReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomSummary(ctx, &lark.BatchGetMeetingRoomSummaryReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomBuildingList(ctx, &lark.GetMeetingRoomBuildingListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuilding(ctx, &lark.BatchGetMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomRoomList(ctx, &lark.GetMeetingRoomRoomListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoom(ctx, &lark.BatchGetMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomFreebusy(ctx, &lark.BatchGetMeetingRoomFreebusyReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ReplyMeetingRoomInstance(ctx, &lark.ReplyMeetingRoomInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomBuilding(ctx, &lark.CreateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomBuilding(ctx, &lark.UpdateMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomBuilding(ctx, &lark.DeleteMeetingRoomBuildingReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomBuildingID(ctx, &lark.BatchGetMeetingRoomBuildingIDReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateMeetingRoomRoom(ctx, &lark.CreateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateMeetingRoomRoom(ctx, &lark.UpdateMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteMeetingRoomRoom(ctx, &lark.DeleteMeetingRoomRoomReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetMeetingRoomRoomID(ctx, &lark.BatchGetMeetingRoomRoomIDReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomCountryList(ctx, &lark.GetMeetingRoomCountryListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetMeetingRoomDistrictList(ctx, &lark.GetMeetingRoomDistrictListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
