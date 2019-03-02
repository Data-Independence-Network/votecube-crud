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

// DesignPattern is an object representing the database table.
type DesignPattern struct {
	DesignPatternID   int64  `boil:"design_pattern_id" json:"design_pattern_id" toml:"design_pattern_id" yaml:"design_pattern_id"`
	DesignPatternName string `boil:"design_pattern_name" json:"design_pattern_name" toml:"design_pattern_name" yaml:"design_pattern_name"`

	R *designPatternR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L designPatternL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DesignPatternColumns = struct {
	DesignPatternID   string
	DesignPatternName string
}{
	DesignPatternID:   "design_pattern_id",
	DesignPatternName: "design_pattern_name",
}

// DesignPatternRels is where relationship names are stored.
var DesignPatternRels = struct {
	Directions                string
	PollsDimensionsDirections string
}{
	Directions:                "Directions",
	PollsDimensionsDirections: "PollsDimensionsDirections",
}

// designPatternR is where relationships are stored.
type designPatternR struct {
	Directions                DirectionSlice
	PollsDimensionsDirections PollsDimensionsDirectionSlice
}

// NewStruct creates a new relationship struct
func (*designPatternR) NewStruct() *designPatternR {
	return &designPatternR{}
}

// designPatternL is where Load methods for each relationship are stored.
type designPatternL struct{}

var (
	designPatternColumns               = []string{"design_pattern_id", "design_pattern_name"}
	designPatternColumnsWithoutDefault = []string{"design_pattern_id", "design_pattern_name"}
	designPatternColumnsWithDefault    = []string{}
	designPatternPrimaryKeyColumns     = []string{"design_pattern_id"}
)

type (
	// DesignPatternSlice is an alias for a slice of pointers to DesignPattern.
	// This should generally be used opposed to []DesignPattern.
	DesignPatternSlice []*DesignPattern
	// DesignPatternHook is the signature for custom DesignPattern hook methods
	DesignPatternHook func(context.Context, boil.ContextExecutor, *DesignPattern) error

	designPatternQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	designPatternType                 = reflect.TypeOf(&DesignPattern{})
	designPatternMapping              = queries.MakeStructMapping(designPatternType)
	designPatternPrimaryKeyMapping, _ = queries.BindMapping(designPatternType, designPatternMapping, designPatternPrimaryKeyColumns)
	designPatternInsertCacheMut       sync.RWMutex
	designPatternInsertCache          = make(map[string]insertCache)
	designPatternUpdateCacheMut       sync.RWMutex
	designPatternUpdateCache          = make(map[string]updateCache)
	designPatternUpsertCacheMut       sync.RWMutex
	designPatternUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var designPatternBeforeInsertHooks []DesignPatternHook
var designPatternBeforeUpdateHooks []DesignPatternHook
var designPatternBeforeDeleteHooks []DesignPatternHook
var designPatternBeforeUpsertHooks []DesignPatternHook

var designPatternAfterInsertHooks []DesignPatternHook
var designPatternAfterSelectHooks []DesignPatternHook
var designPatternAfterUpdateHooks []DesignPatternHook
var designPatternAfterDeleteHooks []DesignPatternHook
var designPatternAfterUpsertHooks []DesignPatternHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DesignPattern) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DesignPattern) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DesignPattern) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DesignPattern) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DesignPattern) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DesignPattern) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DesignPattern) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DesignPattern) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DesignPattern) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range designPatternAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDesignPatternHook registers your hook function for all future operations.
func AddDesignPatternHook(hookPoint boil.HookPoint, designPatternHook DesignPatternHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		designPatternBeforeInsertHooks = append(designPatternBeforeInsertHooks, designPatternHook)
	case boil.BeforeUpdateHook:
		designPatternBeforeUpdateHooks = append(designPatternBeforeUpdateHooks, designPatternHook)
	case boil.BeforeDeleteHook:
		designPatternBeforeDeleteHooks = append(designPatternBeforeDeleteHooks, designPatternHook)
	case boil.BeforeUpsertHook:
		designPatternBeforeUpsertHooks = append(designPatternBeforeUpsertHooks, designPatternHook)
	case boil.AfterInsertHook:
		designPatternAfterInsertHooks = append(designPatternAfterInsertHooks, designPatternHook)
	case boil.AfterSelectHook:
		designPatternAfterSelectHooks = append(designPatternAfterSelectHooks, designPatternHook)
	case boil.AfterUpdateHook:
		designPatternAfterUpdateHooks = append(designPatternAfterUpdateHooks, designPatternHook)
	case boil.AfterDeleteHook:
		designPatternAfterDeleteHooks = append(designPatternAfterDeleteHooks, designPatternHook)
	case boil.AfterUpsertHook:
		designPatternAfterUpsertHooks = append(designPatternAfterUpsertHooks, designPatternHook)
	}
}

