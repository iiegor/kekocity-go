package types

import "time"

const (
  DEBUG bool = true

  // Service port
  SERVICE int = 8080

  // Time allowed to write a message to the peer.
  WRITE_WAIT = 10 * time.Second

  // Time allowed to read the next pong message from the peer.
  PONG_WAIT = 60 * time.Second

  // Send pings to peer with this period. Must be less than pongWait.
  PING_PERIOD = (PONG_WAIT * 9) / 10

  // Maximum message size allowed from peer.
  MAX_MESSAGE_SIZE = 512
)
