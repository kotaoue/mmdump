package mmdump

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Dump(input, output string) error {
	if err := extract(input); err != nil {
		return err
	}

	return exec.Command("mmdc", "--input", mmdName(input), "--output", output).Run()
}

func extract(name string) error {
	input, err := os.Open(name)
	if err != nil {
		return nil
	}
	defer input.Close()

	output, err := os.Create(mmdName(name))
	if err != nil {
		return err
	}
	defer output.Close()

	r := bufio.NewReader(input)
	w := bufio.NewWriter(output)

	mm := false
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}

		if mm && strings.HasPrefix(line, "```") {
			mm = false
		}

		if mm {
			_, err = w.WriteString(line)
			if err != nil {
				return err
			}
		}

		if !mm && strings.HasPrefix(line, "```mermaid") {
			mm = true
		}

	}
	return w.Flush()
}

func mmdName(md string) string {
	return strings.ReplaceAll(md, ".md", ".mmd")
}
