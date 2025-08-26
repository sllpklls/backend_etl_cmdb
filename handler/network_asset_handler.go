package handler

import (
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sllpklls/template-backend-go/model"
	"github.com/sllpklls/template-backend-go/repository"
)

type NetworkAssetHandler struct {
	NetworkAssetRepo repository.NetworkAssetRepo
}

func NewNetworkAssetHandler(networkAssetRepo repository.NetworkAssetRepo) *NetworkAssetHandler {
	return &NetworkAssetHandler{
		NetworkAssetRepo: networkAssetRepo,
	}
}

func (h *NetworkAssetHandler) GetAllNetworkAssets(c echo.Context) error {
	// Parse query parameters
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	// Get assets
	assets, err := h.NetworkAssetRepo.GetAllNetworkAssets(c.Request().Context(), page, limit)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get network assets",
			Data:       nil,
		})
	}

	// Get total count
	total, err := h.NetworkAssetRepo.GetTotalNetworkAssets(c.Request().Context())
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get total count",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ListResponseAsset{
		StatusCode: http.StatusOK,
		Message:    "Lấy danh sách network assets thành công",
		Data:       assets,
		Total:      total,
		Page:       page,
		Limit:      limit,
	})
}

func (h *NetworkAssetHandler) GetNetworkAssetByName(c echo.Context) error {
	name := c.Param("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    "Name parameter is required",
			Data:       nil,
		})
	}

	asset, err := h.NetworkAssetRepo.GetNetworkAssetByName(c.Request().Context(), name)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusNotFound, model.ResponseAsset{
			StatusCode: http.StatusNotFound,
			Message:    "Network asset not found",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ResponseAsset{
		StatusCode: http.StatusOK,
		Message:    "Lấy thông tin network asset thành công",
		Data:       asset,
	})
}

func (h *NetworkAssetHandler) SearchByDNSHostName(c echo.Context) error {
	dnsHostName := c.QueryParam("dns_host_name")

	if dnsHostName == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "DNS hostname parameter is required",
			Data:       nil,
		})
	}

	// Parse pagination parameters
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	// Get assets by DNS hostname
	assets, err := h.NetworkAssetRepo.GetNetworkAssetsByDNSHostName(c.Request().Context(), dnsHostName, page, limit)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to search by DNS hostname",
			Data:       nil,
		})
	}

	// Get total count
	total, err := h.NetworkAssetRepo.GetTotalNetworkAssetsByDNSHostName(c.Request().Context(), dnsHostName)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get total count",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ListResponse{
		StatusCode: http.StatusOK,
		Message:    "Tìm kiếm theo DNS hostname thành công",
		Data:       assets,
		Total:      total,
		Page:       page,
		Limit:      limit,
	})
}
func (h *NetworkAssetHandler) CheckExistByDNSHostName(c echo.Context) error {
	dnsHostName := c.QueryParam("dns_hostname")

	if dnsHostName == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "DNS hostname parameter is required",
			Data:       nil,
		})
	}

	// Get assets by DNS hostname
	isExist, err := h.NetworkAssetRepo.GetIPEndpointByDNSHostName(c.Request().Context(), dnsHostName)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to search by DNS hostname",
			Data:       nil,
		})
	}
	if isExist {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "DNS hostname exists",
			Data:       nil,
		})
	} else {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "DNS hostname does not exist",
			Data:       nil,
		})
	}
}

func (h *NetworkAssetHandler) SearchNetworkAssets(c echo.Context) error {
	var filter model.NetworkAssetFilter

	// Bind query parameters
	if err := c.Bind(&filter); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid query parameters",
			Data:       nil,
		})
	}

	// Set default values
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	// Get filtered assets
	assets, err := h.NetworkAssetRepo.GetNetworkAssetsByFilter(c.Request().Context(), filter)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to search network assets",
			Data:       nil,
		})
	}

	// Get total count for filtered results
	total, err := h.NetworkAssetRepo.GetTotalNetworkAssetsByFilter(c.Request().Context(), filter)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get total count",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ListResponseAsset{
		StatusCode: http.StatusOK,
		Message:    "Tìm kiếm network assets thành công",
		Data:       assets,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
	})
}

func (h *NetworkAssetHandler) CreateNetworkAsset(c echo.Context) error {
	var asset model.NetworkAsset

	if err := c.Bind(&asset); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON format",
			Data:       nil,
		})
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(asset); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Validate required fields
	if asset.Name == "" || asset.Address == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    "Name and Address are required",
			Data:       nil,
		})
	}

	if err := h.NetworkAssetRepo.CreateNetworkAsset(c.Request().Context(), asset); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create network asset",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusCreated, model.ResponseAsset{
		StatusCode: http.StatusCreated,
		Message:    "Tạo network asset thành công",
		Data:       asset,
	})
}

func (h *NetworkAssetHandler) UpdateNetworkAsset(c echo.Context) error {
	name := c.Param("name")
	var asset model.NetworkAsset

	if err := c.Bind(&asset); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseAsset{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON format",
			Data:       nil,
		})
	}

	if err := h.NetworkAssetRepo.UpdateNetworkAsset(c.Request().Context(), name, asset); err != nil {
		log.Error(err.Error())
		if err.Error() == "network asset not found" {
			return c.JSON(http.StatusNotFound, model.ResponseAsset{
				StatusCode: http.StatusNotFound,
				Message:    "Network asset not found",
				Data:       nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update network asset",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ResponseAsset{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật network asset thành công",
		Data:       asset,
	})
}

func (h *NetworkAssetHandler) DeleteNetworkAsset(c echo.Context) error {
	name := c.Param("name")

	if err := h.NetworkAssetRepo.DeleteNetworkAsset(c.Request().Context(), name); err != nil {
		log.Error(err.Error())
		if err.Error() == "network asset not found" {
			return c.JSON(http.StatusNotFound, model.ResponseAsset{
				StatusCode: http.StatusNotFound,
				Message:    "Network asset not found",
				Data:       nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.ResponseAsset{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete network asset",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.ResponseAsset{
		StatusCode: http.StatusOK,
		Message:    "Xóa network asset thành công",
		Data:       nil,
	})
}