// OneG returns a single designPattern record from the query using the global executor.
func (q designPatternQuery) OneG(ctx context.Context) (*DesignPattern, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single designPattern record from the query.
func (q designPatternQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DesignPattern, error) {
	o := &DesignPattern{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for design_patterns")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all DesignPattern records from the query using the global executor.
func (q designPatternQuery) AllG(ctx context.Context) (DesignPatternSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all DesignPattern records from the query.
func (q designPatternQuery) All(ctx context.Context, exec boil.ContextExecutor) (DesignPatternSlice, error) {
	var o []*DesignPattern

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DesignPattern slice")
	}

	if len(designPatternAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all DesignPattern records in the query, and panics on error.
func (q designPatternQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all DesignPattern records in the query.
func (q designPatternQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count design_patterns rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q designPatternQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q designPatternQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if design_patterns exists")
	}

	return count > 0, nil
}

// Directions retrieves all the direction's Directions with an executor.
func (o *DesignPattern) Directions(mods ...qm.QueryMod) directionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"directions\".\"design_pattern_id\"=?", o.DesignPatternID),
	)

	query := Directions(queryMods...)
	queries.SetFrom(query.Query, "\"directions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"directions\".*"})
	}

	return query
}

// PollsDimensionsDirections retrieves all the polls_dimensions_direction's PollsDimensionsDirections with an executor.
func (o *DesignPattern) PollsDimensionsDirections(mods ...qm.QueryMod) pollsDimensionsDirectionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"polls_dimensions_directions\".\"design_pattern_id\"=?", o.DesignPatternID),
	)

	query := PollsDimensionsDirections(queryMods...)
	queries.SetFrom(query.Query, "\"polls_dimensions_directions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"polls_dimensions_directions\".*"})
	}

	return query
}

// LoadDirections allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (designPatternL) LoadDirections(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDesignPattern interface{}, mods queries.Applicator) error {
	var slice []*DesignPattern
	var object *DesignPattern

	if singular {
		object = maybeDesignPattern.(*DesignPattern)
	} else {
		slice = *maybeDesignPattern.(*[]*DesignPattern)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &designPatternR{}
		}
		args = append(args, object.DesignPatternID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &designPatternR{}
			}

			for _, a := range args {
				if a == obj.DesignPatternID {
					continue Outer
				}
			}

			args = append(args, obj.DesignPatternID)
		}
	}

	query := NewQuery(qm.From(`directions`), qm.WhereIn(`design_pattern_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load directions")
	}

	var resultSlice []*Direction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice directions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on directions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for directions")
	}

	if len(directionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Directions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &directionR{}
			}
			foreign.R.DesignPattern = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DesignPatternID == foreign.DesignPatternID {
				local.R.Directions = append(local.R.Directions, foreign)
				if foreign.R == nil {
					foreign.R = &directionR{}
				}
				foreign.R.DesignPattern = local
				break
			}
		}
	}

	return nil
}

// LoadPollsDimensionsDirections allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (designPatternL) LoadPollsDimensionsDirections(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDesignPattern interface{}, mods queries.Applicator) error {
	var slice []*DesignPattern
	var object *DesignPattern

	if singular {
		object = maybeDesignPattern.(*DesignPattern)
	} else {
		slice = *maybeDesignPattern.(*[]*DesignPattern)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &designPatternR{}
		}
		args = append(args, object.DesignPatternID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &designPatternR{}
			}

			for _, a := range args {
				if a == obj.DesignPatternID {
					continue Outer
				}
			}

			args = append(args, obj.DesignPatternID)
		}
	}

	query := NewQuery(qm.From(`polls_dimensions_directions`), qm.WhereIn(`design_pattern_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load polls_dimensions_directions")
	}

	var resultSlice []*PollsDimensionsDirection
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice polls_dimensions_directions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on polls_dimensions_directions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls_dimensions_directions")
	}

	if len(pollsDimensionsDirectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PollsDimensionsDirections = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pollsDimensionsDirectionR{}
			}
			foreign.R.DesignPattern = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DesignPatternID == foreign.DesignPatternID {
				local.R.PollsDimensionsDirections = append(local.R.PollsDimensionsDirections, foreign)
				if foreign.R == nil {
					foreign.R = &pollsDimensionsDirectionR{}
				}
				foreign.R.DesignPattern = local
				break
			}
		}
	}

	return nil
}

