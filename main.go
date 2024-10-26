package main

import (
	analyzer "P1_202200252/analyzer" // Importa el paquete "bufio" para operaciones de buffer de entrada/salida
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func main() {

	http.HandleFunc("/api/execute", withCORS(ejecutarComando)) // Define el endpoint

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Inicia el servidor
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}

}

// Middleware para manejar CORS
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configurar los encabezados CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todos los orígenes
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			// Responder a las solicitudes preflight con un código 204 (No Content)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Llamar al manejador principal
		next(w, r)
	}
}

func ejecutarComando(w http.ResponseWriter, r *http.Request) {

	var req CommandRequest

	// Decodifica la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input := req.Command
	var results []string

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		if strings.HasPrefix(line, "#") {
			// Separar '#' del texto con un espacio
			line = "# " + strings.TrimPrefix(line, "#")
		}

		if line == "" {
			result := ""
			results = append(results, result)
			continue
		}
		// Llamar a la función Analyzer del paquete analyzer para analizar la línea
		result, err := analyzer.Analyzer(line)
		if err != nil {
			// Si hay un error, almacenar el mensaje de error en lugar del resultado
			result = fmt.Sprintf("Error: %s", err.Error())
		}

		// Acumular los resultados
		results = append(results, result)
	}

	// Crea la respuesta
	consola := strings.Join(results, "\n")
	// Preparar la respuesta
	res := CommandResponse{
		Output: consola,
	}

	// Devuelve la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
