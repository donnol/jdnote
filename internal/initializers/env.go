package initializers

import "os"

const (
	ProjectEnv = "PROJECT_ENV"

	ProjectEnvDev = "PROJECT_ENV_DEV"

	ProjectEnvProd = "PROJECT_ENV_PROD"
)

func GetProjectEnv() string {
	if v, ok := os.LookupEnv(ProjectEnv); ok && v == ProjectEnvProd {
		return ProjectEnvProd
	}
	return ProjectEnvDev
}
