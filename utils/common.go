package utils

import (
	//Importaciones de go (vienen incluidas al instalar)
	"os"
	"fmt"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/secretsmanager"
	//"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
)

var SecretModel entities.SecretRDSJson
var err error

//Lee el secreto de SecretsMananger Conectado a RDS en AWS
func ReadSecret() error {
	SecretModel, err = secundary.GetSecret(os.Getenv("SecretName"))
	return err
}

//Valida El Token Para ver que tenga authorization para ingresar

func ValidoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "/ecommerceEscom/product" && method == "GET") ||
		(path == "/ecommerceEscom/category" && method == "GET") {
		
			fmt.Println("No Necesita Autorización porque es es GET")
			return true, 200, ""
	}

	// Recibimos el token que viene en el headers
	token := headers["Authorization"]
	fmt.Println("Imprimiremos el token que nos dan authorization: ")
	fmt.Println(token)
	fmt.Println("Ya imprimimos el token")

	// Recibimos el token que viene en el headers 02
	token02 := headers["Authorization"]
	fmt.Println("Imprimiremos el token02 que nos dan authorization: ")
	fmt.Println(token02)
	fmt.Println("Ya imprimimos el token02")

	//verificamos que si hayamos recibido la autorización de "authorization"
	if len(token) == 0 {
		return false, 401, "Token Requerido"
	}

	fmt.Println("Se empezará a validar token con VALIDOTOKEN()")
	//Si nos llego el token correctamente validamos el token sea correcto
	todoOK, err, msg := secundary.ValidoToken(token)

	// si algo no estuvo bien, verificamos que fue lo que fallo en el token
	if !todoOK {
		//Verificamos si fallo la verificació del token porque existio un error
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token porque: " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK Yei")

	return true, 200, msg
}
