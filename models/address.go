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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Address is an object representing the database table.
type Address struct {
	AddressID      int64       `boil:"address_id" json:"address_id" toml:"address_id" yaml:"address_id"`
	StreetID       int64       `boil:"street_id" json:"street_id" toml:"street_id" yaml:"street_id"`
	StreetNumber   null.Int16  `boil:"street_number" json:"street_number,omitempty" toml:"street_number" yaml:"street_number,omitempty"`
	BuildingUnit   null.String `boil:"building_unit" json:"building_unit,omitempty" toml:"building_unit" yaml:"building_unit,omitempty"`
	IsBuildingUnit null.Bool   `boil:"is_building_unit" json:"is_building_unit,omitempty" toml:"is_building_unit" yaml:"is_building_unit,omitempty"`
	PostCode       string      `boil:"post_code" json:"post_code" toml:"post_code" yaml:"post_code"`

	R *addressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L addressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AddressColumns = struct {
	AddressID      string
	StreetID       string
	StreetNumber   string
	BuildingUnit   string
	IsBuildingUnit string
	PostCode       string
}{
	AddressID:      "address_id",
	StreetID:       "street_id",
	StreetNumber:   "street_number",
	BuildingUnit:   "building_unit",
	IsBuildingUnit: "is_building_unit",
	PostCode:       "post_code",
}

// AddressRels is where relationship names are stored.
var AddressRels = struct {
	Street            string
	UserPersonalInfos string
}{
	Street:            "Street",
	UserPersonalInfos: "UserPersonalInfos",
}

// addressR is where relationships are stored.
type addressR struct {
	Street            *Street
	UserPersonalInfos UserPersonalInfoSlice
}

// NewStruct creates a new relationship struct
func (*addressR) NewStruct() *addressR {
	return &addressR{}
}

// addressL is where Load methods for each relationship are stored.
type addressL struct{}

var (
	addressColumns               = []string{"address_id", "street_id", "street_number", "building_unit", "is_building_unit", "post_code"}
	addressColumnsWithoutDefault = []string{"address_id", "street_id", "street_number", "building_unit", "is_building_unit", "post_code"}
	addressColumnsWithDefault    = []string{}
	addressPrimaryKeyColumns     = []string{"address_id"}
)

type (
	// AddressSlice is an alias for a slice of pointers to Address.
	// This should generally be used opposed to []Address.
	AddressSlice []*Address
	// AddressHook is the signature for custom Address hook methods
	AddressHook func(context.Context, boil.ContextExecutor, *Address) error

	addressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	addressType                 = reflect.TypeOf(&Address{})
	addressMapping              = queries.MakeStructMapping(addressType)
	addressPrimaryKeyMapping, _ = queries.BindMapping(addressType, addressMapping, addressPrimaryKeyColumns)
	addressInsertCacheMut       sync.RWMutex
	addressInsertCache          = make(map[string]insertCache)
	addressUpdateCacheMut       sync.RWMutex
	addressUpdateCache          = make(map[string]updateCache)
	addressUpsertCacheMut       sync.RWMutex
	addressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var addressBeforeInsertHooks []AddressHook
var addressBeforeUpdateHooks []AddressHook
var addressBeforeDeleteHooks []AddressHook
var addressBeforeUpsertHooks []AddressHook

var addressAfterInsertHooks []AddressHook
var addressAfterSelectHooks []AddressHook
var addressAfterUpdateHooks []AddressHook
var addressAfterDeleteHooks []AddressHook
var addressAfterUpsertHooks []AddressHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Address) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Address) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Address) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Address) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Address) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Address) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Address) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Address) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Address) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range addressAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAddressHook registers your hook function for all future operations.
func AddAddressHook(hookPoint boil.HookPoint, addressHook AddressHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		addressBeforeInsertHooks = append(addressBeforeInsertHooks, addressHook)
	case boil.BeforeUpdateHook:
		addressBeforeUpdateHooks = append(addressBeforeUpdateHooks, addressHook)
	case boil.BeforeDeleteHook:
		addressBeforeDeleteHooks = append(addressBeforeDeleteHooks, addressHook)
	case boil.BeforeUpsertHook:
		addressBeforeUpsertHooks = append(addressBeforeUpsertHooks, addressHook)
	case boil.AfterInsertHook:
		addressAfterInsertHooks = append(addressAfterInsertHooks, addressHook)
	case boil.AfterSelectHook:
		addressAfterSelectHooks = append(addressAfterSelectHooks, addressHook)
	case boil.AfterUpdateHook:
		addressAfterUpdateHooks = append(addressAfterUpdateHooks, addressHook)
	case boil.AfterDeleteHook:
		addressAfterDeleteHooks = append(addressAfterDeleteHooks, addressHook)
	case boil.AfterUpsertHook:
		addressAfterUpsertHooks = append(addressAfterUpsertHooks, addressHook)
	}
}

