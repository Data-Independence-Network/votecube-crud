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

// State is an object representing the database table.
type State struct {
	StateID       int64  `boil:"state_id" json:"state_id" toml:"state_id" yaml:"state_id"`
	CountryID     int64  `boil:"country_id" json:"country_id" toml:"country_id" yaml:"country_id"`
	TimezoneID    int64  `boil:"timezone_id" json:"timezone_id" toml:"timezone_id" yaml:"timezone_id"`
	StateCode     string `boil:"state_code" json:"state_code" toml:"state_code" yaml:"state_code"`
	StateName     string `boil:"state_name" json:"state_name" toml:"state_name" yaml:"state_name"`
	StateFullName string `boil:"state_full_name" json:"state_full_name" toml:"state_full_name" yaml:"state_full_name"`

	R *stateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StateColumns = struct {
	StateID       string
	CountryID     string
	TimezoneID    string
	StateCode     string
	StateName     string
	StateFullName string
}{
	StateID:       "state_id",
	CountryID:     "country_id",
	TimezoneID:    "timezone_id",
	StateCode:     "state_code",
	StateName:     "state_name",
	StateFullName: "state_full_name",
}

// StateRels is where relationship names are stored.
var StateRels = struct {
	Timezone    string
	Country     string
	Counties    string
	PollsStates string
}{
	Timezone:    "Timezone",
	Country:     "Country",
	Counties:    "Counties",
	PollsStates: "PollsStates",
}

// stateR is where relationships are stored.
type stateR struct {
	Timezone    *Timezone
	Country     *Country
	Counties    CountySlice
	PollsStates PollsStateSlice
}

// NewStruct creates a new relationship struct
func (*stateR) NewStruct() *stateR {
	return &stateR{}
}

// stateL is where Load methods for each relationship are stored.
type stateL struct{}

var (
	stateColumns               = []string{"state_id", "country_id", "timezone_id", "state_code", "state_name", "state_full_name"}
	stateColumnsWithoutDefault = []string{"state_id", "country_id", "timezone_id", "state_code", "state_name", "state_full_name"}
	stateColumnsWithDefault    = []string{}
	statePrimaryKeyColumns     = []string{"state_id"}
)

type (
	// StateSlice is an alias for a slice of pointers to State.
	// This should generally be used opposed to []State.
	StateSlice []*State
	// StateHook is the signature for custom State hook methods
	StateHook func(context.Context, boil.ContextExecutor, *State) error

	stateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stateType                 = reflect.TypeOf(&State{})
	stateMapping              = queries.MakeStructMapping(stateType)
	statePrimaryKeyMapping, _ = queries.BindMapping(stateType, stateMapping, statePrimaryKeyColumns)
	stateInsertCacheMut       sync.RWMutex
	stateInsertCache          = make(map[string]insertCache)
	stateUpdateCacheMut       sync.RWMutex
	stateUpdateCache          = make(map[string]updateCache)
	stateUpsertCacheMut       sync.RWMutex
	stateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var stateBeforeInsertHooks []StateHook
var stateBeforeUpdateHooks []StateHook
var stateBeforeDeleteHooks []StateHook
var stateBeforeUpsertHooks []StateHook

var stateAfterInsertHooks []StateHook
var stateAfterSelectHooks []StateHook
var stateAfterUpdateHooks []StateHook
var stateAfterDeleteHooks []StateHook
var stateAfterUpsertHooks []StateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *State) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *State) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *State) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *State) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *State) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *State) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *State) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *State) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *State) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range stateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStateHook registers your hook function for all future operations.
func AddStateHook(hookPoint boil.HookPoint, stateHook StateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stateBeforeInsertHooks = append(stateBeforeInsertHooks, stateHook)
	case boil.BeforeUpdateHook:
		stateBeforeUpdateHooks = append(stateBeforeUpdateHooks, stateHook)
	case boil.BeforeDeleteHook:
		stateBeforeDeleteHooks = append(stateBeforeDeleteHooks, stateHook)
	case boil.BeforeUpsertHook:
		stateBeforeUpsertHooks = append(stateBeforeUpsertHooks, stateHook)
	case boil.AfterInsertHook:
		stateAfterInsertHooks = append(stateAfterInsertHooks, stateHook)
	case boil.AfterSelectHook:
		stateAfterSelectHooks = append(stateAfterSelectHooks, stateHook)
	case boil.AfterUpdateHook:
		stateAfterUpdateHooks = append(stateAfterUpdateHooks, stateHook)
	case boil.AfterDeleteHook:
		stateAfterDeleteHooks = append(stateAfterDeleteHooks, stateHook)
	case boil.AfterUpsertHook:
		stateAfterUpsertHooks = append(stateAfterUpsertHooks, stateHook)
	}
}

