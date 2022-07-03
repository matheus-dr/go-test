package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/matheus-dr/go-test/infra/kafka"
	"github.com/matheus-dr/go-test/infra/repository"
	"github.com/matheus-dr/go-test/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp:(mysql:3306)/matheus-dr")
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository.CourseMysqlRepository{Db: db}

	useCase := usecase.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id":          "app-go",
	}

	topics := []string{"go-test.create-course"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase.CreateCourseInputDTO
		err := json.Unmarshal(msg.Value, &input)
		if err != nil {
			panic(err)
		}

		output, err := useCase.Execute(input)
		if err != nil {
			fmt.Println("Error...", err)
		} else {
			fmt.Println(output)
		}
	}
}
