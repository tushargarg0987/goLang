package main

import (
	"fmt"
	"log"

	"example.com/hello"
)

func main(){
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	message,err := hello.Hello("Tushar")

	// If an error was returned, print it to the console and
    // exit the program.
	if err != nil {
        log.Fatal(err)
    }

	fmt.Println(message)
	// log.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
    messages, err2 := hello.Hellos(names)
    if err2 != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)

}