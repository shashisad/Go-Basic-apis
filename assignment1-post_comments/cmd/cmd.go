package cmd

import (
	"context"
	"flag"
	"fmt"
	"posts/handlers"
	"posts/helpers"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	httpAddr = ":8000" // bind address
)

const (
	httpTimeout = 3 * time.Second // timeouts used to protect the server
)

// run accepts the program arguments and where to send output (default: stdout)
func Run(args []string, _ io.Writer) error {
	var (
		port string
	)

	//	c := cache.CacheFunc()

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&port, "port", "8000", "set port  --port=8000")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if port != "" {
		httpAddr = ":" + port
	}

	log.Println("Starting Application on port", httpAddr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errGrp, egCtx := errgroup.WithContext(ctx)
	s := handlers.NewServer(httpAddr, httpTimeout, httpTimeout)

	// http server
	errGrp.Go(func() error {
		return s.Start(egCtx)
	})

	go ticker(ctx)

	// signal handler
	errGrp.Go(func() error {
		return handleSignals(egCtx, cancel)
	})

	return errGrp.Wait()
}

func ticker(ctx context.Context) error {
	fmt.Println("Ticker running...")
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			_, errReadingPosts := helpers.FetchAllPosts("https://jsonplaceholder.typicode.com/posts")
			if errReadingPosts != nil {
				fmt.Println("Can't cache.. error while reading posts")
				break
			}
			handlers.Caching()

		case <-ctx.Done():
			fmt.Printf("Ctx done")
			return ctx.Err()
		}
	}
}

// handleSignals will handle Interrupts or termination signals
func handleSignals(ctx context.Context, cancel context.CancelFunc) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-sigCh:
		log.Printf("got signal %v, stopping", s)
		cancel()
		return nil
	case <-ctx.Done():
		log.Println("context is done")
		return ctx.Err()
	}
}
