package grpc_service_clients

import (
	"fmt"

	pbe "Booking/api_establishment_booking/genproto/establishment-proto"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"Booking/api_establishment_booking/internal/pkg/config"
)

type ServiceClient interface {
	EstablishmentService() pbe.EstablishmentServiceClient
	Close()
}

type serviceClient struct {
	connections          []*grpc.ClientConn
	establishmentService pbe.EstablishmentServiceClient
}

func New(cfg *config.Config) (ServiceClient, error) {
	// dial to client service
	connEstablishmentService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.EstablishmentService.Host, cfg.EstablishmentService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		establishmentService: pbe.NewEstablishmentServiceClient(connEstablishmentService),
		connections: []*grpc.ClientConn{
			connEstablishmentService,
		},
	}, nil
}

func (s *serviceClient) EstablishmentService() pbe.EstablishmentServiceClient {
	return s.establishmentService
}

func (s *serviceClient) Close() {
	for _, conn := range s.connections {
		if err := conn.Close(); err != nil {
			// should be replaced by logger soon
			fmt.Printf("error while closing grpc connection: %v", err)
		}
	}
}
