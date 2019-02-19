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

// Continent is an object representing the database table.
type Continent struct {
	ContinentID       int64     `boil:"continent_id" json:"continent_id" toml:"continent_id" yaml:"continent_id"`
	ContinentCode     string    `boil:"continent_code" json:"continent_code" toml:"continent_code" yaml:"continent_code"`
	ContinentName     string    `boil:"continent_name" json:"continent_name" toml:"continent_name" yaml:"continent_name"`
	ContinentFullName string    `boil:"continent_full_name" json:"continent_full_name" toml:"continent_full_name" yaml:"continent_full_name"`
	CreatedDate       time.Time `boil:"created_date" json:"created_date" toml:"created_date" yaml:"created_date"`
	ContinentID2      int64     `boil:"continent_id_2" json:"continent_id_2" toml:"continent_id_2" yaml:"continent_id_2"`

	R *continentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L continentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ContinentColumns = struct {
	ContinentID       string
	ContinentCode     string
	ContinentName     string
	ContinentFullName string
	CreatedDate       string
	ContinentID2      string
}{
	ContinentID:       "continent_id",
	ContinentCode:     "continent_code",
	ContinentName:     "continent_name",
	ContinentFullName: "continent_full_name",
	CreatedDate:       "created_date",
	ContinentID2:      "continent_id_2",
}

// ContinentRels is where relationship names are stored.
var ContinentRels = struct {
	Countries       string
	PollsContinents string
}{
	Countries:       "Countries",
	PollsContinents: "PollsContinents",
}

// continentR is where relationships are stored.
type continentR struct {
	Countries       CountrySlice
	PollsContinents PollsContinentSlice
}

// NewStruct creates a new relationship struct
func (*continentR) NewStruct() *continentR {
	return &continentR{}
}

// continentL is where Load methods for each relationship are stored.
type continentL struct{}

var (
	continentColumns               = []string{"continent_id", "continent_code", "continent_name", "continent_full_name", "created_date", "continent_id_2"}
	continentColumnsWithoutDefault = []string{"continent_id", "continent_code", "continent_name", "continent_full_name", "created_date", "continent_id_2"}
	continentColumnsWithDefault    = []string{}
	continentPrimaryKeyColumns     = []string{"continent_id_2"}
)

type (
	// ContinentSlice is an alias for a slice of pointers to Continent.
	// This should generally be used opposed to []Continent.
	ContinentSlice []*Continent
	// ContinentHook is the signature for custom Continent hook methods
	ContinentHook func(context.Context, boil.ContextExecutor, *Continent) error

	continentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	continentType                 = reflect.TypeOf(&Continent{})
	continentMapping              = queries.MakeStructMapping(continentType)
	continentPrimaryKeyMapping, _ = queries.BindMapping(continentType, continentMapping, continentPrimaryKeyColumns)
	continentInsertCacheMut       sync.RWMutex
	continentInsertCache          = make(map[string]insertCache)
	continentUpdateCacheMut       sync.RWMutex
	continentUpdateCache          = make(map[string]updateCache)
	continentUpsertCacheMut       sync.RWMutex
	continentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var continentBeforeInsertHooks []ContinentHook
var continentBeforeUpdateHooks []ContinentHook
var continentBeforeDeleteHooks []ContinentHook
var continentBeforeUpsertHooks []ContinentHook

var continentAfterInsertHooks []ContinentHook
var continentAfterSelectHooks []ContinentHook
var continentAfterUpdateHooks []ContinentHook
var continentAfterDeleteHooks []ContinentHook
var continentAfterUpsertHooks []ContinentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Continent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Continent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Continent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Continent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Continent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Continent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Continent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Continent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Continent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range continentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddContinentHook registers your hook function for all future operations.
func AddContinentHook(hookPoint boil.HookPoint, continentHook ContinentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		continentBeforeInsertHooks = append(continentBeforeInsertHooks, continentHook)
	case boil.BeforeUpdateHook:
		continentBeforeUpdateHooks = append(continentBeforeUpdateHooks, continentHook)
	case boil.BeforeDeleteHook:
		continentBeforeDeleteHooks = append(continentBeforeDeleteHooks, continentHook)
	case boil.BeforeUpsertHook:
		continentBeforeUpsertHooks = append(continentBeforeUpsertHooks, continentHook)
	case boil.AfterInsertHook:
		continentAfterInsertHooks = append(continentAfterInsertHooks, continentHook)
	case boil.AfterSelectHook:
		continentAfterSelectHooks = append(continentAfterSelectHooks, continentHook)
	case boil.AfterUpdateHook:
		continentAfterUpdateHooks = append(continentAfterUpdateHooks, continentHook)
	case boil.AfterDeleteHook:
		continentAfterDeleteHooks = append(continentAfterDeleteHooks, continentHook)
	case boil.AfterUpsertHook:
		continentAfterUpsertHooks = append(continentAfterUpsertHooks, continentHook)
	}
}

// OneG returns a single continent record from the query using the global executor.
func (q continentQuery) OneG(ctx context.Context) (*Continent, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single continent record from the query.
func (q continentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Continent, error) {
	o := &Continent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for continent")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Continent records from the query using the global executor.
func (q continentQuery) AllG(ctx context.Context) (ContinentSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Continent records from the query.
func (q continentQuery) All(ctx context.Context, exec boil.ContextExecutor) (ContinentSlice, error) {
	var o []*Continent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Continent slice")
	}

	if len(continentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Continent records in the query, and panics on error.
func (q continentQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Continent records in the query.
func (q continentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count continent rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q continentQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q continentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if continent exists")
	}

	return count > 0, nil
}

// Countries retrieves all the country's Countries with an executor.
func (o *Continent) Countries(mods ...qm.QueryMod) countryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"country\".\"continent_id\"=?", o.ContinentID2),
	)

	query := Countries(queryMods...)
	queries.SetFrom(query.Query, "\"country\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"country\".*"})
	}

	return query
}

