package cmd

import (
	"log"
	"net/http"

	"github.com/joranmulderij/pocketbase/apis"
	"github.com/joranmulderij/pocketbase/core"
	"github.com/spf13/cobra"
)

// NewServeCommand creates and returns new command responsible for
// starting the default PocketBase web server.
func NewServeCommand(app core.App, showStartBanner bool) *cobra.Command {
	var allowedOrigins []string
	var tlsOrigins []string
	var httpAddr string
	var httpsAddr string

	command := &cobra.Command{
		Use:   "serve",
		Short: "Starts the web server (default to 127.0.0.1:8090)",
		Run: func(command *cobra.Command, args []string) {
			_, err := apis.Serve(app, apis.ServeConfig{
				HttpAddr:        httpAddr,
				HttpsAddr:       httpsAddr,
				ShowStartBanner: showStartBanner,
				AllowedOrigins:  allowedOrigins,
				TLSOrigins:      tlsOrigins,
			})

			if err != http.ErrServerClosed {
				log.Fatalln(err)
			}
		},
	}

	command.PersistentFlags().StringSliceVar(
		&allowedOrigins,
		"origins",
		[]string{"*"},
		"CORS allowed domain origins list",
	)

	command.PersistentFlags().StringSliceVar(
		&tlsOrigins,
		"tls-origins",
		[]string{},
		"TLS allowed domain origins list",
	)

	command.PersistentFlags().StringVar(
		&httpAddr,
		"http",
		"127.0.0.1:8090",
		"api HTTP server address",
	)

	command.PersistentFlags().StringVar(
		&httpsAddr,
		"https",
		"",
		"api HTTPS server address (auto TLS via Let's Encrypt)\nthe incoming --http address traffic also will be redirected to this address",
	)

	return command
}
