// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PollsCountry is an object representing the database table.
type PollsCountry struct {
	PollCountryID int64 `boil:"poll_country_id" json:"poll_country_id" toml:"poll_country_id" yaml:"poll_country_id"`
	PollID        int64 `boil:"poll_id" json:"poll_id" toml:"poll_id" yaml:"poll_id"`
	CountryID     int64 `boil:"country_id" json:"country_id" toml:"country_id" yaml:"country_id"`

	R *pollsCountryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pollsCountryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PollsCountryColumns = struct {
	PollCountryID string
	PollID        string
	CountryID     string
}{
	PollCountryID: "poll_country_id",
	PollID:        "poll_id",
	CountryID:     "country_id",
}

// PollsCountryRels is where relationship names are stored.
var PollsCountryRels = struct {
	Poll    string
	Country string
}{
	Poll:    "Poll",
	Country: "Country",
}

// pollsCountryR is where relationships are stored.
type pollsCountryR struct {
	Poll    *Poll
	Country *Country
}

// NewStruct creates a new relationship struct
func (*pollsCountryR) NewStruct() *pollsCountryR {
	return &pollsCountryR{}
}

// pollsCountryL is where Load methods for each relationship are stored.
type pollsCountryL struct{}

var (
	pollsCountryColumns               = []string{"poll_country_id", "poll_id", "country_id"}
	pollsCountryColumnsWithoutDefault = []string{"poll_country_id", "poll_id", "country_id"}
	pollsCountryColumnsWithDefault    = []string{}
	pollsCountryPrimaryKeyColumns     = []string{"poll_country_id"}
)

type (
	// PollsCountrySlice is an alias for a slice of pointers to PollsCountry.
	// This should generally be used opposed to []PollsCountry.
	PollsCountrySlice []*PollsCountry
	// PollsCountryHook is the signature for custom PollsCountry hook methods
	PollsCountryHook func(context.Context, boil.ContextExecutor, *PollsCountry) error

	pollsCountryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pollsCountryType                 = reflect.TypeOf(&PollsCountry{})
	pollsCountryMapping              = queries.MakeStructMapping(pollsCountryType)
	pollsCountryPrimaryKeyMapping, _ = queries.BindMapping(pollsCountryType, pollsCountryMapping, pollsCountryPrimaryKeyColumns)
	pollsCountryInsertCacheMut       sync.RWMutex
	pollsCountryInsertCache          = make(map[string]insertCache)
	pollsCountryUpdateCacheMut       sync.RWMutex
	pollsCountryUpdateCache          = make(map[string]updateCache)
	pollsCountryUpsertCacheMut       sync.RWMutex
	pollsCountryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var pollsCountryBeforeInsertHooks []PollsCountryHook
var pollsCountryBeforeUpdateHooks []PollsCountryHook
var pollsCountryBeforeDeleteHooks []PollsCountryHook
var pollsCountryBeforeUpsertHooks []PollsCountryHook

var pollsCountryAfterInsertHooks []PollsCountryHook
var pollsCountryAfterSelectHooks []PollsCountryHook
var pollsCountryAfterUpdateHooks []PollsCountryHook
var pollsCountryAfterDeleteHooks []PollsCountryHook
var pollsCountryAfterUpsertHooks []PollsCountryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PollsCountry) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PollsCountry) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PollsCountry) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PollsCountry) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PollsCountry) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PollsCountry) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PollsCountry) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PollsCountry) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PollsCountry) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsCountryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPollsCountryHook registers your hook function for all future operations.
func AddPollsCountryHook(hookPoint boil.HookPoint, pollsCountryHook PollsCountryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pollsCountryBeforeInsertHooks = append(pollsCountryBeforeInsertHooks, pollsCountryHook)
	case boil.BeforeUpdateHook:
		pollsCountryBeforeUpdateHooks = append(pollsCountryBeforeUpdateHooks, pollsCountryHook)
	case boil.BeforeDeleteHook:
		pollsCountryBeforeDeleteHooks = append(pollsCountryBeforeDeleteHooks, pollsCountryHook)
	case boil.BeforeUpsertHook:
		pollsCountryBeforeUpsertHooks = append(pollsCountryBeforeUpsertHooks, pollsCountryHook)
	case boil.AfterInsertHook:
		pollsCountryAfterInsertHooks = append(pollsCountryAfterInsertHooks, pollsCountryHook)
	case boil.AfterSelectHook:
		pollsCountryAfterSelectHooks = append(pollsCountryAfterSelectHooks, pollsCountryHook)
	case boil.AfterUpdateHook:
		pollsCountryAfterUpdateHooks = append(pollsCountryAfterUpdateHooks, pollsCountryHook)
	case boil.AfterDeleteHook:
		pollsCountryAfterDeleteHooks = append(pollsCountryAfterDeleteHooks, pollsCountryHook)
	case boil.AfterUpsertHook:
		pollsCountryAfterUpsertHooks = append(pollsCountryAfterUpsertHooks, pollsCountryHook)
	}
}

