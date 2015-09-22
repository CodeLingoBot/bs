// Copyright 2015 bs authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import "github.com/fsouza/go-dockerclient"

func statsToMetricsMap(s *docker.Stats) (map[string]float, error) {
	previousCPU := s.PreCPUStats.CPUUsage.TotalUsage
	previousSystem := s.PreCPUStats.SystemCPUUsage
	cpuPercent := calculateCPUPercent(previousCPU, previousSystem, s)
	memPercent := float64(s.MemoryStats.Usage) / float64(s.MemoryStats.Limit) * 100.0
	stats := map[string]float{
		"cpu_max":     float(cpuPercent),
		"mem_max":     float(s.MemoryStats.Usage),
		"mem_pct_max": float(memPercent),
	}
	return stats, nil
}

func calculateCPUPercent(previousCPU, previousSystem uint64, s *docker.Stats) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(s.CPUStats.CPUUsage.TotalUsage - previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(s.CPUStats.SystemCPUUsage - previousSystem)
	)
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(s.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}
