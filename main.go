package main

import "os/exec"
import . "fmt"

func main() {
    app := "wall"

    arg0 := "pwned :("

    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()

    if err != nil {
        Println(err.Error())
        return
    }

    Print(string(stdout))
}

