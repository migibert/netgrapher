package command

import (
  "bytes"
  "log"
  "os/exec"
)

func ExecuteCommand(command string, args ...string) string {
    cmd := exec.Command(command, args...)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return out.String()
}