// One returns a single state record from the query.
func (q stateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*State, error) {
	o := &State{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for state")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all State records from the query.
func (q stateQuery) All(ctx context.Context, exec boil.ContextExecutor) (StateSlice, error) {
	var o []*State

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to State slice")
	}

	if len(stateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all State records in the query.
func (q stateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count state rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if state exists")
	}

	return count > 0, nil
}

// Timezone pointed to by the foreign key.
func (o *State) Timezone(mods ...qm.QueryMod) timezoneQuery {
	queryMods := []qm.QueryMod{
		qm.Where("timezone_id=?", o.TimezoneID),
	}

	queryMods = append(queryMods, mods...)

	query := Timezones(queryMods...)
	queries.SetFrom(query.Query, "\"timezone\"")

	return query
}

// Country pointed to by the foreign key.
func (o *State) Country(mods ...qm.QueryMod) countryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("country_id=?", o.CountryID),
	}

	queryMods = append(queryMods, mods...)

	query := Countries(queryMods...)
	queries.SetFrom(query.Query, "\"country\"")

	return query
}

// Counties retrieves all the county's Counties with an executor.
func (o *State) Counties(mods ...qm.QueryMod) countyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"county\".\"state_id\"=?", o.StateID),
	)

	query := Counties(queryMods...)
	queries.SetFrom(query.Query, "\"county\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"county\".*"})
	}

	return query
}

// PollsStates retrieves all the polls_state's PollsStates with an executor.
func (o *State) PollsStates(mods ...qm.QueryMod) pollsStateQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"polls_state\".\"state_id\"=?", o.StateID),
	)

	query := PollsStates(queryMods...)
	queries.SetFrom(query.Query, "\"polls_state\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"polls_state\".*"})
	}

	return query
}