// OneG returns a single address record from the query using the global executor.
func (q addressQuery) OneG(ctx context.Context) (*Address, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single address record from the query.
func (q addressQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Address, error) {
	o := &Address{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for address")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Address records from the query using the global executor.
func (q addressQuery) AllG(ctx context.Context) (AddressSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Address records from the query.
func (q addressQuery) All(ctx context.Context, exec boil.ContextExecutor) (AddressSlice, error) {
	var o []*Address

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Address slice")
	}

	if len(addressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Address records in the query, and panics on error.
func (q addressQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Address records in the query.
func (q addressQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count address rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q addressQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q addressQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if address exists")
	}

	return count > 0, nil
}

// Street pointed to by the foreign key.
func (o *Address) Street(mods ...qm.QueryMod) streetQuery {
	queryMods := []qm.QueryMod{
		qm.Where("street_id=?", o.StreetID),
	}

	queryMods = append(queryMods, mods...)

	query := Streets(queryMods...)
	queries.SetFrom(query.Query, "\"street\"")

	return query
}

// UserPersonalInfos retrieves all the user_personal_info's UserPersonalInfos with an executor.
func (o *Address) UserPersonalInfos(mods ...qm.QueryMod) userPersonalInfoQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_personal_info\".\"address_id\"=?", o.AddressID),
	)

	query := UserPersonalInfos(queryMods...)
	queries.SetFrom(query.Query, "\"user_personal_info\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_personal_info\".*"})
	}

	return query
}

// LoadStreet allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (addressL) LoadStreet(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAddress interface{}, mods queries.Applicator) error {
	var slice []*Address
	var object *Address

	if singular {
		object = maybeAddress.(*Address)
	} else {
		slice = *maybeAddress.(*[]*Address)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &addressR{}
		}
		args = append(args, object.StreetID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &addressR{}
			}

			for _, a := range args {
				if a == obj.StreetID {
					continue Outer
				}
			}

			args = append(args, obj.StreetID)
		}
	}

	query := NewQuery(qm.From(`street`), qm.WhereIn(`street_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Street")
	}

	var resultSlice []*Street
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Street")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for street")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for street")
	}

	if len(addressAfterSelectHooks) != 0 {
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
		object.R.Street = foreign
		if foreign.R == nil {
			foreign.R = &streetR{}
		}
		foreign.R.Addresses = append(foreign.R.Addresses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StreetID == foreign.StreetID {
				local.R.Street = foreign
				if foreign.R == nil {
					foreign.R = &streetR{}
				}
				foreign.R.Addresses = append(foreign.R.Addresses, local)
				break
			}
		}
	}

	return nil
}

// LoadUserPersonalInfos allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (addressL) LoadUserPersonalInfos(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAddress interface{}, mods queries.Applicator) error {
	var slice []*Address
	var object *Address

	if singular {
		object = maybeAddress.(*Address)
	} else {
		slice = *maybeAddress.(*[]*Address)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &addressR{}
		}
		args = append(args, object.AddressID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &addressR{}
			}

			for _, a := range args {
				if a == obj.AddressID {
					continue Outer
				}
			}

			args = append(args, obj.AddressID)
		}
	}

	query := NewQuery(qm.From(`user_personal_info`), qm.WhereIn(`address_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_personal_info")
	}

	var resultSlice []*UserPersonalInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_personal_info")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_personal_info")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_personal_info")
	}

	if len(userPersonalInfoAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserPersonalInfos = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userPersonalInfoR{}
			}
			foreign.R.Address = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AddressID == foreign.AddressID {
				local.R.UserPersonalInfos = append(local.R.UserPersonalInfos, foreign)
				if foreign.R == nil {
					foreign.R = &userPersonalInfoR{}
				}
				foreign.R.Address = local
				break
			}
		}
	}

	return nil
}

// SetStreetG of the address to the related item.
// Sets o.R.Street to related.
// Adds o to related.R.Addresses.
// Uses the global database handle.
func (o *Address) SetStreetG(ctx context.Context, insert bool, related *Street) error {
	return o.SetStreet(ctx, boil.GetContextDB(), insert, related)
}

// SetStreet of the address to the related item.
// Sets o.R.Street to related.
// Adds o to related.R.Addresses.
func (o *Address) SetStreet(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Street) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"address\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"street_id"}),
		strmangle.WhereClause("\"", "\"", 2, addressPrimaryKeyColumns),
	)
	values := []interface{}{related.StreetID, o.AddressID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StreetID = related.StreetID
	if o.R == nil {
		o.R = &addressR{
			Street: related,
		}
	} else {
		o.R.Street = related
	}

	if related.R == nil {
		related.R = &streetR{
			Addresses: AddressSlice{o},
		}
	} else {
		related.R.Addresses = append(related.R.Addresses, o)
	}

	return nil
}

