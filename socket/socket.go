package socket

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/centrifugal/centrifuge-go"
	jwt "github.com/dgrijalva/jwt-go"
)

// SocketClient ...
var SocketClient *centrifuge.Client

type eventHandler struct{}

func (h *eventHandler) OnPrivateSub(c *centrifuge.Client, e centrifuge.PrivateSubEvent) (string, error) {
	token := subscribeToken(e.Channel, e.ClientID, time.Now().Unix()+10)
	return token, nil
}

func (h *eventHandler) OnConnect(c *centrifuge.Client, e centrifuge.ConnectEvent) {
	log.Println("Connected")
}

func (h *eventHandler) OnError(c *centrifuge.Client, e centrifuge.ErrorEvent) {
	log.Println("Error", e.Message)
}

func (h *eventHandler) OnDisconnect(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
	log.Println("Disconnected", e.Reason)
}

type subEventHandler struct{}

func (h *subEventHandler) OnSubscribeSuccess(sub *centrifuge.Subscription, e centrifuge.SubscribeSuccessEvent) {
	log.Println(fmt.Sprintf("Successfully subscribed to private channel %s", sub.Channel()))
}

func (h *subEventHandler) OnSubscribeError(sub *centrifuge.Subscription, e centrifuge.SubscribeErrorEvent) {
	log.Println(fmt.Sprintf("Error subscribing to private channel %s: %v", sub.Channel(), e.Error))
}

func (h *subEventHandler) OnUnsubscribe(sub *centrifuge.Subscription, e centrifuge.UnsubscribeEvent) {
	log.Println(fmt.Sprintf("Unsubscribed from private channel %s", sub.Channel()))
}

func (h *subEventHandler) OnPublish(sub *centrifuge.Subscription, e centrifuge.PublishEvent) {
	log.Println(fmt.Sprintf("New message received from channel %s: %s", sub.Channel(), string(e.Data)))
}

func connToken(user string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SOCKET_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}

func subscribeToken(channel string, client string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"channel": channel, "client": client}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SOCKET_SECRET_ADMIN")))
	if err != nil {
		panic(err)
	}
	return t
}

// InitSocketEngine ...
func InitSocketEngine(username string, expiration int64) {
	wsURL := "ws://" + os.Getenv("SOCKET_ADDR") + "/connection/websocket?format=protobuf"
	c := centrifuge.New(wsURL, centrifuge.DefaultConfig())
	c.SetToken(connToken(username, expiration))
	handler := &eventHandler{}
	c.OnPrivateSub(handler)
	c.OnDisconnect(handler)
	c.OnConnect(handler)
	c.OnError(handler)

	err := c.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	SocketClient = c
}
