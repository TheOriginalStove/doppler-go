package dopplergo

import "time"

// TODO: Finish this comment
// DopplerConfig interface will implement the various config gathering functions
// and can be used for mocking in tests.
type DopplerConfig interface {
	ListConfigs()
	CreateConfig()
	RetrieveConfig()
	UpdateConfig()
	DeleteConfig()
	CloneConfig()
	LockConfig()
	UnlockConfig()
}

type ConfigClient struct {
}

type Config struct {
	Name           string    `json:"name"`
	Project        string    `json:"project"`
	Environment    string    `json:"environment"`
	CreatedAt      time.Time `json:"created_at"`
	InitialFetchAt time.Time `json:"intial_fetch_at"`
	LastFetchAt    time.Time `json:"last_fetch_at"`
	Root           bool      `json:"root"`
	Locked         bool      `json:"locked"`
}

func (c *ConfigClient) ListConfigs() {
}

func (c *ConfigClient) CreateConfig() {
}

func (c *ConfigClient) RetrieveConfig() {
}

func (c *ConfigClient) UpdateConfig() {
}

func (c *ConfigClient) DeleteConfig() {
}

func (c *ConfigClient) CloneConfig() {
}

func (c *ConfigClient) LockConfig() {
}

func (c *ConfigClient) UnlockConfig() {
}