// PollsContinents retrieves all the polls_continent's PollsContinents with an executor.
func (o *Continent) PollsContinents(mods ...qm.QueryMod) pollsContinentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"polls_continent\".\"continent_id\"=?", o.ContinentID2),
	)

	query := PollsContinents(queryMods...)
	queries.SetFrom(query.Query, "\"polls_continent\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"polls_continent\".*"})
	}

	return query
}

// LoadCountries allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (continentL) LoadCountries(ctx context.Context, e boil.ContextExecutor, singular bool, maybeContinent interface{}, mods queries.Applicator) error {
	var slice []*Continent
	var object *Continent

	if singular {
		object = maybeContinent.(*Continent)
	} else {
		slice = *maybeContinent.(*[]*Continent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &continentR{}
		}
		args = append(args, object.ContinentID2)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &continentR{}
			}

			for _, a := range args {
				if a == obj.ContinentID2 {
					continue Outer
				}
			}

			args = append(args, obj.ContinentID2)
		}
	}

	query := NewQuery(qm.From(`country`), qm.WhereIn(`continent_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load country")
	}

	var resultSlice []*Country
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice country")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on country")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for country")
	}

	if len(countryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Countries = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &countryR{}
			}
			foreign.R.Continent = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContinentID2 == foreign.ContinentID {
				local.R.Countries = append(local.R.Countries, foreign)
				if foreign.R == nil {
					foreign.R = &countryR{}
				}
				foreign.R.Continent = local
				break
			}
		}
	}

	return nil
}

// LoadPollsContinents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (continentL) LoadPollsContinents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeContinent interface{}, mods queries.Applicator) error {
	var slice []*Continent
	var object *Continent

	if singular {
		object = maybeContinent.(*Continent)
	} else {
		slice = *maybeContinent.(*[]*Continent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &continentR{}
		}
		args = append(args, object.ContinentID2)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &continentR{}
			}

			for _, a := range args {
				if a == obj.ContinentID2 {
					continue Outer
				}
			}

			args = append(args, obj.ContinentID2)
		}
	}

	query := NewQuery(qm.From(`polls_continent`), qm.WhereIn(`continent_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load polls_continent")
	}

	var resultSlice []*PollsContinent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice polls_continent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on polls_continent")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls_continent")
	}

	if len(pollsContinentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PollsContinents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pollsContinentR{}
			}
			foreign.R.Continent = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContinentID2 == foreign.ContinentID {
				local.R.PollsContinents = append(local.R.PollsContinents, foreign)
				if foreign.R == nil {
					foreign.R = &pollsContinentR{}
				}
				foreign.R.Continent = local
				break
			}
		}
	}

	return nil
}

// AddCountriesG adds the given related objects to the existing relationships
// of the continent, optionally inserting them as new records.
// Appends related to o.R.Countries.
// Sets related.R.Continent appropriately.
// Uses the global database handle.
func (o *Continent) AddCountriesG(ctx context.Context, insert bool, related ...*Country) error {
	return o.AddCountries(ctx, boil.GetContextDB(), insert, related...)
}

// AddCountries adds the given related objects to the existing relationships
// of the continent, optionally inserting them as new records.
// Appends related to o.R.Countries.
// Sets related.R.Continent appropriately.
func (o *Continent) AddCountries(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Country) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ContinentID = o.ContinentID2
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"country\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"continent_id"}),
				strmangle.WhereClause("\"", "\"", 2, countryPrimaryKeyColumns),
			)
			values := []interface{}{o.ContinentID2, rel.CountryID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ContinentID = o.ContinentID2
		}
	}

	if o.R == nil {
		o.R = &continentR{
			Countries: related,
		}
	} else {
		o.R.Countries = append(o.R.Countries, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &countryR{
				Continent: o,
			}
		} else {
			rel.R.Continent = o
		}
	}
	return nil
}

