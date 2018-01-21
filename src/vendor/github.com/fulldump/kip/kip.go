package kip

type Kip struct {
	collections map[string]*Collection
}

func NewKip() *Kip {
	return &Kip{
		collections: map[string]*Collection{},
	}
}

func (k *Kip) Define(c *Collection) {
	name := c.Name

	// Check name is available
	_, exists := k.collections[name]
	if exists {
		panic("Collection `" + name + "` already defined")
	}

	// Check mandatory callback
	if nil == c.OnCreate {
		panic("Mandatory callback `OnCreate` is needed for `" + c.Name + "`")
	}

	k.collections[name] = c
}

func (k *Kip) NewDao(name string, database *Database) *Dao {

	// Check name is defined
	c, exists := k.collections[name]
	if !exists {
		panic("Try to Dao `" + name + "` but it is not defined")
	}

	// Create Dao
	i := &Dao{
		Collection: c,
		Database:   database,
	}

	db := database.Clone()
	defer db.Close()

	// Ensure indexes
	for _, index := range c.Indexes {
		if nil != db.C(c.Name).EnsureIndex(index) {
			panic("Unable to ensure index for `" + name + "`")
		}
	}

	return i
}
