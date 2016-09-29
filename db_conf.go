package db

type Config struct {
	recLevel int
	netLevel int

	readLevel  int
	writeLevel int
}

func DefaultConfig() *Config {
	return &Config{
		recLevel:   RecordLevelMemory,
		netLevel:   NetLevelNone,
		readLevel:  ReadLevelStable,
		writeLevel: WriteLevelSaved,
	}
}

func (c *Config) SetRecordLevel(level int) {
	c.recLevel = level
}

func (c *Config) SetNetLevel(level int) {
	c.netLevel = level
}

func (c *Config) SetReadLevel(level int) {
	c.readLevel = level
}

func (c *Config) SetWriteLevel(level int) {
	c.writeLevel = level
}
