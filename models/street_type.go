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

// StreetType is an object representing the database table.
type StreetType struct {
	StreetTypeID int64  `boil:"street_type_id" json:"street_type_id" toml:"street_type_id" yaml:"street_type_id"`
	StreetType   string `boil:"street_type" json:"street_type" toml:"street_type" yaml:"street_type"`

	R *streetTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L streetTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StreetTypeColumns = struct {
	StreetTypeID string
	StreetType   string
}{
	StreetTypeID: "street_type_id",
	StreetType:   "street_type",
}

// StreetTypeRels is where relationship names are stored.
var StreetTypeRels = struct {
	Streets string
}{
	Streets: "Streets",
}

// streetTypeR is where relationships are stored.
type streetTypeR struct {
	Streets StreetSlice
}

// NewStruct creates a new relationship struct
func (*streetTypeR) NewStruct() *streetTypeR {
	return &streetTypeR{}
}

// streetTypeL is where Load methods for each relationship are stored.
type streetTypeL struct{}

var (
	streetTypeColumns               = []string{"street_type_id", "street_type"}
	streetTypeColumnsWithoutDefault = []string{"street_type_id", "street_type"}
	streetTypeColumnsWithDefault    = []string{}
	streetTypePrimaryKeyColumns     = []string{"street_type_id"}
)

type (
	// StreetTypeSlice is an alias for a slice of pointers to StreetType.
	// This should generally be used opposed to []StreetType.
	StreetTypeSlice []*StreetType
	// StreetTypeHook is the signature for custom StreetType hook methods
	StreetTypeHook func(context.Context, boil.ContextExecutor, *StreetType) error

	streetTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	streetTypeType                 = reflect.TypeOf(&StreetType{})
	streetTypeMapping              = queries.MakeStructMapping(streetTypeType)
	streetTypePrimaryKeyMapping, _ = queries.BindMapping(streetTypeType, streetTypeMapping, streetTypePrimaryKeyColumns)
	streetTypeInsertCacheMut       sync.RWMutex
	streetTypeInsertCache          = make(map[string]insertCache)
	streetTypeUpdateCacheMut       sync.RWMutex
	streetTypeUpdateCache          = make(map[string]updateCache)
	streetTypeUpsertCacheMut       sync.RWMutex
	streetTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var streetTypeBeforeInsertHooks []StreetTypeHook
var streetTypeBeforeUpdateHooks []StreetTypeHook
var streetTypeBeforeDeleteHooks []StreetTypeHook
var streetTypeBeforeUpsertHooks []StreetTypeHook

var streetTypeAfterInsertHooks []StreetTypeHook
var streetTypeAfterSelectHooks []StreetTypeHook
var streetTypeAfterUpdateHooks []StreetTypeHook
var streetTypeAfterDeleteHooks []StreetTypeHook
var streetTypeAfterUpsertHooks []StreetTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StreetType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StreetType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StreetType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StreetType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StreetType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StreetType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StreetType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StreetType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StreetType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range streetTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStreetTypeHook registers your hook function for all future operations.
func AddStreetTypeHook(hookPoint boil.HookPoint, streetTypeHook StreetTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		streetTypeBeforeInsertHooks = append(streetTypeBeforeInsertHooks, streetTypeHook)
	case boil.BeforeUpdateHook:
		streetTypeBeforeUpdateHooks = append(streetTypeBeforeUpdateHooks, streetTypeHook)
	case boil.BeforeDeleteHook:
		streetTypeBeforeDeleteHooks = append(streetTypeBeforeDeleteHooks, streetTypeHook)
	case boil.BeforeUpsertHook:
		streetTypeBeforeUpsertHooks = append(streetTypeBeforeUpsertHooks, streetTypeHook)
	case boil.AfterInsertHook:
		streetTypeAfterInsertHooks = append(streetTypeAfterInsertHooks, streetTypeHook)
	case boil.AfterSelectHook:
		streetTypeAfterSelectHooks = append(streetTypeAfterSelectHooks, streetTypeHook)
	case boil.AfterUpdateHook:
		streetTypeAfterUpdateHooks = append(streetTypeAfterUpdateHooks, streetTypeHook)
	case boil.AfterDeleteHook:
		streetTypeAfterDeleteHooks = append(streetTypeAfterDeleteHooks, streetTypeHook)
	case boil.AfterUpsertHook:
		streetTypeAfterUpsertHooks = append(streetTypeAfterUpsertHooks, streetTypeHook)
	}
}

// One returns a single streetType record from the query.
func (q streetTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StreetType, error) {
	o := &StreetType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for street_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all StreetType records from the query.
func (q streetTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (StreetTypeSlice, error) {
	var o []*StreetType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StreetType slice")
	}

	if len(streetTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all StreetType records in the query.
func (q streetTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count street_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q streetTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if street_type exists")
	}

	return count > 0, nil
}

// Streets retrieves all the street's Streets with an executor.
func (o *StreetType) Streets(mods ...qm.QueryMod) streetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"street\".\"street_type_id\"=?", o.StreetTypeID),
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
func (streetTypeL) LoadStreets(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStreetType interface{}, mods queries.Applicator) error {
	var slice []*StreetType
	var object *StreetType

	if singular {
		object = maybeStreetType.(*StreetType)
	} else {
		slice = *maybeStreetType.(*[]*StreetType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &streetTypeR{}
		}
		args = append(args, object.StreetTypeID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &streetTypeR{}
			}

			for _, a := range args {
				if a == obj.StreetTypeID {
					continue Outer
				}
			}

			args = append(args, obj.StreetTypeID)
		}
	}

	query := NewQuery(qm.From(`street`), qm.WhereIn(`street_type_id in ?`, args...))
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
			foreign.R.StreetType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StreetTypeID == foreign.StreetTypeID {
				local.R.Streets = append(local.R.Streets, foreign)
				if foreign.R == nil {
					foreign.R = &streetR{}
				}
				foreign.R.StreetType = local
				break
			}
		}
	}

	return nil
}

