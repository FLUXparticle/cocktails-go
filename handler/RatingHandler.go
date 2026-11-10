package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"cocktails-go/repository"
	"github.com/gin-gonic/gin"
)

// RatingHandler verbindet die Bewertungs-API mit SQLite und MongoDB.
type RatingHandler struct {
	cocktails *repository.CocktailRepository
	ratings   *repository.RatingRepository
}

// NewRatingHandler übernimmt die beiden Repositorys aus Fx.
func NewRatingHandler(cocktails *repository.CocktailRepository, ratings *repository.RatingRepository) *RatingHandler {
	return &RatingHandler{cocktails: cocktails, ratings: ratings}
}

// Create prüft die Eingabe und speichert eine neue Bewertung.
func (h *RatingHandler) Create(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize-1)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Cocktail-ID."})
		return
	}
	var input struct {
		Stars   int    `json:"stars" binding:"required,min=1,max=5"`
		Comment string `json:"comment"`
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 8192)
	if err := c.ShouldBindJSON(&input); err != nil || utf8.RuneCountInString(input.Comment) > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bitte 1 bis 5 Sterne und höchstens 500 Zeichen eingeben."})
		return
	}

	cocktails := h.cocktails.FindByIDs([]uint{uint(id)})
	if len(cocktails) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cocktail nicht gefunden."})
		return
	}
	rating := repository.Rating{
		CocktailID: uint(id), CocktailName: cocktails[0].Name,
		Stars: input.Stars, Comment: strings.TrimSpace(input.Comment),
	}
	if err := h.ratings.Insert(c.Request.Context(), &rating); err != nil {
		c.Error(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Bewertung konnte nicht gespeichert werden."})
		return
	}
	c.JSON(http.StatusCreated, rating)
}

// List liefert die jüngsten Bewertungen für den Start des Feeds.
func (h *RatingHandler) List(c *gin.Context) {
	ratings, err := h.ratings.Recent(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Bewertungen konnten nicht geladen werden."})
		return
	}
	c.JSON(http.StatusOK, ratings)
}

// Stream übersetzt neue MongoDB-Dokumente in Server-Sent Events.
func (h *RatingHandler) Stream(c *gin.Context) {
	ctx := c.Request.Context()
	stream, err := h.ratings.WatchInserts(ctx)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Der Live-Feed benötigt ein MongoDB-Replica-Set."})
		return
	}
	defer func() {
		// Zum Aufräumen ist der Request-Kontext möglicherweise bereits abgebrochen.
		closeCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		stream.Close(closeCtx)
	}()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	// Erst jetzt lädt der Browser die Historie, während der Stream neue Inserts bereits erfasst.
	c.SSEvent("ready", true)
	c.Writer.Flush()
	for stream.Next(ctx) {
		var change struct {
			Rating repository.Rating `bson:"fullDocument"`
		}
		if err := stream.Decode(&change); err != nil {
			c.Error(err)
			continue
		}
		c.SSEvent("rating", change.Rating)
		c.Writer.Flush()
	}
	if err := stream.Err(); err != nil && ctx.Err() == nil {
		c.Error(err)
	}
}
