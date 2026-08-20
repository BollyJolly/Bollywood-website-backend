package routes

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/amanhasnainy/bingo-backend/internal/websocket"
	"github.com/amanhasnainy/bingo-backend/internal/room"
	"github.com/amanhasnainy/bingo-backend/internal/auth"
	"github.com/amanhasnainy/bingo-backend/internal/config"
	"github.com/amanhasnainy/bingo-backend/internal/session"
	"github.com/amanhasnainy/bingo-backend/internal/user"
	"github.com/amanhasnainy/bingo-backend/internal/playlist"
	"github.com/amanhasnainy/bingo-backend/internal/feedback"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins: strings.Split(
					config.GetEnv("ALLOWED_ORIGINS"),
					",",
				),

				AllowMethods: []string{
					"GET",
					"POST",
					"PUT",
					"PATCH",
					"DELETE",
					"OPTIONS",
				},

				AllowHeaders: []string{
					"Origin",
					"Content-Type",
					"Authorization",
				},

				AllowCredentials: true,

				MaxAge: 12 * time.Hour,
			},
		),
	)

	userRepository := &user.Repository{}
	sessionRepository := &session.Repository{}
	playlistRepository := &playlist.Repository{}
	feedbackRepository := &feedback.Repository{}

	feedbackService := feedback.NewService(
	feedbackRepository,
)

feedbackHandler := feedback.NewHandler(
	feedbackService,
)

playlistService := playlist.NewService(
	playlistRepository,
)

playlistHandler := playlist.NewHandler(
	playlistService,
)

	userService := user.NewService(
		userRepository,
	)

	userHandler := user.NewHandler(
		userService,
	)

	authService := auth.NewService(
		userRepository,
		sessionRepository,
	)

	authHandler := auth.NewHandler(
		authService,
	)

	roomRepository := &room.Repository{}

roomService := room.NewService(
	roomRepository,
)



roomHandler := room.NewHandler(
	roomService,
)

hub := websocket.NewHub()

websocketHandler := websocket.NewHandler(
	hub,
)
	router.GET("/", func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "Bolly Bingo Backend Running",
			},
		)
	})

	api := router.Group("/api/v1")
	roomRoutes := api.Group(
		"/rooms",
		auth.Middleware(),
	)

	playlistRoutes := api.Group(
	"/playlists",
)

{
	playlistRoutes.GET(
		"",
		playlistHandler.GetAll,
	)

	
}

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST(
			"/register",
			authHandler.Register,
		)

		authRoutes.POST(
			"/login",
			authHandler.Login,
		)

		authRoutes.POST(
			"/refresh",
			authHandler.Refresh,
		)
	}

	userRoutes := api.Group("/users")
    {
		userRoutes.GET(
			"/me",
			auth.Middleware(),
			userHandler.GetMe,
		)
    }

	
	{
		roomRoutes.GET(
	"/:code/ws",
	websocketHandler.Connect,
)
		roomRoutes.POST(
			"",
			roomHandler.Create,
		)

		roomRoutes.GET(
			"",
			roomHandler.GetAll,
		)

		roomRoutes.GET(
			"/:code",
			roomHandler.GetByCode,
		)

		roomRoutes.POST(
			"/:code/join",
			roomHandler.Join,
		)

		roomRoutes.POST(
			"/:code/start",
			roomHandler.Start,
		)

		roomRoutes.POST(
			"/:code/next",
			roomHandler.Next,
		)

		roomRoutes.POST(
			"/:code/leave",
			roomHandler.Leave,
		)
	}

	feedbackRoutes := api.Group(
	"/feedback",
	auth.Middleware(),
)

feedbackRoutes.POST(
	"",
	feedbackHandler.Create,
)

	return router
}