package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HelloIndexResponse 疎通確認レスポンス用json構造体
type HelloIndexResponse struct {
	Messaage string `json:"message"`
}

// HelloIndex 疎通確認用
func HelloIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexResponse := HelloIndexResponse{Messaage: "Hello Go API!!"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(indexResponse); err != nil {
		panic(err)
	}
}
