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
	"context"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/ithyl/chaosblade-spec-go/channel"
	"github.com/ithyl/chaosblade-spec-go/spec"
)

const KillProcessBin = "chaos_ccbCommand"

type ExecCommandActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewExecCommandActionCommandSpec() spec.ExpActionCommandSpec {
	return &ExecCommandActionCommandSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "command",
					Desc: "origin command",
				},
				&spec.ExpFlag{
					Name: "user",
					Desc: "exec command by user",
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &CommandExecutor{},
			ActionExample: `
# Kill the process that contains the SimpleHTTPServer keyword
blade create ccbCommand bash --command ls --user joe`,
			ActionPrograms:   []string{KillProcessBin},
			ActionCategories: []string{category.SystemProcess},
		},
	}
}

func (*ExecCommandActionCommandSpec) Name() string {
	return "bash"
}

func (*ExecCommandActionCommandSpec) Aliases() []string {
	return []string{"bash"}
}

func (*ExecCommandActionCommandSpec) ShortDesc() string {
	return "command"
}

func (k *ExecCommandActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "execute command what you want"
}

func (*ExecCommandActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CommandExecutor struct {
	channel spec.Channel
}

func (kpe *CommandExecutor) Name() string {
	return "Exec command"
}

func (kpe *CommandExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnSuccess(uid)
	}
	command := model.ActionFlags["command"]
	user := model.ActionFlags["user"]
	if user == "root" {
		return kpe.channel.Run(ctx, command, "")
	} else {
		return channel.ExecScriptBySomeOne(ctx, command, "", user)
	}
	//resp := getPids(ctx, kpe.channel, model, uid)
	//if !resp.Success {
	//	return resp
	//}
	//pids := resp.Result.(string)
	//signal := model.ActionFlags["signal"]
	//if signal == "" {
	//	log.Errorf(ctx, "less signal flag value")
	//	return spec.ResponseFailWithFlags(spec.ParameterLess, "signal")
	//}
	//return kpe.channel.Run(ctx, "kill", fmt.Sprintf("-%s %s", signal, pids))
	return &spec.Response{}
}

func (kpe *CommandExecutor) SetChannel(channel spec.Channel) {
	kpe.channel = channel
}
