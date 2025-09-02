package xmpp

import "fmt"

// SendRaw memungkinkan mengirim stanza XML mentah ke server XMPP
func (c *Client) SendRaw(stanza string) error {
    if c == nil || c.conn == nil {
        return fmt.Errorf("xmpp client not connected")
    }
    _, err := c.conn.Write([]byte(stanza))
    return err
}
