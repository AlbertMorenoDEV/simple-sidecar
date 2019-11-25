package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/config"
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
		conf := config.New()

		var params map[string]*parameter.Parameter

		repo := inmemory.NewParameterRepository(params)
		s := server.New(repo)

		srv := &http.Server{
			Addr:         "0.0.0.0:" + conf.Parameters.Port,
			WriteTimeout: time.Second * time.Duration(conf.Parameters.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(conf.Parameters.ReadTimeout),
			IdleTimeout:  time.Second * time.Duration(conf.Parameters.IdleTimeout),
			Handler:      s.Router(),
		}

		log.Printf("Listening on port %v..\n", conf.Parameters.Port)

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

		wait := time.Second * time.Duration(conf.Parameters.GracefulTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		srv.Shutdown(ctx)
		log.Println("shutting down")
		os.Exit(0)
	},
}

func init() {
	config.Setup(rootCmd)
}

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
