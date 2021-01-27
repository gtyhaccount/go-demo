package main

import (
	"crypto/tls"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type ServerBuilder struct {
	Server
}

func (sb ServerBuilder) Create(addr string,port int) ServerBuilder {
	sb.Addr=addr
	sb.Port=port
	return sb
}

func (sb ServerBuilder) UpdateProtocol(protocol string) ServerBuilder {
	sb.Protocol = protocol
	return sb
}

func (sb ServerBuilder) UpdateMaxConns(maxConns int) ServerBuilder {
	sb.MaxConns = maxConns
	return sb
}

func TestStrut(t *testing.T) {
	sb := ServerBuilder{Server{}}
	sb.Create("addr",8080)
	sb.UpdateProtocol("protocol")
}




func TestName(t *testing.T) {
	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
}
