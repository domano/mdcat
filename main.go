package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mdcat [file...]",
		Short: "Pretty print markdown files in the terminal",
		RunE:  run,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	// Setup renderer with iTerm2-optimized style
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(int(getTerminalWidth())),
	)
	if err != nil {
		return fmt.Errorf("failed to create renderer: %w", err)
	}

	// If no files specified, read from stdin
	if len(args) == 0 {
		return renderMarkdown(r, os.Stdin)
	}

	// Process each file
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", filename, err)
		}
		defer file.Close()

		if err := renderMarkdown(r, file); err != nil {
			return fmt.Errorf("failed to render %s: %w", filename, err)
		}
	}

	return nil
}

func renderMarkdown(r *glamour.TermRenderer, input io.Reader) error {
	content, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	rendered, err := r.Render(string(content))
	if err != nil {
		return err
	}

	fmt.Print(rendered)
	return nil
}

func getTerminalWidth() float64 {
	if ws, err := getWinsize(); err == nil {
		return float64(ws.Col)
	}
	return 80 // fallback width
}
