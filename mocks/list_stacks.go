package mocks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"strconv"
)

func ListStacks(router *mux.Router, stacks []types.Stack) {
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName := vars["orgName"]

			result := make([]types.Stack, 0)
			for _, stack := range stacks {
				if stack.OrgName == orgName {
					result = append(result, stack)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, rawStackId := vars["orgName"], vars["stackId"]
			stackId, _ := strconv.ParseInt(rawStackId, 10, 64)

			for _, stack := range stacks {
				if stack.OrgName == orgName && stack.Id == stackId {
					raw, _ := json.Marshal(stack)
					w.Write(raw)
					return
				}
			}
		})
}
