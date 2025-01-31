package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Memo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var memos = []Memo{}
var nextID int = 1
var mu sync.Mutex // 複数のリクエストに対応するためのロック

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Dockerized Go!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// メモ一覧を取得（GET /memos）
func getMemosHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memos)
}

// メモを作成（POST /memos）
func createMemoHandler(w http.ResponseWriter, r *http.Request) {
	var newMemo Memo
	if err := json.NewDecoder(r.Body).Decode(&newMemo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	newMemo.ID = nextID
	nextID++
	memos = append(memos, newMemo)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMemo)
}

// メモを更新（PUT /memos/{id}）
func updateMemoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/memos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid memo ID", http.StatusBadRequest)
		return
	}

	var updatedMemo Memo
	if err := json.NewDecoder(r.Body).Decode(&updatedMemo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, memo := range memos {
		if memo.ID == id {
			memos[i].Text = updatedMemo.Text
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(memos[i])
			return
		}
	}

	http.Error(w, "Memo not found", http.StatusNotFound)
}

// メモを削除（DELETE /memos/{id}）
func deleteMemoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/memos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid memo ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, memo := range memos {
		if memo.ID == id {
			memos = append(memos[:i], memos[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Memo deleted successfully"})
			return
		}
	}

	http.Error(w, "Memo not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/memos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getMemosHandler(w, r)
		} else if r.Method == http.MethodPost {
			createMemoHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/memos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			updateMemoHandler(w, r)
		} else if r.Method == http.MethodDelete {
			deleteMemoHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
