package cmd

import (
	"context"
	pb "github.com/FeelDat/go-gopher-grpc/pkg/gopher"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect %s", err)
		}
		defer conn.Close()

		client := pb.NewGopherClient(conn)

		var name string

		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.GetGopher(ctx, &pb.GopherRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet %v", err)
		}
		log.Printf("URL: %s", r.GetMessage())

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

}
