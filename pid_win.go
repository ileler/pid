// +build windows

package pid

import (
    "strconv"
    "os/exec"
    "golang.org/x/sys/windows"
)

func PID() *string {
    _pid := windows.GetCurrentProcessId()
    pid := strconv.Itoa(int(_pid))
    return &pid
}

func PIDExists(pid string) bool {
    out, err := exec.Command("cmd", "/c", "tasklist", "/FI", "PID eq " + pid, "2>&1", "|", "findstr", pid).Output()
    if err != nil {
        return false
    }
    return out != nil
}

/**
import (
    "unsafe"
    "strconv"
    "syscall"
    "os/exec"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
    dwSize              ulong
    cntUsage            ulong
    th32ProcessID       ulong
    th32DefaultHeapID   ulong_ptr
    th32ModuleID        ulong
    cntThreads          ulong
    th32ParentProcessID ulong
    pcPriClassBase      ulong
    dwFlags             ulong
    szExeFile           [260]byte
}

func PID() *string {
    kernel32 := syscall.NewLazyDLL("kernel32.dll")
    CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
    pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
    if int(pHandle) == -1 {
        return nil
    }
    Process32Next := kernel32.NewProc("Process32Next")
    var pid *string
    for {
        var proc PROCESSENTRY32
        proc.dwSize = ulong(unsafe.Sizeof(proc))
        if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
            _p := strconv.Itoa(int(proc.th32ProcessID))
            pid = &_p
        } else {
            break
        }
    }
    CloseHandle := kernel32.NewProc("CloseHandle")
    _, _, _ = CloseHandle.Call(pHandle)
    return pid
}
*/
