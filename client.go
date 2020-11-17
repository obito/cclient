package cclient

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, enableJar bool, proxyUrl ...string) (http.Client, error) {
	if len(proxyUrl) > 0 && len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{}, err
		}

		if enableJar {
			jar, err := cookiejar.New(nil)
			if err != nil {
				return http.Client{}, err
			}

			return http.Client{
				Transport: newRoundTripper(clientHello, dialer),
				Jar:       jar,
			}, nil
		}

		return http.Client{
			Transport: newRoundTripper(clientHello, dialer),
		}, nil
	}

	if enableJar {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return http.Client{}, err
		}

		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
			Jar:       jar,
		}, nil
	}

	return http.Client{
		Transport: newRoundTripper(clientHello, proxy.Direct),
	}, nil

}
