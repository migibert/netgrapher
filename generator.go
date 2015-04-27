package main

import (
  "fmt"
  "strings"
  "github.com/migibert/netgrapher/command"
  "github.com/migibert/netgrapher/types"
)

func main() {
    namespaces := findNamespaces()
    fmt.Printf("%q\n", namespaces)
}

func findNamespaces() []types.Namespace {
    cmd := command.ExecuteCommand("ip", "netns", "show")
    splittedCmdResult := strings.Split(cmd, "\n")
    namespaces := make([]types.Namespace, len(splittedCmdResult)-1)

    for index := range namespaces {
        namespaces[index] = types.Namespace{ Name: splittedCmdResult[index] }
    }
    return namespaces
}
