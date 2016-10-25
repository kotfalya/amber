package amber

type Config struct {
	persist    int
	readLevel  int
	writeLevel int
}

func DefaultConfig() *Config {
	return &Config{
		persist:    PersistNone,
		readLevel:  LevelLocal,
		writeLevel: LevelMaster,
	}
}

func (c *Config) SetPersist(persist int) {
	c.persist = persist
}

func (c *Config) SetReadLevel(level int) {
	c.readLevel = level
}

func (c *Config) SetWriteLevel(level int) {
	c.writeLevel = level
}
