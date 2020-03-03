package pid

import (
    "os"
    "fmt"
    "strings"
    "runtime"
    "io/ioutil"
    "path/filepath"
)

func NewPIDFile() (*string, *string, error) {
    return NewPIDFileWithName(".pid")
}

func NewPIDFileWithName(name string) (*string, *string, error) {
    if name == "" {
        return NewPIDFile()
    }
    var path string
    if runtime.GOOS == "windows" {
        path = Getenv("ProgramData", "C:\\ProgramData")
    } else {
        path = "/opt/"
    }
    return NewPIDFileWithPath(path + "/" + name)
}

func NewPIDFileWithPath(path string) (*string, *string, error) {
    if err := pfExists(path); err != nil {
        return nil, nil, err
    }

    if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
        return nil, nil, err
    }
    pid := PID()
    if err := ioutil.WriteFile(path, []byte(*pid), 0644); err != nil {
        return nil, nil, err
    }
    return pid, &path, nil
}

func pfExists(path string) error {
    if pidByte, err := ioutil.ReadFile(path); err == nil {
        pid := strings.TrimSpace(string(pidByte))
        if pid != "" && PIDExists(pid) {
            return fmt.Errorf("process has already started:[%s]-[%s]", pid, path)
        }
    }
    return nil
}

func Getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
