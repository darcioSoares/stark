package config

import (
	"os"
)

var (
	PrivateKey string
	IDProject  string
)

func LoadEnvVars() {
	PrivateKey = os.Getenv("PRIVATE_KEY")
	IDProject = os.Getenv("ID_PROJECT")

	if PrivateKey == "" {
		panic("Erro: PRIVATE_KEY não definido")
	}
	if IDProject == "" {
		panic("Erro: ID_PROJECT não definido")
	}
}
