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

// PollsMessage is an object representing the database table.
type PollsMessage struct {
	PollMessageID int64 `boil:"poll_message_id" json:"poll_message_id" toml:"poll_message_id" yaml:"poll_message_id"`
	MessageID     int64 `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	PollID        int64 `boil:"poll_id" json:"poll_id" toml:"poll_id" yaml:"poll_id"`

	R *pollsMessageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pollsMessageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PollsMessageColumns = struct {
	PollMessageID string
	MessageID     string
	PollID        string
}{
	PollMessageID: "poll_message_id",
	MessageID:     "message_id",
	PollID:        "poll_id",
}

// PollsMessageRels is where relationship names are stored.
var PollsMessageRels = struct {
	Poll    string
	Message string
}{
	Poll:    "Poll",
	Message: "Message",
}

// pollsMessageR is where relationships are stored.
type pollsMessageR struct {
	Poll    *Poll
	Message *Message
}

// NewStruct creates a new relationship struct
func (*pollsMessageR) NewStruct() *pollsMessageR {
	return &pollsMessageR{}
}

// pollsMessageL is where Load methods for each relationship are stored.
type pollsMessageL struct{}

var (
	pollsMessageColumns               = []string{"poll_message_id", "message_id", "poll_id"}
	pollsMessageColumnsWithoutDefault = []string{"poll_message_id", "message_id", "poll_id"}
	pollsMessageColumnsWithDefault    = []string{}
	pollsMessagePrimaryKeyColumns     = []string{"poll_message_id"}
)

type (
	// PollsMessageSlice is an alias for a slice of pointers to PollsMessage.
	// This should generally be used opposed to []PollsMessage.
	PollsMessageSlice []*PollsMessage
	// PollsMessageHook is the signature for custom PollsMessage hook methods
	PollsMessageHook func(context.Context, boil.ContextExecutor, *PollsMessage) error

	pollsMessageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pollsMessageType                 = reflect.TypeOf(&PollsMessage{})
	pollsMessageMapping              = queries.MakeStructMapping(pollsMessageType)
	pollsMessagePrimaryKeyMapping, _ = queries.BindMapping(pollsMessageType, pollsMessageMapping, pollsMessagePrimaryKeyColumns)
	pollsMessageInsertCacheMut       sync.RWMutex
	pollsMessageInsertCache          = make(map[string]insertCache)
	pollsMessageUpdateCacheMut       sync.RWMutex
	pollsMessageUpdateCache          = make(map[string]updateCache)
	pollsMessageUpsertCacheMut       sync.RWMutex
	pollsMessageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var pollsMessageBeforeInsertHooks []PollsMessageHook
var pollsMessageBeforeUpdateHooks []PollsMessageHook
var pollsMessageBeforeDeleteHooks []PollsMessageHook
var pollsMessageBeforeUpsertHooks []PollsMessageHook

var pollsMessageAfterInsertHooks []PollsMessageHook
var pollsMessageAfterSelectHooks []PollsMessageHook
var pollsMessageAfterUpdateHooks []PollsMessageHook
var pollsMessageAfterDeleteHooks []PollsMessageHook
var pollsMessageAfterUpsertHooks []PollsMessageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PollsMessage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PollsMessage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PollsMessage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PollsMessage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PollsMessage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PollsMessage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PollsMessage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PollsMessage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PollsMessage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pollsMessageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPollsMessageHook registers your hook function for all future operations.
func AddPollsMessageHook(hookPoint boil.HookPoint, pollsMessageHook PollsMessageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pollsMessageBeforeInsertHooks = append(pollsMessageBeforeInsertHooks, pollsMessageHook)
	case boil.BeforeUpdateHook:
		pollsMessageBeforeUpdateHooks = append(pollsMessageBeforeUpdateHooks, pollsMessageHook)
	case boil.BeforeDeleteHook:
		pollsMessageBeforeDeleteHooks = append(pollsMessageBeforeDeleteHooks, pollsMessageHook)
	case boil.BeforeUpsertHook:
		pollsMessageBeforeUpsertHooks = append(pollsMessageBeforeUpsertHooks, pollsMessageHook)
	case boil.AfterInsertHook:
		pollsMessageAfterInsertHooks = append(pollsMessageAfterInsertHooks, pollsMessageHook)
	case boil.AfterSelectHook:
		pollsMessageAfterSelectHooks = append(pollsMessageAfterSelectHooks, pollsMessageHook)
	case boil.AfterUpdateHook:
		pollsMessageAfterUpdateHooks = append(pollsMessageAfterUpdateHooks, pollsMessageHook)
	case boil.AfterDeleteHook:
		pollsMessageAfterDeleteHooks = append(pollsMessageAfterDeleteHooks, pollsMessageHook)
	case boil.AfterUpsertHook:
		pollsMessageAfterUpsertHooks = append(pollsMessageAfterUpsertHooks, pollsMessageHook)
	}
}

// One returns a single pollsMessage record from the query.
func (q pollsMessageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PollsMessage, error) {
	o := &PollsMessage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for polls_messages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PollsMessage records from the query.
func (q pollsMessageQuery) All(ctx context.Context, exec boil.ContextExecutor) (PollsMessageSlice, error) {
	var o []*PollsMessage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PollsMessage slice")
	}

	if len(pollsMessageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PollsMessage records in the query.
func (q pollsMessageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count polls_messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pollsMessageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if polls_messages exists")
	}

	return count > 0, nil
}

// Poll pointed to by the foreign key.
func (o *PollsMessage) Poll(mods ...qm.QueryMod) pollQuery {
	queryMods := []qm.QueryMod{
		qm.Where("poll_id=?", o.PollID),
	}

	queryMods = append(queryMods, mods...)

	query := Polls(queryMods...)
	queries.SetFrom(query.Query, "\"polls\"")

	return query
}

// Message pointed to by the foreign key.
func (o *PollsMessage) Message(mods ...qm.QueryMod) messageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("message_id=?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	query := Messages(queryMods...)
	queries.SetFrom(query.Query, "\"messages\"")

	return query
}

// LoadPoll allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pollsMessageL) LoadPoll(ctx context.Context, e boil.ContextExecutor, singular bool, maybePollsMessage interface{}, mods queries.Applicator) error {
	var slice []*PollsMessage
	var object *PollsMessage

	if singular {
		object = maybePollsMessage.(*PollsMessage)
	} else {
		slice = *maybePollsMessage.(*[]*PollsMessage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pollsMessageR{}
		}
		args = append(args, object.PollID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pollsMessageR{}
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

	if len(pollsMessageAfterSelectHooks) != 0 {
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
		foreign.R.PollsMessages = append(foreign.R.PollsMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PollID == foreign.PollID {
				local.R.Poll = foreign
				if foreign.R == nil {
					foreign.R = &pollR{}
				}
				foreign.R.PollsMessages = append(foreign.R.PollsMessages, local)
				break
			}
		}
	}

	return nil
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pollsMessageL) LoadMessage(ctx context.Context, e boil.ContextExecutor, singular bool, maybePollsMessage interface{}, mods queries.Applicator) error {
	var slice []*PollsMessage
	var object *PollsMessage

	if singular {
		object = maybePollsMessage.(*PollsMessage)
	} else {
		slice = *maybePollsMessage.(*[]*PollsMessage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pollsMessageR{}
		}
		args = append(args, object.MessageID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pollsMessageR{}
			}

			for _, a := range args {
				if a == obj.MessageID {
					continue Outer
				}
			}

			args = append(args, obj.MessageID)
		}
	}

	query := NewQuery(qm.From(`messages`), qm.WhereIn(`message_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Message")
	}

	var resultSlice []*Message
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Message")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for messages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for messages")
	}

	if len(pollsMessageAfterSelectHooks) != 0 {
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
		object.R.Message = foreign
		if foreign.R == nil {
			foreign.R = &messageR{}
		}
		foreign.R.PollsMessages = append(foreign.R.PollsMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MessageID == foreign.MessageID {
				local.R.Message = foreign
				if foreign.R == nil {
					foreign.R = &messageR{}
				}
				foreign.R.PollsMessages = append(foreign.R.PollsMessages, local)
				break
			}
		}
	}

	return nil
}

