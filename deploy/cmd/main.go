// Command deploy deploys omniagent-panel to Kubernetes.
//
// Usage:
//
//	# Deploy to Kubernetes
//	export KUBECONFIG=~/.kube/config
//	go run ./deploy/cmd
//
//	# Deploy to specific namespace
//	export KUBERNETES_NAMESPACE=production
//	go run ./deploy/cmd
//
//	# Destroy deployment
//	go run ./deploy/cmd -destroy
//
//	# Preview changes
//	export AGENTKIT_DEPLOY_DRY_RUN=true
//	go run ./deploy/cmd
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/plexusone/agentkit/deploy"

	// Import Kubernetes provider (auto-registers via init())
	_ "github.com/plexusone/agentkit-k8s-pulumi/deploy/providers/kubernetes"
)

func main() {
	destroy := flag.Bool("destroy", false, "Destroy the deployment")
	configPath := flag.String("config", "deploy/deploy.yaml", "Path to deployment config")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nCancelling...")
		cancel()
	}()

	// Load deployment configuration
	cfg, err := deploy.LoadDeployConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	// Get provider (respects AGENTKIT_DEPLOY_PROVIDER env var)
	provider, err := deploy.GetProvider(cfg)
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}
	defer func() { _ = provider.Close() }()

	fmt.Printf("Provider: %s\n", provider.Name())
	fmt.Printf("Stack:    %s\n", cfg.Stack.Name)

	if *destroy {
		// Destroy deployment
		fmt.Println("\nDestroying deployment...")
		if err := provider.Destroy(ctx, cfg.Stack.Name); err != nil {
			log.Fatalf("Destroy failed: %v", err)
		}
		fmt.Println("Deployment destroyed successfully.")
		return
	}

	// Deploy
	fmt.Println("\nDeploying...")
	status, err := provider.Deploy(ctx, cfg)
	if err != nil {
		log.Fatalf("Deployment failed: %v", err)
	}

	// Print results
	fmt.Printf("\nDeployment %s!\n", status.State)
	fmt.Printf("Duration: %s\n", status.Duration)
	fmt.Println("\nOutputs:")
	for k, v := range status.Outputs {
		fmt.Printf("  %s: %s\n", k, v)
	}

	if len(status.Resources) > 0 {
		fmt.Println("\nResources:")
		for _, r := range status.Resources {
			fmt.Printf("  %s: %s (%s)\n", r.Type, r.Name, r.State)
		}
	}
}
