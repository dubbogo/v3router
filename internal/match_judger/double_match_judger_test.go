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
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/dubbogo/v3router/config"
)

func TestDoubleMatchJudger(t *testing.T) {
	assert.True(t, newDoubleMatchJudger(&config.DoubleMatch{
		Exact: 3.14159,
	}).Judge(3.14159))

	assert.False(t, newDoubleMatchJudger(&config.DoubleMatch{
		Exact: 3.14159,
	}).Judge(3.14155927))

	assert.True(t, newDoubleMatchJudger(&config.DoubleMatch{
		Range: &config.DoubleRangeMatch{
			Start: 1.0,
			End:   1.5,
		},
	}).Judge(1.3))

	assert.False(t, newDoubleMatchJudger(&config.DoubleMatch{
		Range: &config.DoubleRangeMatch{
			Start: 1.0,
			End:   1.5,
		},
	}).Judge(1.9))

	assert.False(t, newDoubleMatchJudger(&config.DoubleMatch{
		Range: &config.DoubleRangeMatch{
			Start: 1.0,
			End:   1.5,
		},
	}).Judge(0.9))
}
