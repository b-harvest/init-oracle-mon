package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bharvest.io/init-oracle-mon/store"
	"bharvest.io/init-oracle-mon/utils"
)

func Run(listenPort int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.GlobalState)
	})

	addr := fmt.Sprintf(":%d", listenPort)
	utils.Info(fmt.Sprintf("server listening on %s", addr))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		utils.Error(err)
	}
}
