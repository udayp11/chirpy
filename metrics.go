package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})

}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Chirpy Admin</title>
  <style>
    body {
      background-color: #f8f9fa;
      font-family: Arial, sans-serif;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      height: 100vh;
      margin: 0;
      text-align: center;
      color: #333;
    }

    h1 {
      font-size: 2.5rem;
      color: #1e90ff;
      margin-bottom: 10px;
    }

    p {
      font-size: 1.2rem;
      color: #555;
      background-color: #e6f0ff;
      padding: 10px 20px;
      border-radius: 8px;
      box-shadow: 1px 1px 5px rgba(0,0,0,0.1);
    }
  </style>
</head>
<body>
  <h1>Welcome, Chirpy Admin</h1>
  <p>Chirpy has been visited %d times!</p>
</body>
</html>

	`, cfg.fileserverHits.Load())))
}