// SetPoll of the pollsMessage to the related item.
// Sets o.R.Poll to related.
// Adds o to related.R.PollsMessages.
func (o *PollsMessage) SetPoll(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Poll) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"polls_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"poll_id"}),
		strmangle.WhereClause("\"", "\"", 2, pollsMessagePrimaryKeyColumns),
	)
	values := []interface{}{related.PollID, o.PollMessageID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PollID = related.PollID
	if o.R == nil {
		o.R = &pollsMessageR{
			Poll: related,
		}
	} else {
		o.R.Poll = related
	}

	if related.R == nil {
		related.R = &pollR{
			PollsMessages: PollsMessageSlice{o},
		}
	} else {
		related.R.PollsMessages = append(related.R.PollsMessages, o)
	}

	return nil
}

// SetMessage of the pollsMessage to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.PollsMessages.
func (o *PollsMessage) SetMessage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Message) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"polls_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, pollsMessagePrimaryKeyColumns),
	)
	values := []interface{}{related.MessageID, o.PollMessageID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID = related.MessageID
	if o.R == nil {
		o.R = &pollsMessageR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &messageR{
			PollsMessages: PollsMessageSlice{o},
		}
	} else {
		related.R.PollsMessages = append(related.R.PollsMessages, o)
	}

	return nil
}

