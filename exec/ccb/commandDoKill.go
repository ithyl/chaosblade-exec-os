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
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/ithyl/chaosblade-spec-go/channel"
	"github.com/ithyl/chaosblade-spec-go/spec"
)

const KillProcess = "chaos_ccbCommand_kill"

type CommandDoKillActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCommandDoKillActionCommandSpec() spec.ExpActionCommandSpec {
	return &CommandDoKillActionCommandSpec{
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
				&spec.ExpFlag{
					Name: "mode",
					Desc: "exec mode",
				},
				&spec.ExpFlag{
					Name: "pid",
					Desc: "exec command by user",
				},
				&spec.ExpFlag{
					Name: "name",
					Desc: "process name",
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &CommandKillExecutor{},
			ActionExample: `
# Kill the process that contains the SimpleHTTPServer keyword
blade create ccbCommand bash --command ls --user joe`,
			ActionPrograms:   []string{KillProcess},
			ActionCategories: []string{category.SystemProcess},
		},
	}
}

func (*CommandDoKillActionCommandSpec) Name() string {
	return "bash"
}

func (*CommandDoKillActionCommandSpec) Aliases() []string {
	return []string{"bash"}
}

func (*CommandDoKillActionCommandSpec) ShortDesc() string {
	return "command"
}

func (k *CommandDoKillActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "execute command what you want"
}

func (*CommandDoKillActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CommandKillExecutor struct {
	channel spec.Channel
}

func (kpe *CommandKillExecutor) Name() string {
	return "Exec command"
}

func (kpe *CommandKillExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnSuccess(uid)
	}
	command := model.ActionFlags["command"]
	user := model.ActionFlags["user"]
	signal := model.ActionFlags["signal"]
	mode := model.ActionFlags["mode"]
	name := model.ActionFlags["name"]
	pid := model.ActionFlags["pid"]
	if mode == "name" {
		command = fmt.Sprintf("ps -ef|grep -v grep |grep %s |awk '{print $2}'|xargs kill -%s", name, signal)
	} else {
		command = fmt.Sprintf("kill -%s %s", signal, pid)
	}
	resp := &spec.Response{}
	if user == "root" {
		resp = kpe.channel.Run(ctx, command, "")
	} else {
		resp = channel.ExecScriptBySomeOne(ctx, command, "", user)
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
	if resp.Code == spec.OK.Code {
		// return directly
		resp.Code = spec.ReturnOKDirectly.Code
	}
	return resp
}

func (kpe *CommandKillExecutor) SetChannel(channel spec.Channel) {
	kpe.channel = channel
}
