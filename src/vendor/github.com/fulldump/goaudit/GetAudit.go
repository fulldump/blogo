package goaudit

import (
	"github.com/fulldump/golax"
)

// GetAudit retrieve Audit object from context.
//
// Typical usage:
//
//	func MyHandler(c *golax.Context) {
//	    // ...
//	    audit := goaudit.GetAudit(c)
//	    // ...
//	}
func GetAudit(c *golax.Context) *Audit {
	v, exists := c.Get(CONTEXT_KEY)

	if !exists {
		return nil
	}
	return v.(*Audit)
}
