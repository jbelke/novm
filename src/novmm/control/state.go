package control

import (
    "novmm/machine"
    "novmm/platform"
)

//
// State.
//

type State struct {
    // Our device state.
    // Note that we only encode state associated with
    // specific devices. The model type is a generic wrapped
    // around devices which *may* encode additional state,
    // but all the state associated with the model should be
    // regenerated on startup.
    Devices []machine.DeviceInfo `json:"devices,omitempty"`

    // Our vcpu state.
    // Similarly, this should encode all the state associated
    // with the underlying VM. If we have other internal platform
    // devices (such as APICs or PITs) then these should be somehow
    // encoded as generic devices.
    Vcpus []platform.VcpuInfo `json:"vcpus,omitempty"`
}

func (control *Control) State(nop *Nop, res *State) error {

    // Pause all vcpus.
    vcpus := control.vm.Vcpus()
    for _, vcpu := range vcpus {
        err := vcpu.Pause(false)
        if err != nil {
            return err
        }
        defer vcpu.Unpause(false)
    }

    // Grab our vcpu states.
    res.Vcpus = control.vm.VcpuInfo()

    // Grab our devices.
    res.Devices = control.model.DeviceInfo()

    // That's it.
    // We let the serialization handle the rest.
    return nil
}
