package main

import (
	"errors"
	"fmt"
	"net/http"
)

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type SimpleDataStore struct{
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool){
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore{
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Jorma",
			"3": "Pat",
		},
	}
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string){
	lg(message)
}

func LogOutput(message string){
	fmt.Println(message)
}

type SimpleLogic struct {
	l Logger
	ds DataStore
}

func NewSipleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l: l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

func (sl SimpleLogic) SayHello(userID string) (string,error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

type Controller struct{
	l Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request){
	c.l.Log("In SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller{
	return Controller{
		l: l,
		logic: logic,
	}
}

func main(){
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSipleLogic(l, ds)
	c := NewController(l, logic)

	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
