package client

import (
	"context"
	"log"

	pb "github.com/michaelhendraw/stockbit_microservice/grpc/proto"
	"google.golang.org/grpc"
)

// GetClient for GRPC
func GetClient(o *Options) (*Client, error) {

	var (
		conn *grpc.ClientConn
		err  error
	)

	// Open GRPC connection
	if o.WithInterceptor == true {
		conn, err = grpc.Dial(o.Address, grpc.WithInsecure(), withClientInterceptor())
	} else {
		conn, err = grpc.Dial(o.Address, grpc.WithInsecure())
	}

	if err != nil {
		log.Println("func GetClient", err)
		return nil, err
	}

	newClient := &Client{
		conn:     conn,
		stockbit: pb.NewStockbitClient(conn),
	}

	return newClient, nil
}

// Close GRPC connection
func (c *Client) Close() {
	c.conn.Close()
}

// Search func
func (c *Client) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return c.stockbit.Search(ctx, req)
}
