package bolt

import (
	"dolphin/api"
	"dolphin/api/bolt/internal"

	"github.com/boltdb/bolt"
)

// TeamService represents a service for managing teams.
type TeamService struct {
	store *Store
}

// Team returns a Team by ID
func (service *TeamService) Team(ID dockm.TeamID) (*dockm.Team, error) {
	var data []byte
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))
		value := bucket.Get(internal.Itob(int(ID)))
		if value == nil {
			return dockm.ErrTeamNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)
		return nil
	})
	if err != nil {
		return nil, err
	}

	var team dockm.Team
	err = internal.UnmarshalTeam(data, &team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// TeamByName returns a team by name.
func (service *TeamService) TeamByName(name string) (*dockm.Team, error) {
	var team *dockm.Team

	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var t dockm.Team
			err := internal.UnmarshalTeam(v, &t)
			if err != nil {
				return err
			}
			if t.Name == name {
				team = &t
			}
		}

		if team == nil {
			return dockm.ErrTeamNotFound
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return team, nil
}

// Teams return an array containing all the teams.
func (service *TeamService) Teams() ([]dockm.Team, error) {
	var teams = make([]dockm.Team, 0)
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var team dockm.Team
			err := internal.UnmarshalTeam(v, &team)
			if err != nil {
				return err
			}
			teams = append(teams, team)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// UpdateTeam saves a Team.
func (service *TeamService) UpdateTeam(ID dockm.TeamID, team *dockm.Team) error {
	data, err := internal.MarshalTeam(team)
	if err != nil {
		return err
	}

	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))
		err = bucket.Put(internal.Itob(int(ID)), data)

		if err != nil {
			return err
		}
		return nil
	})
}

// CreateTeam creates a new Team.
func (service *TeamService) CreateTeam(team *dockm.Team) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))

		id, _ := bucket.NextSequence()
		team.ID = dockm.TeamID(id)

		data, err := internal.MarshalTeam(team)
		if err != nil {
			return err
		}

		err = bucket.Put(internal.Itob(int(team.ID)), data)
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteTeam deletes a Team.
func (service *TeamService) DeleteTeam(ID dockm.TeamID) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(teamBucketName))
		err := bucket.Delete(internal.Itob(int(ID)))
		if err != nil {
			return err
		}
		return nil
	})
}
