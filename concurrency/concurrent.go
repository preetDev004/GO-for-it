package main

import (
	"fmt"
	"bufio"
	"os"
)

func addToQueue(emails []string) chan string {
	// channel with buffer size to store 100 integers.
	emailsToSend := make(chan string, len(emails)) // optional passing - len(emails)

	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		fmt.Println("Sending Email: ", <-ch)
		
	}
	// range works on channel
	// for email := range ch { // loops until channel is empty and closed
	// 	fmt.Println("Sending Email: ", email)
	// }
}

// Read-only channel
// func readCh(ch <-chan int){
	// we can only read from 
	// ch in this function.
// }

// Write-only channel
// func writeCh(ch chan<- int){
	// we can only write to 
	// ch in this function.
// }

func main() {

	ch := make(chan int)
	// send data to channel (BLOCKING OPERATION)
	// -> If we are sending data to the channel and there is no other go routine to read it out.
	//    The code will stop on below go routine untill there is a reader ready.

	// ch <- 89 // will end-up in deadlock as sending and receive is not happening in same go routine.

	// FIX
	go func() {
		ch <- 42
	}()

	// retrieve data from channel (BLOCKING OPERATION)
	// -> If we are receiving data from the channel and there is no other go routine sending the data in.
	//    The code will stop on below go routine untill there is a sender sends.
	v := <-ch
	fmt.Println(v)

	emails := []string{
		"preet@gmail.com",
		"preetpatel@gmail.com",
		"patel@gmail.com",
		"smit@gmail.com",
		"smitpatel@gmail.com",
		"parv@gmail.com",
		"parvpatel@gmail.com",
		"madhav@gmail.com",
		"madhavRajpal@gmail.com",
		"maddy@gmail.com",
	}
	emailChan := addToQueue(emails)

	want:=true
	for want{
		sendEmails(5,emailChan)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Want more:(yes/no) ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if text!="yes\n"{
			fmt.Println("Thanks for sending emails, channel has been closed!")
			close(emailChan)
			want=false
		}else if len(emailChan) == 0 {
			fmt.Println("No more emails are left!")
			want=false
		}
		fmt.Println("")
	} 
}