package webapi

import (
	"net/http"

	"github.com/Jackabc911/webApi/internal/app/middleware"
	"github.com/Jackabc911/webApi/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Type for WebApiServer object for instancing server
type WebApiServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *gin.Engine
	storage *storage.Storage
}

// WebApiServer constructor
func New(config *Config) *WebApiServer {
	return &WebApiServer{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

// Start http server and connection to db and logger confingurations
func (s *WebApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting web api server at port :", s.config.BindAddr)
	s.configureRouter()
	if err := s.configureStorage(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// A func for configure logger, should be unexported
func (s *WebApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)

	return nil
}

// A func for configure Router
func (s *WebApiServer) configureRouter() {
	// Add routers
	s.router.GET("/", s.GetHome)
	s.router.GET("/login", s.GetLogin)
	s.router.POST("/login", s.PostLogin)
	s.router.GET("/logout", s.GetLogout)
	s.router.GET("/register", s.GetRegister)
	s.router.POST("/register", s.PostRegister)
	s.router.GET("/users/", middleware.AuthenticateMiddleware, s.GetAllUsers)
	s.router.GET("/users/:id", s.GetUserById)
}

// A func for configure Storage
func (s *WebApiServer) configureStorage() error {
	st := storage.New(s.config.Storage)
	if err := st.Open(); err != nil {
		return err
	}
	s.storage = st
	return nil
}
