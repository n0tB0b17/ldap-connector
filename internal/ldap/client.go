package ldap

import (
	"fmt"
	"sync"

	"github.com/go-ldap/ldap/v3"
	"github.com/n0tB0b17/ldapconnector/internal/config"
)

type Client struct {
	ldapConn *ldap.Conn
	cfg      config.Cfg
	mu       sync.Mutex
}

func NewClient(cfg config.Cfg) (*Client, error) {
	ldapFullAddr := fmt.Sprintf("%s:%d", cfg.LDAPHost, cfg.LDAPPort)
	conn, err := ldap.DialURL(ldapFullAddr)
	if err != nil {
		fmt.Printf("error while dialing to LDAP server: %v\n", err)
		return nil, err
	}

	err = conn.Bind(cfg.LDAPUser, cfg.LDAPPassword)
	if err != nil {
		conn.Close()
		fmt.Printf("error while binding to LDAP server: %v\n", err)
		return nil, err
	}

	return &Client{ldapConn: conn, cfg: cfg}, nil
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ldapConn != nil {
		c.ldapConn.Close()
		c.ldapConn = nil
	}
}
