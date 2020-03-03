// +build !windows

package pid

import (
    "os"
    "fmt"
    "path/filepath"
)

func PID() *string {
    pid := fmt.Sprintf("%d", os.Getpid())
    return &pid
}

func PIDExists(pid string) bool {
    if _, err := os.Stat(filepath.Join("/proc", pid)); err == nil {
        return true
    }
    return false
}
