package main

import (
	"flag"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"site/pkg/models/repository"
)

type Config struct {
	Addr      string
	StaticDir string
}

// Создаем структуру `application` для хранения зависимостей всего веб-приложения.
// Пока, что мы добавим поля только для двух логгеров, но
// мы будем расширять данную структуру по мере усложнения приложения.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	webhooks *repository.WebhookModel
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./src/site/ui/static", "Path to static assets")
	flag.Parse()

	db, err := sqlx.Connect("postgres", "user=dev password=dev host=localhost port=5435 dbname=paneldb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Инициализируем новую структуру с зависимостями приложения.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		webhooks: &repository.WebhookModel{DB: db},
	}

	// Инициализируем новую структуру http.Server. Мы устанавливаем поля Addr и Handler, так
	// что сервер использует тот же сетевой адрес и маршруты, что и раньше, и назначаем
	// поле ErrorLog, чтобы сервер использовал наш логгер
	// при возникновении проблем.
	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Вызов нового метода app.routes()
	}

	infoLog.Printf("%s%s","Запуск сервера на http://127.0.0.1", cfg.Addr)
	// Вызываем метод ListenAndServe() от нашей новой структуры http.Server
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}