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

// EventV2CorehrPersonUpdatedV1 员工在飞书人事的「个人信息被更新」时将触发此事件, 个人信息的创建和删除不会触发该事件。注: 籍贯、政治面貌、户口类型、户口所在地、工龄变化不会触发该事件{使用示例}(url=/api/tools/api_explore/api_explore_config?project=corehr&version=v1&resource=person&event=updated)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/events/updated
func (r *EventCallbackService) HandlerEventV2CorehrPersonUpdatedV1(f EventV2CorehrPersonUpdatedV1Handler) {
	r.cli.eventHandler.eventV2CorehrPersonUpdatedV1Handler = f
}

// EventV2CorehrPersonUpdatedV1Handler event EventV2CorehrPersonUpdatedV1 handler
type EventV2CorehrPersonUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrPersonUpdatedV1) (string, error)

// EventV2CorehrPersonUpdatedV1 ...
type EventV2CorehrPersonUpdatedV1 struct {
	PersonID     string   `json:"person_id,omitempty"`     // 被更新个人信息的 ID
	FieldChanges []string `json:"field_changes,omitempty"` // 发生变更的字段
}
