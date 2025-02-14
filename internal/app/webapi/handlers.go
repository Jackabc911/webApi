package webapi

import (
	"net/http"

	"github.com/Jackabc911/webApi/internal/app/middleware"
	"github.com/Jackabc911/webApi/internal/app/models"
	"github.com/gin-gonic/gin"
)

const textTwo string = "Вы зарегистрированы с логином:"
const textOne string = "Вы вошли с логином:"

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// GET Home page handler func
func (s *WebApiServer) GetHome(ctx *gin.Context) {
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/welcome.html")
	ctx.HTML(http.StatusOK, "layout.html", gin.H{})
}

// GET all users handler func
func (s *WebApiServer) GetAllUsers(ctx *gin.Context) {
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/list.html")
	u, err := s.storage.User().SelectAll()
	if err != nil {
		s.logger.Info("Troubles while creating new user:", err)
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Users": "'none"})
		return
	}
	ctx.HTML(http.StatusOK, "layout.html", gin.H{"Users": u})
}

// GET Logout handler func
func (s *WebApiServer) GetLogout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	//ctx.Redirect(http.StatusSeeOther, "/")
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/logout.html")
	ctx.HTML(http.StatusOK, "layout.html", gin.H{})
}

// GET Login page handler func
func (s *WebApiServer) GetLogin(ctx *gin.Context) {
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/loginFirst.html")
	ctx.HTML(http.StatusOK, "layout.html", gin.H{})
}

// POST for Login page handler func
func (s *WebApiServer) PostLogin(ctx *gin.Context) {
	login := ctx.PostForm("login")
	password := ctx.PostForm("password")

	var errors []string
	var error string

	u, ok, err := s.storage.User().FindByLogin(login)
	if err != nil {
		s.logger.Info("Troubles while creating new user:", err)
		errors = append(errors, "We have some troubles to accessing database. Try again")
		s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/login.html")
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Errors": errors})
		return
	}
	if !ok {
		error = "Incorrect login or password"
		s.logger.Info(error)
		errors = append(errors, error)
		//s.logger.Info("1")
		s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/login.html")
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Errors": errors})
		return
	}
	if u.HashedPassword != password {
		error = "Incorrect login or password"
		s.logger.Info(error)
		errors = append(errors, error)
		//s.logger.Info("2")
		s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/login.html")
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Errors": errors})
		return
	}

	tokenString, err := middleware.CreateToken(u.Login)
	if err != nil {
		error = "Incorrect login or password"
		s.logger.Info(error)
		errors = append(errors, error)
		s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/login.html")
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Errors": errors})
		return
	}

	s.logger.Info("Token created: ", tokenString)

	ctx.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/thanks.html")
	ctx.HTML(http.StatusSeeOther, "layout.html", gin.H{"TextMsg": textOne, "UserLogin": u.Login})
}

// GET Register page
func (s *WebApiServer) GetRegister(ctx *gin.Context) {
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/register.html")
	ctx.HTML(http.StatusOK, "layout.html", gin.H{"Errors": ""})
}

// POST for Register page
func (s *WebApiServer) PostRegister(ctx *gin.Context) {
	var user models.User
	var errors []string
	var error string

	// parse login from form
	user.Login = ctx.PostForm("login")
	if user.Login == "" {
		error = "Empty login field"
		s.logger.Info(error)
		errors = append(errors, error)
	}
	// parse password from form
	pass1 := ctx.PostForm("password")
	pass2 := ctx.PostForm("passwordrety")
	if pass1 != pass2 {
		error = "Password fields no equal"
		s.logger.Info(error)
		errors = append(errors, error)
	}
	user.HashedPassword = pass1
	if pass1 == "" {
		error = "Empty password field"
		s.logger.Info(error)
		errors = append(errors, error)
	}
	// parse field from form
	user.SecretNumber = ctx.PostForm("secretnumber")
	if user.SecretNumber == "" {
		error = "Empty SecretNumber field"
		s.logger.Info(error)
		errors = append(errors, error)
	}
	// if slice errors not empty - send errors
	if len(errors) != 0 {
		ctx.HTML(http.StatusOK, "layout.html", gin.H{"Errors": errors})
		return
	}
	// if slice errors empty - create user
	a, err := s.storage.User().Create(&user)
	if err != nil {
		s.logger.Info("Troubles while creating new user:", err)
		errors = append(errors, "We have some troubles to accessing database. Try again")
		ctx.HTML(http.StatusInternalServerError, "layout.html", gin.H{"Errors": errors})
		return
	}
	s.router.LoadHTMLFiles("./website/template/layout.html", "./website/template/thanks.html")
	ctx.HTML(http.StatusOK, "layout.html", gin.H{"TextMsg": textTwo, "UserLogin": a.Login})
}

// GET user by ID
func (s *WebApiServer) GetUserById(ctx *gin.Context) {

}
