package main

import (
    "fmt"
    "os"
    "crypto/sha1"
    "strconv"
)

func main() {
    arg := os.Args[1]

    switch arg {
        case "init":
            pwd := get_dir()
            init_repo(pwd)
        case "hash-object":
            hash("test")
        default:
            fmt.Println("not a proper command")
    }
}

// Will attempt to make a .git repo for the given file path
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

func hash(input string) {
    hash := sha1.New()

    //to_hash := "blob " + strconv.Itoa(len(input) + 1)  + "\x00" + input
    hash.Write([]byte("blob"))
    hash.Write([]byte(" "))
    hash.Write([]byte(strconv.Itoa(len(input) + 1)))
    hash.Write([]byte{0})
    hash.Write([]byte(input))

    fmt.Printf("%x \n", hash.Sum(nil))
}

// Gets the current working directory 
func get_dir() string {
    path, err := os.Getwd()
    if err != nil{
        panic(err)
    }
    return path
}
