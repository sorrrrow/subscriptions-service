// @title Subscriptions API
// @version 1.0
// @description API для управления онлайн-подписками пользователей
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "subscriptions/docs"

	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var config Config

type Subscription struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// @Summary Создание подписки
// @Description Добавляет новую подписку в БД
// @Accept  json
// @Produce  json
// @Param subscription body Subscription true "Данные подписки"
// @Success 200 {object} Subscription
// @Failure 400 {string} string "Неверный JSON"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /subscriptions [post]
func createSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("Создание подписки")
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	if err := db.Create(&sub).Error; err != nil {
		http.Error(w, "Ошибка сохранения в БД", http.StatusInternalServerError)
		return
	}

	log.Printf("Подписка создана: %+v\n", sub)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

// @Summary Получение всех подписок
// @Description Возвращает список всех подписок из БД
// @Produce  json
// @Success 200 {array} Subscription
// @Failure 500 {string} string "Ошибка сервера"
// @Router /subscriptions [get]
func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	log.Println("Получение списка подписок")
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	userID := query.Get("user_id")
	serviceName := query.Get("service_name")
	from := query.Get("from")
	to := query.Get("to")

	var subs []Subscription
	tx := db.Model(&Subscription{})

	if userID != "" {
		tx = tx.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		tx = tx.Where("service_name = ?", serviceName)
	}
	if from != "" && to != "" {
		tx = tx.Where("start_date >= ? AND start_date <= ?", from, to)
	}

	if err := tx.Find(&subs).Error; err != nil {
		http.Error(w, "Ошибка запроса к БД", http.StatusInternalServerError)
		return
	}

	log.Printf("Найдено подписок: %d\n", len(subs))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

// @Summary Сумма подписок по фильтрам
// @Description Возвращает сумму цен подписок по user_id, дате и имени сервиса
// @Produce  json
// @Param user_id query string true "ID пользователя"
// @Param service_name query string false "Название сервиса"
// @Param from query string true "Дата начала периода (YYYY-MM-DD)"
// @Param to query string true "Дата конца периода (YYYY-MM-DD)"
// @Success 200 {object} map[string]int
// @Failure 400 {string} string "Ошибка параметров"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /subscriptions/summary [get]
func getSummary(w http.ResponseWriter, r *http.Request) {
	log.Println("Расчёт суммы подписок")
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	userID := query.Get("user_id")
	serviceName := query.Get("service_name")
	from := query.Get("from")
	to := query.Get("to")

	if userID == "" || from == "" || to == "" {
		http.Error(w, "Нужно указать user_id, from и to", http.StatusBadRequest)
		return
	}

	tx := db.Model(&Subscription{}).
		Select("SUM(price)").
		Where("user_id = ?", userID).
		Where("start_date >= ? AND start_date <= ?", from, to)

	if serviceName != "" {
		tx = tx.Where("service_name = ?", serviceName)
	}

	var total sql.NullInt64
	err := tx.Scan(&total).Error
	if err != nil {
		http.Error(w, "Ошибка запроса к базе", http.StatusInternalServerError)
		return
	}

	result := int64(0)
	if total.Valid {
		result = total.Int64
	}

	log.Printf("Итоговая сумма: %d\n", result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"total": result})
}

// @Summary Обновление подписки
// @Description Обновляет существующую подписку по ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID подписки"
// @Param subscription body Subscription true "Обновлённые данные подписки"
// @Success 200 {object} Subscription
// @Failure 400 {string} string "Неверный JSON"
// @Failure 404 {string} string "Подписка не найдена"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /subscriptions/{id} [put]
func updateSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("Обновление подписки")
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/subscriptions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	result := db.Model(&Subscription{}).Where("id = ?", id).Updates(sub)
	if result.Error != nil {
		http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Подписка не найдена", http.StatusNotFound)
		return
	}

	log.Printf("Подписка с ID %d обновлена\n", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

// @Summary Удаление подписки
// @Description Удаляет подписку по ID
// @Produce  json
// @Param id path int true "ID подписки"
// @Success 204 {string} string "Успешно удалено"
// @Failure 404 {string} string "Подписка не найдена"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /subscriptions/{id} [delete]
func deleteSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("Удаление подписки")
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/subscriptions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	result := db.Delete(&Subscription{}, id)
	if result.Error != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Подписка не найдена", http.StatusNotFound)
		return
	}

	log.Printf("Подписка с ID %d удалена\n", id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	config = LoadConfig()
	dsn := config.DSN()

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	http.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createSubscription(w, r)
		} else if r.Method == http.MethodGet {
			getSubscriptions(w, r)
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/subscriptions/summary", getSummary)
	http.HandleFunc("/subscriptions/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/subscriptions/") {
			if r.Method == http.MethodPut {
				updateSubscription(w, r)
			} else if r.Method == http.MethodDelete {
				deleteSubscription(w, r)
			} else {
				http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			}
		} else {
			http.NotFound(w, r)
		}
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Printf("Сервер запущен на http://localhost:%s\n", config.ServerPort)
	http.ListenAndServe(":"+config.ServerPort, nil)
}
