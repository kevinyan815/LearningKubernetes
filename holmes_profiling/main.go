package main

import (
	"context"
	"fmt"
	"github.com/mosn/holmes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type indexHandler struct {}
type healthCheckHandler struct {}
type memProfileTestHandler struct {}
type cpuExHandler struct {}
type chanBlockHandler struct {}

func (*indexHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s", hostname)
}

func (*healthCheckHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func (*memProfileTestHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var a = make([]byte, 1073741824)
	_ = a
}

func (*cpuExHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go func() {
		for {
			time.Sleep(time.Millisecond)
		}
	}()
}

var nilCh chan int
func (*chanBlockHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nilCh <- 1
}

func main() {
	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})
	mux.Handle("/health-check", &healthCheckHandler{})
	mux.Handle("/1gb-slice", &memProfileTestHandler{})
	mux.Handle("/cpu-ex", &cpuExHandler{})
	mux.Handle("/chan-block", &chanBlockHandler{})

	srv := &http.Server{
		Addr:              ":3000",
		Handler:           mux,
		TLSConfig:         nil,
	}
	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		// start serve
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	// 开启Continuous Profiling
	h, _ := holmes.New(
		holmes.WithCollectInterval("10s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath("/tmp", "profile.log"),
		holmes.WithTextDump(),
		holmes.WithMemDump(30, 25, 80),
		holmes.WithCPUDump(20, 25, 80),
		holmes.WithGoroutineDump(10, 25, 2000),
		holmes.WithCGroup(true),
	)
	h.EnableMemDump().EnableCPUDump().EnableGoroutineDump().Start()

	select {
	case err := <-errChan:
		panic(err)
	case <-stopChan:
		// 关闭Profiling
		h.Stop()
		// 关闭Server
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown error %+w", err)
		}
		fmt.Println("Server shutdown...")
	}
}
