package cmd_test

import (
	"os"
	"testing"

	"github.com/jenkins-x/jx/pkg/config"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
)

func TestGetPreviewValuesConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		opts               cmd.PreviewOptions
		env                map[string]string
		domain             string
		expectedYAMLConfig string
	}{
		{
			opts: cmd.PreviewOptions{
				HelmValuesConfig: config.HelmValuesConfig{
					ExposeController: &config.ExposeController{},
				},
			},
			env: map[string]string{
				cmd.DOCKER_REGISTRY: "my.registry",
				cmd.ORG:             "my-org",
				cmd.APP_NAME:        "my-app",
				cmd.PREVIEW_VERSION: "1.0.0",
			},
			expectedYAMLConfig: `expose: {}
preview:
  image:
    repository: my.registry/my-org/my-app
    tag: 1.0.0
`,
		},
		{
			opts: cmd.PreviewOptions{
				HelmValuesConfig: config.HelmValuesConfig{
					ExposeController: &config.ExposeController{
						Config: config.ExposeControllerConfig{
							HTTP:    "false",
							TLSAcme: "true",
						},
					},
				},
			},
			env: map[string]string{
				cmd.DOCKER_REGISTRY: "my.registry",
				cmd.ORG:             "my-org",
				cmd.APP_NAME:        "my-app",
				cmd.PREVIEW_VERSION: "1.0.0",
			},
			domain: "jenkinsx.io",
			expectedYAMLConfig: `expose:
  config:
    domain: jenkinsx.io
    http: "false"
    tlsacme: "true"
preview:
  image:
    repository: my.registry/my-org/my-app
    tag: 1.0.0
`,
		},
	}

	for i, test := range tests {
		for k, v := range test.env {
			os.Setenv(k, v)
		}

		config, err := test.opts.GetPreviewValuesConfig(test.domain)
		if err != nil {
			t.Errorf("[%d] got unexpected err: %v", i, err)
			continue
		}

		configYAML, err := config.String()
		if err != nil {
			t.Errorf("[%d] %v", i, err)
			continue
		}

		if test.expectedYAMLConfig != configYAML {
			t.Errorf("[%d] expected %#v but got %#v", i, test.expectedYAMLConfig, configYAML)
		}
	}
}
