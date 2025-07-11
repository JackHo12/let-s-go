package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Use the flag package to parse command-line flags. We define a string flag
	// called "addr" with a default value of ":4000" and a description.
	// The flag.Parse() function is called to parse the command-line flags.
	// The value of the "addr" flag will be stored in the addr variable.
	// The flag package will automatically handle the command-line arguments
	// and set the value of the addr variable based on the provided flag.
	// ./web -addr=:4000
	// addr := flag.String("addr", ":4000", "HTTP network address")
	type config struct {
		addr      string
		staticDir string
	}
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to serve static files")
	addr := &cfg.addr
	flag.Parse()
	// os.Stdout is the standard output stream, which is where log messages will be written.
	// slog.HandlerOptions is used to configure the logging behavior.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Set the minimum log level to Debug
		// AddSource: true,            // Add source code file and line number information to log entries
	}))
	app := &application{
		logger: logger,
	}
	mux := app.routes(cfg.staticDir)

	// Print a log message to say that the server is starting.
	// log.Printf("starting server on %s", *addr)
	logger.Info("starting server", slog.String("addr", *addr))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(*addr, mux)

	// // if you pass nil as the second parameter, it will use the default servemux
	// err := http.ListenAndServe(":4000", nil)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1) // Exit the program with a non-zero status code
}

func neuter(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
