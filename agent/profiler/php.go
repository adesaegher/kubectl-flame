package profiler

import (
	"fmt"
	"bytes"
	"github.com/adesaegher/kubectl-flame/agent/details"
	"github.com/adesaegher/kubectl-flame/agent/utils"
	"os/exec"
	"strconv"
	"os"
)

const (
	phpSpyLocation = "/app/phpspy"
	phpOutputFileName = "/tmp/php.svg"
	flameGraphPHPScriptLocation = "/app/FlameGraph/flamegraph.pl"
	flameGraphPHPOutputLocation = "/tmp/flamegraph.svg"
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

	duration := strconv.Itoa(int(job.Duration.Seconds() * 1000))
	cmd := exec.Command(phpSpyLocation, "--buffer-size=40000", "--limit=50000", "-p", pid, "-o", phpOutputFileName, "-i", duration )
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(phpSpyLocation, "--buffer-size=40000", "--limit=50000", "-p", pid, "-o", phpOutputFileName, "-i", duration )
		fmt.Println(err)
		return err
	}

	err = p.generateFlameGraph()
	if err != nil {
		return fmt.Errorf("flamegraph generation failed: %s", err)
	}

	return utils.PublishFlameGraph(flameGraphPHPOutputLocation)
}

func (p *PhpProfiler) generateFlameGraph() error {
	inputFile, err := os.Open(phpOutputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(flameGraphPHPOutputLocation)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	flameGraphCmd := exec.Command(flameGraphPHPScriptLocation)
	flameGraphCmd.Stdin = inputFile
	flameGraphCmd.Stdout = outputFile

	return flameGraphCmd.Run()
}