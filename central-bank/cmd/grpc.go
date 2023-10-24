/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/JonatasMSantos/codepix/application/grpc"
	"github.com/JonatasMSantos/codepix/infra/db"

	"github.com/spf13/cobra"
)

var PortNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start a gRPC Server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("grpc called")
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, PortNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&PortNumber, "port", "p", 50051, "gRPC Server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
