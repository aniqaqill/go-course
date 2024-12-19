package main

// import "fmt"

// type user struct {
// 	name   string
// 	number int
// }

// type messageToSend struct {
// 	phoneNumber int
// 	message     string
// 	recepient   user
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.recepient.name == "" {
// 		return false
// 	}
// 	if mToSend.recepient.number == 0 {
// 		return false
// 	}
// 	return true
// }

// func test(mToSend messageToSend) {
// 	fmt.Println("Message sent to:", mToSend.recepient.name)
// }

// func main() {
// 	msg := messageToSend{
// 		phoneNumber: 12131313,
// 		message:     "Thanks for signing up",
// 		recepient: user{
// 			name:   "",
// 			number: 12131313,
// 		},
// 	}

// 	if canSendMessage(msg) {
// 		test(msg)
// 	} else {
// 		fmt.Println("Cannot send message. Invalid recipient details.")
// 	}
// }
