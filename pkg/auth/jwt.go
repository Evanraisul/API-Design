package auth

import (
	"encoding/json"
	"fmt"
	"github.com/evanraisul/book_api/model"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var secretKey = []byte("SecretKey")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token\n")
	}

	return nil
}

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := fmt.Fprint(w, "Missing authorization header\n")
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := fmt.Fprint(w, "Invalid token\n")
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		/*
			_, er := fmt.Fprint(w, "Welcome to the the protected area\n")
			if er != nil {
				fmt.Println(err)
			}
		*/
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The user request value %v\n", u)

	if u.Username == "admin" && u.Password == "admin" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := fmt.Errorf("no username found\n")
			if err != nil {
				fmt.Println(err)
			}
		}
		w.WriteHeader(http.StatusOK)
		_, er := fmt.Fprint(w, tokenString)

		if er != nil {
			fmt.Println(er)
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprint(w, "Invalid credentials\n")

		if err != nil {
			fmt.Println(err)
		}
	}
}
