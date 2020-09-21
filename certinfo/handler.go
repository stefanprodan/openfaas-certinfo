package function

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	handler "github.com/openfaas/templates-sdk/go-http"
)

func Handle(req handler.Request) (handler.Response, error) {
	var err error

	request := strings.ToLower(string(req.Body))

	u, err := url.Parse(request)
	if err != nil {
		return handler.Response{
			Body:       []byte(err.Error()),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	address := u.Hostname() + ":443"
	ipConn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return handler.Response{
			Body:       []byte(fmt.Sprintf("SSL/TLS not enabed on %v\nDial error: %v", u.Hostname(), err)),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	defer ipConn.Close()
	conn := tls.Client(ipConn, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         u.Hostname(),
	})
	if err = conn.Handshake(); err != nil {
		return handler.Response{
			Body:       []byte(fmt.Sprintf("Invalid SSL/TLS for %v\nHandshake error: %v", address, err)),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	defer conn.Close()
	addr := conn.RemoteAddr()
	host, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		return handler.Response{
			Body:       []byte(err.Error()),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	cert := conn.ConnectionState().PeerCertificates[0]
	asJson := os.Getenv("Http_Query")

	if len(asJson) > 0 && asJson == "output=json" {
		res := struct {
			Host          string
			Port          string
			Issuer        string
			CommonName    string
			NotBefore     time.Time
			NotAfter      time.Time
			NotAfterUnix  int64
			SANs          []string
			TimeRemaining string
		}{
			host,
			port,
			cert.Issuer.CommonName,
			cert.Subject.CommonName,
			cert.NotBefore,
			cert.NotAfter,
			cert.NotAfter.Unix(),
			cert.DNSNames,
			humanize.Time(cert.NotAfter),
		}

		b, err := json.Marshal(res)
		if err != nil {
			return handler.Response{
				Body:       []byte(err.Error()),
				StatusCode: http.StatusInternalServerError,
			}, err
		}

		return handler.Response{
			Body:       b,
			StatusCode: http.StatusOK,
		}, err
	}

	message := fmt.Sprintf("Host %v\nPort %v\nIssuer %v\nCommonName %v\nNotBefore %v\nNotAfter %v\nNotAfterUnix %v\nSANs %v\nTimeRemaining %v",
		host, port, cert.Issuer.CommonName, cert.Subject.CommonName, cert.NotBefore, cert.NotAfter, cert.NotAfter.Unix(), cert.DNSNames, humanize.Time(cert.NotAfter))
	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