// PollsMessages retrieves all the records using an executor.
func PollsMessages(mods ...qm.QueryMod) pollsMessageQuery {
	mods = append(mods, qm.From("\"polls_messages\""))
	return pollsMessageQuery{NewQuery(mods...)}
}

// FindPollsMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPollsMessage(ctx context.Context, exec boil.ContextExecutor, pollMessageID int64, selectCols ...string) (*PollsMessage, error) {
	pollsMessageObj := &PollsMessage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"polls_messages\" where \"poll_message_id\"=$1", sel,
	)

	q := queries.Raw(query, pollMessageID)

	err := q.Bind(ctx, exec, pollsMessageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from polls_messages")
	}

	return pollsMessageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PollsMessage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls_messages provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pollsMessageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pollsMessageInsertCacheMut.RLock()
	cache, cached := pollsMessageInsertCache[key]
	pollsMessageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pollsMessageColumns,
			pollsMessageColumnsWithDefault,
			pollsMessageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pollsMessageType, pollsMessageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pollsMessageType, pollsMessageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"polls_messages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"polls_messages\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into polls_messages")
	}

	if !cached {
		pollsMessageInsertCacheMut.Lock()
		pollsMessageInsertCache[key] = cache
		pollsMessageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PollsMessage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PollsMessage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pollsMessageUpdateCacheMut.RLock()
	cache, cached := pollsMessageUpdateCache[key]
	pollsMessageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pollsMessageColumns,
			pollsMessagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update polls_messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"polls_messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pollsMessagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pollsMessageType, pollsMessageMapping, append(wl, pollsMessagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update polls_messages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for polls_messages")
	}

	if !cached {
		pollsMessageUpdateCacheMut.Lock()
		pollsMessageUpdateCache[key] = cache
		pollsMessageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pollsMessageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for polls_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for polls_messages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PollsMessageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"polls_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pollsMessagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pollsMessage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pollsMessage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PollsMessage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls_messages provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pollsMessageColumnsWithDefault, o)

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

	pollsMessageUpsertCacheMut.RLock()
	cache, cached := pollsMessageUpsertCache[key]
	pollsMessageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pollsMessageColumns,
			pollsMessageColumnsWithDefault,
			pollsMessageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pollsMessageColumns,
			pollsMessagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert polls_messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pollsMessagePrimaryKeyColumns))
			copy(conflict, pollsMessagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"polls_messages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pollsMessageType, pollsMessageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pollsMessageType, pollsMessageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert polls_messages")
	}

	if !cached {
		pollsMessageUpsertCacheMut.Lock()
		pollsMessageUpsertCache[key] = cache
		pollsMessageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PollsMessage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PollsMessage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PollsMessage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pollsMessagePrimaryKeyMapping)
	sql := "DELETE FROM \"polls_messages\" WHERE \"poll_message_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from polls_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for polls_messages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pollsMessageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pollsMessageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from polls_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for polls_messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PollsMessageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PollsMessage slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(pollsMessageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"polls_messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollsMessagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pollsMessage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for polls_messages")
	}

	if len(pollsMessageAfterDeleteHooks) != 0 {
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
func (o *PollsMessage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPollsMessage(ctx, exec, o.PollMessageID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PollsMessageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PollsMessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollsMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"polls_messages\".* FROM \"polls_messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollsMessagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PollsMessageSlice")
	}

	*o = slice

	return nil
}

// PollsMessageExists checks if the PollsMessage row exists.
func PollsMessageExists(ctx context.Context, exec boil.ContextExecutor, pollMessageID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"polls_messages\" where \"poll_message_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, pollMessageID)
	}

	row := exec.QueryRowContext(ctx, sql, pollMessageID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if polls_messages exists")
	}

	return exists, nil
}
