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

// Emoji is an object representing the database table.
type Emoji struct {
	EmojiID  int64  `boil:"emoji_id" json:"emoji_id" toml:"emoji_id" yaml:"emoji_id"`
	CSSClass string `boil:"css_class" json:"css_class" toml:"css_class" yaml:"css_class"`

	R *emojiR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L emojiL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmojiColumns = struct {
	EmojiID  string
	CSSClass string
}{
	EmojiID:  "emoji_id",
	CSSClass: "css_class",
}

// EmojiRels is where relationship names are stored.
var EmojiRels = struct {
	Directions                string
	PollsDimensionsDirections string
}{
	Directions:                "Directions",
	PollsDimensionsDirections: "PollsDimensionsDirections",
}

// emojiR is where relationships are stored.
type emojiR struct {
	Directions                DirectionSlice
	PollsDimensionsDirections PollsDimensionsDirectionSlice
}

// NewStruct creates a new relationship struct
func (*emojiR) NewStruct() *emojiR {
	return &emojiR{}
}

// emojiL is where Load methods for each relationship are stored.
type emojiL struct{}

var (
	emojiColumns               = []string{"emoji_id", "css_class"}
	emojiColumnsWithoutDefault = []string{"emoji_id", "css_class"}
	emojiColumnsWithDefault    = []string{}
	emojiPrimaryKeyColumns     = []string{"emoji_id"}
)

type (
	// EmojiSlice is an alias for a slice of pointers to Emoji.
	// This should generally be used opposed to []Emoji.
	EmojiSlice []*Emoji
	// EmojiHook is the signature for custom Emoji hook methods
	EmojiHook func(context.Context, boil.ContextExecutor, *Emoji) error

	emojiQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	emojiType                 = reflect.TypeOf(&Emoji{})
	emojiMapping              = queries.MakeStructMapping(emojiType)
	emojiPrimaryKeyMapping, _ = queries.BindMapping(emojiType, emojiMapping, emojiPrimaryKeyColumns)
	emojiInsertCacheMut       sync.RWMutex
	emojiInsertCache          = make(map[string]insertCache)
	emojiUpdateCacheMut       sync.RWMutex
	emojiUpdateCache          = make(map[string]updateCache)
	emojiUpsertCacheMut       sync.RWMutex
	emojiUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var emojiBeforeInsertHooks []EmojiHook
var emojiBeforeUpdateHooks []EmojiHook
var emojiBeforeDeleteHooks []EmojiHook
var emojiBeforeUpsertHooks []EmojiHook

var emojiAfterInsertHooks []EmojiHook
var emojiAfterSelectHooks []EmojiHook
var emojiAfterUpdateHooks []EmojiHook
var emojiAfterDeleteHooks []EmojiHook
var emojiAfterUpsertHooks []EmojiHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Emoji) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Emoji) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Emoji) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Emoji) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Emoji) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Emoji) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Emoji) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Emoji) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Emoji) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emojiAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEmojiHook registers your hook function for all future operations.
func AddEmojiHook(hookPoint boil.HookPoint, emojiHook EmojiHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		emojiBeforeInsertHooks = append(emojiBeforeInsertHooks, emojiHook)
	case boil.BeforeUpdateHook:
		emojiBeforeUpdateHooks = append(emojiBeforeUpdateHooks, emojiHook)
	case boil.BeforeDeleteHook:
		emojiBeforeDeleteHooks = append(emojiBeforeDeleteHooks, emojiHook)
	case boil.BeforeUpsertHook:
		emojiBeforeUpsertHooks = append(emojiBeforeUpsertHooks, emojiHook)
	case boil.AfterInsertHook:
		emojiAfterInsertHooks = append(emojiAfterInsertHooks, emojiHook)
	case boil.AfterSelectHook:
		emojiAfterSelectHooks = append(emojiAfterSelectHooks, emojiHook)
	case boil.AfterUpdateHook:
		emojiAfterUpdateHooks = append(emojiAfterUpdateHooks, emojiHook)
	case boil.AfterDeleteHook:
		emojiAfterDeleteHooks = append(emojiAfterDeleteHooks, emojiHook)
	case boil.AfterUpsertHook:
		emojiAfterUpsertHooks = append(emojiAfterUpsertHooks, emojiHook)
	}
}

// One returns a single emoji record from the query.
func (q emojiQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Emoji, error) {
	o := &Emoji{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for emoji")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Emoji records from the query.
func (q emojiQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmojiSlice, error) {
	var o []*Emoji

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Emoji slice")
	}

	if len(emojiAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Emoji records in the query.
func (q emojiQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count emoji rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q emojiQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if emoji exists")
	}

	return count > 0, nil
}

// Directions retrieves all the direction's Directions with an executor.
func (o *Emoji) Directions(mods ...qm.QueryMod) directionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"directions\".\"emoji_id\"=?", o.EmojiID),
	)

	query := Directions(queryMods...)
	queries.SetFrom(query.Query, "\"directions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"directions\".*"})
	}

	return query
}

