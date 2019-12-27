// +build !linux,!windows

package gost

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/go-log/log"
	"github.com/songgao/water"
)

func createTun(cfg TunConfig) (conn net.Conn, ipNet *net.IPNet, err error) {
	ip, ipNet, err := net.ParseCIDR(cfg.Addr)
	if err != nil {
		return
	}

	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return
	}

	mtu := cfg.MTU
	if mtu <= 0 {
		mtu = DefaultMTU
	}

	cmd := fmt.Sprintf("ifconfig %s inet %s mtu %d up", ifce.Name(), cfg.Addr, mtu)
	log.Log("[tun]", cmd)
	args := strings.Split(cmd, " ")
	if er := exec.Command(args[0], args[1:]...).Run(); er != nil {
		err = fmt.Errorf("%s: %v", cmd, er)
		return
	}

	if err = addRoutes(ifce.Name(), cfg.Routes...); err != nil {
		return
	}

	conn = &tunConn{
		ifce: ifce,
		addr: &net.IPAddr{IP: ip},
	}
	return
}

func addRoutes(ifName string, routes ...string) error {
	for _, route := range routes {
		if route == "" {
			continue
		}
		cmd := fmt.Sprintf("route add -net %s -interface %s", route, ifName)
		log.Log("[tun]", cmd)
		args := strings.Split(cmd, " ")
		if er := exec.Command(args[0], args[1:]...).Run(); er != nil {
			return fmt.Errorf("%s: %v", cmd, er)
		}
	}
	return nil
}