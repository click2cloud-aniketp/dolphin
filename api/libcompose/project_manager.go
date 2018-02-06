package libcompose

import (
	"path"
	"path/filepath"

	"dolphin/api"

	"github.com/orcaman/concurrent-map"
	"github.com/Click2Cloud/libcompose/config"
	"github.com/Click2Cloud/libcompose/docker"
	"github.com/Click2Cloud/libcompose/docker/client"
	"github.com/Click2Cloud/libcompose/docker/ctx"
	"github.com/Click2Cloud/libcompose/lookup"
	"github.com/Click2Cloud/libcompose/project"
)

// ProjectManager represents a service for managing libcompose projects.
type ProjectManager struct {
	projects cmap.ConcurrentMap
}

// NewProjectManager initializes a new ProjectManager service.
func NewProjectManager() *ProjectManager {
	return &ProjectManager{
		projects: cmap.New(),
	}
}

// GetProject will return a project associated to a stack inside an endpoint.
// If the project does not exists, it will be created.
func (manager *ProjectManager) GetProject(stack *dockm.Stack, endpoint *dockm.Endpoint) (project.APIProject, error) {
	proj, ok := manager.projects.Get(string(stack.ID))
	if !ok {
		return manager.createAndRegisterProject(stack, endpoint)
	}
	return proj.(project.APIProject), nil
}

func (manager *ProjectManager) createAndRegisterProject(stack *dockm.Stack, endpoint *dockm.Endpoint) (project.APIProject, error) {

	// TODO: APIVersion should be retrieved from the endpoint
	clientFactory, err := client.NewDefaultFactory(client.Options{
		TLS:         endpoint.TLS,
		TLSVerify:   endpoint.TLS,
		Host:        endpoint.URL,
		TLSCAFile:   endpoint.TLSCACertPath,
		TLSCertFile: endpoint.TLSCertPath,
		TLSKeyFile:  endpoint.TLSKeyPath,
		APIVersion:  "1.30",
	})
	if err != nil {
		return nil, err
	}

	composeFilePath := path.Join(stack.ProjectPath, "docker-compose.yml")
	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{composeFilePath},
			EnvironmentLookup: &lookup.ComposableEnvLookup{
				Lookups: []config.EnvironmentLookup{
					&lookup.EnvfileLookup{
						Path: filepath.Join(stack.ProjectPath, ".env"),
					},
					&lookup.OsEnvLookup{},
				},
			},
			ProjectName: stack.Name,
		},
		ClientFactory: clientFactory,
	}, nil)

	if err != nil {
		return nil, err
	}

	manager.projects.Set(string(stack.ID), project)

	return project, nil
}