// PollsDimensionsDirections retrieves all the polls_dimensions_direction's PollsDimensionsDirections with an executor.
func (o *Emoji) PollsDimensionsDirections(mods ...qm.QueryMod) pollsDimensionsDirectionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"polls_dimensions_directions\".\"emoji_id\"=?", o.EmojiID),
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
func (emojiL) LoadDirections(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmoji interface{}, mods queries.Applicator) error {
	var slice []*Emoji
	var object *Emoji

	if singular {
		object = maybeEmoji.(*Emoji)
	} else {
		slice = *maybeEmoji.(*[]*Emoji)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emojiR{}
		}
		args = append(args, object.EmojiID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emojiR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.EmojiID) {
					continue Outer
				}
			}

			args = append(args, obj.EmojiID)
		}
	}

	query := NewQuery(qm.From(`directions`), qm.WhereIn(`emoji_id in ?`, args...))
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
			foreign.R.Emoji = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.EmojiID, foreign.EmojiID) {
				local.R.Directions = append(local.R.Directions, foreign)
				if foreign.R == nil {
					foreign.R = &directionR{}
				}
				foreign.R.Emoji = local
				break
			}
		}
	}

	return nil
}

// LoadPollsDimensionsDirections allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (emojiL) LoadPollsDimensionsDirections(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmoji interface{}, mods queries.Applicator) error {
	var slice []*Emoji
	var object *Emoji

	if singular {
		object = maybeEmoji.(*Emoji)
	} else {
		slice = *maybeEmoji.(*[]*Emoji)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emojiR{}
		}
		args = append(args, object.EmojiID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emojiR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.EmojiID) {
					continue Outer
				}
			}

			args = append(args, obj.EmojiID)
		}
	}

	query := NewQuery(qm.From(`polls_dimensions_directions`), qm.WhereIn(`emoji_id in ?`, args...))
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
			foreign.R.Emoji = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.EmojiID, foreign.EmojiID) {
				local.R.PollsDimensionsDirections = append(local.R.PollsDimensionsDirections, foreign)
				if foreign.R == nil {
					foreign.R = &pollsDimensionsDirectionR{}
				}
				foreign.R.Emoji = local
				break
			}
		}
	}

	return nil
}

// AddDirections adds the given related objects to the existing relationships
// of the emoji, optionally inserting them as new records.
// Appends related to o.R.Directions.
// Sets related.R.Emoji appropriately.
func (o *Emoji) AddDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Direction) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.EmojiID, o.EmojiID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"directions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"emoji_id"}),
				strmangle.WhereClause("\"", "\"", 2, directionPrimaryKeyColumns),
			)
			values := []interface{}{o.EmojiID, rel.DirectionID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.EmojiID, o.EmojiID)
		}
	}

	if o.R == nil {
		o.R = &emojiR{
			Directions: related,
		}
	} else {
		o.R.Directions = append(o.R.Directions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &directionR{
				Emoji: o,
			}
		} else {
			rel.R.Emoji = o
		}
	}
	return nil
}

// SetDirections removes all previously related items of the
// emoji replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Emoji's Directions accordingly.
// Replaces o.R.Directions with related.
// Sets related.R.Emoji's Directions accordingly.
func (o *Emoji) SetDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Direction) error {
	query := "update \"directions\" set \"emoji_id\" = null where \"emoji_id\" = $1"
	values := []interface{}{o.EmojiID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Directions {
			queries.SetScanner(&rel.EmojiID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Emoji = nil
		}

		o.R.Directions = nil
	}
	return o.AddDirections(ctx, exec, insert, related...)
}

// RemoveDirections relationships from objects passed in.
// Removes related items from R.Directions (uses pointer comparison, removal does not keep order)
// Sets related.R.Emoji.
func (o *Emoji) RemoveDirections(ctx context.Context, exec boil.ContextExecutor, related ...*Direction) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.EmojiID, nil)
		if rel.R != nil {
			rel.R.Emoji = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("emoji_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Directions {
			if rel != ri {
				continue
			}

			ln := len(o.R.Directions)
			if ln > 1 && i < ln-1 {
				o.R.Directions[i] = o.R.Directions[ln-1]
			}
			o.R.Directions = o.R.Directions[:ln-1]
			break
		}
	}

	return nil
}

// AddPollsDimensionsDirections adds the given related objects to the existing relationships
// of the emoji, optionally inserting them as new records.
// Appends related to o.R.PollsDimensionsDirections.
// Sets related.R.Emoji appropriately.
func (o *Emoji) AddPollsDimensionsDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsDimensionsDirection) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.EmojiID, o.EmojiID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"polls_dimensions_directions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"emoji_id"}),
				strmangle.WhereClause("\"", "\"", 2, pollsDimensionsDirectionPrimaryKeyColumns),
			)
			values := []interface{}{o.EmojiID, rel.PollDimensionDirectionID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.EmojiID, o.EmojiID)
		}
	}

	if o.R == nil {
		o.R = &emojiR{
			PollsDimensionsDirections: related,
		}
	} else {
		o.R.PollsDimensionsDirections = append(o.R.PollsDimensionsDirections, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pollsDimensionsDirectionR{
				Emoji: o,
			}
		} else {
			rel.R.Emoji = o
		}
	}
	return nil
}

