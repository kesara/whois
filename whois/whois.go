// Package whois provides WHOIS query capabilities and data parsing
package whois

import (
    "bytes"
    "io"
    "net"
)

type WhoisServer struct {
    Server string
    Port string
}

// Query makes an WHOIS query to the given host
// Returns the raw output as string and error
func Query(host WhoisServer, domain string) (string, error) {
    conn, err := net.Dial("tcp", net.JoinHostPort(host.Server, host.Port))
    if err != nil {
        return "", err
    }

    _, err = conn.Write(append([]byte(domain), '\r', '\n'))
    if err != nil {
        return "", err
    }

    var buffer bytes.Buffer
    io.Copy(&buffer, conn)
    return buffer.String(), nil
}
