package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sllpklls/template-backend-go/handler"
	"github.com/sllpklls/template-backend-go/middleware"
)

type API struct {
	Echo                *echo.Echo
	UserHandler         handler.UserHandler
	NetworkAssetHandler handler.NetworkAssetHandler
}

func (api *API) SetupRouter() {
	// Route không yêu cầu xác thực JWT
	api.Echo.POST("/user/sign-in", api.UserHandler.HandlerSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())

	// api.Echo.GET("/list/ci", api.UserHandler.ListCI, middleware.JWTMiddleware())
	v1 := api.Echo.Group("/api/v1")
	v1.Use(middleware.JWTMiddleware()) // Uncomment nếu cần JWT protection

	v1.GET("/network-assets", api.NetworkAssetHandler.GetAllNetworkAssets)
	v1.GET("/network-assets/search", api.NetworkAssetHandler.SearchNetworkAssets)
	v1.GET("/network-assets/search-dns", api.NetworkAssetHandler.SearchByDNSHostName)
	v1.GET("/network-assets/:name", api.NetworkAssetHandler.GetNetworkAssetByName)
	v1.POST("/network-assets", api.NetworkAssetHandler.CreateNetworkAsset)
	v1.PUT("/network-assets/:name", api.NetworkAssetHandler.UpdateNetworkAsset)
	v1.DELETE("/network-assets/:name", api.NetworkAssetHandler.DeleteNetworkAsset)
}
