// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// +build go1.3

package lxd

import (
	"os/exec"
	"strconv"

	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/packaging/commands"
	"github.com/juju/utils/packaging/manager"
	"github.com/juju/utils/series"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/testing"
)

type InitialiserSuite struct {
	testing.BaseSuite
	calledCmds []string
}

var _ = gc.Suite(&InitialiserSuite{})

// getMockRunCommandWithRetry is a helper function which returns a function
// with an identical signature to manager.RunCommandWithRetry which saves each
// command it recieves in a slice and always returns no output, error code 0
// and a nil error.
func getMockRunCommandWithRetry(calledCmds *[]string) func(string) (string, int, error) {
	return func(cmd string) (string, int, error) {
		*calledCmds = append(*calledCmds, cmd)
		return "", 0, nil
	}
}

func (s *InitialiserSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)
	s.calledCmds = []string{}
	s.PatchValue(&manager.RunCommandWithRetry, getMockRunCommandWithRetry(&s.calledCmds))
	s.PatchValue(&configureZFS, func() {})
	s.PatchValue(&configureLXDBridge, func() error { return nil })
}

func (s *InitialiserSuite) TestLTSSeriesPackages(c *gc.C) {
	// Momentarily, the only series with a dedicated cloud archive is precise,
	// which we will use for the following test:
	paccmder, err := commands.NewPackageCommander("trusty")
	c.Assert(err, jc.ErrorIsNil)

	s.PatchValue(&series.HostSeries, func() string { return "trusty" })
	container := NewContainerInitialiser("trusty")

	err = container.Initialise()
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(s.calledCmds, gc.DeepEquals, []string{
		paccmder.InstallCmd("--target-release", "trusty-backports", "lxd"),
	})
}

func (s *InitialiserSuite) TestNoSeriesPackages(c *gc.C) {
	// Here we want to test for any other series whilst avoiding the
	// possibility of hitting a cloud archive-requiring release.
	// As such, we simply pass an empty series.
	paccmder, err := commands.NewPackageCommander("xenial")
	c.Assert(err, jc.ErrorIsNil)

	container := NewContainerInitialiser("")

	err = container.Initialise()
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(s.calledCmds, gc.DeepEquals, []string{
		paccmder.InstallCmd("lxd"),
	})
}

func (s *InitialiserSuite) TestEditLXDBridgeFile(c *gc.C) {
	input := `# WARNING: Don't modify this file by hand, it is generated by debconf!
# To update those values, please run "dpkg-reconfigure lxd"

# Whether to setup a new bridge
USE_LXD_BRIDGE="true"
EXISTING_BRIDGE=""

# Bridge name
LXD_BRIDGE="lxdbr0"

# dnsmasq configuration path
LXD_CONFILE=""

# dnsmasq domain
LXD_DOMAIN="lxd"

# IPv4
LXD_IPV4_ADDR="10.0.4.1"
LXD_IPV4_NETMASK="255.255.255.0"
LXD_IPV4_NETWORK="10.0.4.1/24"
LXD_IPV4_DHCP_RANGE="10.0.4.2,10.0.4.100"
LXD_IPV4_DHCP_MAX="50"
LXD_IPV4_NAT="true"

# IPv6
LXD_IPV6_ADDR="2001:470:b2b5:9999::1"
LXD_IPV6_MASK="64"
LXD_IPV6_NETWORK="2001:470:b2b5:9999::1/64"
LXD_IPV6_NAT="true"

# Proxy server
LXD_IPV6_PROXY="true"
`
	expected := `# WARNING: Don't modify this file by hand, it is generated by debconf!
# To update those values, please run "dpkg-reconfigure lxd"

# Whether to setup a new bridge
USE_LXD_BRIDGE="true"
EXISTING_BRIDGE=""

# Bridge name
LXD_BRIDGE="lxdbr0"

# dnsmasq configuration path
LXD_CONFILE=""

# dnsmasq domain
LXD_DOMAIN="lxd"

# IPv4
LXD_IPV4_ADDR="10.0.19.1"
LXD_IPV4_NETMASK="255.255.255.0"
LXD_IPV4_NETWORK="10.0.19.1/24"
LXD_IPV4_DHCP_RANGE="10.0.19.2,10.0.19.254"
LXD_IPV4_DHCP_MAX="253"
LXD_IPV4_NAT="true"

# IPv6
LXD_IPV6_ADDR="2001:470:b2b5:9999::1"
LXD_IPV6_MASK="64"
LXD_IPV6_NETWORK="2001:470:b2b5:9999::1/64"
LXD_IPV6_NAT="true"

# Proxy server
LXD_IPV6_PROXY="false"

`
	result := editLXDBridgeFile(input, "19")
	c.Assert(result, jc.DeepEquals, expected)
}

