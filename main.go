package main

import (
	"fmt"
	"net/http"
)

func getIP(ip string) (string, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	// Выполняем запрос
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return "sd", nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "FILIkR API collection: ip")
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("ip")

	if name == "" {
		fmt.Fprintln(w, "Not IP")
		return
	}

	fmt.Fprintf(w, "Привет, %s!\n", name)
}

func main() {
	// Регистрируем обработчики для разных путей
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/about", aboutHandler)

	// Запуск HTTP сервера на порту 8080
	// fmt.Println("Сервер запущен на http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	url := fmt.Sprintf("http://ip-api.com/json/%s", "37.212.92.167")

	// Выполняем запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(resp)

}
