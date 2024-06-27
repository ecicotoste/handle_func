package handles

import (
	"encoding/json"
	"exemplo_api_handle_func/internal/entity"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	var jsonMensagem entity.JsonHello
	jsonMensagem.Mensagem = "GET Hello Edney"
	json.NewEncoder(w).Encode(jsonMensagem)
}

func PostHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var jsonMensagem entity.JsonHello
	jsonMensagem.Mensagem = "POST Hello Edney"
	json.NewEncoder(w).Encode(jsonMensagem)
}
