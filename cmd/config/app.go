package config

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"ashish.com/m/internal/utils"
	v1 "ashish.com/m/pb/ashish.com/v1"
	api "ashish.com/m/pkg/api/v1"
	app "ashish.com/m/pkg/app"

	//httpic "ashish.com/m/pkg/interceptors/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	stdsvc "ashish.com/m/pkg/services/standard"
	pgstorage "ashish.com/m/pkg/storage/postgres"
)

// App represents Test api application.
type App struct {
	cfg               Config
	db                *gorm.DB
	sqlDB             *sql.DB
	isRunning         bool
	grpcServer        *grpc.Server
	httpServer        *http.Server
	grpcServerRunning uint32
	httpServerRunning uint32
}

// New creates a new instance of application.
func New(cfg Config) *App {
	return &App{
		cfg:               cfg,
		db:                nil,
		sqlDB:             nil,
		isRunning:         false,
		grpcServer:        nil,
		httpServer:        nil,
		grpcServerRunning: 0,
		httpServerRunning: 0,
	}
}

// Init initializes the application.
func (a *App) Init(ctx context.Context) {
	// open db connection
	db, err := app.OpenDatabase(ctx, a.cfg.DBEndpoint)
	if err != nil {
		log.WithContext(ctx).Fatalf("error while connecting to database: %v", err)
	}
	a.db = db
	a.sqlDB, _ = db.DB()

	// instantiate dependencies - storage
	emplpoyeeStore := pgstorage.NewEmployeeStore(db)
	//	personStore := pgstorage.NewPersonStore(db)

	// instantiate dependencies - services
	employeeService := stdsvc.NewEmployeeService(emplpoyeeStore)
	//personService := stdsvc.NewPersonService(personStore)

	// instantiate dependencies - grpc and http api

	employeeGRPCGRPCService := api.NewEmployeeGRPCService(employeeService)
	//personHTTPService := api.NewPersonHTTPService(personService)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	v1.RegisterEmployeeServiceServer(grpcServer, employeeGRPCGRPCService)
	a.grpcServer = grpcServer

	// configure grpc gateway
	grpcEndpoint := fmt.Sprintf(":%d", a.cfg.GRPCPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	_ = v1.RegisterEmployeeServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)

	// // configure custom http endpoints
	// getPerson := httpic.InterceptHTTPServer(personHTTPService.GetPerson)

	// //_ = mux.HandlePath(http.MethodGet, "/v1beta2/{name=persons/*}", getPerson)
	// //"/v1beta2/{name=persons/*}"
	// patOpt := runtime.PatternOpt
	// pat,_  := runtime.NewPattern(1, []int{}, []string{}, "Get", )
	// pattern := runtime.MustPattern(runtime.NewPattern(runtime.WithPathPattern("/v1beta2/{name=persons/*}")))
	// mux.Handle("http.MethodGet", pattern,  myHandler)

	// instantiate http server
	// httpServer := &http.Server{
	// 	Addr:              fmt.Sprintf(":%d", a.cfg.HTTPPort),
	// 	ReadHeaderTimeout: time.Second * 30,
	// 	Handler:           mux,
	// }

	// a.httpServer = httpServer
	// log.WithContext(ctx).Info("application initialization done")
}

func myHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// Your code to handle the HTTP request goes here
	// You can read request parameters, write a response, etc.
	w.Write([]byte("Hello, World!"))
}

// Run starts the application.
func (a *App) Run(ctx context.Context) {
	// don't run if already running
	if a.isRunning {
		return
	}

	// set internal state as running
	a.isRunning = true

	// use waitgroup for synchronization
	wg := new(sync.WaitGroup)

	// run grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer atomic.StoreUint32(&a.grpcServerRunning, 0)
		atomic.StoreUint32(&a.grpcServerRunning, 1)

		// run grpc server
		runGRPCServer(ctx, a.grpcServer, a.cfg.GRPCPort)
	}()

	// wait for all goroutines to finish
	time.Sleep(utils.WaitTiny)
	log.WithContext(ctx).Info("started application and all its components")
	wg.Wait()

	// set internal state as not running
	a.isRunning = false
}

// Shutdown stops the application.
func (a *App) Shutdown(ctx context.Context) {
	// don't proceed if not running
	if !a.isRunning {
		return
	}

	// stop grpc server
	if atomic.LoadUint32(&a.grpcServerRunning) == 1 {
		a.grpcServer.GracefulStop()
	}

	// check for shutdown status
	for {
		if !a.isRunning {
			break
		}
		time.Sleep(utils.WaitBlink)
	}

	// close database connection
	if err := a.sqlDB.Close(); err != nil {
		log.WithContext(ctx).Errorf("error while closing db connection: %v", err)
	}

	// log success message for application shutdown
	log.WithContext(ctx).Info("application has been shutdown")
}

func runGRPCServer(ctx context.Context, server *grpc.Server, port uint16) {
	// create tcp listener
	tcpListner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.WithContext(ctx).Errorf("error on starting tcp listener: %v", err)
		return
	}

	// start grpc server
	log.WithContext(ctx).WithField("port", port).Info("started grpc server")
	if err := server.Serve(tcpListner); err != nil {
		log.WithContext(ctx).Errorf("error on serving grpc server: %v", err)
	}
	log.WithContext(ctx).Info("stopped grpc server")
}
