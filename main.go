package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LogEntry struct {
	Time    string `json:"time"`
	Action  string `json:"action"`
	Result  string `json:"result"`
}

func main() {
	// Создаём папку для данных если её нет
	os.MkdirAll("/app/data", 0755)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Многофункциональное приложение</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; }
        .section { border: 1px solid #ccc; padding: 15px; margin: 10px 0; border-radius: 5px; }
        button { padding: 10px 15px; margin: 5px; background: #007bff; color: white; border: none; border-radius: 5px; cursor: pointer; }
        .result { margin-top: 10px; font-weight: bold; font-size: 18px; }
        .success { color: green; }
        .history { background: #f5f5f5; padding: 10px; border-radius: 5px; margin-top: 10px; }
    </style>
</head>
<body>
    <h1>🎯 Многофункциональное приложение</h1>
    
    <div class="section">
        <h2>🎲 Случайное число</h2>
        <button onclick="getRandom()">Получить случайное число (1-10)</button>
        <div class="result" id="randomResult"></div>
    </div>

    <div class="section">
        <h2>🪙 Монетка</h2>
        <button onclick="flipCoin()">Подбросить монетку</button>
        <div class="result" id="coinResult"></div>
    </div>

    <div class="section">
        <h2>📊 История операций</h2>
        <button onclick="showHistory()">Показать историю</button>
        <button onclick="clearHistory()">Очистить историю</button>
        <div class="history" id="historyResult"></div>
    </div>

    <script>
        async function getRandom() {
            const response = await fetch('/random');
            const result = await response.json();
            document.getElementById('randomResult').innerHTML = "🎲 " + result.message;
            document.getElementById('randomResult').className = "result success";
        }

        async function flipCoin() {
            const response = await fetch('/coin');
            const result = await response.json();
            document.getElementById('coinResult').innerHTML = "🪙 " + result.message;
            document.getElementById('coinResult').className = "result success";
        }

        async function showHistory() {
            const response = await fetch('/history');
            const result = await response.json();
            if (result.success) {
                document.getElementById('historyResult').innerHTML = result.message;
            } else {
                document.getElementById('historyResult').innerHTML = "История пуста";
            }
        }

        async function clearHistory() {
            const response = await fetch('/clear-history', {method: 'POST'});
            const result = await response.json();
            document.getElementById('historyResult').innerHTML = result.message;
        }
    </script>
</body>
</html>`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/coin", coinHandler)
	http.HandleFunc("/history", historyHandler)
	http.HandleFunc("/clear-history", clearHistoryHandler)

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10) + 1
	
	// Логируем операцию
	logOperation("random", fmt.Sprintf("Случайное число: %d", num))
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("Случайное число: %d", num),
	})
}

func coinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	side := "🦅 Орёл"
	if rand.Intn(2) == 1 {
		side = "🐍 Решка"
	}
	
	// Логируем операцию
	logOperation("coin", fmt.Sprintf("Монетка: %s", side))
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("Монетка: %s", side),
	})
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	history, err := readHistory()
	if err != nil || len(history) == 0 {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "История операций пуста",
		})
		return
	}
	
	// Форматируем историю
	var historyText string
	for i, entry := range history {
		historyText += fmt.Sprintf("%d. [%s] %s - %s<br>", i+1, entry.Time, entry.Action, entry.Result)
	}
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: historyText,
	})
}

func clearHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	err := clearHistory()
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Ошибка очистки истории",
		})
		return
	}
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "История очищена",
	})
}

// Функции для работы с томом (файловой системой)

func logOperation(action, result string) {
	history, _ := readHistory()
	
	entry := LogEntry{
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Action: action,
		Result: result,
	}
	
	history = append(history, entry)
	
	// Сохраняем в файл
	file, err := os.Create("/app/data/history.json")
	if err != nil {
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.Encode(history)
}

func readHistory() ([]LogEntry, error) {
	var history []LogEntry
	
	file, err := os.Open("/app/data/history.json")
	if err != nil {
		return history, err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&history)
	
	return history, err
}

func clearHistory() error {
	return os.Remove("/app/data/history.json")
}