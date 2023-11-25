package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

// EventStore представляет хранилище событий пользователя.
type EventStore struct {
	events map[int]map[int]Event
}

// NewEventStore создает новый экземпляр EventStore.
func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[int]map[int]Event),
	}
}

// CreateEvent создает событие для указанного пользователя.
func (store *EventStore) CreateEvent(userID int, event Event) {
	if _, ok := store.events[userID]; !ok {
		store.events[userID] = make(map[int]Event)
	}
	store.events[userID][event.ID] = event
}

// GetUserEvents возвращает все события указанного пользователя.
func (store *EventStore) GetUserEvents(userID int) ([]Event, error) {
	userEvents, ok := store.events[userID]
	if !ok {
		return nil, nil
	}

	events := make([]Event, 0, len(userEvents))
	for _, event := range userEvents {
		events = append(events, event)
	}

	return events, nil
}

// ContainsEvent проверяет, содержится ли указанное событие у указанного пользователя.
func (store *EventStore) ContainsEvent(userID, eventID int) bool {
	if userEvents, ok := store.events[userID]; ok {
		_, ok := userEvents[eventID]
		return ok
	}
	return false
}

// createUserEventHandler обрабатывает запрос на создание нового пользователя.
func createUserEventHandler(store *EventStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userID int
		if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		if _, ok := store.events[userID]; !ok {
			store.events[userID] = make(map[int]Event)
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// createEventHandler обрабатывает запрос на создание нового события.
func createEventHandler(store *EventStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type CreateEventRequest struct {
			UserID int   `json:"userId"`
			Event  Event `json:"event"`
		}

		var req CreateEventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		store.CreateEvent(req.UserID, req.Event)
		w.WriteHeader(http.StatusCreated)
	}
}

// getUserEventsHandler обрабатывает запрос на получение событий пользователя.
func getUserEventsHandler(store *EventStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("userId")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		events, err := store.GetUserEvents(userID)
		if err != nil {
			http.Error(w, "Failed to get user events", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(events)
	}
}

// containsEventHandler обрабатывает запрос на проверку наличия события у пользователя.
func containsEventHandler(store *EventStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("userId")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		eventIDStr := r.URL.Query().Get("eventId")
		eventID, err := strconv.Atoi(eventIDStr)
		if err != nil {
			http.Error(w, "Invalid event ID", http.StatusBadRequest)
			return
		}

		containsEvent := store.ContainsEvent(userID, eventID)
		json.NewEncoder(w).Encode(containsEvent)
	}
}

func main() {
	store := NewEventStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/users/create", createUserEventHandler(store))
	mux.HandleFunc("/events/create", createEventHandler(store))
	mux.HandleFunc("/events/get", getUserEventsHandler(store))
	mux.HandleFunc("/events/contains", containsEventHandler(store))

	log.Println("Server started")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
