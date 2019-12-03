package dev

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jim-minter/rp/pkg/env/shared"
)

type dev struct {
	*shared.Shared
}

func New(ctx context.Context, log *logrus.Entry) (*dev, error) {
	for _, key := range []string{
		"LOCATION",
		"RESOURCEGROUP",
	} {
		if _, found := os.LookupEnv(key); !found {
			return nil, fmt.Errorf("environment variable %q unset", key)
		}
	}

	d := &dev{}

	var err error
	d.Shared, err = shared.NewShared(ctx, log, os.Getenv("AZURE_TENANT_ID"), os.Getenv("AZURE_SUBSCRIPTION_ID"), os.Getenv("RESOURCEGROUP"))
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *dev) ListenTLS(ctx context.Context) (net.Listener, error) {
	key, cert, err := d.GetSecret(ctx, "tls")
	if err != nil {
		return nil, err
	}

	// no TLS client cert verification in dev mode, but we'll only listen on
	// localhost
	return tls.Listen("tcp", "localhost:8443", &tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{
					cert.Raw,
				},
				PrivateKey: key,
			},
		},
		MinVersion: tls.VersionTLS12,
	})
}

func (d *dev) Authenticated(h http.Handler) http.Handler {
	return h
}

func (d *dev) IsReady() bool {
	return true
}

func (d *dev) Location() string {
	return os.Getenv("LOCATION")
}

func (d *dev) ResourceGroup() string {
	return os.Getenv("RESOURCEGROUP")
}
