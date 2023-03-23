package task

import (
	"codego/gtask/util/client"
	"encoding/json"
	"net/http"
	"strings"
)

type JobFunc func(rule *Rule)

func HttpJobFunc() JobFunc {
	return func(r *Rule) {
		r.Method = strings.ToUpper(r.Method)
		if r.Method == http.MethodGet {
			params := make(map[string]string)
			err := json.Unmarshal([]byte(r.Params), &params)
			if err == nil {
				client.Get(r.Url, params)
			}
		} else if r.Method == http.MethodPost {

		}
	}
}

func TcpJobFunc() JobFunc {
	return func(r *Rule) {

	}
}

func ShellJobFunc() JobFunc {
	return func(r *Rule) {

	}
}
