package trace

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http/httptrace"
	"net/textproto"
	"os"
	"strings"
)

func Tracer(context string) *httptrace.ClientTrace {
	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", context), 0)
	return &httptrace.ClientTrace{
		GetConn:              nil,
		GotConn:              nil,
		PutIdleConn:          nil,
		GotFirstResponseByte: nil,
		Got100Continue:       nil,
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			logger.Printf("Got1xxResponse (code=%d header=%s)", code, header)
			return nil
		},
		DNSStart:     nil,
		DNSDone:      nil,
		ConnectStart: nil,
		ConnectDone: func(network, addr string, err error) {
			logger.Printf("Connection established (network=%s addr=%s err=%s)\n", network, addr, err)
		},
		TLSHandshakeStart: nil,
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			logger.Printf("tls handshake done (server=%s, err=%s)\n", state.ServerName, err)
		},
		WroteHeaderField: func(key string, values []string) {
			logger.Printf("%s: %s\n", key, strings.Join(values, ","))
		},
		WroteHeaders:    func() {},
		Wait100Continue: nil,
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			logger.Printf("Request written (err=%s)\n", info.Err)
		},
	}
}
