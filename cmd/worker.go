package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func workerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "worker example",
		Run: func(cmd *cobra.Command, args []string) {
			const tickInterval = 2 * time.Second

			for {
				select {
				case <-cmd.Context().Done():
					return
				case t := <-time.Tick(tickInterval):
					fmt.Println("Tick at", t)
				}
			}
		},
	}
}
