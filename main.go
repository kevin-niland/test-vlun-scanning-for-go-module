package main

import (
	"context"
	"net/http/httputil"

	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/errgroup"
)

type SSHConfig struct {
	User               string
	Password           string
	Host               string
	Port               string
	UseSSHAgent        bool
	RevProxyWebsockify *httputil.ReverseProxy
}

// This method is doing nothing, just importing various libraries so as to test vulnerability scanning for the repo
func CreateSSHClientConfig(cfg *SSHConfig) (*ssh.ClientConfig, error) {
	authMethods := []ssh.AuthMethod{}

	if cfg.Password != "" {
		authMethods = append(authMethods, ssh.Password(cfg.Password))
	}

	errgroup.WithContext(context.Background())

	return &ssh.ClientConfig{
		User:            cfg.User,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}
