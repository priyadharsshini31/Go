package main

import (
	"fmt"
)

type User struct {
	name       string
	phonenumer int
	email      string
	notifier UserNotifier
}

type UserNotifier interface {
	sendMessage(user *User, message string)
}


type EmailNotifier struct{}

type SmsNotifier struct{}

func (u *User) Notify(message string) {
	u.notifier.sendMessage(u, message)
}

func (e EmailNotifier) sendMessage(user *User, message string) {
	fmt.Println(user.name, user.email)
}

func (s SmsNotifier) sendMessage(user *User, message string) {
	fmt.Println(user.name, user.phonenumer)
}

func main() {
a := &User{"priya",123,"priyagmailc.om", EmailNotifier{}}
b := &User{"priya",123,"priyagmailc.om", SmsNotifier{}}

a.Notify("hello")
b.Notify("sms")	

}
