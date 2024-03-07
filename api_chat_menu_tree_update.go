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

// UpdateChatMenuTree 修改某个一级菜单或者二级菜单的元信息, 包括群菜单的图标、名称、国际化名称和跳转链接。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 机器人必须在群里。
// - 该API暂时不支持在一级菜单上添加或者删除二级菜单。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_item/patch
// new doc: https://open.feishu.cn/document/server-docs/group/chat-menu_tree/patch
func (r *ChatService) UpdateChatMenuTree(ctx context.Context, request *UpdateChatMenuTreeReq, options ...MethodOptionFunc) (*UpdateChatMenuTreeResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatMenuTree != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatMenuTree mock enable")
		return r.cli.mock.mockChatUpdateChatMenuTree(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatMenuTree",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/menu_items/:menu_item_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateChatMenuTreeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChatMenuTree mock ChatUpdateChatMenuTree method
func (r *Mock) MockChatUpdateChatMenuTree(f func(ctx context.Context, request *UpdateChatMenuTreeReq, options ...MethodOptionFunc) (*UpdateChatMenuTreeResp, *Response, error)) {
	r.mockChatUpdateChatMenuTree = f
}

// UnMockChatUpdateChatMenuTree un-mock ChatUpdateChatMenuTree method
func (r *Mock) UnMockChatUpdateChatMenuTree() {
	r.mockChatUpdateChatMenuTree = nil
}

// UpdateChatMenuTreeReq ...
type UpdateChatMenuTreeReq struct {
	ChatID       string                             `path:"chat_id" json:"-"`         // 群ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`的群ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MenuItemID   string                             `path:"menu_item_id" json:"-"`    // 一级或二级菜单ID, 通过 [获取群菜单](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_tree/get) 接口通过群ID获取菜单ID, 示例值: "7156553273518882844"
	UpdateFields []string                           `json:"update_fields,omitempty"`  // 要修改的字段, 示例值: ["ICON"], 可选值有: ICON: 图标, NAME: 名称, I18N_NAME: 国际化名称, REDIRECT_LINK: 跳转链接
	ChatMenuItem *UpdateChatMenuTreeReqChatMenuItem `json:"chat_menu_item,omitempty"` // 元信息
}

// UpdateChatMenuTreeReqChatMenuItem ...
type UpdateChatMenuTreeReqChatMenuItem struct {
	ActionType   *string                                        `json:"action_type,omitempty"`   // 菜单类型, 注意, 如果一级菜单有二级菜单时, 则此一级菜单的值必须为NONE, 示例值: "NONE", 可选值有: NONE: 无类型, REDIRECT_LINK: 跳转链接类型
	RedirectLink *UpdateChatMenuTreeReqChatMenuItemRedirectLink `json:"redirect_link,omitempty"` // 跳转链接
	ImageKey     *string                                        `json:"image_key,omitempty"`     // 图片的key值。通过 [上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) 接口上传message类型图片获取image_key, 注意, 如果一级菜单有二级菜单, 则此一级菜单不能有图标, 示例值: "img_v2_b0fbe905-7988-4282-b882-82edd010336j"
	Name         *string                                        `json:"name,omitempty"`          // 菜单名称, 注意, 一级、二级菜单名称字符数要在1到120范围内, 示例值: "群聊"
	I18nNames    *I18nNames                                     `json:"i18n_names,omitempty"`    // 菜单国际化名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
}

// UpdateChatMenuTreeReqChatMenuItemRedirectLink ...
type UpdateChatMenuTreeReqChatMenuItemRedirectLink struct {
	CommonURL  *string `json:"common_url,omitempty"`  // 公用跳转链接, 必须以http开头, 示例值: "https://open.feishu.cn/"
	IosURL     *string `json:"ios_url,omitempty"`     // IOS端跳转链接, 当该字段不设置时, IOS端会使用common_url。必须以http开头, 示例值: "https://open.feishu.cn/"
	AndroidURL *string `json:"android_url,omitempty"` // Android端跳转链接, 当该字段不设置时, Android端会使用common_url。必须以http开头, 示例值: "https://open.feishu.cn/"
	PcURL      *string `json:"pc_url,omitempty"`      // PC端跳转链接, 当该字段不设置时, PC端会使用common_url。必须以http开头。在PC端点击群菜单后, 如果需要url对应的页面在飞书侧边栏展开, 可以在url前加上https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=, 比如https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=https://open.feishu.cn/, 示例值: "https://open.feishu.cn/"
	WebURL     *string `json:"web_url,omitempty"`     // Web端跳转链接, 当该字段不设置时, Web端会使用common_url。必须以http开头, 示例值: "https://open.feishu.cn/"
}

// UpdateChatMenuTreeResp ...
type UpdateChatMenuTreeResp struct {
	ChatMenuItem *UpdateChatMenuTreeRespChatMenuItem `json:"chat_menu_item,omitempty"` // 修改后的菜单元信息
}

// UpdateChatMenuTreeRespChatMenuItem ...
type UpdateChatMenuTreeRespChatMenuItem struct {
	ActionType   string                                          `json:"action_type,omitempty"`   // 菜单类型, 注意, 如果一级菜单有二级菜单时, 则此一级菜单的值必须为NONE, 可选值有: NONE: 无类型, REDIRECT_LINK: 跳转链接类型
	RedirectLink *UpdateChatMenuTreeRespChatMenuItemRedirectLink `json:"redirect_link,omitempty"` // 跳转链接
	ImageKey     string                                          `json:"image_key,omitempty"`     // 图片的key值。通过 [上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) 接口上传message类型图片获取image_key, 注意, 如果一级菜单有二级菜单, 则此一级菜单不能有图标。
	Name         string                                          `json:"name,omitempty"`          // 菜单名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
	I18nNames    *I18nNames                                      `json:"i18n_names,omitempty"`    // 菜单国际化名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
}

// UpdateChatMenuTreeRespChatMenuItemRedirectLink ...
type UpdateChatMenuTreeRespChatMenuItemRedirectLink struct {
	CommonURL  string `json:"common_url,omitempty"`  // 公用跳转链接, 必须以http开头。
	IosURL     string `json:"ios_url,omitempty"`     // IOS端跳转链接, 当该字段不设置时, IOS端会使用common_url。必须以http开头。
	AndroidURL string `json:"android_url,omitempty"` // Android端跳转链接, 当该字段不设置时, Android端会使用common_url。必须以http开头。
	PcURL      string `json:"pc_url,omitempty"`      // PC端跳转链接, 当该字段不设置时, PC端会使用common_url。必须以http开头。在PC端点击群菜单后, 如果需要url对应的页面在飞书侧边栏展开, 可以在url前加上https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=, 比如https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=https://open.feishu.cn/
	WebURL     string `json:"web_url,omitempty"`     // Web端跳转链接, 当该字段不设置时, Web端会使用common_url。必须以http开头。
}

// updateChatMenuTreeResp ...
type updateChatMenuTreeResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateChatMenuTreeResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
