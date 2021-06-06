package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaimsExample struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func CreateJwt() {
	// Make a sample token
	// In a real world situation, this token will have been acquired from
	// some other API call (see Example_getTokenViaHTTP)
	serverPort := 8080
	token, err := createToken("foo")
	fatal(err)

	// Make request.  See func restrictedHandler for example request processor
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/restricted", serverPort), nil)
	fatal(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)

	// Read the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	fmt.Println(buf.String())

}

func ParseJwt() (interface{}, error) {

	verifyKey := "key"
	// See func authHandler for an example auth handler that produces a token
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/authenticate", "8080"), url.Values{
		"user": {"test"},
		"pass": {"known"},
	})
	if err != nil {
		fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token out of the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	tokenString := strings.TrimSpace(buf.String())

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	fatal(err)

	claims := token.Claims.(*CustomClaimsExample)
	return claims, nil
}

func fatal(err error) {

}

func createToken(key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{})
	return token.Raw, nil
}
