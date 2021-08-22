package engine

// This source file is for the special case of serving a single file.

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xyproto/algernon/utils"
	"github.com/xyproto/datablock"
)

const (
	defaultStaticCacheSize = 128 * utils.MiB

	maxAttemptsAtIncreasingPortNumber = 128

	delayBeforeLaunchingBrowser = time.Millisecond * 200
)

// nextPort increases the port number by 1
func nextPort(colonPort string) (string, error) {
	if !strings.HasPrefix(colonPort, ":") {
		return colonPort, errors.New("colonPort does not start with a colon! \"" + colonPort + "\"")
	}
	num, err := strconv.Atoi(colonPort[1:])
	if err != nil {
		return colonPort, errors.New("Could not convert port number to string: \"" + colonPort[1:] + "\"")
	}
	// Increase the port number by 1, add a colon, convert to string and return
	return ":" + strconv.Itoa(num+1), nil
}

// This is a bit hacky, but it's only used when serving a single static file
func (ac *Config) openAfter(wait time.Duration, hostname, colonPort string, https bool, cancelChannel chan bool) {
	// Wait a bit
	time.Sleep(wait)
	select {
	case <-cancelChannel:
		// Got a message on the cancelChannel:
		// don't open the URL with an external application.
		return
	case <-time.After(delayBeforeLaunchingBrowser):
		// Got timeout, assume the port was not busy
		ac.OpenURL(hostname, colonPort, https)
	}
}

// shortInfo outputs a short string about which file is served where
func (ac *Config) shortInfoAndOpen(filename, colonPort string, cancelChannel chan bool) {
	hostname := "localhost"
	if ac.serverHost != "" {
		hostname = ac.serverHost
	}
	log.Info("Serving " + filename + " on http://" + hostname + colonPort)

	if ac.openURLAfterServing {
		go ac.openAfter(delayBeforeLaunchingBrowser, hostname, colonPort, false, cancelChannel)
	}
}

// ServeStaticFile is a convenience function for serving only a single file.
// It can be used as a quick and easy way to view a README.md file.
func (ac *Config) ServeStaticFile(filename, colonPort string) error {
	log.Info("Single file mode. Not using the regular parameters.")

	cancelChannel := make(chan bool, 1)

	ac.shortInfoAndOpen(filename, colonPort, cancelChannel)

	mux := http.NewServeMux()
	// 64 MiB cache, use cache compression, no per-file size limit, use best gzip compression, compress for size not for speed
	ac.cache = datablock.NewFileCache(defaultStaticCacheSize, true, 0, false, 0)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Server", ac.versionString)
		ac.FilePage(w, req, filename, ac.defaultLuaDataFilename)
	})
	HTTPserver := ac.NewGracefulServer(mux, false, ac.serverHost+colonPort)

	// Attempt to serve just the single file
	if errServe := HTTPserver.ListenAndServe(); errServe != nil {
		// If it fails, try several times, increasing the port by 1 each time
		for i := 0; i < maxAttemptsAtIncreasingPortNumber; i++ {
			if errServe = HTTPserver.ListenAndServe(); errServe != nil {
				cancelChannel <- true
				if !strings.HasSuffix(errServe.Error(), "already in use") {
					// Not a problem with address already being in use
					ac.fatalExit(errServe)
				}
				log.Warn("Address already in use. Using next port number.")
				if newPort, errNext := nextPort(colonPort); errNext != nil {
					ac.fatalExit(errNext)
				} else {
					colonPort = newPort
				}

				// Make a new cancel channel, and use the new URL
				cancelChannel = make(chan bool, 1)
				ac.shortInfoAndOpen(filename, colonPort, cancelChannel)

				HTTPserver = ac.NewGracefulServer(mux, false, ac.serverHost+colonPort)
			}
		}
		// Several attempts failed
		return errServe
		//ac.fatalExit(errServe)
	}
	return nil
}
