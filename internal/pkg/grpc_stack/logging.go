package grpc_stack

import (
	"context"

	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
	"google.golang.org/grpc"
)

func LoggingUnaryInterceptor(logger logging.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		logger.Info("received request", lf.String("method", info.FullMethod))

		resp, err := handler(ctx, req)
		loggingFn, msg, fields := getLoggingFn(err, logger)
		fields = append(fields, lf.String("method", info.FullMethod))
		loggingFn(msg, fields...)

		return resp, err
	}
}

func LoggingStreamInterceptor(logger logging.Logger) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger.Info("received stream request", lf.String("method", info.FullMethod))

		err := handler(srv, stream)
		loggingFn, msg, fields := getLoggingFn(err, logger)
		fields = append(fields, lf.String("method", info.FullMethod))
		loggingFn(msg, fields...)

		return err
	}
}

func getLoggingFn(err error, logger logging.Logger) (func(string, ...lf.Field), string, []lf.Field) {
	loggingFn := logger.Info
	msg := "request completed"
	var fields []lf.Field
	if err != nil {
		loggingFn = logger.Error
		msg = "request failed"
		fields = append(fields, lf.Err(err))
	}
	return loggingFn, msg, fields
}
