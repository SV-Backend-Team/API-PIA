package jwt_config

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JWT_Config struct {
	Key string
}

func (jwt_config JWT_Config) GetToken(c echo.Context) error {

	log.Println("get token")
	app := c.FormValue("app")
	if app == "" {
		return c.String(http.StatusBadRequest, "Must send App Name")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["app"] = app

	secret, err := token.SigningString()
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, "Failed to generate singing string")
	}
	secret = "6AB11EE40BC4545298064C0A739CD3CF"
	jwt_config.Key = secret

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, "Failed to create token")
	}

	log.Printf("Key: %s\n", jwt_config.Key)
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}
