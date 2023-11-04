package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

// Middleware logging for RPC
func loggingMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(startTime)

	log.Printf("RPC method=%s duration=%s error=%v", info.FullMethod, duration, err)

	return resp, err
}

// Middleware logging for HTTP
func loggingHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(startTime)
			log.Printf("HTTP method=%s path=%s duration=%s", r.Method, r.URL.Path, duration)
	})
}