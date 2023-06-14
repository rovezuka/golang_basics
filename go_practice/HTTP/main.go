package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", calculateArea)
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func calculateArea(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра "radius" из URL
	radiusStr := r.URL.Query().Get("radius")
	if radiusStr == "" {
		http.Error(w, "Missing 'radius' parameter", http.StatusBadRequest)
		return
	}

	// Преобразуем значение радиуса в число
	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid 'radius' parameter", http.StatusBadRequest)
		return
	}

	// Вычисляем площадь круга
	area := 3.14 * radius * radius

	// Отправляем площадь круга в ответе
	fmt.Fprintf(w, "Area of the circle with radius %.2f is %.2f", radius, area)
}
