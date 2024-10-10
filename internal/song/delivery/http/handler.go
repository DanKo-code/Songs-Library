package http

import (
	"SongsLibrary/internal/db/models"
	"SongsLibrary/internal/song"
	"SongsLibrary/internal/song/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Handler struct {
	useCase  song.UseCase
	validate *validator.Validate
}

func NewHandler(useCase song.UseCase, validate *validator.Validate) *Handler {
	return &Handler{
		useCase:  useCase,
		validate: validate,
	}
}

func (h *Handler) GetSongs(c *gin.Context) {
	var gsdto dtos.GetSongsDTO

	if err := c.ShouldBindQuery(&gsdto); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.validate.Struct(gsdto)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(),
		})
		return
	}

	if gsdto.Page == 0 {
		gsdto.Page = song.DefaultPage
	}
	if gsdto.PageSize == 0 {
		gsdto.PageSize = song.DefaultPageSize
	}

	songs, err := h.useCase.GetSongs(&gsdto)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func (h *Handler) DeleteSong(c *gin.Context) {
	id := c.Param("id")

	convertedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Song ID format"})
		return
	}

	deleteSong, err := h.useCase.DeleteSong(convertedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleteSong": deleteSong})
}

func (h *Handler) UpdateSong(c *gin.Context) {
	var fieldsToUpdate models.Song
	if err := c.ShouldBindJSON(&fieldsToUpdate); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateSong, err := h.useCase.UpdateSong(&fieldsToUpdate)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated Song": updateSong})
}

func (h *Handler) CreateSong(c *gin.Context) {
	var createSongDTO dtos.CreateSongDTO
	if err := c.ShouldBindJSON(&createSongDTO); err != nil {
		log.Println(err)
		return
	}

	createSong, err := h.useCase.CreateSong(createSongDTO.Group, createSongDTO.Song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Created Song": createSong})
}

func (h *Handler) GetSongLyrics(c *gin.Context) {
	id := c.Param("id")

	convertedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Song ID format"})
		return
	}

	var gsldtp dtos.GetSongLyricsDTO
	if err := c.ShouldBindQuery(&gsldtp); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gsldtp.Id = convertedId
	if gsldtp.Page == 0 {
		gsldtp.Page = song.DefaultLyricsPage
	}
	if gsldtp.PageSize == 0 {
		gsldtp.PageSize = song.DefaultLyricsPageSize
	}

	lyrics, err := h.useCase.GetSongLyrics(&gsldtp)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"lyrics": lyrics})
}
