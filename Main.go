package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error en el erchivo .env file")
	}
	//v := os.Getenv("V_PATH")
	v_ip := os.Getenv("V_IP")

	fmt.Println("1")
	fmt.Println(v_ip)
	fmt.Println("2")

	//file, err := os.Create("config.ini")
	//if err != nil {
	//	panic(err)
	//}
	//v_directorio := "c:/microsip datos"
	//data := []byte(v_directorio)
	//file.Write(data)
	//file.WriteString("este es un string \n")
	//file.WriteString("este es un string2")
	//file.Close()
}
