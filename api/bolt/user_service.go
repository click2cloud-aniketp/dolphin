package bolt

import (
	"dolphin/api"
	"dolphin/api/bolt/internal"

	"github.com/boltdb/bolt"
)

// UserService represents a service for managing users.
type UserService struct {
	store *Store
}

// User returns a user by ID
func (service *UserService) User(ID dockm.UserID) (*dockm.User, error) {
	var data []byte
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))
		value := bucket.Get(internal.Itob(int(ID)))
		if value == nil {
			return dockm.ErrUserNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)
		return nil
	})
	if err != nil {
		return nil, err
	}

	var user dockm.User
	err = internal.UnmarshalUser(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UserByUsername returns a user by username.
func (service *UserService) UserByUsername(username string) (*dockm.User, error) {
	var user *dockm.User

	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var u dockm.User
			err := internal.UnmarshalUser(v, &u)
			if err != nil {
				return err
			}
			if u.Username == username {
				user = &u
			}
		}

		if user == nil {
			return dockm.ErrUserNotFound
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Users return an array containing all the users.
func (service *UserService) Users() ([]dockm.User, error) {
	var users = make([]dockm.User, 0)
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user dockm.User
			err := internal.UnmarshalUser(v, &user)
			if err != nil {
				return err
			}
			users = append(users, user)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UsersByRole return an array containing all the users with the specified role.
func (service *UserService) UsersByRole(role dockm.UserRole) ([]dockm.User, error) {
	var users = make([]dockm.User, 0)
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user dockm.User
			err := internal.UnmarshalUser(v, &user)
			if err != nil {
				return err
			}
			if user.Role == role {
				users = append(users, user)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser saves a user.
func (service *UserService) UpdateUser(ID dockm.UserID, user *dockm.User) error {
	data, err := internal.MarshalUser(user)
	if err != nil {
		return err
	}

	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))
		err = bucket.Put(internal.Itob(int(ID)), data)

		if err != nil {
			return err
		}
		return nil
	})
}

// CreateUser creates a new user.
func (service *UserService) CreateUser(user *dockm.User) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))

		id, _ := bucket.NextSequence()
		user.ID = dockm.UserID(id)

		data, err := internal.MarshalUser(user)
		if err != nil {
			return err
		}

		err = bucket.Put(internal.Itob(int(user.ID)), data)
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteUser deletes a user.
func (service *UserService) DeleteUser(ID dockm.UserID) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucketName))
		err := bucket.Delete(internal.Itob(int(ID)))
		if err != nil {
			return err
		}
		return nil
	})
}