// LoadTimezone allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stateL) LoadTimezone(ctx context.Context, e boil.ContextExecutor, singular bool, maybeState interface{}, mods queries.Applicator) error {
	var slice []*State
	var object *State

	if singular {
		object = maybeState.(*State)
	} else {
		slice = *maybeState.(*[]*State)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stateR{}
		}
		args = append(args, object.TimezoneID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stateR{}
			}

			for _, a := range args {
				if a == obj.TimezoneID {
					continue Outer
				}
			}

			args = append(args, obj.TimezoneID)
		}
	}

	query := NewQuery(qm.From(`timezone`), qm.WhereIn(`timezone_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Timezone")
	}

	var resultSlice []*Timezone
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Timezone")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for timezone")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for timezone")
	}

	if len(stateAfterSelectHooks) != 0 {
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
		object.R.Timezone = foreign
		if foreign.R == nil {
			foreign.R = &timezoneR{}
		}
		foreign.R.States = append(foreign.R.States, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TimezoneID == foreign.TimezoneID {
				local.R.Timezone = foreign
				if foreign.R == nil {
					foreign.R = &timezoneR{}
				}
				foreign.R.States = append(foreign.R.States, local)
				break
			}
		}
	}

	return nil
}

// LoadCountry allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stateL) LoadCountry(ctx context.Context, e boil.ContextExecutor, singular bool, maybeState interface{}, mods queries.Applicator) error {
	var slice []*State
	var object *State

	if singular {
		object = maybeState.(*State)
	} else {
		slice = *maybeState.(*[]*State)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stateR{}
		}
		args = append(args, object.CountryID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stateR{}
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

	if len(stateAfterSelectHooks) != 0 {
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
		foreign.R.States = append(foreign.R.States, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CountryID == foreign.CountryID {
				local.R.Country = foreign
				if foreign.R == nil {
					foreign.R = &countryR{}
				}
				foreign.R.States = append(foreign.R.States, local)
				break
			}
		}
	}

	return nil
}

// LoadCounties allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stateL) LoadCounties(ctx context.Context, e boil.ContextExecutor, singular bool, maybeState interface{}, mods queries.Applicator) error {
	var slice []*State
	var object *State

	if singular {
		object = maybeState.(*State)
	} else {
		slice = *maybeState.(*[]*State)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stateR{}
		}
		args = append(args, object.StateID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stateR{}
			}

			for _, a := range args {
				if a == obj.StateID {
					continue Outer
				}
			}

			args = append(args, obj.StateID)
		}
	}

	query := NewQuery(qm.From(`county`), qm.WhereIn(`state_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load county")
	}

	var resultSlice []*County
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice county")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on county")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for county")
	}

	if len(countyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Counties = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &countyR{}
			}
			foreign.R.State = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StateID == foreign.StateID {
				local.R.Counties = append(local.R.Counties, foreign)
				if foreign.R == nil {
					foreign.R = &countyR{}
				}
				foreign.R.State = local
				break
			}
		}
	}

	return nil
}

