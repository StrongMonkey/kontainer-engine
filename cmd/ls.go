package cmd

import (
	"github.com/rancher/kontainer-engine/store"
	"github.com/rancher/kontainer-engine/utils"
	"github.com/urfave/cli"
)

func LsCommand() cli.Command {
	return cli.Command{
		Name:      "list",
		ShortName: "ls",
		Usage:     "list kubernetes clusters",
		Action:    lsCluster,
		Flags:     []cli.Flag{},
	}
}

func lsCluster(ctx *cli.Context) error {
	// todo: add filter support
	clusters, err := store.GetAllClusterFromStore()
	if err != nil {
		return err
	}

	writer := utils.NewTableWriter([][]string{
		{"NAME", "Name"},
		{"DRIVER", "DriverName"},
		{"VERISON", "Version"},
		{"ENDPOINT", "Endpoint"},
		{"NODE_COUNT", "NodeCount"},
	}, ctx)
	defer writer.Close()
	for _, cluster := range clusters {
		writer.Write(cluster)
	}
	return writer.Err()
}
