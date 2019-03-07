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

// EmailDomain is an object representing the database table.
type EmailDomain struct {
	EmailDomainID   int64     `boil:"email_domain_id" json:"email_domain_id" toml:"email_domain_id" yaml:"email_domain_id"`
	EmailDomainName string    `boil:"email_domain_name" json:"email_domain_name" toml:"email_domain_name" yaml:"email_domain_name"`
	UserAccountID   int64     `boil:"user_account_id" json:"user_account_id" toml:"user_account_id" yaml:"user_account_id"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *emailDomainR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L emailDomainL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmailDomainColumns = struct {
	EmailDomainID   string
	EmailDomainName string
	UserAccountID   string
	CreatedAt       string
}{
	EmailDomainID:   "email_domain_id",
	EmailDomainName: "email_domain_name",
	UserAccountID:   "user_account_id",
	CreatedAt:       "created_at",
}

// EmailDomainRels is where relationship names are stored.
var EmailDomainRels = struct {
	UserAccount    string
	EmailAddresses string
}{
	UserAccount:    "UserAccount",
	EmailAddresses: "EmailAddresses",
}

// emailDomainR is where relationships are stored.
type emailDomainR struct {
	UserAccount    *UserAccount
	EmailAddresses EmailAddressSlice
}

// NewStruct creates a new relationship struct
func (*emailDomainR) NewStruct() *emailDomainR {
	return &emailDomainR{}
}

// emailDomainL is where Load methods for each relationship are stored.
type emailDomainL struct{}

var (
	emailDomainColumns               = []string{"email_domain_id", "email_domain_name", "user_account_id", "created_at"}
	emailDomainColumnsWithoutDefault = []string{"email_domain_id", "email_domain_name", "user_account_id", "created_at"}
	emailDomainColumnsWithDefault    = []string{}
	emailDomainPrimaryKeyColumns     = []string{"email_domain_id"}
)

type (
	// EmailDomainSlice is an alias for a slice of pointers to EmailDomain.
	// This should generally be used opposed to []EmailDomain.
	EmailDomainSlice []*EmailDomain
	// EmailDomainHook is the signature for custom EmailDomain hook methods
	EmailDomainHook func(context.Context, boil.ContextExecutor, *EmailDomain) error

	emailDomainQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	emailDomainType                 = reflect.TypeOf(&EmailDomain{})
	emailDomainMapping              = queries.MakeStructMapping(emailDomainType)
	emailDomainPrimaryKeyMapping, _ = queries.BindMapping(emailDomainType, emailDomainMapping, emailDomainPrimaryKeyColumns)
	emailDomainInsertCacheMut       sync.RWMutex
	emailDomainInsertCache          = make(map[string]insertCache)
	emailDomainUpdateCacheMut       sync.RWMutex
	emailDomainUpdateCache          = make(map[string]updateCache)
	emailDomainUpsertCacheMut       sync.RWMutex
	emailDomainUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var emailDomainBeforeInsertHooks []EmailDomainHook
var emailDomainBeforeUpdateHooks []EmailDomainHook
var emailDomainBeforeDeleteHooks []EmailDomainHook
var emailDomainBeforeUpsertHooks []EmailDomainHook

var emailDomainAfterInsertHooks []EmailDomainHook
var emailDomainAfterSelectHooks []EmailDomainHook
var emailDomainAfterUpdateHooks []EmailDomainHook
var emailDomainAfterDeleteHooks []EmailDomainHook
var emailDomainAfterUpsertHooks []EmailDomainHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EmailDomain) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EmailDomain) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EmailDomain) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EmailDomain) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EmailDomain) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EmailDomain) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EmailDomain) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EmailDomain) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EmailDomain) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range emailDomainAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEmailDomainHook registers your hook function for all future operations.
func AddEmailDomainHook(hookPoint boil.HookPoint, emailDomainHook EmailDomainHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		emailDomainBeforeInsertHooks = append(emailDomainBeforeInsertHooks, emailDomainHook)
	case boil.BeforeUpdateHook:
		emailDomainBeforeUpdateHooks = append(emailDomainBeforeUpdateHooks, emailDomainHook)
	case boil.BeforeDeleteHook:
		emailDomainBeforeDeleteHooks = append(emailDomainBeforeDeleteHooks, emailDomainHook)
	case boil.BeforeUpsertHook:
		emailDomainBeforeUpsertHooks = append(emailDomainBeforeUpsertHooks, emailDomainHook)
	case boil.AfterInsertHook:
		emailDomainAfterInsertHooks = append(emailDomainAfterInsertHooks, emailDomainHook)
	case boil.AfterSelectHook:
		emailDomainAfterSelectHooks = append(emailDomainAfterSelectHooks, emailDomainHook)
	case boil.AfterUpdateHook:
		emailDomainAfterUpdateHooks = append(emailDomainAfterUpdateHooks, emailDomainHook)
	case boil.AfterDeleteHook:
		emailDomainAfterDeleteHooks = append(emailDomainAfterDeleteHooks, emailDomainHook)
	case boil.AfterUpsertHook:
		emailDomainAfterUpsertHooks = append(emailDomainAfterUpsertHooks, emailDomainHook)
	}
}

// One returns a single emailDomain record from the query.
func (q emailDomainQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EmailDomain, error) {
	o := &EmailDomain{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for email_domain")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EmailDomain records from the query.
func (q emailDomainQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmailDomainSlice, error) {
	var o []*EmailDomain

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EmailDomain slice")
	}

	if len(emailDomainAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EmailDomain records in the query.
func (q emailDomainQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count email_domain rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q emailDomainQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if email_domain exists")
	}

	return count > 0, nil
}

// UserAccount pointed to by the foreign key.
func (o *EmailDomain) UserAccount(mods ...qm.QueryMod) userAccountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_account_id=?", o.UserAccountID),
	}

	queryMods = append(queryMods, mods...)

	query := UserAccounts(queryMods...)
	queries.SetFrom(query.Query, "\"user_account\"")

	return query
}

// EmailAddresses retrieves all the email_address's EmailAddresses with an executor.
func (o *EmailDomain) EmailAddresses(mods ...qm.QueryMod) emailAddressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"email_address\".\"email_domain_id\"=?", o.EmailDomainID),
	)

	query := EmailAddresses(queryMods...)
	queries.SetFrom(query.Query, "\"email_address\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"email_address\".*"})
	}

	return query
}

// LoadUserAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (emailDomainL) LoadUserAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmailDomain interface{}, mods queries.Applicator) error {
	var slice []*EmailDomain
	var object *EmailDomain

	if singular {
		object = maybeEmailDomain.(*EmailDomain)
	} else {
		slice = *maybeEmailDomain.(*[]*EmailDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emailDomainR{}
		}
		args = append(args, object.UserAccountID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emailDomainR{}
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

	if len(emailDomainAfterSelectHooks) != 0 {
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
		foreign.R.EmailDomains = append(foreign.R.EmailDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserAccountID == foreign.UserAccountID {
				local.R.UserAccount = foreign
				if foreign.R == nil {
					foreign.R = &userAccountR{}
				}
				foreign.R.EmailDomains = append(foreign.R.EmailDomains, local)
				break
			}
		}
	}

	return nil
}

// LoadEmailAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (emailDomainL) LoadEmailAddresses(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmailDomain interface{}, mods queries.Applicator) error {
	var slice []*EmailDomain
	var object *EmailDomain

	if singular {
		object = maybeEmailDomain.(*EmailDomain)
	} else {
		slice = *maybeEmailDomain.(*[]*EmailDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emailDomainR{}
		}
		args = append(args, object.EmailDomainID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emailDomainR{}
			}

			for _, a := range args {
				if a == obj.EmailDomainID {
					continue Outer
				}
			}

			args = append(args, obj.EmailDomainID)
		}
	}

	query := NewQuery(qm.From(`email_address`), qm.WhereIn(`email_domain_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load email_address")
	}

	var resultSlice []*EmailAddress
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice email_address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on email_address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for email_address")
	}

	if len(emailAddressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.EmailAddresses = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &emailAddressR{}
			}
			foreign.R.EmailDomain = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EmailDomainID == foreign.EmailDomainID {
				local.R.EmailAddresses = append(local.R.EmailAddresses, foreign)
				if foreign.R == nil {
					foreign.R = &emailAddressR{}
				}
				foreign.R.EmailDomain = local
				break
			}
		}
	}

	return nil
}

