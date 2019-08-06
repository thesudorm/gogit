package main

import (
    "fmt"
    "os"
)

func main() {
    init_repo("/home/thesudorm/test")
}

func init_repo(repo string) {
    fmt.Println("Initializing repo")
    err := os.Mkdir(repo + "/.git/", os.ModePerm)

    if err != nil { // Trips if repo already exits
        fmt.Print("Error: ")
        fmt.Println(err)
        return
    }

    dirs := [3]string{"objects", "refs", "refs/heads"}
    for _, dir := range dirs {
        err := os.Mkdir(repo + "/.git/" + dir, os.ModePerm)
        if err != nil {
            fmt.Println(err)
        }
    }
    head, err := os.Create(repo + "/.git/HEAD")
    head.WriteString("ref: refs/heads/master")

    fmt.Println("Succesfully initialized \"" + repo + "/.git\"")
}
