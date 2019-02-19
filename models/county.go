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

// County is an object representing the database table.
type County struct {
	CountyID   int64  `boil:"county_id" json:"county_id" toml:"county_id" yaml:"county_id"`
	StateID    int64  `boil:"state_id" json:"state_id" toml:"state_id" yaml:"state_id"`
	CountyCode string `boil:"county_code" json:"county_code" toml:"county_code" yaml:"county_code"`
	CountyName string `boil:"county_name" json:"county_name" toml:"county_name" yaml:"county_name"`

	R *countyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L countyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CountyColumns = struct {
	CountyID   string
	StateID    string
	CountyCode string
	CountyName string
}{
	CountyID:   "county_id",
	StateID:    "state_id",
	CountyCode: "county_code",
	CountyName: "county_name",
}

// CountyRels is where relationship names are stored.
var CountyRels = struct {
	State         string
	PollsCounties string
	Towns         string
}{
	State:         "State",
	PollsCounties: "PollsCounties",
	Towns:         "Towns",
}

// countyR is where relationships are stored.
type countyR struct {
	State         *State
	PollsCounties PollsCountySlice
	Towns         TownSlice
}

// NewStruct creates a new relationship struct
func (*countyR) NewStruct() *countyR {
	return &countyR{}
}

// countyL is where Load methods for each relationship are stored.
type countyL struct{}

var (
	countyColumns               = []string{"county_id", "state_id", "county_code", "county_name"}
	countyColumnsWithoutDefault = []string{"county_id", "state_id", "county_code", "county_name"}
	countyColumnsWithDefault    = []string{}
	countyPrimaryKeyColumns     = []string{"county_id"}
)

type (
	// CountySlice is an alias for a slice of pointers to County.
	// This should generally be used opposed to []County.
	CountySlice []*County
	// CountyHook is the signature for custom County hook methods
	CountyHook func(context.Context, boil.ContextExecutor, *County) error

	countyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	countyType                 = reflect.TypeOf(&County{})
	countyMapping              = queries.MakeStructMapping(countyType)
	countyPrimaryKeyMapping, _ = queries.BindMapping(countyType, countyMapping, countyPrimaryKeyColumns)
	countyInsertCacheMut       sync.RWMutex
	countyInsertCache          = make(map[string]insertCache)
	countyUpdateCacheMut       sync.RWMutex
	countyUpdateCache          = make(map[string]updateCache)
	countyUpsertCacheMut       sync.RWMutex
	countyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var countyBeforeInsertHooks []CountyHook
var countyBeforeUpdateHooks []CountyHook
var countyBeforeDeleteHooks []CountyHook
var countyBeforeUpsertHooks []CountyHook

var countyAfterInsertHooks []CountyHook
var countyAfterSelectHooks []CountyHook
var countyAfterUpdateHooks []CountyHook
var countyAfterDeleteHooks []CountyHook
var countyAfterUpsertHooks []CountyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *County) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *County) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *County) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *County) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *County) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *County) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *County) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *County) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *County) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range countyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCountyHook registers your hook function for all future operations.
func AddCountyHook(hookPoint boil.HookPoint, countyHook CountyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		countyBeforeInsertHooks = append(countyBeforeInsertHooks, countyHook)
	case boil.BeforeUpdateHook:
		countyBeforeUpdateHooks = append(countyBeforeUpdateHooks, countyHook)
	case boil.BeforeDeleteHook:
		countyBeforeDeleteHooks = append(countyBeforeDeleteHooks, countyHook)
	case boil.BeforeUpsertHook:
		countyBeforeUpsertHooks = append(countyBeforeUpsertHooks, countyHook)
	case boil.AfterInsertHook:
		countyAfterInsertHooks = append(countyAfterInsertHooks, countyHook)
	case boil.AfterSelectHook:
		countyAfterSelectHooks = append(countyAfterSelectHooks, countyHook)
	case boil.AfterUpdateHook:
		countyAfterUpdateHooks = append(countyAfterUpdateHooks, countyHook)
	case boil.AfterDeleteHook:
		countyAfterDeleteHooks = append(countyAfterDeleteHooks, countyHook)
	case boil.AfterUpsertHook:
		countyAfterUpsertHooks = append(countyAfterUpsertHooks, countyHook)
	}
}

