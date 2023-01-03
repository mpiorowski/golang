package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	grpcMetadata "google.golang.org/grpc/metadata"
)

type Cache struct {
    token string
    expiry time.Time
}

var cacheByHost = map[string]Cache{}

func Connect(env string, host string) (*grpc.ClientConn, error, context.Context, context.CancelFunc) {
	// Create a context with a timeout that will be used to dial the server.
	ctx, cancel, err := CreateContext(env, host)
	if err != nil {
		return nil, err, ctx, cancel
	}

	// Create a TLS credentials object with the certificate authority.
	host = fmt.Sprintf("%s:443", host)

	// For local development use insecure credentials
	if env != "production" {
		conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("grpc.Dial: %v", err)
			return nil, err, ctx, cancel
		}
		return conn, nil, ctx, cancel
	}

	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Printf("x509.SystemCertPool: %v", err)
		return nil, err, ctx, cancel
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))

	// Dial the server.
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("grpc.Dial: %v", err)
		return nil, err, ctx, cancel
	}
	return conn, nil, ctx, cancel
}

func CreateContext(env string, host string) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if env != "production" {
		return ctx, cancel, nil
	}

    token := cacheByHost[host].token
    expiry := cacheByHost[host].expiry

	if token == "" || time.Now().After(expiry) {
        log.Printf("Token not found in cache or expired")
		// Create an identity token.
		// With a global TokenSource tokens would be reused and auto-refreshed at need.
		// A given TokenSource is specific to the audience.
		tokenSource, err := idtoken.NewTokenSource(ctx, "https://"+host)
		if err != nil {
			log.Printf("idtoken.NewTokenSource: %v", err)
			return ctx, cancel, err
		}
		tokenStruct, err := tokenSource.Token()
		if err != nil {
			log.Printf("tokenSource.Token: %v", err)
			return ctx, cancel, err
		}
		token = tokenStruct.AccessToken
		// Cache token
        cacheByHost[host] = Cache{
            token: token,
            expiry: time.Now().Add(time.Hour),
        }
	} else {
        log.Printf("Token found in cache")
    }

	// Add token to gRPC Request.
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	return ctx, cancel, nil
}
