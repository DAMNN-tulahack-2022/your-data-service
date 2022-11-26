package transport

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/damnn/tulahack/your-supadmin-service/gen/proto"
	"github.com/damnn/tulahack/your-supadmin-service/tools"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	gatewayProxy struct {
		proto.UnimplementedYourAdminServiceServer
	}

	AuthProxy struct {
		// login
		// service supadmin.AuthorizationCore

		log    *tools.YourLogger
		proxy  *gatewayProxy
		config *tools.GenericEnvAppConfig

		// parent
		ctx context.Context

		sig chan os.Signal
	}
)

func NewAuthProxy(ctx context.Context, log *tools.YourLogger, config *tools.GenericEnvAppConfig) *AuthProxy {
	return &AuthProxy{
		ctx:    ctx,
		log:    log,
		config: config,
		sig:    make(chan os.Signal),
		proxy:  &gatewayProxy{},
	}
}

func (ap *AuthProxy) Run() error {
	// rest api server init
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true},
		}),
	)

	if err := proto.RegisterYourAdminServiceHandlerServer(ap.ctx, mux, ap.proxy); err != nil {
		log.Panic(err)
	}

	if err := http.ListenAndServe(ap.config.DataProxyPort, cors(mux)); err != nil {
		return err
	}

	return nil
}

func (ap *AuthProxy) GracefullNotify() error {
	signal.Notify(ap.sig, os.Interrupt, syscall.SIGTERM)
	<-ap.sig

	return nil
}