// LoadPollsStates allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stateL) LoadPollsStates(ctx context.Context, e boil.ContextExecutor, singular bool, maybeState interface{}, mods queries.Applicator) error {
	var slice []*State
	var object *State

	if singular {
		object = maybeState.(*State)
	} else {
		slice = *maybeState.(*[]*State)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stateR{}
		}
		args = append(args, object.StateID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stateR{}
			}

			for _, a := range args {
				if a == obj.StateID {
					continue Outer
				}
			}

			args = append(args, obj.StateID)
		}
	}

	query := NewQuery(qm.From(`polls_state`), qm.WhereIn(`state_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load polls_state")
	}

	var resultSlice []*PollsState
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice polls_state")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on polls_state")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls_state")
	}

	if len(pollsStateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PollsStates = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pollsStateR{}
			}
			foreign.R.State = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StateID == foreign.StateID {
				local.R.PollsStates = append(local.R.PollsStates, foreign)
				if foreign.R == nil {
					foreign.R = &pollsStateR{}
				}
				foreign.R.State = local
				break
			}
		}
	}

	return nil
}

// SetTimezone of the state to the related item.
// Sets o.R.Timezone to related.
// Adds o to related.R.States.
func (o *State) SetTimezone(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Timezone) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"timezone_id"}),
		strmangle.WhereClause("\"", "\"", 2, statePrimaryKeyColumns),
	)
	values := []interface{}{related.TimezoneID, o.StateID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TimezoneID = related.TimezoneID
	if o.R == nil {
		o.R = &stateR{
			Timezone: related,
		}
	} else {
		o.R.Timezone = related
	}

	if related.R == nil {
		related.R = &timezoneR{
			States: StateSlice{o},
		}
	} else {
		related.R.States = append(related.R.States, o)
	}

	return nil
}

// SetCountry of the state to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.States.
func (o *State) SetCountry(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Country) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"country_id"}),
		strmangle.WhereClause("\"", "\"", 2, statePrimaryKeyColumns),
	)
	values := []interface{}{related.CountryID, o.StateID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CountryID = related.CountryID
	if o.R == nil {
		o.R = &stateR{
			Country: related,
		}
	} else {
		o.R.Country = related
	}

	if related.R == nil {
		related.R = &countryR{
			States: StateSlice{o},
		}
	} else {
		related.R.States = append(related.R.States, o)
	}

	return nil
}

// AddCounties adds the given related objects to the existing relationships
// of the state, optionally inserting them as new records.
// Appends related to o.R.Counties.
// Sets related.R.State appropriately.
func (o *State) AddCounties(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*County) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StateID = o.StateID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"county\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"state_id"}),
				strmangle.WhereClause("\"", "\"", 2, countyPrimaryKeyColumns),
			)
			values := []interface{}{o.StateID, rel.CountyID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StateID = o.StateID
		}
	}

	if o.R == nil {
		o.R = &stateR{
			Counties: related,
		}
	} else {
		o.R.Counties = append(o.R.Counties, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &countyR{
				State: o,
			}
		} else {
			rel.R.State = o
		}
	}
	return nil
}

// AddPollsStates adds the given related objects to the existing relationships
// of the state, optionally inserting them as new records.
// Appends related to o.R.PollsStates.
// Sets related.R.State appropriately.
func (o *State) AddPollsStates(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsState) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StateID = o.StateID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"polls_state\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"state_id"}),
				strmangle.WhereClause("\"", "\"", 2, pollsStatePrimaryKeyColumns),
			)
			values := []interface{}{o.StateID, rel.PollStateID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StateID = o.StateID
		}
	}

	if o.R == nil {
		o.R = &stateR{
			PollsStates: related,
		}
	} else {
		o.R.PollsStates = append(o.R.PollsStates, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pollsStateR{
				State: o,
			}
		} else {
			rel.R.State = o
		}
	}
	return nil
}

// States retrieves all the records using an executor.
func States(mods ...qm.QueryMod) stateQuery {
	mods = append(mods, qm.From("\"state\""))
	return stateQuery{NewQuery(mods...)}
}

// FindState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindState(ctx context.Context, exec boil.ContextExecutor, stateID int64, selectCols ...string) (*State, error) {
	stateObj := &State{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"state\" where \"state_id\"=$1", sel,
	)

	q := queries.Raw(query, stateID)

	err := q.Bind(ctx, exec, stateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from state")
	}

	return stateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *State) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no state provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stateInsertCacheMut.RLock()
	cache, cached := stateInsertCache[key]
	stateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stateColumns,
			stateColumnsWithDefault,
			stateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stateType, stateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"state\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"state\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into state")
	}

	if !cached {
		stateInsertCacheMut.Lock()
		stateInsertCache[key] = cache
		stateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the State.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *State) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stateUpdateCacheMut.RLock()
	cache, cached := stateUpdateCache[key]
	stateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stateColumns,
			statePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update state, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"state\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, statePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, append(wl, statePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update state row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for state")
	}

	if !cached {
		stateUpdateCacheMut.Lock()
		stateUpdateCache[key] = cache
		stateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q stateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for state")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, statePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in state slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all state")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *State) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no state provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stateColumnsWithDefault, o)

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

	stateUpsertCacheMut.RLock()
	cache, cached := stateUpsertCache[key]
	stateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stateColumns,
			stateColumnsWithDefault,
			stateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stateColumns,
			statePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert state, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(statePrimaryKeyColumns))
			copy(conflict, statePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"state\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stateType, stateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stateType, stateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert state")
	}

	if !cached {
		stateUpsertCacheMut.Lock()
		stateUpsertCache[key] = cache
		stateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single State record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *State) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no State provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), statePrimaryKeyMapping)
	sql := "DELETE FROM \"state\" WHERE \"state_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for state")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for state")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no State slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(stateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"state\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, statePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from state slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for state")
	}

	if len(stateAfterDeleteHooks) != 0 {
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
func (o *State) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindState(ctx, exec, o.StateID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"state\".* FROM \"state\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, statePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StateSlice")
	}

	*o = slice

	return nil
}

// StateExists checks if the State row exists.
func StateExists(ctx context.Context, exec boil.ContextExecutor, stateID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"state\" where \"state_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stateID)
	}

	row := exec.QueryRowContext(ctx, sql, stateID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if state exists")
	}

	return exists, nil
}
