package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func main() {
	fmt.Println(colorCyan + "========================================" + colorReset)
	fmt.Println(colorGreen + "  Webhook Analyzer Server" + colorReset)
	fmt.Println(colorCyan + "========================================" + colorReset)
	fmt.Printf("%sServer lauscht auf Port %s9999%s\n", colorWhite, colorYellow, colorReset)
	fmt.Printf("%sURL: %shttp://localhost:9999%s\n\n", colorWhite, colorYellow, colorReset)
	fmt.Println(colorPurple + "Warte auf eingehende Webhooks..." + colorReset)
	fmt.Println(strings.Repeat("-", 50))

	http.HandleFunc("/", handleWebhook)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("%sFehler: %s%s\n", colorRed, err.Error(), colorReset)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("\n%s%s%s\n", colorCyan, strings.Repeat("=", 60), colorReset)
	fmt.Printf("%s[%s]%s Neue Anfrage empfangen\n", colorGreen, timestamp, colorReset)
	fmt.Printf("%s%s%s\n", colorCyan, strings.Repeat("=", 60), colorReset)

	// Method und Path
	fmt.Printf("\n%sMethode:%s  %s%s%s\n", colorYellow, colorReset, colorGreen, r.Method, colorReset)
	fmt.Printf("%sPath:%s     %s%s%s\n", colorYellow, colorReset, colorWhite, r.URL.Path, colorReset)

	if r.URL.RawQuery != "" {
		fmt.Printf("%sQuery:%s    %s%s%s\n", colorYellow, colorReset, colorWhite, r.URL.RawQuery, colorReset)
	}

	// Remote Address
	fmt.Printf("%sVon:%s      %s%s%s\n", colorYellow, colorReset, colorWhite, r.RemoteAddr, colorReset)

	// Headers
	fmt.Printf("\n%s--- Headers ---%s\n", colorPurple, colorReset)

	// Sortiere Headers für bessere Lesbarkeit
	var headerKeys []string
	for key := range r.Header {
		headerKeys = append(headerKeys, key)
	}
	sort.Strings(headerKeys)

	for _, key := range headerKeys {
		values := r.Header[key]
		fmt.Printf("%s%s:%s %s\n", colorBlue, key, colorReset, strings.Join(values, ", "))
	}

	// Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%sFehler beim Lesen des Body: %s%s\n", colorRed, err.Error(), colorReset)
	} else if len(body) > 0 {
		fmt.Printf("\n%s--- Body (%d Bytes) ---%s\n", colorPurple, len(body), colorReset)

		// Versuche JSON zu formatieren
		var jsonData interface{}
		if err := json.Unmarshal(body, &jsonData); err == nil {
			prettyJSON, _ := json.MarshalIndent(jsonData, "", "  ")
			fmt.Printf("%s%s%s\n", colorWhite, string(prettyJSON), colorReset)
		} else {
			// Kein JSON, zeige raw
			fmt.Printf("%s%s%s\n", colorWhite, string(body), colorReset)
		}
	} else {
		fmt.Printf("\n%s--- Kein Body ---%s\n", colorPurple, colorReset)
	}

	fmt.Printf("\n%s%s%s\n", colorCyan, strings.Repeat("-", 60), colorReset)

	// Sende Antwort
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":    "received",
		"timestamp": timestamp,
		"message":   "Webhook erfolgreich empfangen",
	}
	json.NewEncoder(w).Encode(response)
}
