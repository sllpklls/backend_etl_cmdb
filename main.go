package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sllpklls/template-backend-go/db"
	"github.com/sllpklls/template-backend-go/handler"
	"github.com/sllpklls/template-backend-go/repository/repo_impl"
	"github.com/sllpklls/template-backend-go/router"
)

func main() {
	// lấy env, nếu không có thì dùng default
	host := getEnv("DB_HOST", "localhost")
	portStr := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "admin")
	pass := getEnv("DB_PASSWORD", "admin123")
	dbName := getEnv("DB_NAME", "mydb")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 5432
	}

	sql := &db.Sql{
		Host:     host,
		Port:     port,
		UserName: user,
		Password: pass,
		DbName:   dbName,
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	// ✅ Bật CORS full (mọi domain, mọi method, mọi header)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	networkAssetHandler := handler.NetworkAssetHandler{
		NetworkAssetRepo: repo_impl.NewNetworkAssetRepo(sql),
	}

	api := router.API{
		Echo:                e,
		UserHandler:         userHandler,
		NetworkAssetHandler: networkAssetHandler, // Thêm này
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}

// helper: lấy env hoặc fallback sang default
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
