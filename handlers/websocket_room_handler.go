package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	ID     string
	Name   string
	Socket *websocket.Conn
}

type Room struct {
	ID    string
	Users map[string]*User
}

var (
	rooms    = make(map[string]*Room)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func HandleConnection(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
		log.Printf("error: %v", err)
	}

	defer socket.Close()

	var roomID, userID string

	for {
		var msg map[string]interface{}
		if err := socket.ReadJSON(&msg); err != nil {
			log.Println("Read error:", err)
			break
		}

		switch msg["type"] {
		case "create-room":
			createRoom(socket)
		case "join-room":
			roomID, userID = joinRoom(socket, msg)
		case "change-name":
			changeNameHandler(roomID, msg)
		case "clear-other-peers":
			clearOtherPeers(roomID, msg)
		case "offer", "answer", "candidate", "send-call":
			broadcastToRoom(rooms[roomID], msg)
		}
	}

	disconnectUser(roomID, userID)
}

func createRoom(socket *websocket.Conn) {
	roomID := uuid.New().String()
	rooms[roomID] = &Room{ID: roomID, Users: make(map[string]*User)}
	socket.WriteJSON(gin.H{"type": "room-created", "roomId": roomID})
	log.Println("Room created:", roomID)
}

func joinRoom(socket *websocket.Conn, msg map[string]interface{}) (string, string) {
	roomID := msg["roomId"].(string)
	userName := msg["userName"].(string)
	userID := uuid.New().String()

	fmt.Println("room:", rooms[roomID])
	room, exists := rooms[roomID]
	if !exists {
		log.Println("Room does not exist:", roomID)
		return "", ""
	}

	room.Users[userID] = &User{ID: userID, Name: userName, Socket: socket}

	broadcastToRoom(room, gin.H{
		"type":     "user-joined",
		"userId":   userID,
		"userName": userName,
	})
	socket.WriteJSON(gin.H{
		"type":         "get-users",
		"roomId":       roomID,
		"participants": getRoomParticipants(room),
	})
	log.Printf("User %s joined room %s", userName, roomID)
	return roomID, userID
}

func changeNameHandler(roomID string, msg map[string]interface{}) {
	room := rooms[roomID]
	userID := msg["peerId"].(string)
	userName := msg["userName"].(string)

	if user, exists := room.Users[userID]; exists {
		user.Name = userName
		broadcastToRoom(room, gin.H{
			"type":     "name-changed",
			"peerId":   userID,
			"userName": userName,
		})
	}
}

func clearOtherPeers(roomID string, msg map[string]interface{}) {
	room := rooms[roomID]
	currentUserID := msg["currentUserId"].(string)

	for peerID := range room.Users {
		if peerID != currentUserID {
			delete(room.Users, peerID)
			broadcastToRoom(room, gin.H{
				"type":   "peer-disconnected",
				"peerId": peerID,
			})
		}
	}
}

func disconnectUser(roomID, userID string) {
	if room, exists := rooms[roomID]; exists {
		delete(room.Users, userID)
		if len(room.Users) == 0 {
			delete(rooms, roomID)
		}
		log.Printf("User %s left room %s", userID, roomID)
	}
}

func broadcastToRoom(room *Room, msg map[string]interface{}) {
	for _, user := range room.Users {
		err := user.Socket.WriteJSON(msg)
		if err != nil {
			log.Println("Broadcast error:", err)
		}
	}
}

func getRoomParticipants(room *Room) map[string]string {
	participants := make(map[string]string)
	for id, user := range room.Users {
		participants[id] = user.Name
	}
	return participants
}
