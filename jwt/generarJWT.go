package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/marfig/twittor/models"
)

// GenerarJWT genera el token
func GenerarJWT(t models.Usuario) (string, error) {
	miClave := []byte("MiClave_Secreta_paraGenerarToken")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, err
}
