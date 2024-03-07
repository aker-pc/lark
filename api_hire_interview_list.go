// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetHireInterviewList 根据投递 ID 或面试时间获取面试信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/interview/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/delivery-process-management/interview/list
func (r *HireService) GetHireInterviewList(ctx context.Context, request *GetHireInterviewListReq, options ...MethodOptionFunc) (*GetHireInterviewListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireInterviewList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireInterviewList mock enable")
		return r.cli.mock.mockHireGetHireInterviewList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireInterviewList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/interviews",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireInterviewListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireInterviewList mock HireGetHireInterviewList method
func (r *Mock) MockHireGetHireInterviewList(f func(ctx context.Context, request *GetHireInterviewListReq, options ...MethodOptionFunc) (*GetHireInterviewListResp, *Response, error)) {
	r.mockHireGetHireInterviewList = f
}

// UnMockHireGetHireInterviewList un-mock HireGetHireInterviewList method
func (r *Mock) UnMockHireGetHireInterviewList() {
	r.mockHireGetHireInterviewList = nil
}

// GetHireInterviewListReq ...
type GetHireInterviewListReq struct {
	PageSize       *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 10, 默认值: `10`, 最大值: `100`
	PageToken      *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: xx
	ApplicationID  *string `query:"application_id" json:"-"`    // 投递 ID, 示例值: 6134134355464633
	InterviewID    *string `query:"interview_id" json:"-"`      // 面试 ID, 示例值: 6888217964693309704
	StartTime      *string `query:"start_time" json:"-"`        // 最早开始时间, 格式为时间戳, 示例值: 1609489908000
	EndTime        *string `query:"end_time" json:"-"`          // 最晚开始时间, 格式为时间戳, 示例值: 1610489908000
	JobLevelIDType *IDType `query:"job_level_id_type" json:"-"` // 此次调用中使用的「职级 ID」的类型, 示例值: 6942778198054125570, 可选值有: people_admin_job_level_id: 「人力系统管理后台」适用的职级 ID。人力系统管理后台逐步下线中, 建议不继续使用此 ID。, job_level_id: 「飞书管理后台」适用的职级 ID, 通过[「获取租户职级列表」](https://open.feishu.cn/document/server-docs/contact-v3/job_level/list)接口获取, 默认值: `people_admin_job_level_id`
	UserIDType     *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireInterviewListResp ...
type GetHireInterviewListResp struct {
	Items     []*GetHireInterviewListRespItem `json:"items,omitempty"`      // 面试列表
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetHireInterviewListRespItem ...
type GetHireInterviewListRespItem struct {
	ID                         string                                          `json:"id,omitempty"`                           // 面试 ID
	BeginTime                  int64                                           `json:"begin_time,omitempty"`                   // 面试开始时间（ms）
	EndTime                    int64                                           `json:"end_time,omitempty"`                     // 面试结束时间（ms）
	Round                      int64                                           `json:"round,omitempty"`                        // 面试轮次
	InterviewRecordList        []*GetHireInterviewListRespItemInterviewRecord  `json:"interview_record_list,omitempty"`        // 面试记录信息
	FeedbackSubmitTime         int64                                           `json:"feedback_submit_time,omitempty"`         // 面试评价提交时间
	StageID                    string                                          `json:"stage_id,omitempty"`                     // 面试关联的投递阶段
	ApplicationID              string                                          `json:"application_id,omitempty"`               // 投递 ID
	Stage                      *GetHireInterviewListRespItemStage              `json:"stage,omitempty"`                        // 阶段信息
	Creator                    *GetHireInterviewListRespItemCreator            `json:"creator,omitempty"`                      // 创建人
	BizCreateTime              int64                                           `json:"biz_create_time,omitempty"`              // 创建时间（ms）
	BizModifyTime              int64                                           `json:"biz_modify_time,omitempty"`              // 最近更新时间（ms）
	InterviewRoundSummary      int64                                           `json:"interview_round_summary,omitempty"`      // 面试状态, 可选值有: 2: 未开始, 3: 全部未评价, 4: 全部通过, 5: 全部淘汰, 6: 爽约, 7: 部分评价且均评价通过, 8: 部分评价且评价中有通过有淘汰的, 9: 部分评价且均评价淘汰, 10: 所有面试官都提交评价且评价中有通过有淘汰的
	InterviewArrangementID     string                                          `json:"interview_arrangement_id,omitempty"`     // 面试安排 ID
	InterviewType              int64                                           `json:"interview_type,omitempty"`               // 面试类型, 可选值有: 1: 现场面试, 2: 电话面试, 3: 视频面试
	TalentTimeZone             *GetHireInterviewListRespItemTalentTimeZone     `json:"talent_time_zone,omitempty"`             // 候选人时区
	ContactUser                *GetHireInterviewListRespItemContactUser        `json:"contact_user,omitempty"`                 // 面试联系人
	ContactMobile              string                                          `json:"contact_mobile,omitempty"`               // 面试联系人电话
	Remark                     string                                          `json:"remark,omitempty"`                       // 备注
	Address                    *GetHireInterviewListRespItemAddress            `json:"address,omitempty"`                      // 面试地点
	VideoType                  int64                                           `json:"video_type,omitempty"`                   // 视频面试工具, 可选值有: 1: Zoom, 2: 牛客技术类型, 3: 牛客非技术类型, 4: 赛码, 5: 飞书, 8: Hackerrank, 9: 飞书（含代码考核）, 100: 不使用系统工具
	ArrangementStatus          int64                                           `json:"arrangement_status,omitempty"`           // 当安排类型为集中面试时, 此值表示集中面试的安排状态, 可选值有: 1: 未开始, 2: 进行中, 3: 已结束
	ArrangementType            int64                                           `json:"arrangement_type,omitempty"`             // 安排类型, 可选值有: 1: 社招单面, 2: 集中面试
	ArrangementAppointmentKind int64                                           `json:"arrangement_appointment_kind,omitempty"` // 安排方式（是否使用自助约面）, 可选值有: 1: 直接安排, 2: 自助约面
	MeetingRoomList            []*GetHireInterviewListRespItemMeetingRoom      `json:"meeting_room_list,omitempty"`            // 面试会议室
	InterviewRoundType         *GetHireInterviewListRespItemInterviewRoundType `json:"interview_round_type,omitempty"`         // 面试轮次类型
}

// GetHireInterviewListRespItemAddress ...
type GetHireInterviewListRespItemAddress struct {
	ID       string                                       `json:"id,omitempty"`       // 地址 ID
	Name     *GetHireInterviewListRespItemAddressName     `json:"name,omitempty"`     // 地址名称
	District *GetHireInterviewListRespItemAddressDistrict `json:"district,omitempty"` // 区域
	City     *GetHireInterviewListRespItemAddressCity     `json:"city,omitempty"`     // 城市
	State    *GetHireInterviewListRespItemAddressState    `json:"state,omitempty"`    // 省
	Country  *GetHireInterviewListRespItemAddressCountry  `json:"country,omitempty"`  // 国家
}

// GetHireInterviewListRespItemAddressCity ...
type GetHireInterviewListRespItemAddressCity struct {
	Code string                                       `json:"code,omitempty"` // 编码
	Name *GetHireInterviewListRespItemAddressCityName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemAddressCityName ...
type GetHireInterviewListRespItemAddressCityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemAddressCountry ...
type GetHireInterviewListRespItemAddressCountry struct {
	Code string                                          `json:"code,omitempty"` // 编码
	Name *GetHireInterviewListRespItemAddressCountryName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemAddressCountryName ...
type GetHireInterviewListRespItemAddressCountryName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemAddressDistrict ...
type GetHireInterviewListRespItemAddressDistrict struct {
	Code string                                           `json:"code,omitempty"` // 编码
	Name *GetHireInterviewListRespItemAddressDistrictName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemAddressDistrictName ...
type GetHireInterviewListRespItemAddressDistrictName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemAddressName ...
type GetHireInterviewListRespItemAddressName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemAddressState ...
type GetHireInterviewListRespItemAddressState struct {
	Code string                                        `json:"code,omitempty"` // 编码
	Name *GetHireInterviewListRespItemAddressStateName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemAddressStateName ...
type GetHireInterviewListRespItemAddressStateName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemContactUser ...
type GetHireInterviewListRespItemContactUser struct {
	ID   string                                       `json:"id,omitempty"`   // ID
	Name *GetHireInterviewListRespItemContactUserName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemContactUserName ...
type GetHireInterviewListRespItemContactUserName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemCreator ...
type GetHireInterviewListRespItemCreator struct {
	ID   string                                   `json:"id,omitempty"`   // ID
	Name *GetHireInterviewListRespItemCreatorName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemCreatorName ...
type GetHireInterviewListRespItemCreatorName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemInterviewRecord ...
type GetHireInterviewListRespItemInterviewRecord struct {
	ID             string                                                     `json:"id,omitempty"`              // 面试记录 ID
	UserID         string                                                     `json:"user_id,omitempty"`         // 面试官用户 ID
	Content        string                                                     `json:"content,omitempty"`         // 系统预设「记录」题目内容
	CommitStatus   int64                                                      `json:"commit_status,omitempty"`   // 提交状态, 可选值有: 1: 已提交, 2: 未提交
	Conclusion     int64                                                      `json:"conclusion,omitempty"`      // 面试结论, 可选值有: 1: 通过, 2: 未通过, 3: 未开始, 4: 未提交, 5: 未到场
	InterviewScore *GetHireInterviewListRespItemInterviewRecordInterviewScore `json:"interview_score,omitempty"` // 面试评分
	Interviewer    *GetHireInterviewListRespItemInterviewRecordInterviewer    `json:"interviewer,omitempty"`     // 面试官信息
}

// GetHireInterviewListRespItemInterviewRecordInterviewScore ...
type GetHireInterviewListRespItemInterviewRecordInterviewScore struct {
	ID            string `json:"id,omitempty"`             // 面试评分 ID
	Level         int64  `json:"level,omitempty"`          // 分数级别
	ZhName        string `json:"zh_name,omitempty"`        // 中文名称
	ZhDescription string `json:"zh_description,omitempty"` // 中文描述
	EnName        string `json:"en_name,omitempty"`        // 英文名称
	EnDescription string `json:"en_description,omitempty"` // 英文描述
}

// GetHireInterviewListRespItemInterviewRecordInterviewer ...
type GetHireInterviewListRespItemInterviewRecordInterviewer struct {
	ID   string                                                      `json:"id,omitempty"`   // ID
	Name *GetHireInterviewListRespItemInterviewRecordInterviewerName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemInterviewRecordInterviewerName ...
type GetHireInterviewListRespItemInterviewRecordInterviewerName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemInterviewRoundType ...
type GetHireInterviewListRespItemInterviewRoundType struct {
	ID   string                                              `json:"id,omitempty"`   // ID
	Name *GetHireInterviewListRespItemInterviewRoundTypeName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemInterviewRoundTypeName ...
type GetHireInterviewListRespItemInterviewRoundTypeName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemMeetingRoom ...
type GetHireInterviewListRespItemMeetingRoom struct {
	RoomID         string `json:"room_id,omitempty"`         // 会议室 ID
	RoomName       string `json:"room_name,omitempty"`       // 会议室名称
	BuildingName   string `json:"building_name,omitempty"`   // 建筑名称
	ReservedStatus int64  `json:"reserved_status,omitempty"` // 会议室预定状态, 可选值有: 1: 预约中, 2: 预约成功, 3: 预约失败
	FloorName      string `json:"floor_name,omitempty"`      // 楼层
}

// GetHireInterviewListRespItemStage ...
type GetHireInterviewListRespItemStage struct {
	ID   string                                 `json:"id,omitempty"`   // ID
	Name *GetHireInterviewListRespItemStageName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemStageName ...
type GetHireInterviewListRespItemStageName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireInterviewListRespItemTalentTimeZone ...
type GetHireInterviewListRespItemTalentTimeZone struct {
	Code string                                          `json:"code,omitempty"` // 编码
	Name *GetHireInterviewListRespItemTalentTimeZoneName `json:"name,omitempty"` // 名称
}

// GetHireInterviewListRespItemTalentTimeZoneName ...
type GetHireInterviewListRespItemTalentTimeZoneName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getHireInterviewListResp ...
type getHireInterviewListResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireInterviewListResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
