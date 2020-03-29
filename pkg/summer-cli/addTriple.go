package grpcclient

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
	summer "summer/pkg/grpc-schema"
	"time"
)

// addCmd represents the add command
var addTripleCmd = &cobra.Command{
	Use:   "addTriple",
	Short: "Calculates sum of three numbers",
	Run: func(addTripleCmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:30541", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		num1, err := addTripleCmd.Flags().GetInt("num1")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		num2, err := addTripleCmd.Flags().GetInt("num2")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		num3, err := addTripleCmd.Flags().GetInt("num3")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := summer.NewSummerClient(conn)
		req := summer.RequestTriple{
			A:                    int32(num1),
			B:                    int32(num2),
			C:                    int32(num3),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		response, err := client.AddTriple(ctx, &req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(response.Sum)
	},
}

func init() {
	rootCmd.AddCommand(addTripleCmd)
	addTripleCmd.Flags().Int("num1", 0,  "First num")
	addTripleCmd.Flags().Int("num2", 0,  "Second num")
	addTripleCmd.Flags().Int("num3", 0,  "Third num")
}
