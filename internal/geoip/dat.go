// SPDX-License-Identifier: GPL-3.0-or-later

package geoip

import (
	"fmt"
	"net"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

// geoip.dat is the xray/v2ray GeoIPList protobuf:
//
//	GeoIPList { repeated GeoIP entry = 1; }
//	GeoIP     { string country_code = 1; repeated CIDR cidr = 2; ... }
//	CIDR      { bytes ip = 1; uint32 prefix = 2; }
//
// We only read the fields we need and skip the rest, so newer additions to
// the schema don't break parsing.

// parseGeoIPDat returns a map of lower-cased category code -> CIDR strings.
func parseGeoIPDat(b []byte) (map[string][]string, error) {
	out := map[string][]string{}
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		b = b[n:]
		if num == 1 && typ == protowire.BytesType {
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return nil, protowire.ParseError(n)
			}
			b = b[n:]
			code, cidrs, err := parseGeoIP(v)
			if err != nil {
				return nil, err
			}
			if code != "" && len(cidrs) > 0 {
				out[code] = append(out[code], cidrs...)
			}
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, b)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		b = b[n:]
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no categories found")
	}
	return out, nil
}

func parseGeoIP(b []byte) (string, []string, error) {
	var code string
	var cidrs []string
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", nil, protowire.ParseError(n)
		}
		b = b[n:]
		switch {
		case num == 1 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", nil, protowire.ParseError(n)
			}
			b = b[n:]
			code = strings.ToLower(string(v))
		case num == 2 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", nil, protowire.ParseError(n)
			}
			b = b[n:]
			cidr, err := parseCIDR(v)
			if err == nil && cidr != "" {
				cidrs = append(cidrs, cidr)
			}
		default:
			n = protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return "", nil, protowire.ParseError(n)
			}
			b = b[n:]
		}
	}
	return code, cidrs, nil
}

func parseCIDR(b []byte) (string, error) {
	var ip []byte
	var prefix uint64
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", protowire.ParseError(n)
		}
		b = b[n:]
		switch {
		case num == 1 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", protowire.ParseError(n)
			}
			b = b[n:]
			ip = v
		case num == 2 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(b)
			if n < 0 {
				return "", protowire.ParseError(n)
			}
			b = b[n:]
			prefix = v
		default:
			n = protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return "", protowire.ParseError(n)
			}
			b = b[n:]
		}
	}
	if len(ip) != net.IPv4len && len(ip) != net.IPv6len {
		return "", fmt.Errorf("bad ip length %d", len(ip))
	}
	ipnet := net.IPNet{IP: net.IP(ip), Mask: net.CIDRMask(int(prefix), len(ip)*8)}
	if ipnet.Mask == nil {
		return "", fmt.Errorf("bad prefix %d", prefix)
	}
	return ipnet.String(), nil
}
