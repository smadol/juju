// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package maas

import (
	"fmt"
	"strings"

	"github.com/juju/gomaasapi"

	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
)

type maas2Controller interface {
	Machines(gomaasapi.MachinesArgs) ([]gomaasapi.Machine, error)
}

type maas2Instance struct {
	machine           gomaasapi.Machine
	constraintMatches gomaasapi.ConstraintMatches
	// maasController provides access to the MAAS 2.0 API.
	maasController maas2Controller
}

var _ maasInstance = (*maas2Instance)(nil)

func (mi *maas2Instance) zone() (string, error) {
	return mi.machine.Zone().Name(), nil
}

func (mi *maas2Instance) hostname() (string, error) {
	return mi.machine.Hostname(), nil
}

func (mi *maas2Instance) hardwareCharacteristics() (*instance.HardwareCharacteristics, error) {
	nodeArch := strings.Split(mi.machine.Architecture(), "/")[0]
	nodeCpuCount := uint64(mi.machine.CPUCount())
	nodeMemoryMB := uint64(mi.machine.Memory())
	// zone can't error on the maas2Instance implementaation.
	zone, _ := mi.zone()
	tags := mi.machine.Tags()
	hc := &instance.HardwareCharacteristics{
		Arch:             &nodeArch,
		CpuCores:         &nodeCpuCount,
		Mem:              &nodeMemoryMB,
		AvailabilityZone: &zone,
		Tags:             &tags,
	}
	return hc, nil
}

func (mi *maas2Instance) String() string {
	return fmt.Sprintf("%s:%s", mi.machine.Hostname(), mi.machine.SystemID())
}

func (mi *maas2Instance) Id() instance.Id {
	return instance.Id(mi.machine.SystemID())
}

func (mi *maas2Instance) Addresses() ([]network.Address, error) {
	machineAddresses := mi.machine.IPAddresses()
	addresses := make([]network.Address, len(machineAddresses))
	for i, address := range machineAddresses {
		addresses[i] = network.NewAddress(address)
	}
	return addresses, nil
}

// Status returns a juju status based on the maas instance returned
// status message.
func (mi *maas2Instance) Status() instance.InstanceStatus {
	args := gomaasapi.MachinesArgs{SystemIDs: []string{mi.machine.SystemID()}}
	machines, err := mi.maasController.Machines(args)
	if err != nil {
		logger.Errorf("obtaining machines: %v", err)
		return convertInstanceStatus("", "", mi.Id())
	}
	if len(machines) != 1 {
		logger.Warningf("1 machine was epected, got %d", len(machines))
		return convertInstanceStatus("", "", mi.Id())
	}
	machine := machines[0]
	statusName := machine.StatusName()
	statusMsg := machine.StatusMessage()
	return convertInstanceStatus(statusName, statusMsg, mi.Id())

}

// MAAS does not do firewalling so these port methods do nothing.
func (mi *maas2Instance) OpenPorts(machineId string, ports []network.PortRange) error {
	logger.Debugf("unimplemented OpenPorts() called")
	return nil
}

func (mi *maas2Instance) ClosePorts(machineId string, ports []network.PortRange) error {
	logger.Debugf("unimplemented ClosePorts() called")
	return nil
}

func (mi *maas2Instance) Ports(machineId string) ([]network.PortRange, error) {
	logger.Debugf("unimplemented Ports() called")
	return nil, nil
}
