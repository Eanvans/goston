package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"gostonc/internal/core"
	"gostonc/internal/gost"
	"net/http"
	"os"
	"runtime"

	ginRouter "gostonc/internal/router"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/go-log/log"
)

var (
	configureFile string
	baseCfg       = &baseConfig{}
	pprofAddr     string
	pprofEnabled  = os.Getenv("PROFILING") != ""
)

func init() {
	gost.SetLogger(&gost.LogLogger{})

	var (
		printVersion bool
	)

	flag.Var(&baseCfg.route.ChainNodes, "F", "forward address, can make a forward chain")
	flag.Var(&baseCfg.route.ServeNodes, "L", "listen address, can listen on multiple ports (required)")
	flag.IntVar(&baseCfg.route.Mark, "M", 0, "Specify out connection mark")
	flag.StringVar(&configureFile, "C", "", "configure file")
	flag.StringVar(&baseCfg.route.Interface, "I", "", "Interface to bind")
	flag.BoolVar(&baseCfg.Debug, "D", false, "enable debug log")
	flag.BoolVar(&printVersion, "V", false, "print version")
	if pprofEnabled {
		flag.StringVar(&pprofAddr, "P", ":6060", "profiling HTTP server address")
	}
	flag.Parse()

	if printVersion {
		fmt.Fprintf(os.Stdout, "gost %s (%s %s/%s)\n",
			gost.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if configureFile != "" {
		_, err := parseBaseConfig(configureFile)
		if err != nil {
			log.Log(err)
			os.Exit(1)
		}
	}
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	core.Init()
}

func main() {
	if pprofEnabled {
		go func() {
			log.Log("profiling server on", pprofAddr)
			log.Log(http.ListenAndServe(pprofAddr, nil))
		}()
	}

	// NOTE: as of 2.6, you can use custom cert/key files to initialize the default certificate.
	tlsConfig, err := tlsConfig(defaultCertFile, defaultKeyFile, "")
	if err != nil {
		// generate random self-signed certificate.
		cert, err := gost.GenCertificate()
		if err != nil {
			log.Log(err)
			os.Exit(1)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	} else {
		log.Log("load TLS certificate files OK")
	}

	gost.DefaultTLSConfig = tlsConfig

	if err := start(); err != nil {
		log.Log(err)
		os.Exit(1)
	}

	go serveWeb()

	select {}
}

func start() error {
	gost.Debug = baseCfg.Debug

	var routers []router
	rts, err := baseCfg.route.GenRouters()
	if err != nil {
		return err
	}
	routers = append(routers, rts...)

	for _, route := range baseCfg.Routes {
		rts, err := route.GenRouters()
		if err != nil {
			return err
		}
		routers = append(routers, rts...)
	}

	if len(routers) == 0 {
		return errors.New("invalid config")
	}
	for i := range routers {
		go routers[i].Serve()
	}

	return nil
}

func serveWeb() {
	gin.SetMode("debug")

	router := ginRouter.NewRouter()

	addr := "127.0.0.1:8888"
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    3000,
		WriteTimeout:   1000,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Fprintf(color.Output, "gin server listen on %s\n",
		color.GreenString(fmt.Sprintf("http://%s", addr)),
	)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to launch server, errs: %v", err)
	}
}
