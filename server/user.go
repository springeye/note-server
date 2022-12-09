package server

import (
	"github.com/go-chi/chi/v5"
	jwtauth "github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/springeye/oplin/db"
	"github.com/springeye/oplin/httputil"
	"github.com/springeye/oplin/model"
	"net/http"
)

var tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key
// @Summary      登录
// @Description  用户登录
// @Tags         用户
// @Accept       json
// @Param        user  body      model.Login  true  "login by  user account"
// @Success      200  {object}  db.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
//
// @Router       /user/login [post]
func login(w http.ResponseWriter, r *http.Request) {
	var inputUser db.User
	if err := render.DecodeJSON(r.Body, &inputUser); err != nil {
		render.JSON(w, r, H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	var outUser db.User
	if err := db.Connection.Find(&outUser, "username = ?", inputUser.Username).Error; err != nil {
		render.JSON(w, r, H{
			"code": 20001,
			"msg":  err.Error(),
		})
		return
	}
	inputPassword := db.CalculatePassword(inputUser.Password, outUser.Salt)
	if inputPassword != outUser.Password {
		render.JSON(w, r, H{
			"code": 401,
			"msg":  "username or password not available",
		})
		return
	}
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"uid":      outUser.ID,
		"username": outUser.Username,
	})
	render.JSON(w, r, H{
		"code": 0,
		"data": tokenString,
	})
}
func userRouter() http.Handler {

	r := chi.NewRouter()

	r.Post("/login", login)
	r.Post("/register", register)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/info", info)
	})

	return r
}

// @Summary      用户信息
// @Description  用户信息
// @Tags         用户
// @Accept       json
// @Success      200   {object}  db.User
//
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
//
// @Router       /user/info [get]
// @Security     BearerAuth
func info(writer http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	httputil.NewResult[map[string]interface{}](writer, r, claims)
}

// @Summary      注册
// @Description  用户注册
// @Tags         用户
// @Accept       json
// @Param        user  body      model.Register  true  "register user account"
// @Success      200   {object}  db.User
//
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
//
// @Router       /user/register [post]
func register(writer http.ResponseWriter, request *http.Request) {
	var req model.Register
	if err := render.DecodeJSON(request.Body, &req); err != nil {
		httputil.NewError(writer, request, 400, err)
		return
	}
	if db.CheckUser(req.Username) {
		render.JSON(writer, request, H{
			"code": 10001,
			"msg":  "user already",
		})

	} else {
		user := db.User{}
		user.Username = req.Username
		user.Salt = uuid.NewString()
		user.Password = db.CalculatePassword(req.Password, user.Salt)
		if err := db.Connection.Create(&user).Error; err != nil {
			render.JSON(writer, request, H{
				"code": 20001,
				"msg":  err.Error(),
			})
			return
		}
		_, token, _ := tokenAuth.Encode(map[string]interface{}{
			"uid":      user.ID,
			"username": user.Username,
		})
		render.JSON(writer, request, H{
			"code": 0,
			"data": token,
		})
	}
}