// AddDirectionsG adds the given related objects to the existing relationships
// of the design_pattern, optionally inserting them as new records.
// Appends related to o.R.Directions.
// Sets related.R.DesignPattern appropriately.
// Uses the global database handle.
func (o *DesignPattern) AddDirectionsG(ctx context.Context, insert bool, related ...*Direction) error {
	return o.AddDirections(ctx, boil.GetContextDB(), insert, related...)
}

// AddDirections adds the given related objects to the existing relationships
// of the design_pattern, optionally inserting them as new records.
// Appends related to o.R.Directions.
// Sets related.R.DesignPattern appropriately.
func (o *DesignPattern) AddDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Direction) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DesignPatternID = o.DesignPatternID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"directions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"design_pattern_id"}),
				strmangle.WhereClause("\"", "\"", 2, directionPrimaryKeyColumns),
			)
			values := []interface{}{o.DesignPatternID, rel.DirectionID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DesignPatternID = o.DesignPatternID
		}
	}

	if o.R == nil {
		o.R = &designPatternR{
			Directions: related,
		}
	} else {
		o.R.Directions = append(o.R.Directions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &directionR{
				DesignPattern: o,
			}
		} else {
			rel.R.DesignPattern = o
		}
	}
	return nil
}

// AddPollsDimensionsDirectionsG adds the given related objects to the existing relationships
// of the design_pattern, optionally inserting them as new records.
// Appends related to o.R.PollsDimensionsDirections.
// Sets related.R.DesignPattern appropriately.
// Uses the global database handle.
func (o *DesignPattern) AddPollsDimensionsDirectionsG(ctx context.Context, insert bool, related ...*PollsDimensionsDirection) error {
	return o.AddPollsDimensionsDirections(ctx, boil.GetContextDB(), insert, related...)
}

// AddPollsDimensionsDirections adds the given related objects to the existing relationships
// of the design_pattern, optionally inserting them as new records.
// Appends related to o.R.PollsDimensionsDirections.
// Sets related.R.DesignPattern appropriately.
func (o *DesignPattern) AddPollsDimensionsDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsDimensionsDirection) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DesignPatternID = o.DesignPatternID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"polls_dimensions_directions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"design_pattern_id"}),
				strmangle.WhereClause("\"", "\"", 2, pollsDimensionsDirectionPrimaryKeyColumns),
			)
			values := []interface{}{o.DesignPatternID, rel.PollDimensionDirectionID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DesignPatternID = o.DesignPatternID
		}
	}

	if o.R == nil {
		o.R = &designPatternR{
			PollsDimensionsDirections: related,
		}
	} else {
		o.R.PollsDimensionsDirections = append(o.R.PollsDimensionsDirections, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pollsDimensionsDirectionR{
				DesignPattern: o,
			}
		} else {
			rel.R.DesignPattern = o
		}
	}
	return nil
}

// DesignPatterns retrieves all the records using an executor.
func DesignPatterns(mods ...qm.QueryMod) designPatternQuery {
	mods = append(mods, qm.From("\"design_patterns\""))
	return designPatternQuery{NewQuery(mods...)}
}

