package utils

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func GetRelease(releaseName string) *release.Release {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), nil); err != nil {
		log.Fatal(err)
	}
	getter := action.NewGet(actionConfig)
	releaseInfo, err := getter.Run(releaseName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return releaseInfo
}
