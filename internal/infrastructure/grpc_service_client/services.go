package grpc_service_clients

import (
	"fmt"

	pbe "Booking/api-service-booking/genproto/establishment-proto"
	pbu "Booking/api-service-booking/genproto/user-proto"
	pbb "Booking/api-service-booking/genproto/booking-proto"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"Booking/api-service-booking/internal/pkg/config"
)

type ServiceClient interface {
	EstablishmentService() pbe.EstablishmentServiceClient
	UserService() pbu.UserServiceClient
	BookingService() pbb.BookingServiceClient
	Close()
}

type serviceClient struct {
	connections          []*grpc.ClientConn
	establishmentService pbe.EstablishmentServiceClient
	userService pbu.UserServiceClient
	bookingService pbb.BookingServiceClient
}

func New(cfg *config.Config) (ServiceClient, error) {
	// dial to establishment service
	connEstablishmentService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.EstablishmentService.Host, cfg.EstablishmentService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}


	// user service

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserService.Host, cfg.UserService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

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
		establishmentService: pbe.NewEstablishmentServiceClient(connEstablishmentService),
		userService: pbu.NewUserServiceClient(connUserService),
		bookingService: pbb.NewBookingServiceClient(connBookingService),
		connections: []*grpc.ClientConn{
			connEstablishmentService,
			connUserService,
			connBookingService,
		},
	}, nil
}

func (s *serviceClient) EstablishmentService() pbe.EstablishmentServiceClient {
	return s.establishmentService
}

func (s *serviceClient) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceClient) BookingService() pbb.BookingServiceClient{
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
