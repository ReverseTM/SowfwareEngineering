package main

import (
	"context"
	"log"
	"software-engineering/internal/mediator"
	"software-engineering/internal/mediator/aggregator"
	"software-engineering/internal/model"
	"software-engineering/internal/observer"
	"software-engineering/internal/observer/logger"
	"software-engineering/internal/storage"
	"software-engineering/internal/storage/postgresql"
	"software-engineering/internal/storage/postgresql/user"
)

func main() {
	ctx := context.Background()

	connString := "postgresql://postgres:postgres@localhost:5432/software-engineering"
	pool, err := postgresql.GetConnPool(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}

	eventAggregator := aggregator.NewEventAggregator()
	auditLogger := logger.NewAuditLogger()
	userStorage := user.NewUserStorage(pool, eventAggregator)

	testing(ctx, eventAggregator, auditLogger, userStorage)
}

func testing(
	ctx context.Context,
	eventAggregator mediator.Mediator,
	auditLogger observer.Observer,
	userStorage *user.Storage,
) {
	log.Println("Subscribing INSERT event to audit logger")
	eventAggregator.Subscribe(storage.Insert, auditLogger)

	log.Println("Subscribing UPDATE event to audit logger")
	eventAggregator.Subscribe(storage.Update, auditLogger)

	log.Println("Subscribing DELETE event to audit logger")
	eventAggregator.Subscribe(storage.Delete, auditLogger)

	firstUser := &model.User{
		ID:      1,
		Name:    "John",
		Surname: "Smith",
		Age:     30,
	}

	log.Printf("Inserting user %+v\n", firstUser)
	err := userStorage.Insert(ctx, firstUser)
	if err != nil {
		log.Fatal(err)
	}

	firstUser.Age = 31

	log.Printf("Updating user %+v\n, new age=%d", firstUser, firstUser.Age)
	err = userStorage.Update(ctx, firstUser)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deliting user %+v\n", firstUser)
	err = userStorage.Delete(ctx, firstUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Unsubscribing UPDATE event from audit logger")
	eventAggregator.Unsubscribe(storage.Update, auditLogger)

	secondUser := &model.User{
		ID:      2,
		Name:    "Jane",
		Surname: "Wilson",
		Age:     25,
	}

	log.Printf("Inserting user %+v\n", secondUser)
	err = userStorage.Insert(ctx, secondUser)
	if err != nil {
		log.Fatal(err)
	}

	secondUser.Surname = "Smith"
	secondUser.Age = 26

	log.Printf("Updating user %+v\n, new surname=%s, new age=%d", secondUser, secondUser.Surname, secondUser.Age)
	err = userStorage.Update(ctx, secondUser)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deliting user %+v\n", secondUser)
	err = userStorage.Delete(ctx, secondUser.ID)
	if err != nil {
		log.Fatal(err)
	}
}
