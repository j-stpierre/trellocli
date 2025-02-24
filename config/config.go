package config

import ()

type GetConfig struct {
	Label string
}

type DeleteConfig struct {
	File string
}

type CredentialConfig struct {
	BoardId string
	APIKey  string
	Token   string
}
