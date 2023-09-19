package main

import (
    "database/sql"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

type Code struct {
    ID          int    `json:"id"`
    Code        string `json:"code"`
    Description string `json:"description"`
    CreatedAt   string `json:"created_at"`
}

func main() {
    connStr := "user=andrem password=abc123 dbname=syngular host=127.0.0.1 port=5442 sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println("Erro ao conectar ao banco de dados:", err)
        return
    }
    defer db.Close()

    router := gin.Default()

    router.GET("/codes", func(c *gin.Context) {
        rows, err := db.Query("SELECT id, code, description, created_at FROM codes")
        if err != nil {
            fmt.Println("Erro ao buscar dados:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados"})
            return
        }
        defer rows.Close()

        var codes []Code

        for rows.Next() {
            var code Code
            if err := rows.Scan(&code.ID, &code.Code, &code.Description, &code.CreatedAt); err != nil {
                fmt.Println("Erro ao escanear dados:", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao escanear dados"})
                return
            }
            codes = append(codes, code)
        }

        c.JSON(http.StatusOK, codes)
    })

    router.Run(":3085")
}
