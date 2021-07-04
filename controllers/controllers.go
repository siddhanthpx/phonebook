package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/siddhanthpx/phonebook/database"
	"github.com/siddhanthpx/phonebook/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = os.Getenv("SECRET_KEY")

func Register(rw http.ResponseWriter, r *http.Request) {

	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		log.Fatal(err)
		return
	}

	user := models.User{
		Name:        data["name"],
		PhoneNumber: data["phone_number"],
		Password:    string(password),
	}

	if err := database.RegisterUser(user); err != nil {
		log.Fatal(err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&user)

	defer r.Body.Close()

}

func Login(rw http.ResponseWriter, r *http.Request) {

	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
		return
	}

	user, err := database.VerifyUser(data["phone_number"])
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		http.Error(rw, "Could not find user", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		http.Error(rw, "Bad credentials", http.StatusUnauthorized)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		http.Error(rw, "Could not log in", http.StatusNotFound)
		return
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{
		"message": "success",
	})

}

func User(rw http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("jwt")
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		http.Error(rw, "Not authenticated Cookie", http.StatusUnauthorized)
		return
	}

	fmt.Println(cookie.Value)

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := database.GetUser(claims.Issuer)
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&user)

}

func Logout(rw http.ResponseWriter, r *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{
		"message": "successfully logged out",
	})

}
