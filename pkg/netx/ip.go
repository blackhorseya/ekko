package netx

import (
	"net"
	"net/url"
	"strings"
)

// GetLocalIP4 gets local ip address.
func GetLocalIP4() (ip string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	_, err = net.InterfaceAddrs()
	if err != nil {
		return
	}

	if len(interfaces) == 2 {
		for _, face := range interfaces {
			if strings.Contains(face.Name, "lo") {
				continue
			}
			addrs, err := face.Addrs()
			if err != nil {
				return
			}

			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						currIP := ipnet.IP.String()
						if !strings.Contains(currIP, ":") && currIP != "127.0.0.1" {
							ip = currIP
						}
					}
				}
			}
		}
	}
	for _, face := range interfaces {
		if strings.Contains(face.Name, "lo") {
			continue
		}
		addrs, err := face.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					currIP := ipnet.IP.String()
					if !strings.Contains(currIP, ":") && currIP != "127.0.0.1" && isIntranetIpv4(currIP) {
						ip = currIP
					}
				}
			}
		}
	}

	return
}

// GetIP4FromHost lookup ip from host to get ipv4
func GetIP4FromHost(from string) ([]string, error) {
	uri, err := url.Parse(from)
	if err != nil {
		return nil, err
	}

	ips, err := net.LookupIP(uri.Hostname())
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, ip := range ips {
		if ip.To4() != nil {
			ret = append(ret, ip.To4().String())
		}
	}

	return ret, nil
}

func isIntranetIpv4(ip string) bool {
	if strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "169.254.") ||
		strings.HasPrefix(ip, "172.") ||
		strings.HasPrefix(ip, "10.30.") ||
		strings.HasPrefix(ip, "10.31.") {
		return true
	}
	return false
}
