package rest

import (
	"CBI/configs"
	"CBI/internal/domain/repository"
	"CBI/internal/domain/usecase"
	"CBI/internal/interfaces/controllers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	// Load configuration
	cfg, err := configs.LoadConfig("E:/github/cbi/configs/config.json")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}
	cfg.Server.Port = 8080

	// Create database connection
	db, err := sql.Open("postgres", "user=postgres password=meongberem dbname=CBI host=localhost sslmode=disable")

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		log.Println("Database Successfully Connected")
	}
	defer db.Close()

	// Create repository and use case
	studentsRepo := repository.NewStudentsRepository(db, cfg)
	studentsUseCase := usecase.NewStudentsUseCase(studentsRepo)

	// Create controller and register routes
	studentsController := controllers.NewStudentsController(studentsUseCase)
	router := gin.Default()
	router.GET("/students/:id", studentsController.GetById)
	router.POST("/students", studentsController.Create)
	router.PUT("/students/:id", studentsController.UpdateById)
	router.DELETE("/students/:id", studentsController.Delete)
	router.GET("/students", studentsController.GetAll)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Starting server at %s...", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v", err)
	}
}
