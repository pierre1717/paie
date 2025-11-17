package main

// Minimal Go backend using Gin that exposes a small REST API for employees and payroll.
// Includes simple CORS configuration so a React frontend can call it.
import (
    "net/http"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

// Employee model (in-memory for demo)
type Employee struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Role string `json:"role"`
    Salary float64 `json:"salary"`
}

func main() {
    r := gin.Default()

    // Allow cross-origin calls from any origin for demo.
    // In production, restrict AllowOrigins to your frontend domain, e.g. https://app.monsite.com
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    }))

    // Simple in-memory dataset
    employees := []Employee{
        {ID: 1, Name: "Aida Diagne", Role: "Comptable", Salary: 450000.0},
        {ID: 2, Name: "Moussa Ndiaye", Role: "Technicien", Salary: 300000.0},
    }

    // API group
    api := r.Group("/api")
    {
        // GET /api/employees - list employees
        api.GET("/employees", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"data": employees})
        })

        // GET /api/health - health check used by load balancers / CI
        api.GET("/health", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"status": "ok"})
        })
    }

    // Run server on port 8080
    r.Run(":8080")
}
