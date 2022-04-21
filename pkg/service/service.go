package service

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	trip "github.com/nvsces/flw-server-go"
	"github.com/nvsces/flw-server-go/pkg/repository"
	"google.golang.org/api/option"
	// "firebase.google.com/go/v4/messaging"
)

type Authorization interface {
	CreateUser(user trip.User) (int, error)
	GenerateToken(user_vk_id int) (string, error)
	ParseToken(token string) (int, error)
	GetUser(user_vk_id int) (trip.User, error)
}

type Profile interface{
	Info(userId int) (trip.User, error)
}


type TripItem interface {
	Create(item trip.TripItem) (int, error)
	GetAll() ([]repository.ObjectOutputJson, error)
	GetById(itemId int) (trip.TripItem, error)
	Delete(userId,itemId int) error
}

type Service struct {
	Authorization
	TripItem
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TripItem: NewTripItemService(repos.TripItem),
		Profile: NewProfileService(repos.Profile),
	}
}

func SendMessage(){

opt := option.WithCredentialsFile("serviceAccountKey.json")
app, err := firebase.NewApp(context.Background(), nil, opt)
if err != nil {
  return 
}

ctx := context.Background()
client, err := app.Messaging(ctx)
if err != nil {
        log.Fatalf("error getting Messaging client: %v\n", err)
}

// This registration token comes from the client FCM SDKs.
registrationToken := "fNIrkFCbQe-AYu8b3fmOEO:APA91bGj8MMM16WLfjxSj2xWOw9-wEqMxdYySsE3ES5LEmXHk2RBe0PRpfmo1N-ta66b8gJWRAd9DxzpArPekiKVMAf_e-jqSuuBsr6MGo_-h9vV93IxYUCh7EpAdrXLSbHRP4Ms3bSB"

// See documentation on defining a message payload.
message := &messaging.Message{
        Data: map[string]string{
                "score": "850",
                "time":  "2:45",
        },
        Token: registrationToken,
}

// Send a message to the device corresponding to the provided
// registration token.
response, err := client.Send(ctx, message)
if err != nil {
        log.Fatalln(err)
}
// Response is a message ID string.
fmt.Println("Successfully sent message:", response)
}