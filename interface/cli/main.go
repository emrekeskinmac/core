package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mesg-foundation/core/commands"
	"github.com/mesg-foundation/core/commands/provider"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/mesg-foundation/core/utils/clierrors"
	"github.com/mesg-foundation/core/utils/pretty"
	"github.com/mesg-foundation/core/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	cfg, err := config.Global()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", pretty.Fail(clierrors.ErrorMessage(err)))
		os.Exit(1)
	}

	// Keep alive prevents Docker network to drop TCP idle connections after 15 minutes.
	// See: https://forum.mesg.com/t/solution-summary-for-docker-dropping-connections-after-15-min/246
	dialKeepaliveOpt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: 5 * time.Minute, // 5 minutes because it's the minimun time of gRPC enforcement policy.
	})
	connection, err := grpc.Dial(cfg.Client.Address, dialKeepaliveOpt, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", pretty.Fail(clierrors.ErrorMessage(err)))
		os.Exit(1)
	}

	c, err := container.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", pretty.Fail(clierrors.ErrorMessage(err)))
		os.Exit(1)
	}

	p := provider.New(coreapi.NewCoreClient(connection), daemon.NewContainerDaemon(cfg, c))
	cmd := commands.Build(p)
	cmd.Version = version.Version
	cmd.Short = cmd.Short + " " + version.Version

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", pretty.Fail(clierrors.ErrorMessage(err)))
		os.Exit(1)
	}
}
