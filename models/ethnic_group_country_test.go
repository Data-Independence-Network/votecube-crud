// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEthnicGroupCountries(t *testing.T) {
	t.Parallel()

	query := EthnicGroupCountries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEthnicGroupCountriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEthnicGroupCountriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EthnicGroupCountries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEthnicGroupCountriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EthnicGroupCountrySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEthnicGroupCountriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EthnicGroupCountryExists(ctx, tx, o.EthnicGroupCountryID)
	if err != nil {
		t.Errorf("Unable to check if EthnicGroupCountry exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EthnicGroupCountryExists to return true, but got false.")
	}
}

func testEthnicGroupCountriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ethnicGroupCountryFound, err := FindEthnicGroupCountry(ctx, tx, o.EthnicGroupCountryID)
	if err != nil {
		t.Error(err)
	}

	if ethnicGroupCountryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEthnicGroupCountriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EthnicGroupCountries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEthnicGroupCountriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EthnicGroupCountries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEthnicGroupCountriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ethnicGroupCountryOne := &EthnicGroupCountry{}
	ethnicGroupCountryTwo := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, ethnicGroupCountryOne, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}
	if err = randomize.Struct(seed, ethnicGroupCountryTwo, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ethnicGroupCountryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ethnicGroupCountryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EthnicGroupCountries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEthnicGroupCountriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ethnicGroupCountryOne := &EthnicGroupCountry{}
	ethnicGroupCountryTwo := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, ethnicGroupCountryOne, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}
	if err = randomize.Struct(seed, ethnicGroupCountryTwo, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ethnicGroupCountryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ethnicGroupCountryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func ethnicGroupCountryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func ethnicGroupCountryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EthnicGroupCountry) error {
	*o = EthnicGroupCountry{}
	return nil
}

func testEthnicGroupCountriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EthnicGroupCountry{}
	o := &EthnicGroupCountry{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry object: %s", err)
	}

	AddEthnicGroupCountryHook(boil.BeforeInsertHook, ethnicGroupCountryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryBeforeInsertHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.AfterInsertHook, ethnicGroupCountryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryAfterInsertHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.AfterSelectHook, ethnicGroupCountryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryAfterSelectHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.BeforeUpdateHook, ethnicGroupCountryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryBeforeUpdateHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.AfterUpdateHook, ethnicGroupCountryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryAfterUpdateHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.BeforeDeleteHook, ethnicGroupCountryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryBeforeDeleteHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.AfterDeleteHook, ethnicGroupCountryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryAfterDeleteHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.BeforeUpsertHook, ethnicGroupCountryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryBeforeUpsertHooks = []EthnicGroupCountryHook{}

	AddEthnicGroupCountryHook(boil.AfterUpsertHook, ethnicGroupCountryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ethnicGroupCountryAfterUpsertHooks = []EthnicGroupCountryHook{}
}

func testEthnicGroupCountriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEthnicGroupCountriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ethnicGroupCountryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEthnicGroupCountryToOneEthnicGroupUsingEthnicGroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EthnicGroupCountry
	var foreign EthnicGroup

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, ethnicGroupDBTypes, false, ethnicGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroup struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EthnicGroupID = foreign.EthnicGroupID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.EthnicGroup().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.EthnicGroupID != foreign.EthnicGroupID {
		t.Errorf("want: %v, got %v", foreign.EthnicGroupID, check.EthnicGroupID)
	}

	slice := EthnicGroupCountrySlice{&local}
	if err = local.L.LoadEthnicGroup(ctx, tx, false, (*[]*EthnicGroupCountry)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.EthnicGroup == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.EthnicGroup = nil
	if err = local.L.LoadEthnicGroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.EthnicGroup == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEthnicGroupCountryToOneCountryUsingCountry(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EthnicGroupCountry
	var foreign Country

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ethnicGroupCountryDBTypes, false, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CountryID = foreign.CountryID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Country().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CountryID != foreign.CountryID {
		t.Errorf("want: %v, got %v", foreign.CountryID, check.CountryID)
	}

	slice := EthnicGroupCountrySlice{&local}
	if err = local.L.LoadCountry(ctx, tx, false, (*[]*EthnicGroupCountry)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Country == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Country = nil
	if err = local.L.LoadCountry(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Country == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEthnicGroupCountryToOneSetOpEthnicGroupUsingEthnicGroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EthnicGroupCountry
	var b, c EthnicGroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ethnicGroupCountryDBTypes, false, strmangle.SetComplement(ethnicGroupCountryPrimaryKeyColumns, ethnicGroupCountryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ethnicGroupDBTypes, false, strmangle.SetComplement(ethnicGroupPrimaryKeyColumns, ethnicGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ethnicGroupDBTypes, false, strmangle.SetComplement(ethnicGroupPrimaryKeyColumns, ethnicGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*EthnicGroup{&b, &c} {
		err = a.SetEthnicGroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.EthnicGroup != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EthnicGroupCountries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EthnicGroupID != x.EthnicGroupID {
			t.Error("foreign key was wrong value", a.EthnicGroupID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EthnicGroupID))
		reflect.Indirect(reflect.ValueOf(&a.EthnicGroupID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EthnicGroupID != x.EthnicGroupID {
			t.Error("foreign key was wrong value", a.EthnicGroupID, x.EthnicGroupID)
		}
	}
}
func testEthnicGroupCountryToOneSetOpCountryUsingCountry(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EthnicGroupCountry
	var b, c Country

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ethnicGroupCountryDBTypes, false, strmangle.SetComplement(ethnicGroupCountryPrimaryKeyColumns, ethnicGroupCountryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Country{&b, &c} {
		err = a.SetCountry(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Country != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EthnicGroupCountries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CountryID != x.CountryID {
			t.Error("foreign key was wrong value", a.CountryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CountryID))
		reflect.Indirect(reflect.ValueOf(&a.CountryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CountryID != x.CountryID {
			t.Error("foreign key was wrong value", a.CountryID, x.CountryID)
		}
	}
}

func testEthnicGroupCountriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEthnicGroupCountriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EthnicGroupCountrySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEthnicGroupCountriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EthnicGroupCountries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ethnicGroupCountryDBTypes = map[string]string{`CountryID`: `int8`, `EthnicGroupCountryID`: `int8`, `EthnicGroupID`: `int8`}
	_                         = bytes.MinRead
)

func testEthnicGroupCountriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ethnicGroupCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ethnicGroupCountryColumns) == len(ethnicGroupCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEthnicGroupCountriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ethnicGroupCountryColumns) == len(ethnicGroupCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EthnicGroupCountry{}
	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ethnicGroupCountryDBTypes, true, ethnicGroupCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ethnicGroupCountryColumns, ethnicGroupCountryPrimaryKeyColumns) {
		fields = ethnicGroupCountryColumns
	} else {
		fields = strmangle.SetComplement(
			ethnicGroupCountryColumns,
			ethnicGroupCountryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EthnicGroupCountrySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEthnicGroupCountriesUpsert(t *testing.T) {
	t.Parallel()

	if len(ethnicGroupCountryColumns) == len(ethnicGroupCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EthnicGroupCountry{}
	if err = randomize.Struct(seed, &o, ethnicGroupCountryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EthnicGroupCountry: %s", err)
	}

	count, err := EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ethnicGroupCountryDBTypes, false, ethnicGroupCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EthnicGroupCountry struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EthnicGroupCountry: %s", err)
	}

	count, err = EthnicGroupCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
