package internal

import (
	"dolphin/api"

	"encoding/binary"
	"encoding/json"
)

// MarshalUser encodes a user to binary format.
func MarshalUser(user *dockm.User) ([]byte, error) {
	return json.Marshal(user)
}

// UnmarshalUser decodes a user from a binary data.
func UnmarshalUser(data []byte, user *dockm.User) error {
	return json.Unmarshal(data, user)
}

// MarshalTeam encodes a team to binary format.
func MarshalTeam(team *dockm.Team) ([]byte, error) {
	return json.Marshal(team)
}

// UnmarshalTeam decodes a team from a binary data.
func UnmarshalTeam(data []byte, team *dockm.Team) error {
	return json.Unmarshal(data, team)
}

// MarshalTeamMembership encodes a team membership to binary format.
func MarshalTeamMembership(membership *dockm.TeamMembership) ([]byte, error) {
	return json.Marshal(membership)
}

// UnmarshalTeamMembership decodes a team membership from a binary data.
func UnmarshalTeamMembership(data []byte, membership *dockm.TeamMembership) error {
	return json.Unmarshal(data, membership)
}

// MarshalEndpoint encodes an endpoint to binary format.
func MarshalEndpoint(endpoint *dockm.Endpoint) ([]byte, error) {
	return json.Marshal(endpoint)
}

// UnmarshalEndpoint decodes an endpoint from a binary data.
func UnmarshalEndpoint(data []byte, endpoint *dockm.Endpoint) error {
	return json.Unmarshal(data, endpoint)
}

// MarshalStack encodes a stack to binary format.
func MarshalStack(stack *dockm.Stack) ([]byte, error) {
	return json.Marshal(stack)
}

// UnmarshalStack decodes a stack from a binary data.
func UnmarshalStack(data []byte, stack *dockm.Stack) error {
	return json.Unmarshal(data, stack)
}

// MarshalRegistry encodes a registry to binary format.
func MarshalRegistry(registry *dockm.Registry) ([]byte, error) {
	return json.Marshal(registry)
}

// UnmarshalRegistry decodes a registry from a binary data.
func UnmarshalRegistry(data []byte, registry *dockm.Registry) error {
	return json.Unmarshal(data, registry)
}

// MarshalResourceControl encodes a resource control object to binary format.
func MarshalResourceControl(rc *dockm.ResourceControl) ([]byte, error) {
	return json.Marshal(rc)
}

// UnmarshalResourceControl decodes a resource control object from a binary data.
func UnmarshalResourceControl(data []byte, rc *dockm.ResourceControl) error {
	return json.Unmarshal(data, rc)
}

// MarshalSettings encodes a settings object to binary format.
func MarshalSettings(settings *dockm.Settings) ([]byte, error) {
	return json.Marshal(settings)
}

// UnmarshalSettings decodes a settings object from a binary data.
func UnmarshalSettings(data []byte, settings *dockm.Settings) error {
	return json.Unmarshal(data, settings)
}

// MarshalDockerHub encodes a Dockerhub object to binary format.
func MarshalDockerHub(settings *dockm.DockerHub) ([]byte, error) {
	return json.Marshal(settings)
}

// UnmarshalDockerHub decodes a Dockerhub object from a binary data.
func UnmarshalDockerHub(data []byte, settings *dockm.DockerHub) error {
	return json.Unmarshal(data, settings)
}

// Itob returns an 8-byte big endian representation of v.
// This function is typically used for encoding integer IDs to byte slices
// so that they can be used as BoltDB keys.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
