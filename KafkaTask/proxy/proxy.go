package proxy

import (
	"bufio"

	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func MakeProxy(c *gin.Context, proxyUrl, path string) (err error) {
	req := c.Request
	proxy, err := url.Parse(proxyUrl)
	if err != nil {

		c.String(http.StatusInternalServerError, "error")
		return
	}

	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	req.URL.Path = path
	transport := http.DefaultTransport

	resp, err := transport.RoundTrip(req)

	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	_, _ = bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}
