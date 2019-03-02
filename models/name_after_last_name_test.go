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

func testNameAfterLastNames(t *testing.T) {
	t.Parallel()

	query := NameAfterLastNames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNameAfterLastNamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
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

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNameAfterLastNamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NameAfterLastNames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNameAfterLastNamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NameAfterLastNameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNameAfterLastNamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NameAfterLastNameExists(ctx, tx, o.NameAfterLastNameID)
	if err != nil {
		t.Errorf("Unable to check if NameAfterLastName exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NameAfterLastNameExists to return true, but got false.")
	}
}

func testNameAfterLastNamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nameAfterLastNameFound, err := FindNameAfterLastName(ctx, tx, o.NameAfterLastNameID)
	if err != nil {
		t.Error(err)
	}

	if nameAfterLastNameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNameAfterLastNamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NameAfterLastNames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNameAfterLastNamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NameAfterLastNames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNameAfterLastNamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nameAfterLastNameOne := &NameAfterLastName{}
	nameAfterLastNameTwo := &NameAfterLastName{}
	if err = randomize.Struct(seed, nameAfterLastNameOne, nameAfterLastNameDBTypes, false, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}
	if err = randomize.Struct(seed, nameAfterLastNameTwo, nameAfterLastNameDBTypes, false, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nameAfterLastNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nameAfterLastNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NameAfterLastNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNameAfterLastNamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nameAfterLastNameOne := &NameAfterLastName{}
	nameAfterLastNameTwo := &NameAfterLastName{}
	if err = randomize.Struct(seed, nameAfterLastNameOne, nameAfterLastNameDBTypes, false, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}
	if err = randomize.Struct(seed, nameAfterLastNameTwo, nameAfterLastNameDBTypes, false, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nameAfterLastNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nameAfterLastNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func nameAfterLastNameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func nameAfterLastNameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NameAfterLastName) error {
	*o = NameAfterLastName{}
	return nil
}

func testNameAfterLastNamesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NameAfterLastName{}
	o := &NameAfterLastName{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName object: %s", err)
	}

	AddNameAfterLastNameHook(boil.BeforeInsertHook, nameAfterLastNameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameBeforeInsertHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.AfterInsertHook, nameAfterLastNameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameAfterInsertHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.AfterSelectHook, nameAfterLastNameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameAfterSelectHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.BeforeUpdateHook, nameAfterLastNameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameBeforeUpdateHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.AfterUpdateHook, nameAfterLastNameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameAfterUpdateHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.BeforeDeleteHook, nameAfterLastNameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameBeforeDeleteHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.AfterDeleteHook, nameAfterLastNameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameAfterDeleteHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.BeforeUpsertHook, nameAfterLastNameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameBeforeUpsertHooks = []NameAfterLastNameHook{}

	AddNameAfterLastNameHook(boil.AfterUpsertHook, nameAfterLastNameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	nameAfterLastNameAfterUpsertHooks = []NameAfterLastNameHook{}
}

func testNameAfterLastNamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNameAfterLastNamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nameAfterLastNameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNameAfterLastNameToManyUserAccounts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NameAfterLastName
	var b, c UserAccount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userAccountDBTypes, false, userAccountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userAccountDBTypes, false, userAccountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.NameAfterLastNameID = a.NameAfterLastNameID
	c.NameAfterLastNameID = a.NameAfterLastNameID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	userAccount, err := a.UserAccounts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userAccount {
		if v.NameAfterLastNameID == b.NameAfterLastNameID {
			bFound = true
		}
		if v.NameAfterLastNameID == c.NameAfterLastNameID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := NameAfterLastNameSlice{&a}
	if err = a.L.LoadUserAccounts(ctx, tx, false, (*[]*NameAfterLastName)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAccounts = nil
	if err = a.L.LoadUserAccounts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userAccount)
	}
}

func testNameAfterLastNameToManyAddOpUserAccounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NameAfterLastName
	var b, c, d, e UserAccount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, nameAfterLastNameDBTypes, false, strmangle.SetComplement(nameAfterLastNamePrimaryKeyColumns, nameAfterLastNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserAccount{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userAccountDBTypes, false, strmangle.SetComplement(userAccountPrimaryKeyColumns, userAccountColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserAccount{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAccounts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.NameAfterLastNameID != first.NameAfterLastNameID {
			t.Error("foreign key was wrong value", a.NameAfterLastNameID, first.NameAfterLastNameID)
		}
		if a.NameAfterLastNameID != second.NameAfterLastNameID {
			t.Error("foreign key was wrong value", a.NameAfterLastNameID, second.NameAfterLastNameID)
		}

		if first.R.NameAfterLastName != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.NameAfterLastName != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAccounts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAccounts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAccounts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testNameAfterLastNamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
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

func testNameAfterLastNamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NameAfterLastNameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNameAfterLastNamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NameAfterLastNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nameAfterLastNameDBTypes = map[string]string{`NameAfterLastName`: `varchar`, `NameAfterLastNameID`: `int8`}
	_                        = bytes.MinRead
)

func testNameAfterLastNamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nameAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nameAfterLastNameColumns) == len(nameAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNameAfterLastNamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nameAfterLastNameColumns) == len(nameAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NameAfterLastName{}
	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nameAfterLastNameDBTypes, true, nameAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nameAfterLastNameColumns, nameAfterLastNamePrimaryKeyColumns) {
		fields = nameAfterLastNameColumns
	} else {
		fields = strmangle.SetComplement(
			nameAfterLastNameColumns,
			nameAfterLastNamePrimaryKeyColumns,
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

	slice := NameAfterLastNameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNameAfterLastNamesUpsert(t *testing.T) {
	t.Parallel()

	if len(nameAfterLastNameColumns) == len(nameAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NameAfterLastName{}
	if err = randomize.Struct(seed, &o, nameAfterLastNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NameAfterLastName: %s", err)
	}

	count, err := NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, nameAfterLastNameDBTypes, false, nameAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NameAfterLastName struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NameAfterLastName: %s", err)
	}

	count, err = NameAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}