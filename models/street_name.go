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

// StreetName is an object representing the database table.
type StreetName struct {
	StreetNameID int64  `boil:"street_name_id" json:"street_name_id" toml:"street_name_id" yaml:"street_name_id"`
	StreetName   string `boil:"street_name" json:"street_name" toml:"street_name" yaml:"street_name"`

	R *streetNameR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L streetNameL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StreetNameColumns = struct {
	StreetNameID string
	StreetName   string
}{
	StreetNameID: "street_name_id",
	StreetName:   "street_name",
}

// StreetNameRels is where relationship names are stored.
var StreetNameRels = struct {
	Streets string
}{
	Streets: "Streets",
}

// streetNameR is where relationships are stored.
type streetNameR struct {
	Streets StreetSlice
}

// NewStruct creates a new relationship struct
func (*streetNameR) NewStruct() *streetNameR {
	return &streetNameR{}
}

// streetNameL is where Load methods for each relationship are stored.
type streetNameL struct{}

var (
	streetNameColumns               = []string{"street_name_id", "street_name"}
	streetNameColumnsWithoutDefault = []string{"street_name_id", "street_name"}
	streetNameColumnsWithDefault    = []string{}
	streetNamePrimaryKeyColumns     = []string{"street_name_id"}
)

type (
	// StreetNameSlice is an alias for a slice of pointers to StreetName.
	// This should generally be used opposed to []StreetName.
	StreetNameSlice []*StreetName
	// StreetNameHook is the signature for custom StreetName hook methods
	StreetNameHook func(context.Context, boil.ContextExecutor, *StreetName) error

	streetNameQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	streetNameType                 = reflect.TypeOf(&StreetName{})
	streetNameMapping              = queries.MakeStructMapping(streetNameType)
	streetNamePrimaryKeyMapping, _ = queries.BindMapping(streetNameType, streetNameMapping, streetNamePrimaryKeyColumns)
	streetNameInsertCacheMut       sync.RWMutex
	streetNameInsertCache          = make(map[string]insertCache)
	streetNameUpdateCacheMut       sync.RWMutex
	streetNameUpdateCache          = make(map[string]updateCache)
	streetNameUpsertCacheMut       sync.RWMutex
	streetNameUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var streetNameBeforeInsertHooks []StreetNameHook
var streetNameBeforeUpdateHooks []StreetNameHook
var streetNameBeforeDeleteHooks []StreetNameHook
var streetNameBeforeUpsertHooks []StreetNameHook

var streetNameAfterInsertHooks []StreetNameHook
var streetNameAfterSelectHooks []StreetNameHook
var streetNameAfterUpdateHooks []StreetNameHook
var streetNameAfterDeleteHooks []StreetNameHook
var streetNameAfterUpsertHooks []StreetNameHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StreetName) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StreetName) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StreetName) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StreetName) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StreetName) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StreetName) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StreetName) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StreetName) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StreetName) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetNameAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStreetNameHook registers your hook function for all future operations.
func AddStreetNameHook(hookPoint boil.HookPoint, streetNameHook StreetNameHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		streetNameBeforeInsertHooks = append(streetNameBeforeInsertHooks, streetNameHook)
	case boil.BeforeUpdateHook:
		streetNameBeforeUpdateHooks = append(streetNameBeforeUpdateHooks, streetNameHook)
	case boil.BeforeDeleteHook:
		streetNameBeforeDeleteHooks = append(streetNameBeforeDeleteHooks, streetNameHook)
	case boil.BeforeUpsertHook:
		streetNameBeforeUpsertHooks = append(streetNameBeforeUpsertHooks, streetNameHook)
	case boil.AfterInsertHook:
		streetNameAfterInsertHooks = append(streetNameAfterInsertHooks, streetNameHook)
	case boil.AfterSelectHook:
		streetNameAfterSelectHooks = append(streetNameAfterSelectHooks, streetNameHook)
	case boil.AfterUpdateHook:
		streetNameAfterUpdateHooks = append(streetNameAfterUpdateHooks, streetNameHook)
	case boil.AfterDeleteHook:
		streetNameAfterDeleteHooks = append(streetNameAfterDeleteHooks, streetNameHook)
	case boil.AfterUpsertHook:
		streetNameAfterUpsertHooks = append(streetNameAfterUpsertHooks, streetNameHook)
	}
}

