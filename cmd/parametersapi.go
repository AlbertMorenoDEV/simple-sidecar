package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter"
	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter/server"
	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter/storage/inmemory"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "parameterapi",
	Short: "Feature flag server",
	Long:  `Feature flag server manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listening...")

		var parameters map[string]*parameter.Parameter

		if withData {
			parameters = inmemory.Parameters
		}

		repo := inmemory.NewParameterRepository(parameters)
		s := server.New(repo)

		srv := &http.Server{
			Addr:         "0.0.0.0:7983",
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      s.Router(),
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()

		c := make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		<-c

		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		srv.Shutdown(ctx)
		log.Println("shutting down")
		os.Exit(0)
	},
}

var withData bool
var wait time.Duration

func init() {
	rootCmd.Flags().BoolVarP(&withData, "with-data", "p", false, "Load sample data")
	rootCmd.Flags().DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
}

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
