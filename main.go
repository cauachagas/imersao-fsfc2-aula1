package main

import (
	"log"

	"github.com/cauachagas/imersao-fsfc2/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	// Gambiarrra. Loop infinito
	for {
		_ = 1
	}

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[1])

}