// SetUserAccount of the emailDomain to the related item.
// Sets o.R.UserAccount to related.
// Adds o to related.R.EmailDomains.
func (o *EmailDomain) SetUserAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserAccount) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"email_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_account_id"}),
		strmangle.WhereClause("\"", "\"", 2, emailDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.UserAccountID, o.EmailDomainID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserAccountID = related.UserAccountID
	if o.R == nil {
		o.R = &emailDomainR{
			UserAccount: related,
		}
	} else {
		o.R.UserAccount = related
	}

	if related.R == nil {
		related.R = &userAccountR{
			EmailDomains: EmailDomainSlice{o},
		}
	} else {
		related.R.EmailDomains = append(related.R.EmailDomains, o)
	}

	return nil
}

// AddEmailAddresses adds the given related objects to the existing relationships
// of the email_domain, optionally inserting them as new records.
// Appends related to o.R.EmailAddresses.
// Sets related.R.EmailDomain appropriately.
func (o *EmailDomain) AddEmailAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*EmailAddress) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EmailDomainID = o.EmailDomainID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"email_address\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"email_domain_id"}),
				strmangle.WhereClause("\"", "\"", 2, emailAddressPrimaryKeyColumns),
			)
			values := []interface{}{o.EmailDomainID, rel.EmailName, rel.EmailDomainID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EmailDomainID = o.EmailDomainID
		}
	}

	if o.R == nil {
		o.R = &emailDomainR{
			EmailAddresses: related,
		}
	} else {
		o.R.EmailAddresses = append(o.R.EmailAddresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &emailAddressR{
				EmailDomain: o,
			}
		} else {
			rel.R.EmailDomain = o
		}
	}
	return nil
}