// OneG returns a single pollsCountry record from the query using the global executor.
func (q pollsCountryQuery) OneG(ctx context.Context) (*PollsCountry, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single pollsCountry record from the query.
func (q pollsCountryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PollsCountry, error) {
	o := &PollsCountry{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for polls_country")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all PollsCountry records from the query using the global executor.
func (q pollsCountryQuery) AllG(ctx context.Context) (PollsCountrySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all PollsCountry records from the query.
func (q pollsCountryQuery) All(ctx context.Context, exec boil.ContextExecutor) (PollsCountrySlice, error) {
	var o []*PollsCountry

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PollsCountry slice")
	}

	if len(pollsCountryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all PollsCountry records in the query, and panics on error.
func (q pollsCountryQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all PollsCountry records in the query.
func (q pollsCountryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count polls_country rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q pollsCountryQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q pollsCountryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if polls_country exists")
	}

	return count > 0, nil
}

// Poll pointed to by the foreign key.
func (o *PollsCountry) Poll(mods ...qm.QueryMod) pollQuery {
	queryMods := []qm.QueryMod{
		qm.Where("poll_id=?", o.PollID),
	}

	queryMods = append(queryMods, mods...)

	query := Polls(queryMods...)
	queries.SetFrom(query.Query, "\"polls\"")

	return query
}

// Country pointed to by the foreign key.
func (o *PollsCountry) Country(mods ...qm.QueryMod) countryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("country_id=?", o.CountryID),
	}

	queryMods = append(queryMods, mods...)

	query := Countries(queryMods...)
	queries.SetFrom(query.Query, "\"country\"")

	return query
}

// LoadPoll allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pollsCountryL) LoadPoll(ctx context.Context, e boil.ContextExecutor, singular bool, maybePollsCountry interface{}, mods queries.Applicator) error {
	var slice []*PollsCountry
	var object *PollsCountry

	if singular {
		object = maybePollsCountry.(*PollsCountry)
	} else {
		slice = *maybePollsCountry.(*[]*PollsCountry)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pollsCountryR{}
		}
		args = append(args, object.PollID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pollsCountryR{}
			}

			for _, a := range args {
				if a == obj.PollID {
					continue Outer
				}
			}

			args = append(args, obj.PollID)
		}
	}

	query := NewQuery(qm.From(`polls`), qm.WhereIn(`poll_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Poll")
	}

	var resultSlice []*Poll
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Poll")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for polls")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls")
	}

	if len(pollsCountryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Poll = foreign
		if foreign.R == nil {
			foreign.R = &pollR{}
		}
		foreign.R.PollsCountries = append(foreign.R.PollsCountries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PollID == foreign.PollID {
				local.R.Poll = foreign
				if foreign.R == nil {
					foreign.R = &pollR{}
				}
				foreign.R.PollsCountries = append(foreign.R.PollsCountries, local)
				break
			}
		}
	}

	return nil
}

// LoadCountry allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pollsCountryL) LoadCountry(ctx context.Context, e boil.ContextExecutor, singular bool, maybePollsCountry interface{}, mods queries.Applicator) error {
	var slice []*PollsCountry
	var object *PollsCountry

	if singular {
		object = maybePollsCountry.(*PollsCountry)
	} else {
		slice = *maybePollsCountry.(*[]*PollsCountry)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pollsCountryR{}
		}
		args = append(args, object.CountryID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pollsCountryR{}
			}

			for _, a := range args {
				if a == obj.CountryID {
					continue Outer
				}
			}

			args = append(args, obj.CountryID)
		}
	}

	query := NewQuery(qm.From(`country`), qm.WhereIn(`country_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Country")
	}

	var resultSlice []*Country
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Country")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for country")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for country")
	}

	if len(pollsCountryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Country = foreign
		if foreign.R == nil {
			foreign.R = &countryR{}
		}
		foreign.R.PollsCountries = append(foreign.R.PollsCountries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CountryID == foreign.CountryID {
				local.R.Country = foreign
				if foreign.R == nil {
					foreign.R = &countryR{}
				}
				foreign.R.PollsCountries = append(foreign.R.PollsCountries, local)
				break
			}
		}
	}

	return nil
}

