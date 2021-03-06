package strumt_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/antham/strumt"
)

func Example_customizePromptOutput() {
	var stdout bytes.Buffer
	buf := "whatever\nyes\n"

	p := strumt.NewPromptsFromReaderAndWriter(bytes.NewBufferString(buf), &stdout)
	p.AddLinePrompter(&AreYouOkPrompt{})
	p.SetFirst("okprompt")
	p.Run()

	for {
		line, err := stdout.ReadString('\n')

		if err == io.EOF {
			break
		}

		fmt.Println(strings.TrimSpace(line))
	}
	// Output:
	// ==> Are you Ok ?
	// An error occurred : You must answer yes or no
	// ==> Are you Ok ?
	//
}

type AreYouOkPrompt struct {
}

func (a *AreYouOkPrompt) ID() string {
	return "okprompt"
}

func (a *AreYouOkPrompt) PromptString() string {
	return "Are you Ok ?"
}

func (a *AreYouOkPrompt) Parse(value string) error {
	if value == "yes" || value == "no" {
		return nil
	}

	return fmt.Errorf("You must answer yes or no")
}

func (a *AreYouOkPrompt) NextOnSuccess(value string) string {
	return ""
}

func (a *AreYouOkPrompt) NextOnError(err error) string {
	return "okprompt"
}

func (a *AreYouOkPrompt) PrintPrompt(w io.Writer, prompt string) {
	fmt.Fprintf(w, "==> %s\n", prompt)
}

func (a *AreYouOkPrompt) PrintError(w io.Writer, err error) {
	fmt.Fprintf(w, "An error occurred : %s", err)
}
