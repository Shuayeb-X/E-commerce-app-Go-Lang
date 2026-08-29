package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt
		// parse header and payload or claims
		// hmac-sha-256 algorithm -> hash hmac(heade,payload,secret key)
		// parse signature part from the jwt
		// if the signature and hash is same => forward to create products
		// others 401 status code with unauthorized

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}

		headerArray := strings.Split(header, " ")
		if len(headerArray) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArray[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// fmt.Println("----------TokenParts_------", tokenParts)

		jwtHeader := tokenParts[0]
		// fmt.Println("JWT header: ", jwtHeader)

		jwtPayload := tokenParts[1]
		// fmt.Println("JWT PayLoad/Claim: ", jwtPayload)

		jwtSignature := tokenParts[2]
		// fmt.Println("JWT Signature: ", jwtSignature)

		message := jwtHeader + "." + jwtPayload

		cnf := config.GetConfig()

		byteArrMessage := []byte(message)

		byteArrSecret := []byte(cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)

		newSignatureB64 := base64UrlEncode(hash)

		if newSignatureB64 != jwtSignature {
			http.Error(w, "Unauthorized ,You are a hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