// SetPollG of the pollsCountry to the related item.
// Sets o.R.Poll to related.
// Adds o to related.R.PollsCountries.
// Uses the global database handle.
func (o *PollsCountry) SetPollG(ctx context.Context, insert bool, related *Poll) error {
	return o.SetPoll(ctx, boil.GetContextDB(), insert, related)
}

// SetPoll of the pollsCountry to the related item.
// Sets o.R.Poll to related.
// Adds o to related.R.PollsCountries.
func (o *PollsCountry) SetPoll(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Poll) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"polls_country\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"poll_id"}),
		strmangle.WhereClause("\"", "\"", 2, pollsCountryPrimaryKeyColumns),
	)
	values := []interface{}{related.PollID, o.PollCountryID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PollID = related.PollID
	if o.R == nil {
		o.R = &pollsCountryR{
			Poll: related,
		}
	} else {
		o.R.Poll = related
	}

	if related.R == nil {
		related.R = &pollR{
			PollsCountries: PollsCountrySlice{o},
		}
	} else {
		related.R.PollsCountries = append(related.R.PollsCountries, o)
	}

	return nil
}

// SetCountryG of the pollsCountry to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.PollsCountries.
// Uses the global database handle.
func (o *PollsCountry) SetCountryG(ctx context.Context, insert bool, related *Country) error {
	return o.SetCountry(ctx, boil.GetContextDB(), insert, related)
}

// SetCountry of the pollsCountry to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.PollsCountries.
func (o *PollsCountry) SetCountry(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Country) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"polls_country\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"country_id"}),
		strmangle.WhereClause("\"", "\"", 2, pollsCountryPrimaryKeyColumns),
	)
	values := []interface{}{related.CountryID, o.PollCountryID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CountryID = related.CountryID
	if o.R == nil {
		o.R = &pollsCountryR{
			Country: related,
		}
	} else {
		o.R.Country = related
	}

	if related.R == nil {
		related.R = &countryR{
			PollsCountries: PollsCountrySlice{o},
		}
	} else {
		related.R.PollsCountries = append(related.R.PollsCountries, o)
	}

	return nil
}

// PollsCountries retrieves all the records using an executor.
func PollsCountries(mods ...qm.QueryMod) pollsCountryQuery {
	mods = append(mods, qm.From("\"polls_country\""))
	return pollsCountryQuery{NewQuery(mods...)}
}

// FindPollsCountryG retrieves a single record by ID.
func FindPollsCountryG(ctx context.Context, pollCountryID int64, selectCols ...string) (*PollsCountry, error) {
	return FindPollsCountry(ctx, boil.GetContextDB(), pollCountryID, selectCols...)
}

