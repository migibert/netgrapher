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

    devices := findDevices()
    fmt.Printf("%q\n", devices)
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

func findDevices() []types.Device {
    cmd := command.ExecuteCommand("ip", "addr", "show")
    lines := strings.Split(cmd, "\n")       

    for index, line := range lines {
        if index % 6 == 0 && len(line) != 0 {
            firstSeparatorIndex := strings.Index(line, ":")
            secondSeparatorIndex := strings.Index(line[firstSeparatorIndex+1:len(line)-1], ":")
            deviceName := line[firstSeparatorIndex+1:firstSeparatorIndex+1+secondSeparatorIndex]
            fmt.Printf("%q\n", deviceName)
        }
    }

    return nil

}