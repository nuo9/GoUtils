package cronicle

import (
	"encoding/json"
	"os"
)

type Input struct {
	ID       string           `json:"id"`
	Hostname string           `json:"hostname"`
	Command  string           `json:"command"`
	Event    string           `json:"event"`
	Now      int              `json:"now"`
	LogFile  string           `json:"log_file"`
	Params   *json.RawMessage `json:"params"`
}

type Output struct {
	Complete    int    `json:"complete"`
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func ReadInput(p interface{}) (in Input, e error) {
	in = Input{}
	e = json.NewDecoder(os.Stdin).Decode(&in)
	if e != nil {
		return
	}

	if p != nil {
		e = json.Unmarshal(*in.Params, p)
	}

	return
}

func OutputComplete(complete bool, alarm bool, code int, description string) Output {
	if !complete {
		return Output{Complete: 0}
	}

	if !alarm {
		return Output{Complete: 1, Code: 0}
	}

	return Output{Complete: 1, Code: code, Description: description}
}

func (o *Output) Write() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(o)
}
