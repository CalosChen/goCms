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
		// return []byte("AllYourBase"), nil
		return verifyKey, nil

	})
	fatal(err)

	claims := token.Claims.(*CustomClaimsExample) //此处为类型转换, type conversion in go using .(type)
	return claims, nil
}

func fatal(err error) {
	test(func(a string) {
	})
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func createToken(key string) (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss, err
}

func test(cb func(s string)) {
	cb("")
}