func (s *InitialiserSuite) TestDetectSubnet(c *gc.C) {
	input := `1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 1c:6f:65:d5:56:98 brd ff:ff:ff:ff:ff:ff
    inet 192.168.0.69/24 brd 192.168.0.255 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fd5d:e5bb:c5f9::c0c/128 scope global dynamic 
       valid_lft 83178sec preferred_lft 83178sec
    inet6 fd5d:e5bb:c5f9:0:1e6f:65ff:fed5:5698/64 scope global noprefixroute dynamic 
       valid_lft 6967sec preferred_lft 1567sec
    inet6 fe80::1e6f:65ff:fed5:5698/64 scope link 
       valid_lft forever preferred_lft forever
3: virbr0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 52:54:00:e4:70:2f brd ff:ff:ff:ff:ff:ff
    inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
       valid_lft forever preferred_lft forever
4: virbr0-nic: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast master virbr0 state DOWN group default qlen 500
    link/ether 52:54:00:e4:70:2f brd ff:ff:ff:ff:ff:ff
5: virbr1: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 52:54:00:fe:04:e6 brd ff:ff:ff:ff:ff:ff
    inet 192.168.100.1/24 brd 192.168.100.255 scope global virbr1
       valid_lft forever preferred_lft forever
6: virbr1-nic: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast master virbr1 state DOWN group default qlen 500
    link/ether 52:54:00:fe:04:e6 brd ff:ff:ff:ff:ff:ff
7: lxcbr0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether fe:d3:9d:e4:ba:90 brd ff:ff:ff:ff:ff:ff
    inet 10.0.3.1/24 scope global lxcbr0
       valid_lft forever preferred_lft forever
    inet6 fe80::a00f:35ff:fe81:f7ed/64 scope link 
       valid_lft forever preferred_lft forever
25: vethOG10XO@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master lxcbr0 state UP group default qlen 1000
    link/ether fe:d3:9d:e4:ba:90 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::fcd3:9dff:fee4:ba90/64 scope link 
       valid_lft forever preferred_lft forever
37: vnet0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master virbr0 state UNKNOWN group default qlen 500
    link/ether fe:54:00:6e:2d:7d brd ff:ff:ff:ff:ff:ff
    inet6 fe80::fc54:ff:fe6e:2d7d/64 scope link 
       valid_lft forever preferred_lft forever
38: vnet1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master virbr0 state UNKNOWN group default qlen 500
    link/ether fe:54:00:3e:80:18 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::fc54:ff:fe3e:8018/64 scope link 
       valid_lft forever preferred_lft forever
39: vnet2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master virbr0 state UNKNOWN group default qlen 500
    link/ether fe:54:00:ee:c7:95 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::fc54:ff:feee:c795/64 scope link 
       valid_lft forever preferred_lft forever
40: vnet3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master virbr0 state UNKNOWN group default qlen 500
    link/ether fe:54:00:30:92:16 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::fc54:ff:fe30:9216/64 scope link 
       valid_lft forever preferred_lft forever
`

	result, err := detectSubnet(input)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, jc.DeepEquals, "4")
}

func (s *InitialiserSuite) TestDetectSubnetLocal(c *gc.C) {
	output, err := exec.Command("ip", "addr", "show").CombinedOutput()
	c.Assert(err, jc.ErrorIsNil)

	subnet, err := detectSubnet(string(output))
	c.Assert(err, jc.ErrorIsNil)

	subnetInt, err := strconv.Atoi(subnet)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(subnetInt, jc.GreaterThan, 0)
	c.Assert(subnetInt, jc.LessThan, 255)

}