// AddPollsContinentsG adds the given related objects to the existing relationships
// of the continent, optionally inserting them as new records.
// Appends related to o.R.PollsContinents.
// Sets related.R.Continent appropriately.
// Uses the global database handle.
func (o *Continent) AddPollsContinentsG(ctx context.Context, insert bool, related ...*PollsContinent) error {
	return o.AddPollsContinents(ctx, boil.GetContextDB(), insert, related...)
}

// AddPollsContinents adds the given related objects to the existing relationships
// of the continent, optionally inserting them as new records.
// Appends related to o.R.PollsContinents.
// Sets related.R.Continent appropriately.
func (o *Continent) AddPollsContinents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsContinent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ContinentID = o.ContinentID2
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"polls_continent\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"continent_id"}),
				strmangle.WhereClause("\"", "\"", 2, pollsContinentPrimaryKeyColumns),
			)
			values := []interface{}{o.ContinentID2, rel.PollContinentID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ContinentID = o.ContinentID2
		}
	}

	if o.R == nil {
		o.R = &continentR{
			PollsContinents: related,
		}
	} else {
		o.R.PollsContinents = append(o.R.PollsContinents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pollsContinentR{
				Continent: o,
			}
		} else {
			rel.R.Continent = o
		}
	}
	return nil
}

// Continents retrieves all the records using an executor.
func Continents(mods ...qm.QueryMod) continentQuery {
	mods = append(mods, qm.From("\"continent\""))
	return continentQuery{NewQuery(mods...)}
}

// FindContinentG retrieves a single record by ID.
func FindContinentG(ctx context.Context, continentID2 int64, selectCols ...string) (*Continent, error) {
	return FindContinent(ctx, boil.GetContextDB(), continentID2, selectCols...)
}

// FindContinent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindContinent(ctx context.Context, exec boil.ContextExecutor, continentID2 int64, selectCols ...string) (*Continent, error) {
	continentObj := &Continent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"continent\" where \"continent_id_2\"=$1", sel,
	)

	q := queries.Raw(query, continentID2)

	err := q.Bind(ctx, exec, continentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from continent")
	}

	return continentObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Continent) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Continent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no continent provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(continentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	continentInsertCacheMut.RLock()
	cache, cached := continentInsertCache[key]
	continentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			continentColumns,
			continentColumnsWithDefault,
			continentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(continentType, continentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(continentType, continentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"continent\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"continent\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into continent")
	}

	if !cached {
		continentInsertCacheMut.Lock()
		continentInsertCache[key] = cache
		continentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Continent record using the global executor.
// See Update for more documentation.
func (o *Continent) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Continent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Continent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	continentUpdateCacheMut.RLock()
	cache, cached := continentUpdateCache[key]
	continentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			continentColumns,
			continentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update continent, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"continent\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, continentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(continentType, continentMapping, append(wl, continentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update continent row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for continent")
	}

	if !cached {
		continentUpdateCacheMut.Lock()
		continentUpdateCache[key] = cache
		continentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q continentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for continent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for continent")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ContinentSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ContinentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), continentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"continent\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, continentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in continent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all continent")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Continent) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Continent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no continent provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(continentColumnsWithDefault, o)

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

	continentUpsertCacheMut.RLock()
	cache, cached := continentUpsertCache[key]
	continentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			continentColumns,
			continentColumnsWithDefault,
			continentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			continentColumns,
			continentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert continent, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(continentPrimaryKeyColumns))
			copy(conflict, continentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"continent\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(continentType, continentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(continentType, continentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert continent")
	}

	if !cached {
		continentUpsertCacheMut.Lock()
		continentUpsertCache[key] = cache
		continentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Continent record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Continent) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Continent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Continent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Continent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), continentPrimaryKeyMapping)
	sql := "DELETE FROM \"continent\" WHERE \"continent_id_2\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from continent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for continent")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q continentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no continentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from continent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for continent")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ContinentSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ContinentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Continent slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(continentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), continentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"continent\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, continentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from continent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for continent")
	}

	if len(continentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Continent) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Continent provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Continent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindContinent(ctx, exec, o.ContinentID2)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContinentSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ContinentSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContinentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ContinentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), continentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"continent\".* FROM \"continent\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, continentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ContinentSlice")
	}

	*o = slice

	return nil
}

// ContinentExistsG checks if the Continent row exists.
func ContinentExistsG(ctx context.Context, continentID2 int64) (bool, error) {
	return ContinentExists(ctx, boil.GetContextDB(), continentID2)
}

// ContinentExists checks if the Continent row exists.
func ContinentExists(ctx context.Context, exec boil.ContextExecutor, continentID2 int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"continent\" where \"continent_id_2\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, continentID2)
	}

	row := exec.QueryRowContext(ctx, sql, continentID2)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if continent exists")
	}

	return exists, nil
}
