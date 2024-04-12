package mocks

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"strconv"
)

func ListBlocks(router *mux.Router, blocks []types.Block) {
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/blocks").
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

			result := make([]types.Block, 0)
			for _, block := range blocks {
				if block.OrgName == orgName && block.StackId == stackId {
					result = append(result, block)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/blocks/{blockId}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, rawStackId, rawBlockId := vars["orgName"], vars["stackId"], vars["blockId"]
			stackId, _ := strconv.ParseInt(rawStackId, 10, 64)
			blockId, _ := strconv.ParseInt(rawBlockId, 10, 64)

			for _, block := range blocks {
				if block.OrgName == orgName && block.StackId == stackId && block.Id == blockId {
					raw, _ := json.Marshal(block)
					w.Write(raw)
					return
				}
			}
		})

	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/apps").
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

			result := make([]types.Block, 0)
			for _, block := range blocks {
				if block.Type == "Application" && block.OrgName == orgName && block.StackId == stackId {
					result = append(result, block)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/stacks/{stackId}/app/{blockId}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, rawStackId, rawBlockId := vars["orgName"], vars["stackId"], vars["blockId"]
			stackId, _ := strconv.ParseInt(rawStackId, 10, 64)
			blockId, _ := strconv.ParseInt(rawBlockId, 10, 64)

			for _, block := range blocks {
				if block.Type == "Application" && block.OrgName == orgName && block.StackId == stackId && block.Id == blockId {
					raw, _ := json.Marshal(block)
					w.Write(raw)
					return
				}
			}
		})
}
