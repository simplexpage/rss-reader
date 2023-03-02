package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/oklog/pkg/group"
	"github.com/simplexpage/rss-reader/internal/reader/config"
	"github.com/simplexpage/rss-reader/internal/reader/delivery/transport"
	"github.com/simplexpage/rss-reader/internal/reader/domain/service"
	"github.com/simplexpage/rss-reader/internal/reader/endpoint"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	cfg := config.GetConfig(logger)

	fs := flag.NewFlagSet("rss-reader", flag.ExitOnError)

	var (
		httpAddr = fs.String("http-addr", fmt.Sprintf(":%s", cfg.ListenHttp.Port), "HTTP listen address")
		//queueDsn = fs.String("rabbitmq-url", fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Rabbit.User, cfg.Rabbit.Password, cfg.Rabbit.Host, cfg.Rabbit.Port), "RabbitMQ DSN")
	)

	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	var (
		readerService     = service.New(logger)
		readerEndpoints   = endpoint.NewServerEndpoints(readerService, logger)
		readerHttpHandler = transport.NewHTTPHandler(readerEndpoints, log.With(logger, "component", "HTTP"))
	)

	var g group.Group
	{
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			level.Info(logger).Log("delivery", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			level.Info(logger).Log("delivery", "HTTP", "addr", *httpAddr)
			return http.Serve(httpListener, readerHttpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	level.Info(logger).Log("exit", g.Run())
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
