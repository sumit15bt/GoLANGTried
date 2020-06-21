package main 

import (
    "fmt"
    "io/ioutil"
    "net/http"
)


func main() {
    fmt.Println(" hello go.. Starting the application...")
    response1, err := http.Get("API1")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data1, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data1))
    }
    fmt.Println("Terminating the application...")
}
