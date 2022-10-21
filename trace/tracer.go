package trace

import (
	"fmt"
	"log"
	"net/http/httptrace"
	"os"
	"strings"
)

func Tracer(context string) *httptrace.ClientTrace {
	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", context), 0)
	return &httptrace.ClientTrace{
		GetConn: nil,
		GotConn: func(info httptrace.GotConnInfo) {
			logger.Printf("Connection established\n")
		},
		PutIdleConn:          nil,
		GotFirstResponseByte: nil,
		Got100Continue:       nil,
		Got1xxResponse:       nil,
		DNSStart:             nil,
		DNSDone:              nil,
		ConnectStart:         nil,
		ConnectDone:          nil,
		TLSHandshakeStart:    nil,
		TLSHandshakeDone:     nil,
		WroteHeaderField: func(key string, values []string) {
			logger.Printf("%s: %s\n", key, strings.Join(values, ","))
		},
		WroteHeaders:    func() {},
		Wait100Continue: nil,
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			logger.Printf("Request written (Err = %s)\n", info.Err)
		},
	}
}
