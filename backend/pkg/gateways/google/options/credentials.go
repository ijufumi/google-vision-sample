package options

import (
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"google.golang.org/api/option"
)

func GetCredentialOption(config *configs.Config) option.ClientOption {
	return option.WithCredentialsJSON([]byte(config.Google.Credential))
}