// AddUserPersonalInfosG adds the given related objects to the existing relationships
// of the address, optionally inserting them as new records.
// Appends related to o.R.UserPersonalInfos.
// Sets related.R.Address appropriately.
// Uses the global database handle.
func (o *Address) AddUserPersonalInfosG(ctx context.Context, insert bool, related ...*UserPersonalInfo) error {
	return o.AddUserPersonalInfos(ctx, boil.GetContextDB(), insert, related...)
}

// AddUserPersonalInfos adds the given related objects to the existing relationships
// of the address, optionally inserting them as new records.
// Appends related to o.R.UserPersonalInfos.
// Sets related.R.Address appropriately.
func (o *Address) AddUserPersonalInfos(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserPersonalInfo) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AddressID = o.AddressID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_personal_info\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"address_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPersonalInfoPrimaryKeyColumns),
			)
			values := []interface{}{o.AddressID, rel.UserPersonalInfoID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.AddressID = o.AddressID
		}
	}

	if o.R == nil {
		o.R = &addressR{
			UserPersonalInfos: related,
		}
	} else {
		o.R.UserPersonalInfos = append(o.R.UserPersonalInfos, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userPersonalInfoR{
				Address: o,
			}
		} else {
			rel.R.Address = o
		}
	}
	return nil
}

// Addresses retrieves all the records using an executor.
func Addresses(mods ...qm.QueryMod) addressQuery {
	mods = append(mods, qm.From("\"address\""))
	return addressQuery{NewQuery(mods...)}
}

// FindAddressG retrieves a single record by ID.
func FindAddressG(ctx context.Context, addressID int64, selectCols ...string) (*Address, error) {
	return FindAddress(ctx, boil.GetContextDB(), addressID, selectCols...)
}

// FindAddress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAddress(ctx context.Context, exec boil.ContextExecutor, addressID int64, selectCols ...string) (*Address, error) {
	addressObj := &Address{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"address\" where \"address_id\"=$1", sel,
	)

	q := queries.Raw(query, addressID)

	err := q.Bind(ctx, exec, addressObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from address")
	}

	return addressObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Address) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Address) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no address provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(addressColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	addressInsertCacheMut.RLock()
	cache, cached := addressInsertCache[key]
	addressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			addressColumns,
			addressColumnsWithDefault,
			addressColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(addressType, addressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(addressType, addressMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"address\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"address\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into address")
	}

	if !cached {
		addressInsertCacheMut.Lock()
		addressInsertCache[key] = cache
		addressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Address record using the global executor.
// See Update for more documentation.
func (o *Address) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Address.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Address) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	addressUpdateCacheMut.RLock()
	cache, cached := addressUpdateCache[key]
	addressUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			addressColumns,
			addressPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update address, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"address\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, addressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(addressType, addressMapping, append(wl, addressPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update address row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for address")
	}

	if !cached {
		addressUpdateCacheMut.Lock()
		addressUpdateCache[key] = cache
		addressUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q addressQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for address")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AddressSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AddressSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), addressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"address\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, addressPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in address slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all address")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Address) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Address) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no address provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(addressColumnsWithDefault, o)

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

	addressUpsertCacheMut.RLock()
	cache, cached := addressUpsertCache[key]
	addressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			addressColumns,
			addressColumnsWithDefault,
			addressColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			addressColumns,
			addressPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert address, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(addressPrimaryKeyColumns))
			copy(conflict, addressPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"address\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(addressType, addressMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(addressType, addressMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert address")
	}

	if !cached {
		addressUpsertCacheMut.Lock()
		addressUpsertCache[key] = cache
		addressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Address record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Address) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Address record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Address) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Address provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), addressPrimaryKeyMapping)
	sql := "DELETE FROM \"address\" WHERE \"address_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for address")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q addressQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no addressQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for address")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AddressSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AddressSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Address slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(addressBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), addressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"address\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, addressPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from address slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for address")
	}

	if len(addressAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Address) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Address provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Address) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAddress(ctx, exec, o.AddressID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AddressSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty AddressSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AddressSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AddressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), addressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"address\".* FROM \"address\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, addressPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AddressSlice")
	}

	*o = slice

	return nil
}

// AddressExistsG checks if the Address row exists.
func AddressExistsG(ctx context.Context, addressID int64) (bool, error) {
	return AddressExists(ctx, boil.GetContextDB(), addressID)
}

// AddressExists checks if the Address row exists.
func AddressExists(ctx context.Context, exec boil.ContextExecutor, addressID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"address\" where \"address_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, addressID)
	}

	row := exec.QueryRowContext(ctx, sql, addressID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if address exists")
	}

	return exists, nil
}
