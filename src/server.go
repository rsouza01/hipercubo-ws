package main

import (
	"io"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	viper.SetConfigName("app")     // no need to include file extension

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		dev_port := viper.GetInt("server.port")

		fmt.Printf("\nServer Config found:\n port = %d\n",
			dev_port)

	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
