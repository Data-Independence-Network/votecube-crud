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

// UserAccountSuffix is an object representing the database table.
type UserAccountSuffix struct {
	UserAccountSuffixID int64 `boil:"user_account_suffix_id" json:"user_account_suffix_id" toml:"user_account_suffix_id" yaml:"user_account_suffix_id"`
	SuffixID            int64 `boil:"suffix_id" json:"suffix_id" toml:"suffix_id" yaml:"suffix_id"`
	UserAccountID       int64 `boil:"user_account_id" json:"user_account_id" toml:"user_account_id" yaml:"user_account_id"`
	SuffixPosition      int16 `boil:"suffix_position" json:"suffix_position" toml:"suffix_position" yaml:"suffix_position"`

	R *userAccountSuffixR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userAccountSuffixL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserAccountSuffixColumns = struct {
	UserAccountSuffixID string
	SuffixID            string
	UserAccountID       string
	SuffixPosition      string
}{
	UserAccountSuffixID: "user_account_suffix_id",
	SuffixID:            "suffix_id",
	UserAccountID:       "user_account_id",
	SuffixPosition:      "suffix_position",
}

// UserAccountSuffixRels is where relationship names are stored.
var UserAccountSuffixRels = struct {
	UserAccount string
	Suffix      string
}{
	UserAccount: "UserAccount",
	Suffix:      "Suffix",
}

// userAccountSuffixR is where relationships are stored.
type userAccountSuffixR struct {
	UserAccount *UserAccount
	Suffix      *Suffix
}

// NewStruct creates a new relationship struct
func (*userAccountSuffixR) NewStruct() *userAccountSuffixR {
	return &userAccountSuffixR{}
}

// userAccountSuffixL is where Load methods for each relationship are stored.
type userAccountSuffixL struct{}

var (
	userAccountSuffixColumns               = []string{"user_account_suffix_id", "suffix_id", "user_account_id", "suffix_position"}
	userAccountSuffixColumnsWithoutDefault = []string{"user_account_suffix_id", "suffix_id", "user_account_id", "suffix_position"}
	userAccountSuffixColumnsWithDefault    = []string{}
	userAccountSuffixPrimaryKeyColumns     = []string{"user_account_suffix_id"}
)

type (
	// UserAccountSuffixSlice is an alias for a slice of pointers to UserAccountSuffix.
	// This should generally be used opposed to []UserAccountSuffix.
	UserAccountSuffixSlice []*UserAccountSuffix
	// UserAccountSuffixHook is the signature for custom UserAccountSuffix hook methods
	UserAccountSuffixHook func(context.Context, boil.ContextExecutor, *UserAccountSuffix) error

	userAccountSuffixQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userAccountSuffixType                 = reflect.TypeOf(&UserAccountSuffix{})
	userAccountSuffixMapping              = queries.MakeStructMapping(userAccountSuffixType)
	userAccountSuffixPrimaryKeyMapping, _ = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, userAccountSuffixPrimaryKeyColumns)
	userAccountSuffixInsertCacheMut       sync.RWMutex
	userAccountSuffixInsertCache          = make(map[string]insertCache)
	userAccountSuffixUpdateCacheMut       sync.RWMutex
	userAccountSuffixUpdateCache          = make(map[string]updateCache)
	userAccountSuffixUpsertCacheMut       sync.RWMutex
	userAccountSuffixUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var userAccountSuffixBeforeInsertHooks []UserAccountSuffixHook
var userAccountSuffixBeforeUpdateHooks []UserAccountSuffixHook
var userAccountSuffixBeforeDeleteHooks []UserAccountSuffixHook
var userAccountSuffixBeforeUpsertHooks []UserAccountSuffixHook

var userAccountSuffixAfterInsertHooks []UserAccountSuffixHook
var userAccountSuffixAfterSelectHooks []UserAccountSuffixHook
var userAccountSuffixAfterUpdateHooks []UserAccountSuffixHook
var userAccountSuffixAfterDeleteHooks []UserAccountSuffixHook
var userAccountSuffixAfterUpsertHooks []UserAccountSuffixHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserAccountSuffix) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserAccountSuffix) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserAccountSuffix) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserAccountSuffix) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserAccountSuffix) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserAccountSuffix) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserAccountSuffix) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserAccountSuffix) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserAccountSuffix) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userAccountSuffixAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserAccountSuffixHook registers your hook function for all future operations.
func AddUserAccountSuffixHook(hookPoint boil.HookPoint, userAccountSuffixHook UserAccountSuffixHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userAccountSuffixBeforeInsertHooks = append(userAccountSuffixBeforeInsertHooks, userAccountSuffixHook)
	case boil.BeforeUpdateHook:
		userAccountSuffixBeforeUpdateHooks = append(userAccountSuffixBeforeUpdateHooks, userAccountSuffixHook)
	case boil.BeforeDeleteHook:
		userAccountSuffixBeforeDeleteHooks = append(userAccountSuffixBeforeDeleteHooks, userAccountSuffixHook)
	case boil.BeforeUpsertHook:
		userAccountSuffixBeforeUpsertHooks = append(userAccountSuffixBeforeUpsertHooks, userAccountSuffixHook)
	case boil.AfterInsertHook:
		userAccountSuffixAfterInsertHooks = append(userAccountSuffixAfterInsertHooks, userAccountSuffixHook)
	case boil.AfterSelectHook:
		userAccountSuffixAfterSelectHooks = append(userAccountSuffixAfterSelectHooks, userAccountSuffixHook)
	case boil.AfterUpdateHook:
		userAccountSuffixAfterUpdateHooks = append(userAccountSuffixAfterUpdateHooks, userAccountSuffixHook)
	case boil.AfterDeleteHook:
		userAccountSuffixAfterDeleteHooks = append(userAccountSuffixAfterDeleteHooks, userAccountSuffixHook)
	case boil.AfterUpsertHook:
		userAccountSuffixAfterUpsertHooks = append(userAccountSuffixAfterUpsertHooks, userAccountSuffixHook)
	}
}

