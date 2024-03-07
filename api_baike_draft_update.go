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

// CreateBaikeUpdate 根据 draft_id 更新草稿内容, 已审批的草稿无法编辑。
//
// 为了更好地提升接口文档的的易理解性, 我们对文档进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/draft/update)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/update
// new doc: https://open.feishu.cn/document/server-docs/baike-v1/draft/update
//
// Deprecated
func (r *BaikeService) CreateBaikeUpdate(ctx context.Context, request *CreateBaikeUpdateReq, options ...MethodOptionFunc) (*CreateBaikeUpdateResp, *Response, error) {
	if r.cli.mock.mockBaikeCreateBaikeUpdate != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Baike#CreateBaikeUpdate mock enable")
		return r.cli.mock.mockBaikeCreateBaikeUpdate(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "CreateBaikeUpdate",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/drafts/:draft_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBaikeUpdateResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeCreateBaikeUpdate mock BaikeCreateBaikeUpdate method
func (r *Mock) MockBaikeCreateBaikeUpdate(f func(ctx context.Context, request *CreateBaikeUpdateReq, options ...MethodOptionFunc) (*CreateBaikeUpdateResp, *Response, error)) {
	r.mockBaikeCreateBaikeUpdate = f
}

// UnMockBaikeCreateBaikeUpdate un-mock BaikeCreateBaikeUpdate method
func (r *Mock) UnMockBaikeCreateBaikeUpdate() {
	r.mockBaikeCreateBaikeUpdate = nil
}

// CreateBaikeUpdateReq ...
type CreateBaikeUpdateReq struct {
	DraftID     string                           `path:"draft_id" json:"-"`      // 草稿 ID, 示例值: "5347"
	UserIDType  *IDType                          `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ID          *string                          `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）, 示例值: "enterprise_40217521"
	MainKeys    []*CreateBaikeUpdateReqMainKey   `json:"main_keys,omitempty"`    // 词条名, 最大长度: `1`
	Aliases     []*CreateBaikeUpdateReqAliase    `json:"aliases,omitempty"`      // 别名, 最大长度: `10`
	Description *string                          `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001, 示例值: "词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 最大长度: `5000` 字符
	RelatedMeta *CreateBaikeUpdateReqRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	RichText    *string                          `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 最大长度: `5000` 字符
}

// CreateBaikeUpdateReqAliase ...
type CreateBaikeUpdateReqAliase struct {
	Key           string                                   `json:"key,omitempty"`            // 名称的值, 示例值: "飞书词典"
	DisplayStatus *CreateBaikeUpdateReqAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeUpdateReqAliaseDisplayStatus ...
type CreateBaikeUpdateReqAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeUpdateReqMainKey ...
type CreateBaikeUpdateReqMainKey struct {
	Key           string                                    `json:"key,omitempty"`            // 名称的值, 示例值: "飞书词典"
	DisplayStatus *CreateBaikeUpdateReqMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeUpdateReqMainKeyDisplayStatus ...
type CreateBaikeUpdateReqMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeUpdateReqRelatedMeta ...
type CreateBaikeUpdateReqRelatedMeta struct {
	Users           []*CreateBaikeUpdateReqRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeUpdateReqRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeUpdateReqRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeUpdateReqRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeUpdateReqRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeUpdateReqRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeUpdateReqRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateBaikeUpdateReqRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片, 最大长度: `10`
}

// CreateBaikeUpdateReqRelatedMetaAbbreviation ...
type CreateBaikeUpdateReqRelatedMetaAbbreviation struct {
	ID *string `json:"id,omitempty"` // 相关词条 ID, 示例值: "enterprise_51527260"
}

// CreateBaikeUpdateReqRelatedMetaChat ...
type CreateBaikeUpdateReqRelatedMetaChat struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeUpdateReqRelatedMetaClassification ...
type CreateBaikeUpdateReqRelatedMetaClassification struct {
	ID       string  `json:"id,omitempty"`        // 二级分类 ID, 示例值: "704960692637761"
	FatherID *string `json:"father_id,omitempty"` // 对应一级分类 ID, 示例值: "704960692637777"
}

// CreateBaikeUpdateReqRelatedMetaDoc ...
type CreateBaikeUpdateReqRelatedMetaDoc struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeUpdateReqRelatedMetaImage ...
type CreateBaikeUpdateReqRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token, 示例值: "boxbcEcmKiDia3evgqWTpvdc7jc"
}

// CreateBaikeUpdateReqRelatedMetaLink ...
type CreateBaikeUpdateReqRelatedMetaLink struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeUpdateReqRelatedMetaOncall ...
type CreateBaikeUpdateReqRelatedMetaOncall struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeUpdateReqRelatedMetaUser ...
type CreateBaikeUpdateReqRelatedMetaUser struct {
	ID    string  `json:"id,omitempty"`    // 对应相关信息 ID, 示例值: "格式请看请求体示例"
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
}

// CreateBaikeUpdateResp ...
type CreateBaikeUpdateResp struct {
	Draft *CreateBaikeUpdateRespDraft `json:"draft,omitempty"` // 草稿
}

// CreateBaikeUpdateRespDraft ...
type CreateBaikeUpdateRespDraft struct {
	DraftID string                            `json:"draft_id,omitempty"` // 草稿 ID
	Entity  *CreateBaikeUpdateRespDraftEntity `json:"entity,omitempty"`   // 词条信息
}

// CreateBaikeUpdateRespDraftEntity ...
type CreateBaikeUpdateRespDraftEntity struct {
	ID          string                                       `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）
	MainKeys    []*CreateBaikeUpdateRespDraftEntityMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*CreateBaikeUpdateRespDraftEntityAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                       `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001
	CreateTime  string                                       `json:"create_time,omitempty"`  // 词条创建时间
	UpdateTime  string                                       `json:"update_time,omitempty"`  // 词条最近更新时间
	RelatedMeta *CreateBaikeUpdateRespDraftEntityRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	Statistics  *CreateBaikeUpdateRespDraftEntityStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *CreateBaikeUpdateRespDraftEntityOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    string                                       `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
}

// CreateBaikeUpdateRespDraftEntityAliase ...
type CreateBaikeUpdateRespDraftEntityAliase struct {
	Key           string                                               `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeUpdateRespDraftEntityAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeUpdateRespDraftEntityAliaseDisplayStatus ...
type CreateBaikeUpdateRespDraftEntityAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeUpdateRespDraftEntityMainKey ...
type CreateBaikeUpdateRespDraftEntityMainKey struct {
	Key           string                                                `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeUpdateRespDraftEntityMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeUpdateRespDraftEntityMainKeyDisplayStatus ...
type CreateBaikeUpdateRespDraftEntityMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeUpdateRespDraftEntityOuterInfo ...
type CreateBaikeUpdateRespDraftEntityOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}

// CreateBaikeUpdateRespDraftEntityRelatedMeta ...
type CreateBaikeUpdateRespDraftEntityRelatedMeta struct {
	Users           []*CreateBaikeUpdateRespDraftEntityRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeUpdateRespDraftEntityRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeUpdateRespDraftEntityRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeUpdateRespDraftEntityRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeUpdateRespDraftEntityRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeUpdateRespDraftEntityRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeUpdateRespDraftEntityRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateBaikeUpdateRespDraftEntityRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaAbbreviation ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 ID
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaChat ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaClassification ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaDoc ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaImage ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaLink ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaOncall ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeUpdateRespDraftEntityRelatedMetaUser ...
type CreateBaikeUpdateRespDraftEntityRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeUpdateRespDraftEntityStatistics ...
type CreateBaikeUpdateRespDraftEntityStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 累计点赞
	DislikeCount int64 `json:"dislike_count,omitempty"` // 当前词条版本收到的负反馈数量
}

// createBaikeUpdateResp ...
type createBaikeUpdateResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateBaikeUpdateResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
