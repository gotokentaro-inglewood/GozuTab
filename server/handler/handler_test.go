package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gotokentaro-inglewood/GozuTab/database"
	"github.com/gotokentaro-inglewood/GozuTab/handler"
	"github.com/gotokentaro-inglewood/GozuTab/models"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB = database.InitDB()
	database.CreateUsersTable(testDB)
	database.CreateTabsTable(testDB)
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func clearTables(t *testing.T) {
	t.Helper()
	_, err := testDB.Exec("DELETE FROM tabs")
	if err != nil {
		t.Fatalf("failed to clear tabs: %v", err)
	}
	_, err = testDB.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("failed to clear users: %v", err)
	}
}

func insertUser(t *testing.T, name, email string) int {
	t.Helper()
	var id int
	err := testDB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		name, email,
	).Scan(&id)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}
	return id
}

func insertTab(t *testing.T, userID int, title, content string) int {
	t.Helper()
	var id int
	err := testDB.QueryRow(
		"INSERT INTO tabs (user_id, title, content) VALUES ($1, $2, $3) RETURNING id",
		userID, title, content,
	).Scan(&id)
	if err != nil {
		t.Fatalf("failed to insert tab: %v", err)
	}
	return id
}

// ---- Index ----

func TestIndexHandler(t *testing.T) {
	clearTables(t)
	insertUser(t, "Alice", "alice@example.com")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.IndexHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var data models.PageData
	if err := json.NewDecoder(w.Body).Decode(&data); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(data.Users) != 1 {
		t.Errorf("expected 1 user, got %d", len(data.Users))
	}
}

// ---- Users ----

func TestCreateUserHandler(t *testing.T) {
	clearTables(t)

	body, _ := json.Marshal(map[string]string{"name": "Bob", "email": "bob@example.com"})
	req := httptest.NewRequest(http.MethodPost, "/users/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.CreateUserHandler(testDB)(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}

	var user models.User
	if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if user.Name != "Bob" {
		t.Errorf("expected name Bob, got %s", user.Name)
	}
	if user.ID == 0 {
		t.Error("expected non-zero id")
	}
}

func TestUpdateUserHandler(t *testing.T) {
	clearTables(t)
	id := insertUser(t, "Carol", "carol@example.com")

	body, _ := json.Marshal(map[string]string{"name": "Carol Updated", "icon_url": ""})
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/users/update?id=%d", id), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.UpdateUserHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	clearTables(t)
	id := insertUser(t, "Dave", "dave@example.com")

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/users/delete?id=%d", id), nil)
	w := httptest.NewRecorder()
	handler.DeleteUserHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// ---- Tabs ----

func TestTabsHandler(t *testing.T) {
	clearTables(t)
	userID := insertUser(t, "Eve", "eve@example.com")
	insertTab(t, userID, "Wonderwall", "Em G D A")

	req := httptest.NewRequest(http.MethodGet, "/tabs", nil)
	w := httptest.NewRecorder()
	handler.TabsHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var tabs []models.Tab
	if err := json.NewDecoder(w.Body).Decode(&tabs); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(tabs) != 1 {
		t.Errorf("expected 1 tab, got %d", len(tabs))
	}
}

func TestCreateTabHandler(t *testing.T) {
	clearTables(t)
	userID := insertUser(t, "Frank", "frank@example.com")

	body, _ := json.Marshal(map[string]interface{}{"user_id": userID, "title": "Hotel California", "content": "Am E G D"})
	req := httptest.NewRequest(http.MethodPost, "/tabs/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.CreateTabHandler(testDB)(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}

	var tab models.Tab
	if err := json.NewDecoder(w.Body).Decode(&tab); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if tab.Title != "Hotel California" {
		t.Errorf("expected title Hotel California, got %s", tab.Title)
	}
}

func TestUpdateTabHandler(t *testing.T) {
	clearTables(t)
	userID := insertUser(t, "Grace", "grace@example.com")
	tabID := insertTab(t, userID, "Original Title", "C G Am F")

	body, _ := json.Marshal(map[string]string{"title": "Updated Title", "content": "C G Am F (updated)"})
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/tabs/update?id=%d", tabID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.UpdateTabHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteTabHandler(t *testing.T) {
	clearTables(t)
	userID := insertUser(t, "Heidi", "heidi@example.com")
	tabID := insertTab(t, userID, "Delete Me", "")

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/tabs/delete?id=%d", tabID), nil)
	w := httptest.NewRecorder()
	handler.DeleteTabHandler(testDB)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
