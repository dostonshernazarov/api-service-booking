package grpc_service_clients

import (
	"fmt"

	pbb "Booking/api-service-booking/genproto/booking-proto"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"Booking/api-service-booking/internal/pkg/config"
)

type ServiceClient interface {
	BookingService() pbb.BookingServiceClient
	Close()
}

type serviceClient struct {
	connections          []*grpc.ClientConn
	bookingService pbb.BookingServiceClient
}

func New(cfg *config.Config) (ServiceClient, error) {
	// dial to user service
	connBookingService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.BookingService.Host, cfg.BookingService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		bookingService: pbb.NewBookingServiceClient(connBookingService),
		connections: []*grpc.ClientConn{
			connBookingService,
		},
	}, nil
}

func (s *serviceClient) BookingService() pbb.BookingServiceClient {
	return s.bookingService
}

func (s *serviceClient) Close() {
	for _, conn := range s.connections {
		if err := conn.Close(); err != nil {
			// should be replaced by logger soon
			fmt.Printf("error while closing grpc connection: %v", err)
		}
	}
}
