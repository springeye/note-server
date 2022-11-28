package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	jwtauth "github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/springeye/oplin/db"
	"net/http"
)

func userRouter() http.Handler {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key
	r := chi.NewRouter()

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
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
	})
	r.Post("/register", func(writer http.ResponseWriter, request *http.Request) {
		var user db.User
		if err := render.DecodeJSON(request.Body, &user); err != nil {
			render.JSON(writer, request, H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}
		if db.CheckUser(user.Username) {
			render.JSON(writer, request, H{
				"code": 10001,
				"msg":  "user already",
			})

		} else {
			user.Salt = uuid.NewString()
			user.Password = db.CalculatePassword(user.Password, user.Salt)
			if err := db.Connection.Create(&user); err != nil {
				render.JSON(writer, request, H{
					"code": 20001,
					"msg":  err.Error,
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
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["username"])))
		})
	})

	return r
}
