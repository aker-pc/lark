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

// UpdateMessageEdit 编辑已发送的消息内容, 当前支持编辑文本、富文本消息。如需更新消息卡片, 请参考[更新应用发送的消息卡片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch)。
//
// 注意事项:
// - 一条消息最多可编辑20次
// - 仅可编辑操作者自己发送的消息
// - 不可编辑已撤回, 已删除, 超出可编辑时间的消息
// - 操作者必须在消息所属的群中且具备发言权限才可以编辑消息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/update
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message/update
func (r *MessageService) UpdateMessageEdit(ctx context.Context, request *UpdateMessageEditReq, options ...MethodOptionFunc) (*UpdateMessageEditResp, *Response, error) {
	if r.cli.mock.mockMessageUpdateMessageEdit != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#UpdateMessageEdit mock enable")
		return r.cli.mock.mockMessageUpdateMessageEdit(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "UpdateMessageEdit",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateMessageEditResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageUpdateMessageEdit mock MessageUpdateMessageEdit method
func (r *Mock) MockMessageUpdateMessageEdit(f func(ctx context.Context, request *UpdateMessageEditReq, options ...MethodOptionFunc) (*UpdateMessageEditResp, *Response, error)) {
	r.mockMessageUpdateMessageEdit = f
}

// UnMockMessageUpdateMessageEdit un-mock MessageUpdateMessageEdit method
func (r *Mock) UnMockMessageUpdateMessageEdit() {
	r.mockMessageUpdateMessageEdit = nil
}

// UpdateMessageEditReq ...
type UpdateMessageEditReq struct {
	MessageID string  `path:"message_id" json:"-"` // 待编辑的消息的ID, 仅支持文本（text）或富文本（post）消息, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: "om_dc13264520392913993dd051dba21dcf"
	MsgType   MsgType `json:"msg_type,omitempty"`  // 消息的类型, 当前仅支持文本（text）和富文本（post）类型, 示例值: "text"
	Content   string  `json:"content,omitempty"`   // 消息内容, JSON结构序列化后的字符串。不同msg_type对应不同内容, 具体格式说明参考: [发送消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 注意: JSON字符串需进行转义, 如换行符转义后为`\\n`, 文本消息请求体最大不能超过150KB, 富文本消息请求体最大不能超过30KB, 示例值: "{\"text\":\"test content\"}"
}

// UpdateMessageEditResp ...
type UpdateMessageEditResp struct {
	MessageID      string                          `json:"message_id,omitempty"`       // 消息id open_message_id
	RootID         string                          `json:"root_id,omitempty"`          // 根消息id open_message_id
	ParentID       string                          `json:"parent_id,omitempty"`        // 父消息的id open_message_id
	ThreadID       string                          `json:"thread_id,omitempty"`        // 消息所属的话题 ID（不返回说明该消息非话题消息）, 说明参见: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction)
	MsgType        MsgType                         `json:"msg_type,omitempty"`         // 消息类型 text post card image等等
	CreateTime     string                          `json:"create_time,omitempty"`      // 消息生成的时间戳(毫秒)
	UpdateTime     string                          `json:"update_time,omitempty"`      // 消息更新的时间戳
	Deleted        bool                            `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool                            `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string                          `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender                         `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody                    `json:"body,omitempty"`             // 消息内容, JSON结构
	Mentions       []*UpdateMessageEditRespMention `json:"mentions,omitempty"`         // 被艾特的人或应用的id
	UpperMessageID string                          `json:"upper_message_id,omitempty"` // 合并消息的上一层级消息id open_message_id
}

// UpdateMessageEditRespMention ...
type UpdateMessageEditRespMention struct {
	Key       string `json:"key,omitempty"`        // mention key
	ID        string `json:"id,omitempty"`         // 用户或机器人的 open_id
	IDType    IDType `json:"id_type,omitempty"`    // 被@的用户或机器人 id 类型, 目前仅支持 `open_id` ([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id))
	Name      string `json:"name,omitempty"`       // 被at用户的姓名
	TenantKey string `json:"tenant_key,omitempty"` // tenant key
}

// updateMessageEditResp ...
type updateMessageEditResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateMessageEditResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
