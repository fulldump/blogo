package statics

import (
	"github.com/fulldump/golax"
	"strings"
	"path/filepath"
	"io/ioutil"
)

var contentTypes = map[string]string{

	// Web
	".css":  "text/css; charset=utf-8",
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".js":   "text/javascript; charset=UTF-8",
	".json": "application/json; charset=UTF-8",

	// Images
	".ico":  "image/x-icon",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".svg":  "image/svg+xml",

	// Fonts:
	".eot":   "application/vnd.ms-fontobject",
	".otf":   "application/x-font-opentype",
	".sfnt":  "application/font-sfnt",
	".ttf":   "application/x-font-ttf",
	".woff":  "application/font-woff",
	".woff2": "application/font-woff2",
}

func setContentType(c *golax.Context, filename string) {
	ext := strings.ToLower(filepath.Ext(filename))

	value, exists := contentTypes[ext]
	if exists {
		c.Response.Header().Set("Content-Type", value)
	}
}

func readFileInternal(c *golax.Context, filename string) {
	content, exists := Bytes[filename]
	if !exists {
		c.Error(404, "File '"+filename+"' not found")
		return
	}

	setContentType(c, filename)

	c.Response.Write(content)
}

/**
 * Only for develop purposes
 */
func readFileExternal(statics string) func(*golax.Context, string) {
	return func(c *golax.Context, filename string) {

		bytes, err := ioutil.ReadFile(statics + filename)
		if nil != err {
			c.Error(404, "File '"+filename+"' not found")
			return
		}

		setContentType(c, filename)

		c.Response.Write(bytes)
	}

}
