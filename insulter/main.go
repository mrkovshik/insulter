// main.go
package main

import (
	"github.com/mrkovshik/insulter/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func whoAreYouToday(c proto.InsulterClient) (string, error) {
	ctx := context.TODO()
	Swear, err := c.Insult(ctx, &proto.Value_IN{Value:"Kolya" })
	if err != nil {
		return "", err
	}
	return Swear.Value, nil
}

func main() {
	conn, err := grpc.Dial("172.20.166.46:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No insult")
	}
	defer conn.Close()

	client := proto.NewInsulterClient(conn)
	answer, err := whoAreYouToday(client)
	fmt.Printf(answer)
}
