package task

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SENT: agent.42be26ec-4d0e-4773-9923-3d6f2305a66b {"protocol":3,"method":"get_task","arguments":["73d1d365-1908-4c4a-5825-df567491c5e3"],"reply_to":"director.c9f043e7-d7ee-4ac1-a707-e6b0b4cc8be1.2bb3edc5-40ca-4f68-826d-0f84122228b0"}

type AgentRequest struct {
	AgentID   string
	Protocol  int
	Method    string
	Arguments []interface{}
	ReplyTo   string `json:"reply_to"`

	relation Relation
}

var _ Action = &AgentRequest{}

func NewAgentRequest(str string) *AgentRequest {
	pieces := strings.SplitN(str, " ", 3)

	if len(pieces) == 3 && pieces[0] == "SENT:" {
		var req AgentRequest

		err := json.Unmarshal([]byte(pieces[2]), &req)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse agent request: %s", err))
			return nil // todo?
		}

		req.AgentID = strings.TrimPrefix(pieces[1], "agent.")

		return &req
	}

	return nil
}

func (r *AgentRequest) Relation() Relation {
	if r.relation != nil {
		return r.relation
	}
	r.relation = AgentRelation{AgentTaskID: r.agentTaskID(), ReplyTo: r.ReplyTo}
	return r.relation
}

func (r AgentRequest) ShortDescription() string {
	desc := "[agent req] " + r.Method

	switch r.Method {
	case "run_script", "sync_dns", "mount_disk", "unmount_disk":
		desc += " " + r.Arguments[0].(string)
	}

	return desc
}

func (r AgentRequest) agentTaskID() string {
	if r.Method == "get_task" {
		if val, ok := r.Arguments[0].(string); ok {
			return val
		}
	}
	return ""
}

// RECEIVED: director.c9f043e7-d7ee-4ac1-a707-e6b0b4cc8be1.7f254f79-9b41-44ea-bda0-45cc8f88f093 {"value":{"agent_task_id":"3470dce7-9599-4ce9-743d-350d0a3e32ab","state":"running"}}
// RECEIVED: director.c9f043e7-d7ee-4ac1-a707-e6b0b4cc8be1.d4ad832a-b9df-4500-bd29-f59aba9f6510 {"exception":{"message":"Action Failed get_task: ...

type AgentResponse struct {
	Value     interface{}
	Exception interface{}
	ReplyTo   string

	relation Relation
}

var _ Action = &AgentResponse{}

func NewAgentResponse(str string) *AgentResponse {
	pieces := strings.SplitN(str, " ", 3)

	if len(pieces) == 3 && pieces[0] == "RECEIVED:" {
		var resp AgentResponse

		err := json.Unmarshal([]byte(pieces[2]), &resp)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse agent response: %s", err))
			return nil // todo?
		}

		resp.ReplyTo = pieces[1]

		return &resp
	}

	return nil
}

func (r AgentResponse) Relation() Relation {
	if r.relation != nil {
		return r.relation
	}
	r.relation = AgentRelation{AgentTaskID: r.agentTaskID(), ReplyTo: r.ReplyTo}
	return r.relation
}

func (r AgentResponse) ShortDescription() string {
	if r.Exception != nil {
		return "[agent resp] error"
	}
	return "[agent resp]"
}

func (r AgentResponse) agentTaskID() string {
	if m, ok := r.Value.(map[string]interface{}); ok {
		if val, ok := m["agent_task_id"].(string); ok {
			return val
		}
	}
	return ""
}

type AgentRelation struct {
	AgentTaskID string
	ReplyTo     string
}

var _ Relation = AgentRelation{}

func (id1 AgentRelation) Matches(id Relation) bool {
	if id2, ok := id.(AgentRelation); ok {
		return (len(id1.ReplyTo) > 0 && id1.ReplyTo == id2.ReplyTo) ||
			(len(id1.AgentTaskID) > 0 && id1.AgentTaskID == id2.AgentTaskID)
	}
	return false
}
