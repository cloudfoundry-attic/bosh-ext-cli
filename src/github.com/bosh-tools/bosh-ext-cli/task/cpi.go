package task

import (
	"encoding/json"
	"regexp"
)

var (
	// ... DEBUG -- DirectorJobRunner: [external-cpi] [cpi-413000] request: {"method":"set_disk_metadata","arguments":["disk-4fd39e40-913a-4d18-5781-b8baf653612f",{"director":"kube-minikube","deployment":"zookeeper","instance_id":"5b915e42-46f5-421c-addc-1a597c9f1bdd","instance_index":"2","instance_group":"zookeeper","attached_at":"2018-01-21T06:33:07Z"}],"context":{"director_uuid":"efd93817-b0cc-4be0-b362-5dedb1ebae43","request_id":"cpi-413000"}} with command: /var/vcap/jobs/kubernetes_cpi/bin/cpi
	cpiRequest = regexp.MustCompile("\\[external-cpi\\] \\[(cpi-\\d+)\\] request: (.+) with command")
)

type CPIRequest struct {
	Method    string
	Arguments []interface{}

	RequestID string
}

var _ Action = CPIRequest{}

func NewCPIRequest(str string) *CPIRequest {
	if m := cpiRequest.FindStringSubmatch(str); len(m) > 0 {
		var req CPIRequest

		err := json.Unmarshal([]byte(m[2]), &req)
		if err != nil {
			return nil // todo?
		}

		req.RequestID = m[1]

		return &req
	}

	return nil
}

func (r CPIRequest) Relation() Relation {
	return ExactRelation{"cpi:" + r.RequestID}
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
	// ... DEBUG -- DirectorJobRunner: [external-cpi] [cpi-413000] response: {"result":null,"error":{"type":"Bosh::Clouds::NotImplemented","message":"Must call implemented method: Unknown method 'set_disk_metadata'","ok_to_retry":false},"log":""}, err: [File System] 2018/01/21 06:33:07 DEBUG - Reading file /var/vcap/jobs/kubernetes_cpi/config/cpi.json
	cpiResponse = regexp.MustCompile("\\[external-cpi\\] \\[(cpi-\\d+)\\] response: (.+), err")
)

type CPIResponse struct {
	Result string
	Error  interface{}
	Log    string

	RequestID string
}

var _ Action = CPIResponse{}

func NewCPIResponse(str string) *CPIResponse {
	if m := cpiResponse.FindStringSubmatch(str); len(m) > 0 {
		var resp CPIResponse

		err := json.Unmarshal([]byte(m[2]), &resp)
		if err != nil {
			return nil // todo?
		}

		resp.RequestID = m[1]

		return &resp
	}

	return nil
}

func (r CPIResponse) Relation() Relation {
	return ExactRelation{"cpi:" + r.RequestID}
}

func (r CPIResponse) ShortDescription() string {
	return "[cpi resp] " + r.Result
}
