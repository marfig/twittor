package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

// Email valor del Email usado en todos los Endpoints
var Email string

// IDUsuario el ID devuelto del modelo
var IDUsuario string

// ProcesarToken se procesa el token recibido
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MiClave_Secreta_paraGenerarToken")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
