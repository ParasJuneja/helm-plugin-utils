package utils

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func ListReleases(releaseName string, releaseNamespace string) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), releaseNamespace, os.Getenv("HELM_DRIVER"), nil); err != nil {
		log.Fatal(err)
	}
	getter := action.NewGet(actionConfig)
	getter.Run(releaseName)
}
