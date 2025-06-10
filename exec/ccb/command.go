/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ccb

import (
	"github.com/ithyl/chaosblade-spec-go/spec"
)

type CcbCommandModelSpec struct {
	spec.BaseExpModelCommandSpec
}

func NewCcbCommandModelSpec() spec.ExpModelCommandSpec {
	return &CcbCommandModelSpec{
		spec.BaseExpModelCommandSpec{
			ExpFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:   "ignore-not-found",
					Desc:   "Ignore process that cannot be found",
					NoArgs: true,
				},
			},
			ExpActions: []spec.ExpActionCommandSpec{
				NewExecCommandActionCommandSpec(),
			},
		},
	}
}

func (*CcbCommandModelSpec) Name() string {

	return "ccbCommand"
}

func (*CcbCommandModelSpec) ShortDesc() string {
	return "Command experiment"
}

func (*CcbCommandModelSpec) LongDesc() string {
	return "Command experiment, for example, ls"
}
