package bolt

import (
	"github.com/boltdb/bolt"
	"dolphin/api"
	"dolphin/api/bolt/internal"
)

func (m *Migrator) updateResourceControlsToDBVersion2() error {
	legacyResourceControls, err := m.retrieveLegacyResourceControls()
	if err != nil {
		return err
	}

	for _, resourceControl := range legacyResourceControls {
		resourceControl.SubResourceIDs = []string{}
		resourceControl.TeamAccesses = []dockm.TeamResourceAccess{}

		owner, err := m.UserService.User(resourceControl.OwnerID)
		if err != nil {
			return err
		}

		if owner.Role == dockm.AdministratorRole {
			resourceControl.AdministratorsOnly = true
			resourceControl.UserAccesses = []dockm.UserResourceAccess{}
		} else {
			resourceControl.AdministratorsOnly = false
			userAccess := dockm.UserResourceAccess{
				UserID:      resourceControl.OwnerID,
				AccessLevel: dockm.ReadWriteAccessLevel,
			}
			resourceControl.UserAccesses = []dockm.UserResourceAccess{userAccess}
		}

		err = m.ResourceControlService.CreateResourceControl(&resourceControl)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) updateEndpointsToDBVersion2() error {
	legacyEndpoints, err := m.EndpointService.Endpoints()
	if err != nil {
		return err
	}

	for _, endpoint := range legacyEndpoints {
		endpoint.AuthorizedTeams = []dockm.TeamID{}
		err = m.EndpointService.UpdateEndpoint(endpoint.ID, &endpoint)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) retrieveLegacyResourceControls() ([]dockm.ResourceControl, error) {
	legacyResourceControls := make([]dockm.ResourceControl, 0)
	err := m.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("containerResourceControl"))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var resourceControl dockm.ResourceControl
			err := internal.UnmarshalResourceControl(v, &resourceControl)
			if err != nil {
				return err
			}
			resourceControl.Type = dockm.ContainerResourceControl
			legacyResourceControls = append(legacyResourceControls, resourceControl)
		}

		bucket = tx.Bucket([]byte("serviceResourceControl"))
		cursor = bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var resourceControl dockm.ResourceControl
			err := internal.UnmarshalResourceControl(v, &resourceControl)
			if err != nil {
				return err
			}
			resourceControl.Type = dockm.ServiceResourceControl
			legacyResourceControls = append(legacyResourceControls, resourceControl)
		}

		bucket = tx.Bucket([]byte("volumeResourceControl"))
		cursor = bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var resourceControl dockm.ResourceControl
			err := internal.UnmarshalResourceControl(v, &resourceControl)
			if err != nil {
				return err
			}
			resourceControl.Type = dockm.VolumeResourceControl
			legacyResourceControls = append(legacyResourceControls, resourceControl)
		}
		return nil
	})
	return legacyResourceControls, err
}
