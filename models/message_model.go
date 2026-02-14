package models


type Message struct{
	SenderId string `json:"SenderId"`
	RoomId string `json:"RoomId"`
	Msg string `json:"Msg"`

}