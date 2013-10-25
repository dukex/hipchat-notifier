package main

import (
	"github.com/andybons/hipchat"
  "log"
  "os"
)

func main() {
  args := os.Args
  notify(args[1], args[2], args[3])
}

func notify(from, to, message string) {
  c := hipchat.Client{AuthToken: os.Getenv("HIPCHAT_TOKEN") }
  req := hipchat.MessageRequest{
    RoomId:        to,
    From:          from,
    Message:       message,
    Color:         hipchat.ColorGray,
    MessageFormat: hipchat.FormatHTML,
    Notify:        true,
  }

  if err := c.PostMessage(req); err != nil {
    log.Printf("Expected no error, but got %q", err)
  }
}