// FindDesignPatternG retrieves a single record by ID.
func FindDesignPatternG(ctx context.Context, designPatternID int64, selectCols ...string) (*DesignPattern, error) {
	return FindDesignPattern(ctx, boil.GetContextDB(), designPatternID, selectCols...)
}

// FindDesignPattern retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDesignPattern(ctx context.Context, exec boil.ContextExecutor, designPatternID int64, selectCols ...string) (*DesignPattern, error) {
	designPatternObj := &DesignPattern{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"design_patterns\" where \"design_pattern_id\"=$1", sel,
	)

	q := queries.Raw(query, designPatternID)

	err := q.Bind(ctx, exec, designPatternObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from design_patterns")
	}

	return designPatternObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DesignPattern) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DesignPattern) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no design_patterns provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(designPatternColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	designPatternInsertCacheMut.RLock()
	cache, cached := designPatternInsertCache[key]
	designPatternInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			designPatternColumns,
			designPatternColumnsWithDefault,
			designPatternColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(designPatternType, designPatternMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(designPatternType, designPatternMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"design_patterns\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"design_patterns\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into design_patterns")
	}

	if !cached {
		designPatternInsertCacheMut.Lock()
		designPatternInsertCache[key] = cache
		designPatternInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single DesignPattern record using the global executor.
// See Update for more documentation.
func (o *DesignPattern) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the DesignPattern.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DesignPattern) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	designPatternUpdateCacheMut.RLock()
	cache, cached := designPatternUpdateCache[key]
	designPatternUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			designPatternColumns,
			designPatternPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update design_patterns, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"design_patterns\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, designPatternPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(designPatternType, designPatternMapping, append(wl, designPatternPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update design_patterns row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for design_patterns")
	}

	if !cached {
		designPatternUpdateCacheMut.Lock()
		designPatternUpdateCache[key] = cache
		designPatternUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q designPatternQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for design_patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for design_patterns")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DesignPatternSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DesignPatternSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), designPatternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"design_patterns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, designPatternPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in designPattern slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all designPattern")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DesignPattern) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DesignPattern) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no design_patterns provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(designPatternColumnsWithDefault, o)

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

	designPatternUpsertCacheMut.RLock()
	cache, cached := designPatternUpsertCache[key]
	designPatternUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			designPatternColumns,
			designPatternColumnsWithDefault,
			designPatternColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			designPatternColumns,
			designPatternPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert design_patterns, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(designPatternPrimaryKeyColumns))
			copy(conflict, designPatternPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"design_patterns\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(designPatternType, designPatternMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(designPatternType, designPatternMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert design_patterns")
	}

	if !cached {
		designPatternUpsertCacheMut.Lock()
		designPatternUpsertCache[key] = cache
		designPatternUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single DesignPattern record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DesignPattern) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single DesignPattern record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DesignPattern) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DesignPattern provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), designPatternPrimaryKeyMapping)
	sql := "DELETE FROM \"design_patterns\" WHERE \"design_pattern_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from design_patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for design_patterns")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q designPatternQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no designPatternQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from design_patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for design_patterns")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o DesignPatternSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DesignPatternSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DesignPattern slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(designPatternBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), designPatternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"design_patterns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, designPatternPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from designPattern slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for design_patterns")
	}

	if len(designPatternAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DesignPattern) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no DesignPattern provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DesignPattern) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDesignPattern(ctx, exec, o.DesignPatternID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DesignPatternSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty DesignPatternSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DesignPatternSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DesignPatternSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), designPatternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"design_patterns\".* FROM \"design_patterns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, designPatternPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DesignPatternSlice")
	}

	*o = slice

	return nil
}

// DesignPatternExistsG checks if the DesignPattern row exists.
func DesignPatternExistsG(ctx context.Context, designPatternID int64) (bool, error) {
	return DesignPatternExists(ctx, boil.GetContextDB(), designPatternID)
}

// DesignPatternExists checks if the DesignPattern row exists.
func DesignPatternExists(ctx context.Context, exec boil.ContextExecutor, designPatternID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"design_patterns\" where \"design_pattern_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, designPatternID)
	}

	row := exec.QueryRowContext(ctx, sql, designPatternID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if design_patterns exists")
	}

	return exists, nil
}