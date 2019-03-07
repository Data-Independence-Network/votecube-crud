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

// Honor is an object representing the database table.
type Honor struct {
	HonorID   int64  `boil:"honor_id" json:"honor_id" toml:"honor_id" yaml:"honor_id"`
	HonorName string `boil:"honor_name" json:"honor_name" toml:"honor_name" yaml:"honor_name"`

	R *honorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L honorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HonorColumns = struct {
	HonorID   string
	HonorName string
}{
	HonorID:   "honor_id",
	HonorName: "honor_name",
}

// HonorRels is where relationship names are stored.
var HonorRels = struct {
	UserPersonalInfoHonors string
}{
	UserPersonalInfoHonors: "UserPersonalInfoHonors",
}

// honorR is where relationships are stored.
type honorR struct {
	UserPersonalInfoHonors UserPersonalInfoHonorSlice
}

// NewStruct creates a new relationship struct
func (*honorR) NewStruct() *honorR {
	return &honorR{}
}

// honorL is where Load methods for each relationship are stored.
type honorL struct{}

var (
	honorColumns               = []string{"honor_id", "honor_name"}
	honorColumnsWithoutDefault = []string{"honor_id", "honor_name"}
	honorColumnsWithDefault    = []string{}
	honorPrimaryKeyColumns     = []string{"honor_id"}
)

type (
	// HonorSlice is an alias for a slice of pointers to Honor.
	// This should generally be used opposed to []Honor.
	HonorSlice []*Honor
	// HonorHook is the signature for custom Honor hook methods
	HonorHook func(context.Context, boil.ContextExecutor, *Honor) error

	honorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	honorType                 = reflect.TypeOf(&Honor{})
	honorMapping              = queries.MakeStructMapping(honorType)
	honorPrimaryKeyMapping, _ = queries.BindMapping(honorType, honorMapping, honorPrimaryKeyColumns)
	honorInsertCacheMut       sync.RWMutex
	honorInsertCache          = make(map[string]insertCache)
	honorUpdateCacheMut       sync.RWMutex
	honorUpdateCache          = make(map[string]updateCache)
	honorUpsertCacheMut       sync.RWMutex
	honorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var honorBeforeInsertHooks []HonorHook
var honorBeforeUpdateHooks []HonorHook
var honorBeforeDeleteHooks []HonorHook
var honorBeforeUpsertHooks []HonorHook

var honorAfterInsertHooks []HonorHook
var honorAfterSelectHooks []HonorHook
var honorAfterUpdateHooks []HonorHook
var honorAfterDeleteHooks []HonorHook
var honorAfterUpsertHooks []HonorHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Honor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Honor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Honor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Honor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Honor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Honor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Honor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Honor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Honor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range honorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHonorHook registers your hook function for all future operations.
func AddHonorHook(hookPoint boil.HookPoint, honorHook HonorHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		honorBeforeInsertHooks = append(honorBeforeInsertHooks, honorHook)
	case boil.BeforeUpdateHook:
		honorBeforeUpdateHooks = append(honorBeforeUpdateHooks, honorHook)
	case boil.BeforeDeleteHook:
		honorBeforeDeleteHooks = append(honorBeforeDeleteHooks, honorHook)
	case boil.BeforeUpsertHook:
		honorBeforeUpsertHooks = append(honorBeforeUpsertHooks, honorHook)
	case boil.AfterInsertHook:
		honorAfterInsertHooks = append(honorAfterInsertHooks, honorHook)
	case boil.AfterSelectHook:
		honorAfterSelectHooks = append(honorAfterSelectHooks, honorHook)
	case boil.AfterUpdateHook:
		honorAfterUpdateHooks = append(honorAfterUpdateHooks, honorHook)
	case boil.AfterDeleteHook:
		honorAfterDeleteHooks = append(honorAfterDeleteHooks, honorHook)
	case boil.AfterUpsertHook:
		honorAfterUpsertHooks = append(honorAfterUpsertHooks, honorHook)
	}
}

// One returns a single honor record from the query.
func (q honorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Honor, error) {
	o := &Honor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for honors")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Honor records from the query.
func (q honorQuery) All(ctx context.Context, exec boil.ContextExecutor) (HonorSlice, error) {
	var o []*Honor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Honor slice")
	}

	if len(honorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Honor records in the query.
func (q honorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count honors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q honorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if honors exists")
	}

	return count > 0, nil
}

// UserPersonalInfoHonors retrieves all the user_personal_info_honor's UserPersonalInfoHonors with an executor.
func (o *Honor) UserPersonalInfoHonors(mods ...qm.QueryMod) userPersonalInfoHonorQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_personal_info_honors\".\"honor_id\"=?", o.HonorID),
	)

	query := UserPersonalInfoHonors(queryMods...)
	queries.SetFrom(query.Query, "\"user_personal_info_honors\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_personal_info_honors\".*"})
	}

	return query
}