// One returns a single userAccountSuffix record from the query.
func (q userAccountSuffixQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserAccountSuffix, error) {
	o := &UserAccountSuffix{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_account_suffix")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserAccountSuffix records from the query.
func (q userAccountSuffixQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserAccountSuffixSlice, error) {
	var o []*UserAccountSuffix

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserAccountSuffix slice")
	}

	if len(userAccountSuffixAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserAccountSuffix records in the query.
func (q userAccountSuffixQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_account_suffix rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userAccountSuffixQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_account_suffix exists")
	}

	return count > 0, nil
}

// UserAccount pointed to by the foreign key.
func (o *UserAccountSuffix) UserAccount(mods ...qm.QueryMod) userAccountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_account_id=?", o.UserAccountID),
	}

	queryMods = append(queryMods, mods...)

	query := UserAccounts(queryMods...)
	queries.SetFrom(query.Query, "\"user_account\"")

	return query
}

// Suffix pointed to by the foreign key.
func (o *UserAccountSuffix) Suffix(mods ...qm.QueryMod) suffixQuery {
	queryMods := []qm.QueryMod{
		qm.Where("suffix_id=?", o.SuffixID),
	}

	queryMods = append(queryMods, mods...)

	query := Suffixes(queryMods...)
	queries.SetFrom(query.Query, "\"suffix\"")

	return query
}

// LoadUserAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userAccountSuffixL) LoadUserAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserAccountSuffix interface{}, mods queries.Applicator) error {
	var slice []*UserAccountSuffix
	var object *UserAccountSuffix

	if singular {
		object = maybeUserAccountSuffix.(*UserAccountSuffix)
	} else {
		slice = *maybeUserAccountSuffix.(*[]*UserAccountSuffix)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userAccountSuffixR{}
		}
		args = append(args, object.UserAccountID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userAccountSuffixR{}
			}

			for _, a := range args {
				if a == obj.UserAccountID {
					continue Outer
				}
			}

			args = append(args, obj.UserAccountID)
		}
	}

	query := NewQuery(qm.From(`user_account`), qm.WhereIn(`user_account_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserAccount")
	}

	var resultSlice []*UserAccount
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserAccount")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_account")
	}

	if len(userAccountSuffixAfterSelectHooks) != 0 {
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
		object.R.UserAccount = foreign
		if foreign.R == nil {
			foreign.R = &userAccountR{}
		}
		foreign.R.UserAccountSuffixes = append(foreign.R.UserAccountSuffixes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserAccountID == foreign.UserAccountID {
				local.R.UserAccount = foreign
				if foreign.R == nil {
					foreign.R = &userAccountR{}
				}
				foreign.R.UserAccountSuffixes = append(foreign.R.UserAccountSuffixes, local)
				break
			}
		}
	}

	return nil
}

// LoadSuffix allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userAccountSuffixL) LoadSuffix(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserAccountSuffix interface{}, mods queries.Applicator) error {
	var slice []*UserAccountSuffix
	var object *UserAccountSuffix

	if singular {
		object = maybeUserAccountSuffix.(*UserAccountSuffix)
	} else {
		slice = *maybeUserAccountSuffix.(*[]*UserAccountSuffix)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userAccountSuffixR{}
		}
		args = append(args, object.SuffixID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userAccountSuffixR{}
			}

			for _, a := range args {
				if a == obj.SuffixID {
					continue Outer
				}
			}

			args = append(args, obj.SuffixID)
		}
	}

	query := NewQuery(qm.From(`suffix`), qm.WhereIn(`suffix_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Suffix")
	}

	var resultSlice []*Suffix
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Suffix")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for suffix")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for suffix")
	}

	if len(userAccountSuffixAfterSelectHooks) != 0 {
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
		object.R.Suffix = foreign
		if foreign.R == nil {
			foreign.R = &suffixR{}
		}
		foreign.R.UserAccountSuffixes = append(foreign.R.UserAccountSuffixes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SuffixID == foreign.SuffixID {
				local.R.Suffix = foreign
				if foreign.R == nil {
					foreign.R = &suffixR{}
				}
				foreign.R.UserAccountSuffixes = append(foreign.R.UserAccountSuffixes, local)
				break
			}
		}
	}

	return nil
}

// SetUserAccount of the userAccountSuffix to the related item.
// Sets o.R.UserAccount to related.
// Adds o to related.R.UserAccountSuffixes.
func (o *UserAccountSuffix) SetUserAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserAccount) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_account_suffix\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_account_id"}),
		strmangle.WhereClause("\"", "\"", 2, userAccountSuffixPrimaryKeyColumns),
	)
	values := []interface{}{related.UserAccountID, o.UserAccountSuffixID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserAccountID = related.UserAccountID
	if o.R == nil {
		o.R = &userAccountSuffixR{
			UserAccount: related,
		}
	} else {
		o.R.UserAccount = related
	}

	if related.R == nil {
		related.R = &userAccountR{
			UserAccountSuffixes: UserAccountSuffixSlice{o},
		}
	} else {
		related.R.UserAccountSuffixes = append(related.R.UserAccountSuffixes, o)
	}

	return nil
}

// SetSuffix of the userAccountSuffix to the related item.
// Sets o.R.Suffix to related.
// Adds o to related.R.UserAccountSuffixes.
func (o *UserAccountSuffix) SetSuffix(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Suffix) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_account_suffix\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"suffix_id"}),
		strmangle.WhereClause("\"", "\"", 2, userAccountSuffixPrimaryKeyColumns),
	)
	values := []interface{}{related.SuffixID, o.UserAccountSuffixID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SuffixID = related.SuffixID
	if o.R == nil {
		o.R = &userAccountSuffixR{
			Suffix: related,
		}
	} else {
		o.R.Suffix = related
	}

	if related.R == nil {
		related.R = &suffixR{
			UserAccountSuffixes: UserAccountSuffixSlice{o},
		}
	} else {
		related.R.UserAccountSuffixes = append(related.R.UserAccountSuffixes, o)
	}

	return nil
}

// UserAccountSuffixes retrieves all the records using an executor.
func UserAccountSuffixes(mods ...qm.QueryMod) userAccountSuffixQuery {
	mods = append(mods, qm.From("\"user_account_suffix\""))
	return userAccountSuffixQuery{NewQuery(mods...)}
}

