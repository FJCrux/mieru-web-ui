// SPDX-License-Identifier: GPL-3.0-or-later

package geoip

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

// geosite.dat is the xray/v2ray GeoSiteList protobuf:
//
//	GeoSiteList { repeated GeoSite entry = 1; }
//	GeoSite     { string country_code = 1; repeated Domain domain = 2; }
//	Domain      { Type type = 1; string value = 2; ... }
//	Type        { Plain = 0; Regex = 1; Domain = 2; Full = 3; }
//
// mita matches a DomainName exactly or as a dot-separated suffix, with no
// regex or substring support. So we keep the Domain (suffix) and Full (exact,
// treated as a suffix - a minor over-match) entries and drop Plain (keyword)
// and Regex, which mita cannot represent.

// parseGeoSiteDat returns a map of lower-cased category code -> domain suffixes.
func parseGeoSiteDat(b []byte) (map[string][]string, error) {
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
			code, domains, err := parseGeoSite(v)
			if err != nil {
				return nil, err
			}
			if code != "" && len(domains) > 0 {
				out[code] = append(out[code], domains...)
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

func parseGeoSite(b []byte) (string, []string, error) {
	var code string
	var domains []string
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
			if d, ok := parseDomain(v); ok {
				domains = append(domains, d)
			}
		default:
			n = protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return "", nil, protowire.ParseError(n)
			}
			b = b[n:]
		}
	}
	return code, domains, nil
}

// Domain.Type values that map to mita's suffix matching.
const (
	domainTypePlain = 0 // keyword/substring - unsupported
	domainTypeRegex = 1 // regex - unsupported
	domainTypeRoot  = 2 // suffix ("google.com" matches google.com and *.google.com)
	domainTypeFull  = 3 // exact FQDN
)

// parseDomain returns the domain value and whether it is usable by mita
// (Domain/Full types only).
func parseDomain(b []byte) (string, bool) {
	var dtype uint64
	var value string
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", false
		}
		b = b[n:]
		switch {
		case num == 1 && typ == protowire.VarintType:
			v, n := protowire.ConsumeVarint(b)
			if n < 0 {
				return "", false
			}
			b = b[n:]
			dtype = v
		case num == 2 && typ == protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", false
			}
			b = b[n:]
			value = strings.ToLower(strings.TrimSpace(string(v)))
		default:
			n = protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return "", false
			}
			b = b[n:]
		}
	}
	if value == "" {
		return "", false
	}
	if dtype != domainTypeRoot && dtype != domainTypeFull {
		return "", false // Plain and Regex cannot be represented in mita rules
	}
	return value, true
}
