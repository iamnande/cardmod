package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iamnande/cardmod/pkg/api/calculationv1"
)

var (
	host  = flag.String("host", "localhost:9000", "the address to connect to")
	magic = flag.String("magic", "", "name of the magic needed")
	count = flag.Int("count", 100, "number of magic needed")
)

func main() {
	flag.Parse()

	// client: open connection
	conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	// client: initialize client
	client := calculationv1.NewCalculationAPIClient(conn)

	// client: perform the calculation request
	ctx := context.Background()
	res, err := client.Calculate(ctx, &calculationv1.CalculateRequest{Name: *magic, Count: int32(*count)})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		os.Exit(1)
	}

	// client: write results to stdout
	fmt.Printf("%d %s cards needed for %d %s magics\n", res.GetCount(), res.GetName(), *count, *magic)

}
