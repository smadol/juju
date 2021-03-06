// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package unit_test

import (
	"testing"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	coretesting "github.com/juju/juju/testing"
)

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

type ImportSuite struct{}

var _ = gc.Suite(&ImportSuite{})

func (*ImportSuite) TestImports(c *gc.C) {
	found := coretesting.FindJujuCoreImports(c, "github.com/juju/juju/cmd/k8sagent/unit")
	// TODO: review if there are any expected imports!
	c.Assert(found, jc.SameContents, []string{
		"agent",
		"agent/tools",
		"api",
		"api/agent",
		"api/authentication",
		"api/base",
		"api/block",
		"api/caasoperator",
		"api/common",
		"api/common/cloudspec",
		"api/controller",
		"api/instancepoller",
		"api/keyupdater",
		"api/leadership",
		"api/logger",
		"api/logsender",
		"api/machiner",
		"api/migrationflag",
		"api/migrationminion",
		"api/modelmanager",
		"api/proxyupdater",
		"api/reboot",
		"api/retrystrategy",
		"api/unitassigner",
		"api/uniter",
		"api/upgrader",
		"api/usermanager",
		"api/watcher",
		"apiserver/common",
		"apiserver/facade",
		"apiserver/facades/agent/metricsender",
		"apiserver/params",
		"caas",
		"caas/kubernetes/clientconfig",
		"caas/kubernetes/provider",
		"caas/kubernetes/provider/specs",
		"caas/specs",
		"charmstore",
		"cloud",
		"cloudconfig",
		"cloudconfig/cloudinit",
		"cloudconfig/instancecfg",
		"cloudconfig/podcfg",
		"cmd",
		"cmd/juju/common",
		"cmd/juju/interact",
		"cmd/jujud/agent/addons",
		"cmd/jujud/agent/agentconf",
		"cmd/jujud/agent/engine",
		"cmd/jujud/agent/errors",
		"cmd/jujud/util",
		"cmd/modelcmd",
		"cmd/output",
		"controller",
		"core/actions",
		"core/annotations",
		"core/application",
		"core/cache",
		"core/constraints",
		"core/crossmodel",
		"core/devices",
		"core/firewall",
		"core/globalclock",
		"core/instance",
		"core/leadership",
		"core/lease",
		"core/life",
		"core/lxdprofile",
		"core/machinelock",
		"core/migration",
		"core/model",
		"core/multiwatcher",
		"core/network",
		"core/paths",
		"core/paths/transientfile",
		"core/permission",
		"core/presence",
		"core/quota",
		"core/raftlease",
		"core/relation",
		"core/resources",
		"core/series",
		"core/settings",
		"core/snap",
		"core/status",
		"core/watcher",
		"downloader",
		"environs",
		"environs/bootstrap",
		"environs/config",
		"environs/context",
		"environs/filestorage",
		"environs/gui",
		"environs/imagemetadata",
		"environs/instances",
		"environs/simplestreams",
		"environs/storage",
		"environs/sync",
		"environs/tags",
		"environs/tools",
		"environs/utils",
		"feature",
		"juju",
		"juju/keys",
		"juju/names",
		"juju/osenv",
		"juju/sockets",
		"jujuclient",
		"logfwd",
		"logfwd/syslog",
		"mongo",
		"mongo/utils",
		"network",
		"network/debinterfaces",
		"network/netplan",
		"packaging",
		"packaging/dependency",
		"payload",
		"pki",
		"provider/lxd/lxdnames",
		"resource",
		"rpc",
		"rpc/jsoncodec",
		"service",
		"service/common",
		"service/snap",
		"service/systemd",
		"service/upstart",
		"service/windows",
		"state",
		"state/bakerystorage",
		"state/binarystorage",
		"state/cloudimagemetadata",
		"state/globalclock",
		"state/imagestorage",
		"state/migrations",
		"state/raftlease",
		"state/stateenvirons",
		"state/storage",
		"state/upgrade",
		"state/watcher",
		"storage",
		"storage/poolmanager",
		"storage/provider",
		"tools",
		"utils/proxy",
		"utils/scriptrunner",
		"version",
		"worker",
		"worker/agent",
		"worker/apiaddressupdater",
		"worker/apicaller",
		"worker/apiconfigwatcher",
		"worker/common/charmrunner",
		"worker/common/reboot",
		"worker/fortress",
		"worker/introspection",
		"worker/introspection/pprof",
		"worker/leadership",
		"worker/logger",
		"worker/logsender",
		"worker/migrationflag",
		"worker/migrationminion",
		"worker/proxyupdater",
		"worker/retrystrategy",
		"worker/uniter",
		"worker/uniter/actions",
		"worker/uniter/charm",
		"worker/uniter/container",
		"worker/uniter/hook",
		"worker/uniter/leadership",
		"worker/uniter/operation",
		"worker/uniter/relation",
		"worker/uniter/remotestate",
		"worker/uniter/resolver",
		"worker/uniter/runcommands",
		"worker/uniter/runner",
		"worker/uniter/runner/context",
		"worker/uniter/runner/debug",
		"worker/uniter/runner/jujuc",
		"worker/uniter/storage",
		"worker/uniter/upgradeseries",
		"worker/uniter/verifycharmprofile",
		"wrench",
	})
}
