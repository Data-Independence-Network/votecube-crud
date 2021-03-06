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

func testStreetNames(t *testing.T) {
	t.Parallel()

	query := StreetNames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStreetNamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
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

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStreetNamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StreetNames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStreetNamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StreetNameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStreetNamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StreetNameExists(ctx, tx, o.StreetNameID)
	if err != nil {
		t.Errorf("Unable to check if StreetName exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StreetNameExists to return true, but got false.")
	}
}

func testStreetNamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	streetNameFound, err := FindStreetName(ctx, tx, o.StreetNameID)
	if err != nil {
		t.Error(err)
	}

	if streetNameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStreetNamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StreetNames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStreetNamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StreetNames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStreetNamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	streetNameOne := &StreetName{}
	streetNameTwo := &StreetName{}
	if err = randomize.Struct(seed, streetNameOne, streetNameDBTypes, false, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}
	if err = randomize.Struct(seed, streetNameTwo, streetNameDBTypes, false, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = streetNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = streetNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StreetNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStreetNamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	streetNameOne := &StreetName{}
	streetNameTwo := &StreetName{}
	if err = randomize.Struct(seed, streetNameOne, streetNameDBTypes, false, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}
	if err = randomize.Struct(seed, streetNameTwo, streetNameDBTypes, false, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = streetNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = streetNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func streetNameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func streetNameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StreetName) error {
	*o = StreetName{}
	return nil
}

func testStreetNamesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StreetName{}
	o := &StreetName{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, streetNameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StreetName object: %s", err)
	}

	AddStreetNameHook(boil.BeforeInsertHook, streetNameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	streetNameBeforeInsertHooks = []StreetNameHook{}

	AddStreetNameHook(boil.AfterInsertHook, streetNameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	streetNameAfterInsertHooks = []StreetNameHook{}

	AddStreetNameHook(boil.AfterSelectHook, streetNameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	streetNameAfterSelectHooks = []StreetNameHook{}

	AddStreetNameHook(boil.BeforeUpdateHook, streetNameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	streetNameBeforeUpdateHooks = []StreetNameHook{}

	AddStreetNameHook(boil.AfterUpdateHook, streetNameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	streetNameAfterUpdateHooks = []StreetNameHook{}

	AddStreetNameHook(boil.BeforeDeleteHook, streetNameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	streetNameBeforeDeleteHooks = []StreetNameHook{}

	AddStreetNameHook(boil.AfterDeleteHook, streetNameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	streetNameAfterDeleteHooks = []StreetNameHook{}

	AddStreetNameHook(boil.BeforeUpsertHook, streetNameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	streetNameBeforeUpsertHooks = []StreetNameHook{}

	AddStreetNameHook(boil.AfterUpsertHook, streetNameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	streetNameAfterUpsertHooks = []StreetNameHook{}
}

func testStreetNamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStreetNamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(streetNameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStreetNameToManyStreets(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StreetName
	var b, c Street

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, streetDBTypes, false, streetColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, streetDBTypes, false, streetColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.StreetNameID = a.StreetNameID
	c.StreetNameID = a.StreetNameID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	street, err := a.Streets().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range street {
		if v.StreetNameID == b.StreetNameID {
			bFound = true
		}
		if v.StreetNameID == c.StreetNameID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := StreetNameSlice{&a}
	if err = a.L.LoadStreets(ctx, tx, false, (*[]*StreetName)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Streets); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Streets = nil
	if err = a.L.LoadStreets(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Streets); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", street)
	}
}

func testStreetNameToManyAddOpStreets(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StreetName
	var b, c, d, e Street

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, streetNameDBTypes, false, strmangle.SetComplement(streetNamePrimaryKeyColumns, streetNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Street{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, streetDBTypes, false, strmangle.SetComplement(streetPrimaryKeyColumns, streetColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Street{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStreets(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.StreetNameID != first.StreetNameID {
			t.Error("foreign key was wrong value", a.StreetNameID, first.StreetNameID)
		}
		if a.StreetNameID != second.StreetNameID {
			t.Error("foreign key was wrong value", a.StreetNameID, second.StreetNameID)
		}

		if first.R.StreetName != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.StreetName != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Streets[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Streets[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Streets().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testStreetNamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
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

func testStreetNamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StreetNameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStreetNamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StreetNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	streetNameDBTypes = map[string]string{`StreetName`: `varchar`, `StreetNameID`: `int8`}
	_                 = bytes.MinRead
)

func testStreetNamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(streetNamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(streetNameColumns) == len(streetNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStreetNamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(streetNameColumns) == len(streetNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StreetName{}
	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, streetNameDBTypes, true, streetNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(streetNameColumns, streetNamePrimaryKeyColumns) {
		fields = streetNameColumns
	} else {
		fields = strmangle.SetComplement(
			streetNameColumns,
			streetNamePrimaryKeyColumns,
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

	slice := StreetNameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStreetNamesUpsert(t *testing.T) {
	t.Parallel()

	if len(streetNameColumns) == len(streetNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StreetName{}
	if err = randomize.Struct(seed, &o, streetNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StreetName: %s", err)
	}

	count, err := StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, streetNameDBTypes, false, streetNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StreetName struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StreetName: %s", err)
	}

	count, err = StreetNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
