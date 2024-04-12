package mocks

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"strconv"
)

func ListWorkspaces(router *mux.Router, workspaces []types.Workspace) {
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/workspaces").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName := vars["orgName"]
			var stackId int64
			if val, err := strconv.ParseInt(vars["stackId"], 10, 64); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("Invalid stackId: %s", err)))
				return
			} else {
				stackId = val
			}

			result := make([]types.Workspace, 0)
			for _, ws := range workspaces {
				if ws.OrgName == orgName && ws.StackId == stackId {
					result = append(result, ws)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/blocks/{blockId}/envs/{envId}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, rawStackId, rawBlockId, rawEnvId := vars["orgName"], vars["stackId"], vars["blockId"], vars["envId"]
			stackId, _ := strconv.ParseInt(rawStackId, 10, 64)
			blockId, _ := strconv.ParseInt(rawBlockId, 10, 64)
			envId, _ := strconv.ParseInt(rawEnvId, 10, 64)

			for _, ws := range workspaces {
				if ws.OrgName == orgName && ws.StackId == stackId && ws.BlockId == blockId && ws.EnvId == envId {
					raw, _ := json.Marshal(ws)
					w.Write(raw)
					return
				}
			}
		})
}
