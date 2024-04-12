package mocks

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"strconv"
)

func ListEnvironments(router *mux.Router, envs []types.Environment) {
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/envs").
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

			result := make([]types.Environment, 0)
			for _, env := range envs {
				if env.OrgName == orgName && env.StackId == stackId {
					result = append(result, env)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/envs/{envId}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, rawStackId, rawEnvId := vars["orgName"], vars["stackId"], vars["envId"]
			stackId, _ := strconv.ParseInt(rawStackId, 10, 64)
			envId, _ := strconv.ParseInt(rawEnvId, 10, 64)

			for _, env := range envs {
				if env.OrgName == orgName && env.StackId == stackId && env.Id == envId {
					raw, _ := json.Marshal(env)
					w.Write(raw)
					return
				}
			}
		})
}
