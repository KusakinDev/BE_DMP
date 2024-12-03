package login

import (
	dbA "back/db"
	jwtA "back/jwt/generateJWT"
	us "back/struct/userStruct"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user us.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неудалось декодировать JSON", http.StatusBadRequest)
		return
	}

	// Проверка пользователя в базе данных
	query := "SELECT id, password FROM users WHERE name=$1"
	args := user.Name
	row := dbA.DB.QueryRow(query, args)
	var storedPassword string
	var userID int
	err := row.Scan(&userID, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Пользователь не найдет", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Пока что пароль не захеширован, просто сравниваем строки
	if user.Password != storedPassword {
		http.Error(w, "Пароль неверный", http.StatusUnauthorized)
		return
	}

	tokenString, err := jwtA.GenerateJWT(userID, user.Type)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправляем JWT токен в ответе
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}
