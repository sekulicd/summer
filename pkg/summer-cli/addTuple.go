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
var addTupleCmd = &cobra.Command{
	Use:   "addTuple",
	Short: "Calculates sum of two numbers",
	Run: func(addTupleCmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		num1, err := addTupleCmd.Flags().GetInt("num1")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		num2, err := addTupleCmd.Flags().GetInt("num2")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := summer.NewSummerClient(conn)
		req := summer.RequestTuple{
			A:                    int32(num1),
			B:                    int32(num2),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		response, err := client.AddTuple(ctx, &req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(response.Sum)
	},
}

func init() {
	rootCmd.AddCommand(addTupleCmd)
	addTupleCmd.Flags().Int("num1", 0,  "First num")
	addTupleCmd.Flags().Int("num2", 0,  "Second num")
}
