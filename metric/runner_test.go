// Copyright 2015 bs authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"net/http"
	"os"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/fsouza/go-dockerclient/testing"
	"gopkg.in/check.v1"
)

type bogusContainer struct {
	config docker.Config
	state  docker.State
}

func (s *S) TestRunner(c *check.C) {
	bogusContainers := []bogusContainer{
		{config: docker.Config{Image: "tsuru/python", Env: []string{"HOME=/", "TSURU_APPNAME=someapp", "TSURU_PROCESSNAME=myprocess"}}, state: docker.State{Running: true}},
		{config: docker.Config{Image: "tsuru/python", Env: []string{"HOME=/", "TSURU_APPNAME=someapp"}}, state: docker.State{Running: false, ExitCode: -1}},
	}
	dockerServer, conts := s.startDockerServer(bogusContainers, nil, c)
	defer dockerServer.Stop()
	dockerServer.PrepareStats(conts[0].ID, func(id string) docker.Stats {
		s := docker.Stats{}
		s.PreCPUStats.CPUUsage.TotalUsage = 50
		s.PreCPUStats.SystemCPUUsage = 20
		s.CPUStats.CPUUsage.TotalUsage = 100
		s.CPUStats.SystemCPUUsage = 40
		s.CPUStats.CPUUsage.PercpuUsage = []uint64{100}
		return s
	})
	os.Setenv("METRICS_BACKEND", "fake")
	defer os.Unsetenv("METRICS_BACKEND")
	r := NewRunner(dockerServer.URL(), time.Second)
	err := r.Start()
	c.Assert(err, check.IsNil)
	r.Stop()
	var cpuStat *fakeStat
	for i, stat := range fakeStatter.stats {
		if stat.key == "cpu_max" {
			cpuStat = &fakeStatter.stats[i]
		}
	}
	c.Assert(cpuStat, check.NotNil)
	expected := &fakeStat{
		app:      "someapp",
		hostname: conts[0].ID[:12],
		process:  "myprocess",
		key:      "cpu_max",
		value:    float(250),
	}
	c.Assert(cpuStat, check.DeepEquals, expected)
}

func (s *S) startDockerServer(containers []bogusContainer, hook func(*http.Request), c *check.C) (*testing.DockerServer, []docker.Container) {
	server, err := testing.NewServer("127.0.0.1:0", nil, hook)
	c.Assert(err, check.IsNil)
	client, err := docker.NewClient(server.URL())
	c.Assert(err, check.IsNil)
	createdContainers := make([]docker.Container, len(containers))
	for i, bogus := range containers {
		pullOpts := docker.PullImageOptions{Repository: bogus.config.Image}
		err = client.PullImage(pullOpts, docker.AuthConfiguration{})
		c.Assert(err, check.IsNil)
		createOpts := docker.CreateContainerOptions{Config: &bogus.config}
		container, err := client.CreateContainer(createOpts)
		c.Assert(err, check.IsNil)
		err = server.MutateContainer(container.ID, bogus.state)
		c.Assert(err, check.IsNil)
		createdContainers[i] = *container
	}
	return server, createdContainers
}
