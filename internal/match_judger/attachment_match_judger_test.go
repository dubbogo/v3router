/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package match_judger

import (
	"testing"
)

import (
	"github.com/dubbogo/v3router/config"
)

import (
	"github.com/apache/dubbo-go/protocol/invocation"
	"github.com/stretchr/testify/assert"
)

func TestAttachmentMatchJudger(t *testing.T) {
	dubboCtxMap := make(map[string]*config.StringMatch)
	dubboIvkMap := make(map[string]interface{})
	dubboCtxMap["test-key"] = &config.StringMatch{
		Exact: "abc",
	}
	dubboIvkMap["test-key"] = "abc"
	assert.True(t, NewAttachmentMatchJudger(&config.DubboAttachmentMatch{
		DubboContext: dubboCtxMap,
	}).Judge(invocation.NewRPCInvocation("method", nil, dubboIvkMap)))

	dubboIvkMap["test-key"] = "abd"
	assert.False(t, NewAttachmentMatchJudger(&config.DubboAttachmentMatch{
		DubboContext: dubboCtxMap,
	}).Judge(invocation.NewRPCInvocation("method", nil, dubboIvkMap)))

}
