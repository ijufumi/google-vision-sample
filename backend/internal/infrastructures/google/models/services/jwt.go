package services

import (
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"io"
	"os"
)

func GetJWTConfig(config *configs.Config) (*jwt.Config, error) {
	var jsonBytes []byte
	if len(config.Google.CredentialFile) != 0 {
		file, err := os.Open(config.Google.CredentialFile)
		if err == nil {
			bytes, err := io.ReadAll(file)
			if err != nil {
				return nil, err
			}
			jsonBytes = bytes
		}
	} else {
		jsonBytes = []byte(config.Google.Credential)
	}

	return google.JWTConfigFromJSON(jsonBytes)
}
