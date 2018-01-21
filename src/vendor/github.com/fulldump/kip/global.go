package kip

var global *Kip = nil

func init() {
	global = NewKip()
}

func Global() *Kip {
	return global
}

func Define(c *Collection) {
	global.Define(c)
}

func NewDao(name string, db *Database) *Dao {
	return global.NewDao(name, db)
}
