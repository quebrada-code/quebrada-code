package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"solution-worker/internal/domain/events"
	"syscall"
)

type SolutionSubmitHandler struct {
	outputSolutionHandler OutSolutionHandler
}

func NewSolutionSubmitHandler(outputSolutionHandler OutSolutionHandler) *SolutionSubmitHandler {
	return &SolutionSubmitHandler{
		outputSolutionHandler: outputSolutionHandler,
	}
}

func (h SolutionSubmitHandler) Handler(event events.SolutionSubmittedEvent) error {

	path := "/tmp/app/"
	err := os.RemoveAll(path)
	if err != nil {
		e, ok := err.(*os.PathError)
		if ok && e.Err != syscall.ENOENT {
			return err
		}
	}

	err = h.writeFile(path, "solution.py", event.GetSolutionCodeDecoded())
	if err != nil {
		return err
	}

	err = h.writeFile(path, "test_solution.py", event.GetTestCodeDecoded())
	if err != nil {
		return err
	}

	err = h.executeTest(path)
	if err != nil {
		return err
	}

	result, err := h.getOutputResult(path)
	if err != nil {
		return err
	}

	err = h.outputSolutionHandler.Handler(result)
	if err != nil {
		return err
	}
	return nil
}

func (h SolutionSubmitHandler) executeTest(path string) error {
	cmd := exec.Command("pytest", path, "--json-report", fmt.Sprintf("--json-report-file %s", path), "-vvv")
	fmt.Printf("Running command: %q\n", cmd.String())

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	fmt.Printf("translated phrase: %q\n", out.String())
	if err != nil {
		return err
	}
	return nil
}

func (h SolutionSubmitHandler) getOutputResult(path string) (events.OutputSolutionEvent, error) {
	jsonFile, err := os.Open(fmt.Sprintf("%s.report.json", path))
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var result events.OutputSolutionEvent
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return events.OutputSolutionEvent{}, err
	}
	return result, nil
}

func (h SolutionSubmitHandler) writeFile(path string, filename string, content string) error {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s%s", path, filename))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}
	return nil
}
