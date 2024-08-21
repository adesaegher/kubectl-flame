package profiler

import (
	"bytes"
	"github.com/VerizonMedia/kubectl-flame/agent/details"
	"github.com/VerizonMedia/kubectl-flame/agent/utils"
	"os/exec"
	"strconv"
)

const (
	phpSpyLocation        = "/app/phpspy"
	phpOutputFileName = "/tmp/php.svg"
)

type PhpProfiler struct{}

func (p *PhpProfiler) SetUp(job *details.ProfilingJob) error {
	return nil
}

func (p *PhpProfiler) Invoke(job *details.ProfilingJob) error {
	pid, err := utils.FindRootProcessId(job)
	if err != nil {
		return err
	}

	duration := strconv.Itoa(int(job.Duration.Seconds()))
	cmd := exec.Command(phpSpyLocation, "--buffer-size=40000", "--limit=50000", "-p=", pid, "-o", phpOutputFileName, "-i", duration )
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return utils.PublishFlameGraph(phpOutputFileName)
}
