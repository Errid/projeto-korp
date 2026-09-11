package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		resposta := struct {
			Nome    string `json:"nome"`
			Horario string `json:"horario"`
		}{
			Nome:    "Projeto Korp",
			Horario: time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resposta); err != nil {
			log.Printf("Erro ao escrever resposta: %v", err)
		}
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("http-server-projeto-korp iniciado na porta 8080")
	log.Fatal(server.ListenAndServe())
}