// AddStreets adds the given related objects to the existing relationships
// of the street_type, optionally inserting them as new records.
// Appends related to o.R.Streets.
// Sets related.R.StreetType appropriately.
func (o *StreetType) AddStreets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Street) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StreetTypeID = o.StreetTypeID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"street\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"street_type_id"}),
				strmangle.WhereClause("\"", "\"", 2, streetPrimaryKeyColumns),
			)
			values := []interface{}{o.StreetTypeID, rel.StreetID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StreetTypeID = o.StreetTypeID
		}
	}

	if o.R == nil {
		o.R = &streetTypeR{
			Streets: related,
		}
	} else {
		o.R.Streets = append(o.R.Streets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &streetR{
				StreetType: o,
			}
		} else {
			rel.R.StreetType = o
		}
	}
	return nil
}

// StreetTypes retrieves all the records using an executor.
func StreetTypes(mods ...qm.QueryMod) streetTypeQuery {
	mods = append(mods, qm.From("\"street_type\""))
	return streetTypeQuery{NewQuery(mods...)}
}

// FindStreetType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStreetType(ctx context.Context, exec boil.ContextExecutor, streetTypeID int64, selectCols ...string) (*StreetType, error) {
	streetTypeObj := &StreetType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"street_type\" where \"street_type_id\"=$1", sel,
	)

	q := queries.Raw(query, streetTypeID)

	err := q.Bind(ctx, exec, streetTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from street_type")
	}

	return streetTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StreetType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no street_type provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(streetTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	streetTypeInsertCacheMut.RLock()
	cache, cached := streetTypeInsertCache[key]
	streetTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			streetTypeColumns,
			streetTypeColumnsWithDefault,
			streetTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(streetTypeType, streetTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(streetTypeType, streetTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"street_type\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"street_type\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into street_type")
	}

	if !cached {
		streetTypeInsertCacheMut.Lock()
		streetTypeInsertCache[key] = cache
		streetTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the StreetType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StreetType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	streetTypeUpdateCacheMut.RLock()
	cache, cached := streetTypeUpdateCache[key]
	streetTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			streetTypeColumns,
			streetTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update street_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"street_type\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, streetTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(streetTypeType, streetTypeMapping, append(wl, streetTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update street_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for street_type")
	}

	if !cached {
		streetTypeUpdateCacheMut.Lock()
		streetTypeUpdateCache[key] = cache
		streetTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q streetTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for street_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for street_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StreetTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"street_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, streetTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in streetType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all streetType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StreetType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no street_type provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(streetTypeColumnsWithDefault, o)

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

	streetTypeUpsertCacheMut.RLock()
	cache, cached := streetTypeUpsertCache[key]
	streetTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			streetTypeColumns,
			streetTypeColumnsWithDefault,
			streetTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			streetTypeColumns,
			streetTypePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert street_type, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(streetTypePrimaryKeyColumns))
			copy(conflict, streetTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"street_type\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(streetTypeType, streetTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(streetTypeType, streetTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert street_type")
	}

	if !cached {
		streetTypeUpsertCacheMut.Lock()
		streetTypeUpsertCache[key] = cache
		streetTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single StreetType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StreetType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StreetType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), streetTypePrimaryKeyMapping)
	sql := "DELETE FROM \"street_type\" WHERE \"street_type_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from street_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for street_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q streetTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no streetTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from street_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for street_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StreetTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StreetType slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(streetTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"street_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, streetTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from streetType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for street_type")
	}

	if len(streetTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StreetType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStreetType(ctx, exec, o.StreetTypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StreetTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StreetTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), streetTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"street_type\".* FROM \"street_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, streetTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StreetTypeSlice")
	}

	*o = slice

	return nil
}

// StreetTypeExists checks if the StreetType row exists.
func StreetTypeExists(ctx context.Context, exec boil.ContextExecutor, streetTypeID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"street_type\" where \"street_type_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, streetTypeID)
	}

	row := exec.QueryRowContext(ctx, sql, streetTypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if street_type exists")
	}

	return exists, nil
}
