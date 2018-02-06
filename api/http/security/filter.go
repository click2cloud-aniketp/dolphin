package security

import "dolphin/api"

// FilterUserTeams filters teams based on user role.
// non-administrator users only have access to team they are member of.
func FilterUserTeams(teams []dockm.Team, context *RestrictedRequestContext) []dockm.Team {
	filteredTeams := teams

	if !context.IsAdmin {
		filteredTeams = make([]dockm.Team, 0)
		for _, membership := range context.UserMemberships {
			for _, team := range teams {
				if team.ID == membership.TeamID {
					filteredTeams = append(filteredTeams, team)
					break
				}
			}
		}
	}

	return filteredTeams
}

// FilterLeaderTeams filters teams based on user role.
// Team leaders only have access to team they lead.
func FilterLeaderTeams(teams []dockm.Team, context *RestrictedRequestContext) []dockm.Team {
	filteredTeams := teams

	if context.IsTeamLeader {
		filteredTeams = make([]dockm.Team, 0)
		for _, membership := range context.UserMemberships {
			for _, team := range teams {
				if team.ID == membership.TeamID && membership.Role == dockm.TeamLeader {
					filteredTeams = append(filteredTeams, team)
					break
				}
			}
		}
	}

	return filteredTeams
}

// FilterUsers filters users based on user role.
// Non-administrator users only have access to non-administrator users.
func FilterUsers(users []dockm.User, context *RestrictedRequestContext) []dockm.User {
	filteredUsers := users

	if !context.IsAdmin {
		filteredUsers = make([]dockm.User, 0)

		for _, user := range users {
			if user.Role != dockm.AdministratorRole {
				filteredUsers = append(filteredUsers, user)
			}
		}
	}

	return filteredUsers
}

// FilterRegistries filters registries based on user role and team memberships.
// Non administrator users only have access to authorized endpoints.
func FilterRegistries(registries []dockm.Registry, context *RestrictedRequestContext) ([]dockm.Registry, error) {

	filteredRegistries := registries
	if !context.IsAdmin {
		filteredRegistries = make([]dockm.Registry, 0)

		for _, registry := range registries {
			if isRegistryAccessAuthorized(&registry, context.UserID, context.UserMemberships) {
				filteredRegistries = append(filteredRegistries, registry)
			}
		}
	}

	return filteredRegistries, nil
}

// FilterEndpoints filters endpoints based on user role and team memberships.
// Non administrator users only have access to authorized endpoints.
func FilterEndpoints(endpoints []dockm.Endpoint, context *RestrictedRequestContext) ([]dockm.Endpoint, error) {
	filteredEndpoints := endpoints

	if !context.IsAdmin {
		filteredEndpoints = make([]dockm.Endpoint, 0)

		for _, endpoint := range endpoints {
			if isEndpointAccessAuthorized(&endpoint, context.UserID, context.UserMemberships) {
				filteredEndpoints = append(filteredEndpoints, endpoint)
			}
		}
	}

	return filteredEndpoints, nil
}

func isRegistryAccessAuthorized(registry *dockm.Registry, userID dockm.UserID, memberships []dockm.TeamMembership) bool {
	for _, authorizedUserID := range registry.AuthorizedUsers {
		if authorizedUserID == userID {
			return true
		}
	}
	for _, membership := range memberships {
		for _, authorizedTeamID := range registry.AuthorizedTeams {
			if membership.TeamID == authorizedTeamID {
				return true
			}
		}
	}
	return false
}

func isEndpointAccessAuthorized(endpoint *dockm.Endpoint, userID dockm.UserID, memberships []dockm.TeamMembership) bool {
	for _, authorizedUserID := range endpoint.AuthorizedUsers {
		if authorizedUserID == userID {
			return true
		}
	}
	for _, membership := range memberships {
		for _, authorizedTeamID := range endpoint.AuthorizedTeams {
			if membership.TeamID == authorizedTeamID {
				return true
			}
		}
	}
	return false
}