// EmailDomains retrieves all the records using an executor.
func EmailDomains(mods ...qm.QueryMod) emailDomainQuery {
	mods = append(mods, qm.From("\"email_domain\""))
	return emailDomainQuery{NewQuery(mods...)}
}

// FindEmailDomain retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmailDomain(ctx context.Context, exec boil.ContextExecutor, emailDomainID int64, selectCols ...string) (*EmailDomain, error) {
	emailDomainObj := &EmailDomain{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"email_domain\" where \"email_domain_id\"=$1", sel,
	)

	q := queries.Raw(query, emailDomainID)

	err := q.Bind(ctx, exec, emailDomainObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from email_domain")
	}

	return emailDomainObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EmailDomain) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_domain provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailDomainColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	emailDomainInsertCacheMut.RLock()
	cache, cached := emailDomainInsertCache[key]
	emailDomainInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			emailDomainColumns,
			emailDomainColumnsWithDefault,
			emailDomainColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(emailDomainType, emailDomainMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(emailDomainType, emailDomainMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"email_domain\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"email_domain\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into email_domain")
	}

	if !cached {
		emailDomainInsertCacheMut.Lock()
		emailDomainInsertCache[key] = cache
		emailDomainInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EmailDomain.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EmailDomain) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	emailDomainUpdateCacheMut.RLock()
	cache, cached := emailDomainUpdateCache[key]
	emailDomainUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			emailDomainColumns,
			emailDomainPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update email_domain, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"email_domain\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, emailDomainPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(emailDomainType, emailDomainMapping, append(wl, emailDomainPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update email_domain row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for email_domain")
	}

	if !cached {
		emailDomainUpdateCacheMut.Lock()
		emailDomainUpdateCache[key] = cache
		emailDomainUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q emailDomainQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for email_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for email_domain")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmailDomainSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"email_domain\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, emailDomainPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in emailDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all emailDomain")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EmailDomain) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_domain provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailDomainColumnsWithDefault, o)

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

	emailDomainUpsertCacheMut.RLock()
	cache, cached := emailDomainUpsertCache[key]
	emailDomainUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			emailDomainColumns,
			emailDomainColumnsWithDefault,
			emailDomainColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			emailDomainColumns,
			emailDomainPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert email_domain, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(emailDomainPrimaryKeyColumns))
			copy(conflict, emailDomainPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"email_domain\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(emailDomainType, emailDomainMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(emailDomainType, emailDomainMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert email_domain")
	}

	if !cached {
		emailDomainUpsertCacheMut.Lock()
		emailDomainUpsertCache[key] = cache
		emailDomainUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EmailDomain record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EmailDomain) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmailDomain provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), emailDomainPrimaryKeyMapping)
	sql := "DELETE FROM \"email_domain\" WHERE \"email_domain_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from email_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for email_domain")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q emailDomainQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no emailDomainQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from email_domain")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_domain")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmailDomainSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmailDomain slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(emailDomainBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"email_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, emailDomainPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from emailDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_domain")
	}

	if len(emailDomainAfterDeleteHooks) != 0 {
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
func (o *EmailDomain) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmailDomain(ctx, exec, o.EmailDomainID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmailDomainSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmailDomainSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"email_domain\".* FROM \"email_domain\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, emailDomainPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmailDomainSlice")
	}

	*o = slice

	return nil
}

// EmailDomainExists checks if the EmailDomain row exists.
func EmailDomainExists(ctx context.Context, exec boil.ContextExecutor, emailDomainID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"email_domain\" where \"email_domain_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, emailDomainID)
	}

	row := exec.QueryRowContext(ctx, sql, emailDomainID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if email_domain exists")
	}

	return exists, nil
}
