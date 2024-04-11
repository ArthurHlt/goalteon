package goalteaon

import (
	transport "github.com/ArthurHlt/basicauth-transport"
	"net/http"
	"sync"
)

type clientOpt func(c *Client)

func WithLocker(locker sync.Locker) clientOpt {
	return func(c *Client) {
		c.locker = locker
	}
}

func WithInsecureSkipVerify(insecureSkipVerify bool) clientOpt {
	return func(c *Client) {
		c.httpClient.Transport.(*transport.BasicAuthTransport).
			WrapTransport.(*http.Transport).
			TLSClientConfig.InsecureSkipVerify = insecureSkipVerify
	}
}

func WithNoAutoApplySaveSync() clientOpt {
	return func(c *Client) {
		c.noAutoApplySaveSync = true
	}
}

func WithNoLock() clientOpt {
	return WithLocker(&NoLocker{})
}

type NoLocker struct{}

func (*NoLocker) Lock()   {}
func (*NoLocker) Unlock() {}
