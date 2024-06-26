package offline

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-roster"
)

// type ListJobsCallback is a function type for custom processing of jobs.
type ListJobsCallback func(context.Context, *Job) error

// type Database is an interface for storing and retrieving `Job` instance for future processing.
// Importantly the interface does not define any logic for how or when jobs are processed.
type Database interface {
	// AddJob adds a `Job` instance to the database.
	AddJob(context.Context, *Job) error
	// Retrieve a specific `Job` instance from the database using its unique identifier.
	GetJob(context.Context, int64) (*Job, error)
	// Update a specific `Job` instance in the database.
	UpdateJob(context.Context, *Job) error
	// Remove a specific `Job` instance from the database.
	RemoveJob(context.Context, *Job) error
	// Prune zero or more `Job` instances matching a specific `Status` type and created date from the database.
	PruneJobs(context.Context, Status, int64) error
	// List all of the `Jobs` in the database.
	ListJobs(context.Context, ListJobsCallback) error
	// Close closes any underlying database connections
	Close(context.Context) error
}

var database_roster roster.Roster

// DatabaseInitializationFunc is a function defined by individual database package and used to create
// an instance of that database
type DatabaseInitializationFunc func(ctx context.Context, uri string) (Database, error)

// RegisterDatabase registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Database` instances by the `NewDatabase` method.
func RegisterDatabase(ctx context.Context, scheme string, init_func DatabaseInitializationFunc) error {

	err := ensureDatabaseRoster()

	if err != nil {
		return err
	}

	return database_roster.Register(ctx, scheme, init_func)
}

func ensureDatabaseRoster() error {

	if database_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		database_roster = r
	}

	return nil
}

// NewDatabase returns a new `Database` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `DatabaseInitializationFunc`
// function used to instantiate the new `Database`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterDatabase` method.
func NewDatabase(ctx context.Context, uri string) (Database, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := database_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(DatabaseInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func Schemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureDatabaseRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range database_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
