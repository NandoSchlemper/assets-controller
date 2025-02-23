package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var myUltraSecretHyperKey = []byte("lol")

func GenerateJwt(payload any) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "payload": payload, 
    })

    tokenString, err := token.SignedString(myUltraSecretHyperKey)

    if err != nil {
        fmt.Printf("Erro ao gerar o token JWT; %s", err.Error())
        return "", err
    }

    return tokenString, nil
}


func ValidateJwt(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(
        tokenString, func(token *jwt.Token) (any, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Metodo de Assinatura errado; %v", token.Header["alg"])
            }
            return myUltraSecretHyperKey, nil
        },
    )
    if err != nil {
        return nil, err
    }
    return token, nil
}

// Sim, eu sei que parece dificil, mas NÃO É, eu juro...
// o Jwt.Parse recebe o token e decodifica ele, a decodificação é feita através
// de uma função callback, o JWT é composto por 3 partes, Header, Payload e Signature
// a funçaõ callback vai fazer uma comparação para se de fato, a assinatura do token
// é a mesma, garantindo que ele não tenha sido modificado
// Nós retornamos a chave secreta para o JWT conseguir utilizar, mas de acordo com minhas pesquisas,
// Não é perigoso, já que não será exposto publicamente! basicamente isso!

