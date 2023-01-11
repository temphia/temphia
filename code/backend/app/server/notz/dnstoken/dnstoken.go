package dnstoken

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

func DNSReverseResolve(cluster, host string) (string, error) {

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	txts, err := r.LookupTXT(context.Background(), host)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("temphia_%s=", cluster)
	for _, txt := range txts {
		if !strings.HasPrefix(txt, key) {
			continue
		}

		return strings.TrimLeft(txt, key), nil
	}

	return "", easyerr.NotFound()

}
