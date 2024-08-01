# Network Optimization Library for Go

This Go library provides a set of functions to optimize network performance by configuring various system settings. It includes functionalities for Receive Side Scaling (RSS), kernel tuning, Network Interface Card (NIC) configuration, and IRQ (Interrupt Request) balancing.

## Features

- **RSS (Receive Side Scaling):** Configure RSS queues on a given NIC.
- **Kernel Tuning:** Adjust kernel parameters for improved network performance.
- **NIC Configuration:** Set various NIC settings.
- **IRQ Balancing:** Configure IRQ balancing to optimize interrupt handling.

## Installation 
go get github.com/readytodiss/networkopt

## Usage
Here is a basic example of how to use the networkopt library to perform network optimizations:


    
    import (
        "fmt"
        "log"
        "github.com/readytodiss/networkopt"
    )
    
    func main() {
        // Configure RSS
        if err := networkopt.SetRSS("eth0", 4); err != nil {
            log.Fatalf("Failed to set RSS: %v", err)
        }
        
        // Tune kernel parameters
        if err := networkopt.SetKernelTuning("net.core.rmem_max", "16777216"); err != nil {
            log.Fatalf("Failed to tune kernel parameters: %v", err)
        }
        
        // Configure NIC settings
        if err := networkopt.ConfigureNIC("eth0", "tx", "on"); err != nil {
            log.Fatalf("Failed to configure NIC: %v", err)
        }
        
        // Set IRQ balance
        if err := networkopt.SetIRQBalance("eth0", "0-3"); err != nil {
            log.Fatalf("Failed to set IRQ balance: %v", err)
        }
        
        fmt.Println("Network optimizations applied successfully.")
    }
