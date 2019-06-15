package vresolver

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// ECSMetadataFileEnvVar contains the default env var
// name that points to the container metadata file
const ECSMetadataFileEnvVar = "ECS_CONTAINER_METADATA_FILE"

// ECS extracts the tag of the Docker image in use.
// See AWS ECS Container Metadata File: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-metadata.html
func ECS(file string) string {
	if file == "" {
		return ""
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}

	metadata := struct{ ImageName string }{}

	if err := json.Unmarshal(content, &metadata); err != nil {
		return ""
	}

	parts := strings.Split(metadata.ImageName, ":")
	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}
