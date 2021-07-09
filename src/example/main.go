package main

import (
    "encoding/json"
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/jmoiron/sqlx/types"
    "log"
    "time"

    _ "github.com/lib/pq"
)

type Webhook struct {
    Id int64 `db:"id"`
    AccountId  int64 `db:"account_id"`
    Domain  string `db:"domain"`
    ProductSysname  string `db:"product_sysname"`
    Request types.JSONText `json:"request" db:"request"`
    Status  string `db:"status"`
    Error  types.JSONText `json:"error" db:"error"`
    Info  types.JSONText `json:"info" db:"info"`
    CreatedAt  time.Time `db:"created_at"`
    UpdatedAt  time.Time `db:"updated_at"`
}

type Data struct {
    Id int64
}
// Описываем структуры массива json
type Request struct {
    Account struct {
        Id string `json:"id"`
        Subdomain string `json:"subdomain"`
    } `json:"account"`
    Contacts struct {
        Delete []struct {
            Id string `json:"id"`
            Type string `json:"type"`
        } `json:"delete"`
    } `json:"contacts"`
}

func main() {
    t0 := time.Now()

    db, err := sqlx.Connect("postgres", "user=dev password=dev host=localhost port=5435 dbname=paneldb sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }
    webhooks := []Webhook{}
    db.Select(&webhooks, "SELECT * FROM webhooks ORDER BY id ASC")

    for _, element := range webhooks {
        request := Request{}
        json.Unmarshal([]byte(element.Request), &request)
        if request.Account.Id == "" {
            continue
        }
        fmt.Println(request.Account.Id)

        for  _, contact := range request.Contacts.Delete {
            fmt.Println(contact.Type)
        }
    }
    t1 := time.Now()
    fmt.Println(t0)
    fmt.Println(t1)
}