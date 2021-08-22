package engine

import (
	"html/template"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/xyproto/datablock"
)

// Functions for concurrent use by rendering.go and handlers.go

// Lua2funcMap runs in a Lua file and returns the functions as a template.FuncMap (or an error)
func (ac *Config) Lua2funcMap(w http.ResponseWriter, req *http.Request, filename, luafilename, ext string, errChan chan error, funcMapChan chan template.FuncMap) {

	// Make functions from the given Lua data available
	funcs := make(template.FuncMap)

	// Try reading data.lua, if possible
	luablock, err := ac.cache.Read(luafilename, ac.shouldCache(ext))
	if err != nil {
		// Could not find and/or read data.lua
		luablock = datablock.EmptyDataBlock

		// This only means the file wasn't cached, so just ignore this error
	}

	// luablock can be empty if there was an error or if the file was empty
	if luablock.HasData() {
		// There was Lua code available. Now make the functions and
		// variables available for the template.
		funcs, err = ac.LuaFunctionMap(w, req, luablock.MustData(), luafilename)
		if err != nil {
			funcMapChan <- funcs
			errChan <- err
			return
		}
		if ac.debugMode && ac.verboseMode {
			s := "These functions from " + luafilename
			s += " are useable for " + filename + ": "
			// Create a comma separated list of the available functions
			for key := range funcs {
				s += key + ", "
			}
			// Remove the final comma
			if strings.HasSuffix(s, ", ") {
				s = s[:len(s)-2]
			}
			// Output the message
			log.Info(s)
		}
	}
	funcMapChan <- funcs
	errChan <- err
}
