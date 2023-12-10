package cmd

import (
	"api/api"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type addStorageConfig struct {
	fmsAddr     string
	storageAddr string
}

var addStorageCfg addStorageConfig

// addStorageCmd represents the addStorage command
var addStorageCmd = &cobra.Command{
	Use:   "add-storage",
	Short: "adds storage to FMS",
	RunE: func(cmd *cobra.Command, args []string) error {
		dialCtx, cancel := context.WithTimeout(cmd.Context(), time.Second)
		defer cancel()

		dial, err := grpc.DialContext(dialCtx, addStorageCfg.fmsAddr,
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return fmt.Errorf("cannot connect to fms: %v", err)
		}

		addStorageCtx, cancel := context.WithTimeout(cmd.Context(), time.Second)
		defer cancel()

		fms := api.NewFileManagementAdminServiceClient(dial)
		resp, err := fms.AddStorageV1(addStorageCtx, &api.AddStorageV1Request{Addr: addStorageCfg.storageAddr})
		if err != nil {
			return fmt.Errorf("cannot add storage to FMS: %v", err)
		}

		cmd.Printf("Successfully added storage, id=%d\n", resp.GetId())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addStorageCmd)

	addStorageCmd.Flags().StringVarP(&addStorageCfg.fmsAddr,
		"fms", "f", "localhost:8080",
		"FMS address in the form of <host>:<port>",
	)
	addStorageCmd.Flags().StringVarP(&addStorageCfg.storageAddr,
		"storage", "s", "",
		"Storage address in the form <host>:<port>. Must be accessible from FMS!",
	)
}
