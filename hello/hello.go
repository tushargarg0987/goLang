package hello

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomGreet(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    // Here _ stores the index and name stores the value.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// randomGreet starts with lowercase letter and hence is not exported

func randomGreet() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Hello, %v!",
        "Howdy, %v!",
        "Hey, %v!",
    }

    return formats[rand.Intn(len(formats))]
}