package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Volume total de requisições HTTP",
	})
	isUp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "service_up",
		Help: "Disponibilidade do serviço (1 = UP, 0 = DOWN)",
	})
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestsTotal.Inc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	})
}

func main() {
	isUp.Set(1) // Marca o serviço como disponível

	http.HandleFunc("/projeto-korp", handler)
	http.Handle("/metrics", promhttp.Handler()) // Endpoint para o Prometheus

	http.ListenAndServe(":8080", nil)
}