// OneG returns a single streetName record from the query using the global executor.
func (q streetNameQuery) OneG(ctx context.Context) (*StreetName, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single streetName record from the query.
func (q streetNameQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StreetName, error) {
	o := &StreetName{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for street_name")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all StreetName records from the query using the global executor.
func (q streetNameQuery) AllG(ctx context.Context) (StreetNameSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all StreetName records from the query.
func (q streetNameQuery) All(ctx context.Context, exec boil.ContextExecutor) (StreetNameSlice, error) {
	var o []*StreetName

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StreetName slice")
	}

	if len(streetNameAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all StreetName records in the query, and panics on error.
func (q streetNameQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all StreetName records in the query.
func (q streetNameQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count street_name rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q streetNameQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q streetNameQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if street_name exists")
	}

	return count > 0, nil
}

// Streets retrieves all the street's Streets with an executor.
func (o *StreetName) Streets(mods ...qm.QueryMod) streetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"street\".\"street_name_id\"=?", o.StreetNameID),
	)

	query := Streets(queryMods...)
	queries.SetFrom(query.Query, "\"street\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"street\".*"})
	}

	return query
}

// LoadStreets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (streetNameL) LoadStreets(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStreetName interface{}, mods queries.Applicator) error {
	var slice []*StreetName
	var object *StreetName

	if singular {
		object = maybeStreetName.(*StreetName)
	} else {
		slice = *maybeStreetName.(*[]*StreetName)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &streetNameR{}
		}
		args = append(args, object.StreetNameID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &streetNameR{}
			}

			for _, a := range args {
				if a == obj.StreetNameID {
					continue Outer
				}
			}

			args = append(args, obj.StreetNameID)
		}
	}

	query := NewQuery(qm.From(`street`), qm.WhereIn(`street_name_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load street")
	}

	var resultSlice []*Street
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice street")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on street")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for street")
	}

	if len(streetAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Streets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &streetR{}
			}
			foreign.R.StreetName = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StreetNameID == foreign.StreetNameID {
				local.R.Streets = append(local.R.Streets, foreign)
				if foreign.R == nil {
					foreign.R = &streetR{}
				}
				foreign.R.StreetName = local
				break
			}
		}
	}

	return nil
}

// AddStreetsG adds the given related objects to the existing relationships
// of the street_name, optionally inserting them as new records.
// Appends related to o.R.Streets.
// Sets related.R.StreetName appropriately.
// Uses the global database handle.
func (o *StreetName) AddStreetsG(ctx context.Context, insert bool, related ...*Street) error {
	return o.AddStreets(ctx, boil.GetContextDB(), insert, related...)
}

// AddStreets adds the given related objects to the existing relationships
// of the street_name, optionally inserting them as new records.
// Appends related to o.R.Streets.
// Sets related.R.StreetName appropriately.
func (o *StreetName) AddStreets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Street) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StreetNameID = o.StreetNameID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"street\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"street_name_id"}),
				strmangle.WhereClause("\"", "\"", 2, streetPrimaryKeyColumns),
			)
			values := []interface{}{o.StreetNameID, rel.StreetID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StreetNameID = o.StreetNameID
		}
	}

	if o.R == nil {
		o.R = &streetNameR{
			Streets: related,
		}
	} else {
		o.R.Streets = append(o.R.Streets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &streetR{
				StreetName: o,
			}
		} else {
			rel.R.StreetName = o
		}
	}
	return nil
}

// StreetNames retrieves all the records using an executor.
func StreetNames(mods ...qm.QueryMod) streetNameQuery {
	mods = append(mods, qm.From("\"street_name\""))
	return streetNameQuery{NewQuery(mods...)}
}

// FindStreetNameG retrieves a single record by ID.
func FindStreetNameG(ctx context.Context, streetNameID int64, selectCols ...string) (*StreetName, error) {
	return FindStreetName(ctx, boil.GetContextDB(), streetNameID, selectCols...)
}

// FindStreetName retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStreetName(ctx context.Context, exec boil.ContextExecutor, streetNameID int64, selectCols ...string) (*StreetName, error) {
	streetNameObj := &StreetName{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"street_name\" where \"street_name_id\"=$1", sel,
	)

	q := queries.Raw(query, streetNameID)

	err := q.Bind(ctx, exec, streetNameObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from street_name")
	}

	return streetNameObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StreetName) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StreetName) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no street_name provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(streetNameColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	streetNameInsertCacheMut.RLock()
	cache, cached := streetNameInsertCache[key]
	streetNameInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			streetNameColumns,
			streetNameColumnsWithDefault,
			streetNameColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(streetNameType, streetNameMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(streetNameType, streetNameMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"street_name\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"street_name\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into street_name")
	}

	if !cached {
		streetNameInsertCacheMut.Lock()
		streetNameInsertCache[key] = cache
		streetNameInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single StreetName record using the global executor.
// See Update for more documentation.
func (o *StreetName) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the StreetName.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StreetName) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	streetNameUpdateCacheMut.RLock()
	cache, cached := streetNameUpdateCache[key]
	streetNameUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			streetNameColumns,
			streetNamePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update street_name, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"street_name\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, streetNamePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(streetNameType, streetNameMapping, append(wl, streetNamePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update street_name row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for street_name")
	}

	if !cached {
		streetNameUpdateCacheMut.Lock()
		streetNameUpdateCache[key] = cache
		streetNameUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q streetNameQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for street_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for street_name")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StreetNameSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StreetNameSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"street_name\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, streetNamePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in streetName slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all streetName")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StreetName) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StreetName) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no street_name provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(streetNameColumnsWithDefault, o)

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

	streetNameUpsertCacheMut.RLock()
	cache, cached := streetNameUpsertCache[key]
	streetNameUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			streetNameColumns,
			streetNameColumnsWithDefault,
			streetNameColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			streetNameColumns,
			streetNamePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert street_name, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(streetNamePrimaryKeyColumns))
			copy(conflict, streetNamePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"street_name\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(streetNameType, streetNameMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(streetNameType, streetNameMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert street_name")
	}

	if !cached {
		streetNameUpsertCacheMut.Lock()
		streetNameUpsertCache[key] = cache
		streetNameUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single StreetName record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StreetName) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single StreetName record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StreetName) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StreetName provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), streetNamePrimaryKeyMapping)
	sql := "DELETE FROM \"street_name\" WHERE \"street_name_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from street_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for street_name")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q streetNameQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no streetNameQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from street_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for street_name")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o StreetNameSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StreetNameSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StreetName slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(streetNameBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"street_name\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, streetNamePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from streetName slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for street_name")
	}

	if len(streetNameAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StreetName) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no StreetName provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StreetName) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStreetName(ctx, exec, o.StreetNameID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StreetNameSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty StreetNameSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StreetNameSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StreetNameSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"street_name\".* FROM \"street_name\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, streetNamePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StreetNameSlice")
	}

	*o = slice

	return nil
}

// StreetNameExistsG checks if the StreetName row exists.
func StreetNameExistsG(ctx context.Context, streetNameID int64) (bool, error) {
	return StreetNameExists(ctx, boil.GetContextDB(), streetNameID)
}

// StreetNameExists checks if the StreetName row exists.
func StreetNameExists(ctx context.Context, exec boil.ContextExecutor, streetNameID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"street_name\" where \"street_name_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, streetNameID)
	}

	row := exec.QueryRowContext(ctx, sql, streetNameID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if street_name exists")
	}

	return exists, nil
}
