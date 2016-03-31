package runrunc

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/pivotal-golang/lager"
)

type Status string

const RunningStatus Status = "running"

type State struct {
	Pid    int
	Status Status
}

type Stater struct {
	commandRunner command_runner.CommandRunner
	runc          RuncBinary
}

func NewStater(runner command_runner.CommandRunner, runc RuncBinary) *Stater {
	return &Stater{
		runner, runc,
	}
}

// State gets the state of the bundle
func (r *Stater) State(log lager.Logger, handle string) (state State, err error) {
	log = log.Session("State", lager.Data{"handle": handle})

	log.Info("started")
	defer log.Info("finished")

	buf := new(bytes.Buffer)
	cmd := r.runc.StateCommand(handle)
	cmd.Stdout = buf

	if err := r.commandRunner.Run(cmd); err != nil {
		log.Error("state-cmd-failed", err)
		return State{}, fmt.Errorf("runc state: %s", err)
	}

	if err := json.NewDecoder(buf).Decode(&state); err != nil {
		log.Error("decode-state-failed", err)
		return State{}, fmt.Errorf("runc state: %s", err)
	}

	return state, nil
}