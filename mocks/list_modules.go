package mocks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

func ListModules(router *mux.Router, modules []types.Module) {
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/modules").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName := vars["orgName"]

			result := make([]types.Module, 0)
			for _, mod := range modules {
				if mod.OrgName == orgName {
					result = append(result, mod)
				}
			}

			raw, _ := json.Marshal(result)
			w.Write(raw)
		})
	router.Methods(http.MethodGet).
		Path("/orgs/{orgName}/modules/{moduleName}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			orgName, moduleName := vars["orgName"], vars["moduleName"]

			for _, mod := range modules {
				if mod.OrgName == orgName && mod.Name == moduleName {
					raw, _ := json.Marshal(mod)
					w.Write(raw)
					return
				}
			}
		})
}