// SetPollsDimensionsDirections removes all previously related items of the
// emoji replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Emoji's PollsDimensionsDirections accordingly.
// Replaces o.R.PollsDimensionsDirections with related.
// Sets related.R.Emoji's PollsDimensionsDirections accordingly.
func (o *Emoji) SetPollsDimensionsDirections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PollsDimensionsDirection) error {
	query := "update \"polls_dimensions_directions\" set \"emoji_id\" = null where \"emoji_id\" = $1"
	values := []interface{}{o.EmojiID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.PollsDimensionsDirections {
			queries.SetScanner(&rel.EmojiID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Emoji = nil
		}

		o.R.PollsDimensionsDirections = nil
	}
	return o.AddPollsDimensionsDirections(ctx, exec, insert, related...)
}

// RemovePollsDimensionsDirections relationships from objects passed in.
// Removes related items from R.PollsDimensionsDirections (uses pointer comparison, removal does not keep order)
// Sets related.R.Emoji.
func (o *Emoji) RemovePollsDimensionsDirections(ctx context.Context, exec boil.ContextExecutor, related ...*PollsDimensionsDirection) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.EmojiID, nil)
		if rel.R != nil {
			rel.R.Emoji = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("emoji_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.PollsDimensionsDirections {
			if rel != ri {
				continue
			}

			ln := len(o.R.PollsDimensionsDirections)
			if ln > 1 && i < ln-1 {
				o.R.PollsDimensionsDirections[i] = o.R.PollsDimensionsDirections[ln-1]
			}
			o.R.PollsDimensionsDirections = o.R.PollsDimensionsDirections[:ln-1]
			break
		}
	}

	return nil
}

// Emojis retrieves all the records using an executor.
func Emojis(mods ...qm.QueryMod) emojiQuery {
	mods = append(mods, qm.From("\"emoji\""))
	return emojiQuery{NewQuery(mods...)}
}

// FindEmoji retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmoji(ctx context.Context, exec boil.ContextExecutor, emojiID int64, selectCols ...string) (*Emoji, error) {
	emojiObj := &Emoji{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"emoji\" where \"emoji_id\"=$1", sel,
	)

	q := queries.Raw(query, emojiID)

	err := q.Bind(ctx, exec, emojiObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from emoji")
	}

	return emojiObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Emoji) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no emoji provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emojiColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	emojiInsertCacheMut.RLock()
	cache, cached := emojiInsertCache[key]
	emojiInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			emojiColumns,
			emojiColumnsWithDefault,
			emojiColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(emojiType, emojiMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(emojiType, emojiMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"emoji\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"emoji\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into emoji")
	}

	if !cached {
		emojiInsertCacheMut.Lock()
		emojiInsertCache[key] = cache
		emojiInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Emoji.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Emoji) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	emojiUpdateCacheMut.RLock()
	cache, cached := emojiUpdateCache[key]
	emojiUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			emojiColumns,
			emojiPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update emoji, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"emoji\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, emojiPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(emojiType, emojiMapping, append(wl, emojiPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update emoji row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for emoji")
	}

	if !cached {
		emojiUpdateCacheMut.Lock()
		emojiUpdateCache[key] = cache
		emojiUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q emojiQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for emoji")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for emoji")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmojiSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emojiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"emoji\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, emojiPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in emoji slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all emoji")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Emoji) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no emoji provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emojiColumnsWithDefault, o)

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

	emojiUpsertCacheMut.RLock()
	cache, cached := emojiUpsertCache[key]
	emojiUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			emojiColumns,
			emojiColumnsWithDefault,
			emojiColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			emojiColumns,
			emojiPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert emoji, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(emojiPrimaryKeyColumns))
			copy(conflict, emojiPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"emoji\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(emojiType, emojiMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(emojiType, emojiMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert emoji")
	}

	if !cached {
		emojiUpsertCacheMut.Lock()
		emojiUpsertCache[key] = cache
		emojiUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Emoji record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Emoji) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Emoji provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), emojiPrimaryKeyMapping)
	sql := "DELETE FROM \"emoji\" WHERE \"emoji_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from emoji")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for emoji")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q emojiQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no emojiQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from emoji")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for emoji")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmojiSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Emoji slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(emojiBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emojiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"emoji\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, emojiPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from emoji slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for emoji")
	}

	if len(emojiAfterDeleteHooks) != 0 {
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
func (o *Emoji) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmoji(ctx, exec, o.EmojiID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmojiSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmojiSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emojiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"emoji\".* FROM \"emoji\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, emojiPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmojiSlice")
	}

	*o = slice

	return nil
}

// EmojiExists checks if the Emoji row exists.
func EmojiExists(ctx context.Context, exec boil.ContextExecutor, emojiID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"emoji\" where \"emoji_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, emojiID)
	}

	row := exec.QueryRowContext(ctx, sql, emojiID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if emoji exists")
	}

	return exists, nil
}
