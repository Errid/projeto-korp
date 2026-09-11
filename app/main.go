package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requisicoesTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisicoes HTTP recebidas pelo servico.",
	},
	[]string{"method", "path", "status"},
)

func main() {
	prometheus.MustRegister(requisicoesTotal)

	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			status := http.StatusMethodNotAllowed
			requisicoesTotal.WithLabelValues(
				r.Method,
				"/projeto-korp",
				strconv.Itoa(status),
			).Inc()

			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Metodo nao permitido", status)
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
		requisicoesTotal.WithLabelValues(
			r.Method,
			"/projeto-korp",
			strconv.Itoa(http.StatusOK),
		).Inc()

		if err := json.NewEncoder(w).Encode(resposta); err != nil {
			log.Printf("Erro ao escrever resposta: %v", err)
		}
	})

	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("http-server-projeto-korp iniciado na porta 8080")
	log.Fatal(server.ListenAndServe())
}
