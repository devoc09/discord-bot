package internal

import (
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

type SystemInfo interface {
	MemUtil(interbal time.Duration, percpu bool) *mem.VirtualMemoryStat
	CpuUtil() []float64
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
