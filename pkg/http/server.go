package http

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"person-enrichment-service/server/config"
	"person-enrichment-service/server/repository"
	"person-enrichment-service/server/service"
	"person-enrichment-service/server/swag_handler"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// see this issue https://github.com/swaggo/swag/issues/830#issuecomment-725587162
	// without importing this swagger files will not picked up by swag 
	_ "person-enrichment-service/docs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	// dsn := buildDSN(cfg)
	db, err := connectDB(cfg)
	if err != nil {
		return nil, err
	}

	personRepo := repository.NewPersonRepository(db)

	enrichmentSvc := service.NewEnrichmentService(
		cfg.AgifyURL,
		cfg.GenderizeURL,
		cfg.NationalizeURL,
	)
	personSvc := service.NewPersonService(personRepo, enrichmentSvc)

	personHandler := swag_handler.NewPersonHandler(personSvc)

	router := gin.Default()


	api := router.Group("/api/v1")
	{
		api.POST("/persons", personHandler.CreatePerson)
		api.GET("/persons", personHandler.GetPeople)
		api.GET("/persons/:id", personHandler.GetPerson)
		api.PUT("/persons/:id", personHandler.UpdatePerson)
		api.DELETE("/persons/:id", personHandler.DeletePerson)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// there were problems with the db connection, that is why I added this
	router.GET("/health", func(c *gin.Context) {
		if db != nil {
			sqlDB, err := db.DB()
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"status":  "unhealthy",
					"error":   "database connection error",
					"details": err.Error(),
				})
				return
			}

			if err := sqlDB.Ping(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"status":  "unhealthy",
					"error":   "database ping failed",
					"details": err.Error(),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
			"services": gin.H{
				"database": "connected",
			},
		})
	})
	return &Server{
		router: router,
		config: cfg,
	}, nil
}

func (s *Server) Start() error {
	// return s.router.Run(":" + s.config.ServerPort)
	return s.router.Run("0.0.0.0:" + s.config.ServerPort)
}

func buildDSN(cfg *config.Config) string {
    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBPort,
        cfg.DBSSLMode,
    )
}


func connectDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := buildDSN(cfg)
    
    sqlDB, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, fmt.Errorf("raw connection failed: %w", err)
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
    defer cancel()
    
    err = sqlDB.PingContext(ctx)
    if err != nil {
        return nil, fmt.Errorf("ping failed: %w", err)
    }

    db, err := gorm.Open(postgres.New(postgres.Config{
        Conn: sqlDB,
    }), &gorm.Config{})
    
    return db, err
}