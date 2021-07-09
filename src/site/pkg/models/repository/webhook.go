package repository

import (
	"github.com/jmoiron/sqlx"
	"site/pkg/models"
)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type WebhookModel struct {
	DB *sqlx.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *WebhookModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *WebhookModel) Get(id int) (*models.Webhook, error) {
	return nil, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *WebhookModel) All() ([]*models.Webhook, error) {
	return nil, nil
}