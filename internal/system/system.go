/*
	Autor Andrey Semochkin
*/

package system

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

//getUsedCpuPercent - get a slice from cpu and percentage of their load
func getUsedCpuPercent() ([]float64, error) {

	memoryPercent, err := cpu.Percent(time.Second, false)

	if err != nil {

		return nil, err

	}

	ret := make([]float64, len(memoryPercent))

	for i, v := range memoryPercent {

		ret[i] = v

	}

	return ret, nil

}

//getModelCpu - get slice from cpu models
func getModelCpu() ([]string, error) {

	cpuInformation, err := cpu.Info()

	if err != nil {

		return nil, err

	}

	ret := make([]string, len(cpuInformation))

	for i, v := range cpuInformation {

		ret[i] = v.ModelName

	}

	return ret, nil

}

//getIdleCpuPercent - get a slice per cpu from idle cpu (free)
func getIdleCpuPercent() ([]float64, error) {

	cpuPercent, err := getUsedCpuPercent()

	if err != nil {

		return nil, err

	}

	ret := make([]float64, len(cpuPercent))

	for i, v := range cpuPercent {

		ret[i] = 100 - v

	}

	return ret, nil

}

//getUsedMemoryPercent - gets the percentage of used memory
func getUsedMemoryPercent() (float64, error) {

	memInfo, err := mem.VirtualMemory()

	if err != nil {

		return 0, err

	}

	return memInfo.UsedPercent, nil

}

//getIdleMemoryPercent - get percentage of free memory
func getIdleMemoryPercent() (float64, error) {

	usedMemoryPercent, err := getUsedMemoryPercent()

	if err != nil {

		return 0, err

	}

	return 100 - usedMemoryPercent, nil

}
