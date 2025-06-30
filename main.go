package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Define a struct matching your JSON structure
type Person struct {
    Name   string   `json:"name"`
    Age    int      `json:"age"`
    Skills []string `json:"skills"`
}



func main() {
    // Open the JSON file
    file, err := os.Open("data.json")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read the file's content
    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Parse the JSON
    var person Person
    err = json.Unmarshal(bytes, &person)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Use the parsed data
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}