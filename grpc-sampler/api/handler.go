package api

import (
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// CToF Converter for Celsius to Fahrenheit
func CToF(celsius float32) float32 {
	return celsius*9/5 + 32
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	log.Printf("Received Celsius Temperature of %f", in.TemperatureCelsius)
	return &PingResponse{TemperatureFahrenheit: CToF(in.TemperatureCelsius)}, nil
}
