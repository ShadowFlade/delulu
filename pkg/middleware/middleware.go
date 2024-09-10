package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

type COOKIE_STATE int64

const (
	COOKIE_EMPTY COOKIE_STATE = iota
	COOKIE_VALID
	COOKIE_INVALID
)

func IsSameSite(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		mode := env.Get("MODE", "prod")
		header := c.Request().Header
		var isSameUrl bool

		if mode == "prod" {
            fmt.Println(header,"X-HOST")
            remoteIp := header.Get("X-REAL-IP")
            host := header.Get("X-HOST")
            fmt.Println(host, remoteIp, "remte ip")
            isSameUrl = strings.Contains(host, remoteIp)
		} else {
			isSameUrl = true
		}

		cookieState := validateCookie(c.Response().Writer, c.Request())
		fmt.Println(cookieState,isSameUrl,"cooke state")

		if cookieState == COOKIE_VALID || cookieState == COOKIE_EMPTY {
			return next(c)
		} else {
			return c.String(http.StatusBadRequest, "suck it")
		}
	}
}

func validateCookie(r http.ResponseWriter, req *http.Request) COOKIE_STATE {
	secretKeyHex := env.Get("RANDOM_HEX", "")
	fmt.Println(secretKeyHex, "secret hex")
	var err error
	var secretKey []byte
	secretKey, err = hex.DecodeString(secretKeyHex)

	if err != nil {
		log.Fatal(err)
	}

	cookie, err := readSigned(req, "SID", []byte(secretKey))
	fmt.Println(cookie, " cookie up")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			newCookie := http.Cookie{
				Name:     "SID",
				Value:    "Hello world!",
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			err = writeSigned(r, newCookie, secretKey)

			if err != nil {
				log.Println(err)
				http.Error(r, "server error", http.StatusInternalServerError)
				return COOKIE_INVALID

			}
			fmt.Println(cookie, " empty cookie")
			return COOKIE_EMPTY
		case errors.Is(err, ErrInvalidValue):
			http.Error(r, "invalid cookie", http.StatusBadRequest)
			return COOKIE_INVALID
		default:
			log.Println(err)
			http.Error(r, "server error", http.StatusInternalServerError)
			return COOKIE_INVALID
		}
	}
	fmt.Println(cookie, " cookie")
	return COOKIE_VALID
}

func writeCookie(w http.ResponseWriter, cookie http.Cookie) error {

	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	if len(cookie.String()) > 4096 {
		return ErrValueTooLong
	}

	http.SetCookie(w, &cookie)

	return nil
}

func readCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	fmt.Println(cookie, "raeding cookie")
	if err != nil {
		return "", err
	}

	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	fmt.Println(value, "decoded value")
	if err != nil {
		fmt.Println(err, " decoding error")
		return "", ErrInvalidValue
	}

	return string(value), nil
}

func writeSigned(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)
	cookie.Value = string(signature) + cookie.Value
	return writeCookie(w, cookie)
}

func readSigned(r *http.Request, name string, secretKey []byte) (string, error) {
	signedValue, err := readCookie(r, name)
	fmt.Println("signed value", signedValue)
	if err != nil {
		return "", err
	}
	// A SHA256 HMAC signature has a fixed length of 32 bytes. To avoid a potential
	// 'index out of range' panic in the next step, we need to check sure that the
	// length of the signed cookie value is at least this long. We'll use the
	// sha256.Size constant here, rather than 32, just because it makes our code
	// a bit more understandable at a glance.
	if len(signedValue) <= sha256.Size {
		fmt.Println("cookie too small", sha256.Size, len(signedValue))
		return "", ErrInvalidValue
	}
	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]
	// Recalculate the HMAC signature of the cookie name and original value.
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)
	// Check that the recalculated signature matches the signature we received
	// in the cookie. If they match, we can be confident that the cookie name
	// and value haven't been edited by the client.
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}
	return value, nil
}
