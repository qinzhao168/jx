package cmd_test

import (
	"path"
	"testing"

	"github.com/jenkins-x/jx/pkg/jx/cmd"
	"github.com/stretchr/testify/assert"
)

func TestFindDevPodLabel(t *testing.T) {
	t.Parallel()
	labels := []string{"go", "nodejs", "maven"}

	fileToValues := map[string]string{
		"Jenkinsfile.nodejs": "nodejs",
	}

	for fileName, label := range fileToValues {
		testFile := path.Join("test_data", "jenkinsfiles", fileName)

		answer, err := cmd.FindDevPodLabelFromJenkinsfile(testFile, labels)
		assert.NoError(t, err, "Failed to find label for file %s", testFile)
		if err == nil {
			assert.Equal(t, label, answer, "Failed to find label for file %s", testFile)
		}
	}
}
