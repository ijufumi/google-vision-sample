package options

import (
	"context"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func GetCredentialOption(config *configs.Config) (option.ClientOption, error) {
	credentials, err := google.CredentialsFromJSON(context.Background(), []byte(config.Google.Credential))
	if err != nil {
		return nil, err
	}
	return option.WithCredentials(credentials), nil
}
