package options

import (
	"encoding/json"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"google.golang.org/api/option"
)

func GetCredentialOption(config *configs.Config) option.ClientOption {
	bytes, err := json.Marshal(config.Google.Credential)
	if err != nil {
		panic(err)
	}
	return option.WithCredentialsJSON(bytes)
}
