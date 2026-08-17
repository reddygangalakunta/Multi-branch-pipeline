package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Version   string    `json:"version"`
}

type InfoResponse struct {
	AppName     string `json:"app_name"`
	Environment string `json:"environment"`
	Host        string `json:"host"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:    "UP",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	hostname, _ := os.Hostname()

	resp := InfoResponse{
		AppName:     "Go Multi-Branch Demo App",
		Environment: env,
		Host:        hostname,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web App - Jenkins Multi-Branch Pipeline</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #0f172a 0%%, #1e293b 100%%);
            color: #f8fafc;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: rgba(30, 41, 59, 0.85);
            backdrop-filter: blur(12px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 16px;
            padding: 40px;
            max-width: 600px;
            width: 100%%;
            box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.5), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
            text-align: center;
        }
        .badge {
            display: inline-block;
            background: #3b82f6;
            color: #fff;
            padding: 6px 16px;
            border-radius: 9999px;
            font-size: 0.875rem;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 0.05em;
            margin-bottom: 20px;
        }
        h1 {
            font-size: 2rem;
            margin-bottom: 12px;
            color: #ffffff;
        }
        p {
            color: #94a3b8;
            line-height: 1.6;
            margin-bottom: 24px;
        }
        .status-box {
            background: #0f172a;
            border-radius: 8px;
            padding: 16px;
            display: flex;
            justify-content: space-around;
            margin-bottom: 24px;
            border: 1px solid #334155;
        }
        .status-item {
            display: flex;
            flex-direction: column;
        }
        .status-label {
            font-size: 0.75rem;
            color: #64748b;
            text-transform: uppercase;
        }
        .status-value {
            font-size: 1.1rem;
            font-weight: 600;
            color: #38bdf8;
            margin-top: 4px;
        }
        .links a {
            display: inline-block;
            margin: 0 8px;
            color: #38bdf8;
            text-decoration: none;
            font-weight: 500;
            transition: color 0.2s;
        }
        .links a:hover {
            color: #7dd3fc;
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="badge">Jenkins Multi-Branch CI/CD</div>
        <h1>Go Web Application</h1>
        <p>Simple Go application built and tested with Jenkins Multibranch Pipeline.</p>
        
        <div class="status-box">
            <div class="status-item">
                <span class="status-label">Environment</span>
                <span class="status-value">%s</span>
            </div>
            <div class="status-item">
                <span class="status-label">Status</span>
                <span class="status-value" style="color: #4ade80;">Active</span>
            </div>
            <div class="status-item">
                <span class="status-label">Version</span>
                <span class="status-value">1.0.0</span>
            </div>
        </div>

        <div class="links">
            <a href="/healthz" target="_blank">/healthz API</a>
            <a href="/api/info" target="_blank">/api/info API</a>
        </div>
    </div>
</body>
</html>`, env)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/api/info", infoHandler)
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := setupRouter()
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server starting on port %s...", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
