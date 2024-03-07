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

// GetApprovalInstance 通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/list) 接口获取。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/get
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/instance/get
func (r *ApprovalService) GetApprovalInstance(ctx context.Context, request *GetApprovalInstanceReq, options ...MethodOptionFunc) (*GetApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalInstance != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalInstance mock enable")
		return r.cli.mock.mockApprovalGetApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApprovalInstance",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/:instance_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalGetApprovalInstance mock ApprovalGetApprovalInstance method
func (r *Mock) MockApprovalGetApprovalInstance(f func(ctx context.Context, request *GetApprovalInstanceReq, options ...MethodOptionFunc) (*GetApprovalInstanceResp, *Response, error)) {
	r.mockApprovalGetApprovalInstance = f
}

// UnMockApprovalGetApprovalInstance un-mock ApprovalGetApprovalInstance method
func (r *Mock) UnMockApprovalGetApprovalInstance() {
	r.mockApprovalGetApprovalInstance = nil
}

// GetApprovalInstanceReq ...
type GetApprovalInstanceReq struct {
	InstanceID string  `path:"instance_id" json:"-"`   // 审批实例 Code, 若在创建的时候传了 uuid, 也可以通过传 uuid 获取, 示例值: "81D31358-93AF-92D6-7425-01A5D67C4E71"
	Locale     *string `query:"locale" json:"-"`       // 语言。默认值为[创建审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/create)时在 i18n_resources 字段中配置的语言, 示例值: zh-CN, 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	UserID     *string `query:"user_id" json:"-"`      // 发起审批用户 id, 仅自建应用可返回, 示例值: f7cb567e
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: user_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApprovalInstanceResp ...
type GetApprovalInstanceResp struct {
	ApprovalName         string                             `json:"approval_name,omitempty"`          // 审批名称
	StartTime            string                             `json:"start_time,omitempty"`             // 审批创建时间
	EndTime              string                             `json:"end_time,omitempty"`               // 审批完成时间, 未完成为 0
	UserID               string                             `json:"user_id,omitempty"`                // 发起审批用户
	OpenID               string                             `json:"open_id,omitempty"`                // 发起审批用户 open id
	SerialNumber         string                             `json:"serial_number,omitempty"`          // 审批单编号
	DepartmentID         string                             `json:"department_id,omitempty"`          // 发起审批用户所在部门
	Status               string                             `json:"status,omitempty"`                 // 审批实例状态, 可选值有: PENDING: 审批中, APPROVED: 通过, REJECTED: 拒绝, CANCELED: 撤回, DELETED: 删除
	UUID                 string                             `json:"uuid,omitempty"`                   // 用户的唯一标识 id
	Form                 ApprovalWidgetList                 `json:"form,omitempty"`                   // json 字符串, 控件值详情见下方
	TaskList             []*GetApprovalInstanceRespTask     `json:"task_list,omitempty"`              // 审批任务列表
	CommentList          []*GetApprovalInstanceRespComment  `json:"comment_list,omitempty"`           // 评论列表
	Timeline             []*GetApprovalInstanceRespTimeline `json:"timeline,omitempty"`               // 审批动态
	ModifiedInstanceCode string                             `json:"modified_instance_code,omitempty"` // 修改的原实例 code, 仅在查询修改实例时显示该字段
	RevertedInstanceCode string                             `json:"reverted_instance_code,omitempty"` // 撤销的原实例 code, 仅在查询撤销实例时显示该字段
	ApprovalCode         string                             `json:"approval_code,omitempty"`          // 审批定义 Code
	Reverted             bool                               `json:"reverted,omitempty"`               // 单据是否被撤销
	InstanceCode         string                             `json:"instance_code,omitempty"`          // 审批实例 Code
}

// GetApprovalInstanceRespComment ...
type GetApprovalInstanceRespComment struct {
	ID         string                                `json:"id,omitempty"`          // 评论 id
	UserID     string                                `json:"user_id,omitempty"`     // 发表评论用户
	OpenID     string                                `json:"open_id,omitempty"`     // 发表评论用户 open id
	Comment    string                                `json:"comment,omitempty"`     // 评论内容
	CreateTime string                                `json:"create_time,omitempty"` // 1564590532967
	Files      []*GetApprovalInstanceRespCommentFile `json:"files,omitempty"`       // 评论附件
}

// GetApprovalInstanceRespCommentFile ...
type GetApprovalInstanceRespCommentFile struct {
	URL      string `json:"url,omitempty"`       // 附件路径
	FileSize int64  `json:"file_size,omitempty"` // 附件大小
	Title    string `json:"title,omitempty"`     // 附件标题
	Type     string `json:"type,omitempty"`      // 附件类别
}

// GetApprovalInstanceRespTask ...
type GetApprovalInstanceRespTask struct {
	ID           string `json:"id,omitempty"`             // task id
	UserID       string `json:"user_id,omitempty"`        // 审批人的用户 id, 自动通过、自动拒绝 时为空
	OpenID       string `json:"open_id,omitempty"`        // 审批人 open id
	Status       string `json:"status,omitempty"`         // 任务状态, 可选值有: PENDING: 审批中, APPROVED: 通过, REJECTED: 拒绝, TRANSFERRED: 已转交, DONE: 完成
	NodeID       string `json:"node_id,omitempty"`        // task 所属节点 id
	NodeName     string `json:"node_name,omitempty"`      // task 所属节点名称
	CustomNodeID string `json:"custom_node_id,omitempty"` // task 所属节点自定义 id, 如果没设置自定义 id, 则不返回该字段
	Type         string `json:"type,omitempty"`           // 审批方式, 可选值有: AND: 会签, OR: 或签, AUTO_PASS: 自动通过, AUTO_REJECT: 自动拒绝, SEQUENTIAL: 按顺序
	StartTime    string `json:"start_time,omitempty"`     // task 开始时间
	EndTime      string `json:"end_time,omitempty"`       // task 完成时间, 未完成为 0
}

// GetApprovalInstanceRespTimeline ...
type GetApprovalInstanceRespTimeline struct {
	Type       string                                   `json:"type,omitempty"`         // 动态类型, 不同类型 ext 内的 user_id_list 含义不一样, 可选值有: START: 审批开始, PASS: 通过, REJECT: 拒绝, AUTO_PASS: 自动通过, AUTO_REJECT: 自动拒绝, REMOVE_REPEAT: 去重, TRANSFER: 转交, ADD_APPROVER_BEFORE: 前加签, ADD_APPROVER: 并加签, ADD_APPROVER_AFTER: 后加签, DELETE_APPROVER: 减签, ROLLBACK_SELECTED: 指定回退, ROLLBACK: 全部回退, CANCEL: 撤回, DELETE: 删除, CC: 抄送
	CreateTime string                                   `json:"create_time,omitempty"`  // 发生时间
	UserID     string                                   `json:"user_id,omitempty"`      // 动态产生用户
	OpenID     string                                   `json:"open_id,omitempty"`      // 动态产生用户 open id
	UserIDList []string                                 `json:"user_id_list,omitempty"` // 被抄送人列表
	OpenIDList []string                                 `json:"open_id_list,omitempty"` // 被抄送人列表
	TaskID     string                                   `json:"task_id,omitempty"`      // 产生动态关联的task_id
	Comment    string                                   `json:"comment,omitempty"`      // 理由
	CcUserList []*GetApprovalInstanceRespTimelineCcUser `json:"cc_user_list,omitempty"` // 抄送人列表
	Ext        *GetApprovalInstanceRespTimelineExt      `json:"ext,omitempty"`          // 动态其他信息, json格式, 目前包括 user_id_list, user_id, open_id_list, open_id
	NodeKey    string                                   `json:"node_key,omitempty"`     // 产生 task 的节点 key
	Files      []*GetApprovalInstanceRespTimelineFile   `json:"files,omitempty"`        // 审批附件
}

// GetApprovalInstanceRespTimelineCcUser ...
type GetApprovalInstanceRespTimelineCcUser struct {
	UserID string `json:"user_id,omitempty"` // 抄送人 user id
	CcID   string `json:"cc_id,omitempty"`   // 审批实例内抄送唯一标识
	OpenID string `json:"open_id,omitempty"` // 抄送人 open id
}

// GetApprovalInstanceRespTimelineFile ...
type GetApprovalInstanceRespTimelineFile struct {
	URL      string `json:"url,omitempty"`       // 附件路径
	FileSize int64  `json:"file_size,omitempty"` // 附件大小
	Title    string `json:"title,omitempty"`     // 附件标题
	Type     string `json:"type,omitempty"`      // 附件类别
}

// getApprovalInstanceResp ...
type getApprovalInstanceResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *GetApprovalInstanceResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