// OneG returns a single county record from the query using the global executor.
func (q countyQuery) OneG(ctx context.Context) (*County, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single county record from the query.
func (q countyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*County, error) {
	o := &County{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for county")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all County records from the query using the global executor.
func (q countyQuery) AllG(ctx context.Context) (CountySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all County records from the query.
func (q countyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CountySlice, error) {
	var o []*County

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to County slice")
	}

	if len(countyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all County records in the query, and panics on error.
func (q countyQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all County records in the query.
func (q countyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count county rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q countyQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q countyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if county exists")
	}

	return count > 0, nil
}

// State pointed to by the foreign key.
func (o *County) State(mods ...qm.QueryMod) stateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("state_id=?", o.StateID),
	}

	queryMods = append(queryMods, mods...)

	query := States(queryMods...)
	queries.SetFrom(query.Query, "\"state\"")

	return query
}

// PollsCounties retrieves all the polls_county's PollsCounties with an executor.
func (o *County) PollsCounties(mods ...qm.QueryMod) pollsCountyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"polls_county\".\"county_id\"=?", o.CountyID),
	)

	query := PollsCounties(queryMods...)
	queries.SetFrom(query.Query, "\"polls_county\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"polls_county\".*"})
	}

	return query
}

// Towns retrieves all the town's Towns with an executor.
func (o *County) Towns(mods ...qm.QueryMod) townQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"town\".\"county_id\"=?", o.CountyID),
	)

	query := Towns(queryMods...)
	queries.SetFrom(query.Query, "\"town\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"town\".*"})
	}

	return query
}

// LoadState allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (countyL) LoadState(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCounty interface{}, mods queries.Applicator) error {
	var slice []*County
	var object *County

	if singular {
		object = maybeCounty.(*County)
	} else {
		slice = *maybeCounty.(*[]*County)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &countyR{}
		}
		args = append(args, object.StateID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &countyR{}
			}

			for _, a := range args {
				if a == obj.StateID {
					continue Outer
				}
			}

			args = append(args, obj.StateID)
		}
	}

	query := NewQuery(qm.From(`state`), qm.WhereIn(`state_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load State")
	}

	var resultSlice []*State
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice State")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for state")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for state")
	}

	if len(countyAfterSelectHooks) != 0 {
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
		object.R.State = foreign
		if foreign.R == nil {
			foreign.R = &stateR{}
		}
		foreign.R.Counties = append(foreign.R.Counties, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StateID == foreign.StateID {
				local.R.State = foreign
				if foreign.R == nil {
					foreign.R = &stateR{}
				}
				foreign.R.Counties = append(foreign.R.Counties, local)
				break
			}
		}
	}

	return nil
}

// LoadPollsCounties allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (countyL) LoadPollsCounties(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCounty interface{}, mods queries.Applicator) error {
	var slice []*County
	var object *County

	if singular {
		object = maybeCounty.(*County)
	} else {
		slice = *maybeCounty.(*[]*County)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &countyR{}
		}
		args = append(args, object.CountyID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &countyR{}
			}

			for _, a := range args {
				if a == obj.CountyID {
					continue Outer
				}
			}

			args = append(args, obj.CountyID)
		}
	}

	query := NewQuery(qm.From(`polls_county`), qm.WhereIn(`county_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load polls_county")
	}

	var resultSlice []*PollsCounty
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice polls_county")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on polls_county")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls_county")
	}

	if len(pollsCountyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PollsCounties = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pollsCountyR{}
			}
			foreign.R.County = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CountyID == foreign.CountyID {
				local.R.PollsCounties = append(local.R.PollsCounties, foreign)
				if foreign.R == nil {
					foreign.R = &pollsCountyR{}
				}
				foreign.R.County = local
				break
			}
		}
	}

	return nil
}