// FindPollsCountry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPollsCountry(ctx context.Context, exec boil.ContextExecutor, pollCountryID int64, selectCols ...string) (*PollsCountry, error) {
	pollsCountryObj := &PollsCountry{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"polls_country\" where \"poll_country_id\"=$1", sel,
	)

	q := queries.Raw(query, pollCountryID)

	err := q.Bind(ctx, exec, pollsCountryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from polls_country")
	}

	return pollsCountryObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *PollsCountry) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PollsCountry) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls_country provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pollsCountryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pollsCountryInsertCacheMut.RLock()
	cache, cached := pollsCountryInsertCache[key]
	pollsCountryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pollsCountryColumns,
			pollsCountryColumnsWithDefault,
			pollsCountryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pollsCountryType, pollsCountryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pollsCountryType, pollsCountryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"polls_country\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"polls_country\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into polls_country")
	}

	if !cached {
		pollsCountryInsertCacheMut.Lock()
		pollsCountryInsertCache[key] = cache
		pollsCountryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single PollsCountry record using the global executor.
// See Update for more documentation.
func (o *PollsCountry) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the PollsCountry.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PollsCountry) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pollsCountryUpdateCacheMut.RLock()
	cache, cached := pollsCountryUpdateCache[key]
	pollsCountryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pollsCountryColumns,
			pollsCountryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update polls_country, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"polls_country\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pollsCountryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pollsCountryType, pollsCountryMapping, append(wl, pollsCountryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update polls_country row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for polls_country")
	}

	if !cached {
		pollsCountryUpdateCacheMut.Lock()
		pollsCountryUpdateCache[key] = cache
		pollsCountryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pollsCountryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for polls_country")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for polls_country")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PollsCountrySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PollsCountrySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsCountryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"polls_country\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pollsCountryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pollsCountry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pollsCountry")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *PollsCountry) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PollsCountry) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls_country provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pollsCountryColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	pollsCountryUpsertCacheMut.RLock()
	cache, cached := pollsCountryUpsertCache[key]
	pollsCountryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pollsCountryColumns,
			pollsCountryColumnsWithDefault,
			pollsCountryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pollsCountryColumns,
			pollsCountryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert polls_country, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pollsCountryPrimaryKeyColumns))
			copy(conflict, pollsCountryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"polls_country\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pollsCountryType, pollsCountryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pollsCountryType, pollsCountryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert polls_country")
	}

	if !cached {
		pollsCountryUpsertCacheMut.Lock()
		pollsCountryUpsertCache[key] = cache
		pollsCountryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single PollsCountry record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *PollsCountry) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single PollsCountry record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PollsCountry) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PollsCountry provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pollsCountryPrimaryKeyMapping)
	sql := "DELETE FROM \"polls_country\" WHERE \"poll_country_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from polls_country")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for polls_country")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pollsCountryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pollsCountryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from polls_country")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for polls_country")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o PollsCountrySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PollsCountrySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PollsCountry slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(pollsCountryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsCountryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"polls_country\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollsCountryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pollsCountry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for polls_country")
	}

	if len(pollsCountryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *PollsCountry) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no PollsCountry provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PollsCountry) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPollsCountry(ctx, exec, o.PollCountryID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PollsCountrySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty PollsCountrySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PollsCountrySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PollsCountrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsCountryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"polls_country\".* FROM \"polls_country\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollsCountryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PollsCountrySlice")
	}

	*o = slice

	return nil
}

// PollsCountryExistsG checks if the PollsCountry row exists.
func PollsCountryExistsG(ctx context.Context, pollCountryID int64) (bool, error) {
	return PollsCountryExists(ctx, boil.GetContextDB(), pollCountryID)
}

// PollsCountryExists checks if the PollsCountry row exists.
func PollsCountryExists(ctx context.Context, exec boil.ContextExecutor, pollCountryID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"polls_country\" where \"poll_country_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, pollCountryID)
	}

	row := exec.QueryRowContext(ctx, sql, pollCountryID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if polls_country exists")
	}

	return exists, nil
}
