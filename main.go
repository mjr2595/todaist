package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mjr2595/todaist/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todaist [prompt]",
	Short: "Generate Todoist projects from AI-powered prompts",
	Long: `Todaist uses Gemini AI to generate todo lists from natural language prompts,
creates CSV files, and imports them into new Todoist projects.

Example:
  todaist "plan a wedding"`,
	Args: cobra.ExactArgs(1),
	Run:  runTodaist,
}

func runTodaist(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	prompt := args[0]

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Initialize clients
	clients, err := config.NewClients(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize clients: %v", err)
	}
	defer clients.Close()

	fmt.Printf("🎯 Processing prompt: %s\n", prompt)
	fmt.Println("✅ Configuration loaded successfully")
	fmt.Println("✅ Gemini client initialized")
	
	// TODO: Implement the full workflow in subsequent stages
	fmt.Println("🚧 Full implementation coming in next stages...")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