// LoadTowns allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (countyL) LoadTowns(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCounty interface{}, mods queries.Applicator) error {
	var slice []*County
	var object *County

	if singular {
		object = maybeCounty.(*County)
	} else {
		slice = *maybeCounty.(*[]*County)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &countyR{}
		}
		args = append(args, object.CountyID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &countyR{}
			}

			for _, a := range args {
				if a == obj.CountyID {
					continue Outer
				}
			}

			args = append(args, obj.CountyID)
		}
	}

	query := NewQuery(qm.From(`town`), qm.WhereIn(`county_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load town")
	}

	var resultSlice []*Town
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice town")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on town")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for town")
	}

	if len(townAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Towns = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &townR{}
			}
			foreign.R.County = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CountyID == foreign.CountyID {
				local.R.Towns = append(local.R.Towns, foreign)
				if foreign.R == nil {
					foreign.R = &townR{}
				}
				foreign.R.County = local
				break
			}
		}
	}

	return nil
}

// SetStateG of the county to the related item.
// Sets o.R.State to related.
// Adds o to related.R.Counties.
// Uses the global database handle.
func (o *County) SetStateG(ctx context.Context, insert bool, related *State) error {
	return o.SetState(ctx, boil.GetContextDB(), insert, related)
}

// SetState of the county to the related item.
// Sets o.R.State to related.
// Adds o to related.R.Counties.
func (o *County) SetState(ctx context.Context, exec boil.ContextExecutor, insert bool, related *State) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"county\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"state_id"}),
		strmangle.WhereClause("\"", "\"", 2, countyPrimaryKeyColumns),
	)
	values := []interface{}{related.StateID, o.CountyID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StateID = related.StateID
	if o.R == nil {
		o.R = &countyR{
			State: related,
		}
	} else {
		o.R.State = related
	}

	if related.R == nil {
		related.R = &stateR{
			Counties: CountySlice{o},
		}
	} else {
		related.R.Counties = append(related.R.Counties, o)
	}

	return nil
}

// AddPollsCountiesG adds the given related objects to the existing relationships
// of the county, optionally inserting them as new records.
// Appends related to o.R.PollsCounties.
// Sets related.R.County appropriately.
// Uses the global database handle.
func (o *County) AddPollsCountiesG(ctx context.Context, insert bool, related ...*PollsCounty) error {
	return o.AddPollsCounties(ctx, boil.GetContextDB(), insert, related...)
}

// AddPollsCounties adds the given related objects to the existing relationships
// of the county, optionally inserting them as new records.
// Appends related to o.R.PollsCounties.
// Sets related.R.County appropriately.
func (o *County) AddPollsCounties(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsCounty) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CountyID = o.CountyID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"polls_county\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"county_id"}),
				strmangle.WhereClause("\"", "\"", 2, pollsCountyPrimaryKeyColumns),
			)
			values := []interface{}{o.CountyID, rel.PollCountyID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CountyID = o.CountyID
		}
	}

	if o.R == nil {
		o.R = &countyR{
			PollsCounties: related,
		}
	} else {
		o.R.PollsCounties = append(o.R.PollsCounties, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pollsCountyR{
				County: o,
			}
		} else {
			rel.R.County = o
		}
	}
	return nil
}

// AddTownsG adds the given related objects to the existing relationships
// of the county, optionally inserting them as new records.
// Appends related to o.R.Towns.
// Sets related.R.County appropriately.
// Uses the global database handle.
func (o *County) AddTownsG(ctx context.Context, insert bool, related ...*Town) error {
	return o.AddTowns(ctx, boil.GetContextDB(), insert, related...)
}

// AddTowns adds the given related objects to the existing relationships
// of the county, optionally inserting them as new records.
// Appends related to o.R.Towns.
// Sets related.R.County appropriately.
func (o *County) AddTowns(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Town) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CountyID = o.CountyID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"town\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"county_id"}),
				strmangle.WhereClause("\"", "\"", 2, townPrimaryKeyColumns),
			)
			values := []interface{}{o.CountyID, rel.TownID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CountyID = o.CountyID
		}
	}

	if o.R == nil {
		o.R = &countyR{
			Towns: related,
		}
	} else {
		o.R.Towns = append(o.R.Towns, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &townR{
				County: o,
			}
		} else {
			rel.R.County = o
		}
	}
	return nil
}

// Counties retrieves all the records using an executor.
func Counties(mods ...qm.QueryMod) countyQuery {
	mods = append(mods, qm.From("\"county\""))
	return countyQuery{NewQuery(mods...)}
}

// FindCountyG retrieves a single record by ID.
func FindCountyG(ctx context.Context, countyID int64, selectCols ...string) (*County, error) {
	return FindCounty(ctx, boil.GetContextDB(), countyID, selectCols...)
}

// FindCounty retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCounty(ctx context.Context, exec boil.ContextExecutor, countyID int64, selectCols ...string) (*County, error) {
	countyObj := &County{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"county\" where \"county_id\"=$1", sel,
	)

	q := queries.Raw(query, countyID)

	err := q.Bind(ctx, exec, countyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from county")
	}

	return countyObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *County) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *County) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no county provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	countyInsertCacheMut.RLock()
	cache, cached := countyInsertCache[key]
	countyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			countyColumns,
			countyColumnsWithDefault,
			countyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(countyType, countyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(countyType, countyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"county\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"county\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into county")
	}

	if !cached {
		countyInsertCacheMut.Lock()
		countyInsertCache[key] = cache
		countyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single County record using the global executor.
// See Update for more documentation.
func (o *County) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the County.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *County) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	countyUpdateCacheMut.RLock()
	cache, cached := countyUpdateCache[key]
	countyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			countyColumns,
			countyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update county, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"county\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, countyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(countyType, countyMapping, append(wl, countyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update county row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for county")
	}

	if !cached {
		countyUpdateCacheMut.Lock()
		countyUpdateCache[key] = cache
		countyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q countyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for county")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for county")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CountySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CountySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"county\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, countyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in county slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all county")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *County) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *County) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no county provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countyColumnsWithDefault, o)

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

	countyUpsertCacheMut.RLock()
	cache, cached := countyUpsertCache[key]
	countyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			countyColumns,
			countyColumnsWithDefault,
			countyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			countyColumns,
			countyPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert county, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(countyPrimaryKeyColumns))
			copy(conflict, countyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"county\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(countyType, countyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(countyType, countyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert county")
	}

	if !cached {
		countyUpsertCacheMut.Lock()
		countyUpsertCache[key] = cache
		countyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single County record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *County) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single County record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *County) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no County provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), countyPrimaryKeyMapping)
	sql := "DELETE FROM \"county\" WHERE \"county_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from county")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for county")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q countyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no countyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from county")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for county")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CountySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CountySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no County slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(countyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"county\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from county slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for county")
	}

	if len(countyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *County) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no County provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *County) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCounty(ctx, exec, o.CountyID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CountySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CountySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CountySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CountySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"county\".* FROM \"county\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CountySlice")
	}

	*o = slice

	return nil
}

// CountyExistsG checks if the County row exists.
func CountyExistsG(ctx context.Context, countyID int64) (bool, error) {
	return CountyExists(ctx, boil.GetContextDB(), countyID)
}

// CountyExists checks if the County row exists.
func CountyExists(ctx context.Context, exec boil.ContextExecutor, countyID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"county\" where \"county_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, countyID)
	}

	row := exec.QueryRowContext(ctx, sql, countyID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if county exists")
	}

	return exists, nil
}
