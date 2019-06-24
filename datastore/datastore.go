// Package datastore is responsible for data storage.
// Doesn't matter which datastore has been chosen, there is only on requirement:
// it MUST implement the interface that declared on controller layer.
package datastore

type (
	// Storage represents connectino to datastore.
	Storage struct {
		Employee employeeCollection
	}
)

// New retuns created db instance.
func New(conf string) Storage {
	// Initialise new data store using config.
	return Storage{
		Employee: initEmployeeCollection(),
	}
}