// LoadUserPersonalInfoHonors allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (honorL) LoadUserPersonalInfoHonors(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHonor interface{}, mods queries.Applicator) error {
	var slice []*Honor
	var object *Honor

	if singular {
		object = maybeHonor.(*Honor)
	} else {
		slice = *maybeHonor.(*[]*Honor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &honorR{}
		}
		args = append(args, object.HonorID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &honorR{}
			}

			for _, a := range args {
				if a == obj.HonorID {
					continue Outer
				}
			}

			args = append(args, obj.HonorID)
		}
	}

	query := NewQuery(qm.From(`user_personal_info_honors`), qm.WhereIn(`honor_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_personal_info_honors")
	}

	var resultSlice []*UserPersonalInfoHonor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_personal_info_honors")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_personal_info_honors")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_personal_info_honors")
	}

	if len(userPersonalInfoHonorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserPersonalInfoHonors = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userPersonalInfoHonorR{}
			}
			foreign.R.Honor = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.HonorID == foreign.HonorID {
				local.R.UserPersonalInfoHonors = append(local.R.UserPersonalInfoHonors, foreign)
				if foreign.R == nil {
					foreign.R = &userPersonalInfoHonorR{}
				}
				foreign.R.Honor = local
				break
			}
		}
	}

	return nil
}

// AddUserPersonalInfoHonors adds the given related objects to the existing relationships
// of the honor, optionally inserting them as new records.
// Appends related to o.R.UserPersonalInfoHonors.
// Sets related.R.Honor appropriately.
func (o *Honor) AddUserPersonalInfoHonors(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserPersonalInfoHonor) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.HonorID = o.HonorID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_personal_info_honors\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"honor_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPersonalInfoHonorPrimaryKeyColumns),
			)
			values := []interface{}{o.HonorID, rel.UserPersonalInfoHonorID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.HonorID = o.HonorID
		}
	}

	if o.R == nil {
		o.R = &honorR{
			UserPersonalInfoHonors: related,
		}
	} else {
		o.R.UserPersonalInfoHonors = append(o.R.UserPersonalInfoHonors, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userPersonalInfoHonorR{
				Honor: o,
			}
		} else {
			rel.R.Honor = o
		}
	}
	return nil
}

// Honors retrieves all the records using an executor.
func Honors(mods ...qm.QueryMod) honorQuery {
	mods = append(mods, qm.From("\"honors\""))
	return honorQuery{NewQuery(mods...)}
}

// FindHonor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHonor(ctx context.Context, exec boil.ContextExecutor, honorID int64, selectCols ...string) (*Honor, error) {
	honorObj := &Honor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"honors\" where \"honor_id\"=$1", sel,
	)

	q := queries.Raw(query, honorID)

	err := q.Bind(ctx, exec, honorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from honors")
	}

	return honorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Honor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no honors provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(honorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	honorInsertCacheMut.RLock()
	cache, cached := honorInsertCache[key]
	honorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			honorColumns,
			honorColumnsWithDefault,
			honorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(honorType, honorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(honorType, honorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"honors\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"honors\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into honors")
	}

	if !cached {
		honorInsertCacheMut.Lock()
		honorInsertCache[key] = cache
		honorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Honor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Honor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	honorUpdateCacheMut.RLock()
	cache, cached := honorUpdateCache[key]
	honorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			honorColumns,
			honorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update honors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"honors\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, honorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(honorType, honorMapping, append(wl, honorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update honors row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for honors")
	}

	if !cached {
		honorUpdateCacheMut.Lock()
		honorUpdateCache[key] = cache
		honorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q honorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for honors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for honors")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HonorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), honorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"honors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, honorPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in honor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all honor")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Honor) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no honors provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(honorColumnsWithDefault, o)

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

	honorUpsertCacheMut.RLock()
	cache, cached := honorUpsertCache[key]
	honorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			honorColumns,
			honorColumnsWithDefault,
			honorColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			honorColumns,
			honorPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert honors, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(honorPrimaryKeyColumns))
			copy(conflict, honorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"honors\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(honorType, honorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(honorType, honorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert honors")
	}

	if !cached {
		honorUpsertCacheMut.Lock()
		honorUpsertCache[key] = cache
		honorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Honor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Honor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Honor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), honorPrimaryKeyMapping)
	sql := "DELETE FROM \"honors\" WHERE \"honor_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from honors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for honors")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q honorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no honorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from honors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for honors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HonorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Honor slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(honorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), honorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"honors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, honorPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from honor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for honors")
	}

	if len(honorAfterDeleteHooks) != 0 {
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
func (o *Honor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHonor(ctx, exec, o.HonorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HonorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HonorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), honorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"honors\".* FROM \"honors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, honorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HonorSlice")
	}

	*o = slice

	return nil
}

// HonorExists checks if the Honor row exists.
func HonorExists(ctx context.Context, exec boil.ContextExecutor, honorID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"honors\" where \"honor_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, honorID)
	}

	row := exec.QueryRowContext(ctx, sql, honorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if honors exists")
	}

	return exists, nil
}