// FindUserAccountSuffix retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserAccountSuffix(ctx context.Context, exec boil.ContextExecutor, userAccountSuffixID int64, selectCols ...string) (*UserAccountSuffix, error) {
	userAccountSuffixObj := &UserAccountSuffix{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_account_suffix\" where \"user_account_suffix_id\"=$1", sel,
	)

	q := queries.Raw(query, userAccountSuffixID)

	err := q.Bind(ctx, exec, userAccountSuffixObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_account_suffix")
	}

	return userAccountSuffixObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserAccountSuffix) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_account_suffix provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userAccountSuffixColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userAccountSuffixInsertCacheMut.RLock()
	cache, cached := userAccountSuffixInsertCache[key]
	userAccountSuffixInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAccountSuffixColumns,
			userAccountSuffixColumnsWithDefault,
			userAccountSuffixColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_account_suffix\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_account_suffix\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_account_suffix")
	}

	if !cached {
		userAccountSuffixInsertCacheMut.Lock()
		userAccountSuffixInsertCache[key] = cache
		userAccountSuffixInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserAccountSuffix.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserAccountSuffix) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userAccountSuffixUpdateCacheMut.RLock()
	cache, cached := userAccountSuffixUpdateCache[key]
	userAccountSuffixUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAccountSuffixColumns,
			userAccountSuffixPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_account_suffix, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_account_suffix\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userAccountSuffixPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, append(wl, userAccountSuffixPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_account_suffix row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_account_suffix")
	}

	if !cached {
		userAccountSuffixUpdateCacheMut.Lock()
		userAccountSuffixUpdateCache[key] = cache
		userAccountSuffixUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userAccountSuffixQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_account_suffix")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_account_suffix")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserAccountSuffixSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAccountSuffixPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_account_suffix\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userAccountSuffixPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userAccountSuffix slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userAccountSuffix")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserAccountSuffix) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_account_suffix provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userAccountSuffixColumnsWithDefault, o)

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

	userAccountSuffixUpsertCacheMut.RLock()
	cache, cached := userAccountSuffixUpsertCache[key]
	userAccountSuffixUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userAccountSuffixColumns,
			userAccountSuffixColumnsWithDefault,
			userAccountSuffixColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userAccountSuffixColumns,
			userAccountSuffixPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert user_account_suffix, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userAccountSuffixPrimaryKeyColumns))
			copy(conflict, userAccountSuffixPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"user_account_suffix\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userAccountSuffixType, userAccountSuffixMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_account_suffix")
	}

	if !cached {
		userAccountSuffixUpsertCacheMut.Lock()
		userAccountSuffixUpsertCache[key] = cache
		userAccountSuffixUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserAccountSuffix record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserAccountSuffix) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserAccountSuffix provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userAccountSuffixPrimaryKeyMapping)
	sql := "DELETE FROM \"user_account_suffix\" WHERE \"user_account_suffix_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_account_suffix")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_account_suffix")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userAccountSuffixQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userAccountSuffixQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_account_suffix")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_account_suffix")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserAccountSuffixSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserAccountSuffix slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(userAccountSuffixBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAccountSuffixPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_account_suffix\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userAccountSuffixPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userAccountSuffix slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_account_suffix")
	}

	if len(userAccountSuffixAfterDeleteHooks) != 0 {
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
func (o *UserAccountSuffix) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserAccountSuffix(ctx, exec, o.UserAccountSuffixID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserAccountSuffixSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserAccountSuffixSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAccountSuffixPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_account_suffix\".* FROM \"user_account_suffix\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userAccountSuffixPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserAccountSuffixSlice")
	}

	*o = slice

	return nil
}

// UserAccountSuffixExists checks if the UserAccountSuffix row exists.
func UserAccountSuffixExists(ctx context.Context, exec boil.ContextExecutor, userAccountSuffixID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_account_suffix\" where \"user_account_suffix_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userAccountSuffixID)
	}

	row := exec.QueryRowContext(ctx, sql, userAccountSuffixID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_account_suffix exists")
	}

	return exists, nil
}
