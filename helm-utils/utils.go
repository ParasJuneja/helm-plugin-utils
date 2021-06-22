package utils

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func GetRelease(releaseName string, releaseNamespace string) *release.Release {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "dmp-system", os.Getenv("HELM_DRIVER"), nil); err != nil {
		log.Fatal(err)
	}
	getter := action.NewGet(actionConfig)
	releaseInfo, err := getter.Run("swagger")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return releaseInfo
}
