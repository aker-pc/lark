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

// HighlightLingoEntity 传入一句话, 智能识别句中对应的词条, 并返回词条位置和 entity_id, 可在外部系统中快速实现词条智能高亮。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/highlight
func (r *LingoService) HighlightLingoEntity(ctx context.Context, request *HighlightLingoEntityReq, options ...MethodOptionFunc) (*HighlightLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoHighlightLingoEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Lingo#HighlightLingoEntity mock enable")
		return r.cli.mock.mockLingoHighlightLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "HighlightLingoEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities/highlight",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(highlightLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoHighlightLingoEntity mock LingoHighlightLingoEntity method
func (r *Mock) MockLingoHighlightLingoEntity(f func(ctx context.Context, request *HighlightLingoEntityReq, options ...MethodOptionFunc) (*HighlightLingoEntityResp, *Response, error)) {
	r.mockLingoHighlightLingoEntity = f
}

// UnMockLingoHighlightLingoEntity un-mock LingoHighlightLingoEntity method
func (r *Mock) UnMockLingoHighlightLingoEntity() {
	r.mockLingoHighlightLingoEntity = nil
}

// HighlightLingoEntityReq ...
type HighlightLingoEntityReq struct {
	Text string `json:"text,omitempty"` // 需要识别词条的内容（不超过1000字）, 示例值: "词典是飞书提供的一款知识管理工具", 长度范围: `1` ～ `1000` 字符
}

// HighlightLingoEntityResp ...
type HighlightLingoEntityResp struct {
	Phrases []*HighlightLingoEntityRespPhrase `json:"phrases,omitempty"` // 识别到的词条信息
}

// HighlightLingoEntityRespPhrase ...
type HighlightLingoEntityRespPhrase struct {
	Name      string                              `json:"name,omitempty"`       // 识别到的关键词
	EntityIDs []string                            `json:"entity_ids,omitempty"` // 对应的词条 ID
	Span      *HighlightLingoEntityRespPhraseSpan `json:"span,omitempty"`       // 词条所在位置
}

// HighlightLingoEntityRespPhraseSpan ...
type HighlightLingoEntityRespPhraseSpan struct {
	Start int64 `json:"start,omitempty"` // 关键词开始位置, 从 0 开始计数（编码格式采用 utf-8）
	End   int64 `json:"end,omitempty"`   // 关键词结束位置, 从 0 开始计数（编码格式采用 utf-8）
}

// highlightLingoEntityResp ...
type highlightLingoEntityResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *HighlightLingoEntityResp `json:"data,omitempty"`
}
