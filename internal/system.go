package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/devoc09/discord-bot/internal/thermaldevise"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

type SystemInfo interface {
	MemUtil(interbal time.Duration, percpu bool) *mem.VirtualMemoryStat
	CpuUtil() []float64
    TempUtil()
}

type System struct{}

func (s *System) MemUtil() *mem.VirtualMemoryStat {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return nil
	}
	return vmem
}

func (s *System) CpuUtil(interbal time.Duration, percpu bool) []float64 {
	c, err := cpu.Percent(interbal, percpu)
	if err != nil {
		return nil
	}
	return c
}

func (s *System) TempUtil() ([]thermaldevise.Devise, error) {
    thermalPaths, err := filepath.Glob("/sys/class/thermal/thermal_zone*")
    if err != nil {
        return nil, fmt.Errorf("/sys/class/thermal/thermal_zone* not found: %w", err)
    }
    var Devises []thermaldevise.Devise
    for _, v := range thermalPaths {
        name := filepath.Base(v)
        bytes, err := os.ReadFile(filepath.Join(v, "temp"))
        if err != nil {
            continue
        }
        temp, err := strconv.ParseFloat(strings.TrimSpace(string(bytes)), 64)
        if err != nil {
            return nil, fmt.Errorf("strconv.Atoi Error: %w", err)
        }
        devise := thermaldevise.Devise{
            Name: name,
            Temp: temp,
        }
        Devises = append(Devises, devise)
    }
    return Devises, nil
}
