package rest

import (
	"context"
	"database/sql"
	"developerProfile/internal/adapters/repository/postgres/developer_profile"
	"developerProfile/internal/core/service"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateProfile(ctx context.Context, p service.CreateProfileParams) (developer_profile.DeveloperProfile, error)
	GetProfile(ctx context.Context, id string) (developer_profile.DeveloperProfile, error)
	ListProfiles(ctx context.Context, page, size int) ([]developer_profile.DeveloperProfile, error)
	UpdateProfile(ctx context.Context, id string, p service.UpdateProfileParams) (developer_profile.DeveloperProfile, error)
	DeleteProfile(ctx context.Context, id string) error
}

type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) CreateProfile(c *gin.Context) {
	var body CreateProfileInDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.CreateProfile(c.Request.Context(), service.CreateProfileParams{
		Name:                 body.Name,
		Email:                body.Email,
		Phone:                body.Phone,
		Address:              body.Address,
		AcceptTermsOfService: body.AcceptTermsOfService,
	})
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toProfileOutDTO(profile))
}

func (h *Handler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	profile, err := h.service.GetProfile(c.Request.Context(), id)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, toProfileOutDTO(profile))
}

func (h *Handler) ListProfiles(c *gin.Context) {
	page, err := parsePositiveIntQuery(c, "page", 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	size, err := parsePositiveIntQuery(c, "size", 50)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profiles, err := h.service.ListProfiles(c.Request.Context(), page, size)
	if err != nil {
		writeError(c, err)
		return
	}

	dtos := make([]ProfileOutDTO, len(profiles))
	for i, p := range profiles {
		dtos[i] = toProfileOutDTO(p)
	}
	c.JSON(http.StatusOK, ListProfilesOutDTO{Items: dtos, Page: page, Size: size})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var body UpdateProfileInDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.UpdateProfile(c.Request.Context(), id, service.UpdateProfileParams{
		Name:                 body.Name,
		Email:                body.Email,
		Phone:                body.Phone,
		Address:              body.Address,
		AcceptTermsOfService: body.AcceptTermsOfService,
	})
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, toProfileOutDTO(profile))
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.service.DeleteProfile(c.Request.Context(), id); err != nil {
		writeError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func parsePositiveIntQuery(c *gin.Context, key string, fallback int) (int, error) {
	raw := strings.TrimSpace(c.Query(key))
	if raw == "" {
		return fallback, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil || value < 0 {
		return 0, errors.New("invalid " + key)
	}
	return value, nil
}

func writeError(c *gin.Context, err error) {
	if errors.Is(err, context.Canceled) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request canceled"})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}