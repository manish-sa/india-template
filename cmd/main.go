package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := cobra.Command{}

	cmd.AddCommand(
		apiCommand(),
		workerCommand(),
	)

	if err := cmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}
