package task

import (
	"encoding/json"
	"regexp"
)

var (
	// External CPI sending request: {"method":"create_disk","arguments":[10240,{},"c468c40f-c7d8-4f3f-619d-7ccb5d0e3fc4"],"context":{"director_uuid":"1f8f9d74-a45d-4ab3-831d-1b78563afd24","request_id":"670343"}} with command: /var/vcap/jobs/warden_cpi/bin/cpi
	cpiRequest = regexp.MustCompile("External CPI sending request: (.+) with command")
)

type CPIRequest struct {
	Method    string
	Arguments []interface{}
	Context   CPIRequestContext
	// todo which cpi?
}

type CPIRequestContext struct {
	RequestID string `json:"request_id"`
}

var _ Action = CPIRequest{}

func NewCPIRequest(str string) *CPIRequest {
	if m := cpiRequest.FindStringSubmatch(str); len(m) > 0 {
		var req CPIRequest

		err := json.Unmarshal([]byte(m[1]), &req)
		if err != nil {
			return nil // todo?
		}

		return &req
	}

	return nil
}

func (r CPIRequest) Relation() Relation {
	return ExactRelation{"cpi:" + r.Context.RequestID}
}

func (r CPIRequest) ShortDescription() string {
	desc := "[cpi req] " + r.Method

	switch r.Method {
	case "set_vm_metadata", "set_disk_metadata", "delete_vm", "delete_disk":
		desc += " " + r.Arguments[0].(string)

	case "attach_disk", "detach_disk":
		desc += " " + r.Arguments[0].(string) + " " + r.Arguments[1].(string)
	}

	return desc
}

var (
	// External CPI got response: {"result":null,"error":{"type":"Bosh::Clouds::NotImplemented","message":"Must call implemented method","ok_to_retry":false},"log":""}, err: , exit_status: pid 22178 exit 0
	cpiResponse = regexp.MustCompile("External CPI got response: (.+), err")
)

type CPIResponse struct {
	Result string
	Error  interface{}
	Log    string
	// todo which cpi?
}

var _ Action = CPIResponse{}

func NewCPIResponse(str string) *CPIResponse {
	if m := cpiResponse.FindStringSubmatch(str); len(m) > 0 {
		var resp CPIResponse

		err := json.Unmarshal([]byte(m[1]), &resp)
		if err != nil {
			return nil // todo?
		}

		return &resp
	}

	return nil
}

func (r CPIResponse) Relation() Relation { return NonMatchingRelation{} }

func (r CPIResponse) ShortDescription() string {
	return "[cpi resp] " + r.Result
}
