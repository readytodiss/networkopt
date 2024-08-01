package networkopt

import (
    "fmt"
    "os/exec"
)

// SetRSS configures Receive Side Scaling (RSS) on the given NIC.
func SetRSS(nicName string, numQueues int) error {
    cmd := exec.Command("ethtool", "-L", nicName, "rx", fmt.Sprintf("%d", numQueues))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to set RSS: %v, output: %s", err, string(output))
    }
    return nil
}

// SetKernelTuning configures kernel parameters for network performance.
func SetKernelTuning(param string, value string) error {
    cmd := exec.Command("sysctl", fmt.Sprintf("%s=%s", param, value))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to set kernel tuning: %v, output: %s", err, string(output))
    }
    return nil
}

// ConfigureNIC configures various NIC settings.
func ConfigureNIC(nicName string, setting string, value string) error {
    cmd := exec.Command("ethtool", "-s", nicName, setting, value)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to configure NIC: %v, output: %s", err, string(output))
    }
    return nil
}

// SetIRQBalance configures IRQ balancing.
func SetIRQBalance(nicName string, cpuList string) error {
    cmd := exec.Command("irqbalance", "--cpu-list", cpuList, "--nic", nicName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to set IRQ balance: %v, output: %s", err, string(output))
    }
    return nil
}
