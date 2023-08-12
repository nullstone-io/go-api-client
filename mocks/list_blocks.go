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
}
