package options

import (
	"github.com/ijufumi/google-vision-sample/internal/configs"
	"google.golang.org/api/option"
	"io"
	"os"
)

func GetCredentialOption(config *configs.Config) option.ClientOption {
	if len(config.Google.CredentialFile) != 0 {
		file, err := os.Open(config.Google.CredentialFile)
		if err == nil {
			bytes, err := io.ReadAll(file)
			if err == nil {
				return option.WithCredentialsJSON(bytes)
			}
		}
	}
	return option.WithCredentialsJSON([]byte(config.Google.Credential))
}
