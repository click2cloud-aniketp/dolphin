package bolt

import (
	"dolphin/api"
	"dolphin/api/bolt/internal"

	"github.com/boltdb/bolt"
)

// StackService represents a service for managing stacks.
type StackService struct {
	store *Store
}

// Stack returns a stack object by ID.
func (service *StackService) Stack(ID dockm.StackID) (*dockm.Stack, error) {
	var data []byte
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))
		value := bucket.Get(internal.Itob(int(ID)))
		if value == nil {
			return dockm.ErrStackNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)
		return nil
	})
	if err != nil {
		return nil, err
	}

	var stack dockm.Stack
	err = internal.UnmarshalStack(data, &stack)
	if err != nil {
		return nil, err
	}
	return &stack, nil
}

// Stacks returns an array containing all the stacks.
func (service *StackService) Stacks() ([]dockm.Stack, error) {
	var stacks = make([]dockm.Stack, 0)
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var stack dockm.Stack
			err := internal.UnmarshalStack(v, &stack)
			if err != nil {
				return err
			}
			stacks = append(stacks, stack)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return stacks, nil
}

// StacksByEndpointID return an array containing all the stacks related to the specified endpoint ID.
func (service *StackService) StacksByEndpointID(id dockm.EndpointID) ([]dockm.Stack, error) {
	var stacks = make([]dockm.Stack, 0)
	err := service.store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var stack dockm.Stack
			err := internal.UnmarshalStack(v, &stack)
			if err != nil {
				return err
			}
			if stack.EndpointID == id {
				stacks = append(stacks, stack)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return stacks, nil
}

// CreateStack creates a new stack.
func (service *StackService) CreateStack(stack *dockm.Stack) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))

		id, _ := bucket.NextSequence()
		stack.ID = dockm.StackID(id)

		data, err := internal.MarshalStack(stack)
		if err != nil {
			return err
		}

		err = bucket.Put(internal.Itob(int(stack.ID)), data)
		if err != nil {
			return err
		}
		return nil
	})
}

// UpdateStack updates an stack.
func (service *StackService) UpdateStack(ID dockm.StackID, stack *dockm.Stack) error {
	data, err := internal.MarshalStack(stack)
	if err != nil {
		return err
	}

	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))
		err = bucket.Put(internal.Itob(int(ID)), data)
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteStack deletes an stack.
func (service *StackService) DeleteStack(ID dockm.StackID) error {
	return service.store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(stackBucketName))
		err := bucket.Delete(internal.Itob(int(ID)))
		if err != nil {
			return err
		}
		return nil
	})
}
