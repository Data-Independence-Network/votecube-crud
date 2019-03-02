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

func testUserPersonalInfoTitleAfterLastNames(t *testing.T) {
	t.Parallel()

	query := UserPersonalInfoTitleAfterLastNames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserPersonalInfoTitleAfterLastNamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
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

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitleAfterLastNamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserPersonalInfoTitleAfterLastNames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitleAfterLastNamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPersonalInfoTitleAfterLastNameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitleAfterLastNamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserPersonalInfoTitleAfterLastNameExists(ctx, tx, o.UserPersonalInfoTitleAfterLastNameID)
	if err != nil {
		t.Errorf("Unable to check if UserPersonalInfoTitleAfterLastName exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserPersonalInfoTitleAfterLastNameExists to return true, but got false.")
	}
}

func testUserPersonalInfoTitleAfterLastNamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userPersonalInfoTitleAfterLastNameFound, err := FindUserPersonalInfoTitleAfterLastName(ctx, tx, o.UserPersonalInfoTitleAfterLastNameID)
	if err != nil {
		t.Error(err)
	}

	if userPersonalInfoTitleAfterLastNameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserPersonalInfoTitleAfterLastNamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserPersonalInfoTitleAfterLastNames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserPersonalInfoTitleAfterLastNamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserPersonalInfoTitleAfterLastNames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserPersonalInfoTitleAfterLastNamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userPersonalInfoTitleAfterLastNameOne := &UserPersonalInfoTitleAfterLastName{}
	userPersonalInfoTitleAfterLastNameTwo := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, userPersonalInfoTitleAfterLastNameOne, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}
	if err = randomize.Struct(seed, userPersonalInfoTitleAfterLastNameTwo, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPersonalInfoTitleAfterLastNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPersonalInfoTitleAfterLastNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPersonalInfoTitleAfterLastNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserPersonalInfoTitleAfterLastNamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userPersonalInfoTitleAfterLastNameOne := &UserPersonalInfoTitleAfterLastName{}
	userPersonalInfoTitleAfterLastNameTwo := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, userPersonalInfoTitleAfterLastNameOne, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}
	if err = randomize.Struct(seed, userPersonalInfoTitleAfterLastNameTwo, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPersonalInfoTitleAfterLastNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPersonalInfoTitleAfterLastNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userPersonalInfoTitleAfterLastNameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func userPersonalInfoTitleAfterLastNameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitleAfterLastName) error {
	*o = UserPersonalInfoTitleAfterLastName{}
	return nil
}

func testUserPersonalInfoTitleAfterLastNamesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserPersonalInfoTitleAfterLastName{}
	o := &UserPersonalInfoTitleAfterLastName{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName object: %s", err)
	}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.BeforeInsertHook, userPersonalInfoTitleAfterLastNameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameBeforeInsertHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.AfterInsertHook, userPersonalInfoTitleAfterLastNameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameAfterInsertHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.AfterSelectHook, userPersonalInfoTitleAfterLastNameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameAfterSelectHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.BeforeUpdateHook, userPersonalInfoTitleAfterLastNameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameBeforeUpdateHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.AfterUpdateHook, userPersonalInfoTitleAfterLastNameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameAfterUpdateHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.BeforeDeleteHook, userPersonalInfoTitleAfterLastNameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameBeforeDeleteHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.AfterDeleteHook, userPersonalInfoTitleAfterLastNameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameAfterDeleteHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.BeforeUpsertHook, userPersonalInfoTitleAfterLastNameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameBeforeUpsertHooks = []UserPersonalInfoTitleAfterLastNameHook{}

	AddUserPersonalInfoTitleAfterLastNameHook(boil.AfterUpsertHook, userPersonalInfoTitleAfterLastNameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterLastNameAfterUpsertHooks = []UserPersonalInfoTitleAfterLastNameHook{}
}

func testUserPersonalInfoTitleAfterLastNamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPersonalInfoTitleAfterLastNamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userPersonalInfoTitleAfterLastNameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPersonalInfoTitleAfterLastNameToOneUserPersonalInfoUsingUserPersonalInfo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPersonalInfoTitleAfterLastName
	var foreign UserPersonalInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userPersonalInfoDBTypes, false, userPersonalInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserPersonalInfoID = foreign.UserPersonalInfoID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UserPersonalInfo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserPersonalInfoID != foreign.UserPersonalInfoID {
		t.Errorf("want: %v, got %v", foreign.UserPersonalInfoID, check.UserPersonalInfoID)
	}

	slice := UserPersonalInfoTitleAfterLastNameSlice{&local}
	if err = local.L.LoadUserPersonalInfo(ctx, tx, false, (*[]*UserPersonalInfoTitleAfterLastName)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserPersonalInfo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UserPersonalInfo = nil
	if err = local.L.LoadUserPersonalInfo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserPersonalInfo == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserPersonalInfoTitleAfterLastNameToOneTitleAfterLastNameUsingTitleAfterLastName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPersonalInfoTitleAfterLastName
	var foreign TitleAfterLastName

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, titleAfterLastNameDBTypes, false, titleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TitleAfterLastName struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TitleAfterLastNameID = foreign.TitleAfterLastNameID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TitleAfterLastName().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.TitleAfterLastNameID != foreign.TitleAfterLastNameID {
		t.Errorf("want: %v, got %v", foreign.TitleAfterLastNameID, check.TitleAfterLastNameID)
	}

	slice := UserPersonalInfoTitleAfterLastNameSlice{&local}
	if err = local.L.LoadTitleAfterLastName(ctx, tx, false, (*[]*UserPersonalInfoTitleAfterLastName)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TitleAfterLastName == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TitleAfterLastName = nil
	if err = local.L.LoadTitleAfterLastName(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TitleAfterLastName == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserPersonalInfoTitleAfterLastNameToOneSetOpUserPersonalInfoUsingUserPersonalInfo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPersonalInfoTitleAfterLastName
	var b, c UserPersonalInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPersonalInfoTitleAfterLastNameDBTypes, false, strmangle.SetComplement(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns, userPersonalInfoTitleAfterLastNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userPersonalInfoDBTypes, false, strmangle.SetComplement(userPersonalInfoPrimaryKeyColumns, userPersonalInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userPersonalInfoDBTypes, false, strmangle.SetComplement(userPersonalInfoPrimaryKeyColumns, userPersonalInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserPersonalInfo{&b, &c} {
		err = a.SetUserPersonalInfo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UserPersonalInfo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserPersonalInfoTitleAfterLastNames[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserPersonalInfoID != x.UserPersonalInfoID {
			t.Error("foreign key was wrong value", a.UserPersonalInfoID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserPersonalInfoID))
		reflect.Indirect(reflect.ValueOf(&a.UserPersonalInfoID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserPersonalInfoID != x.UserPersonalInfoID {
			t.Error("foreign key was wrong value", a.UserPersonalInfoID, x.UserPersonalInfoID)
		}
	}
}
func testUserPersonalInfoTitleAfterLastNameToOneSetOpTitleAfterLastNameUsingTitleAfterLastName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPersonalInfoTitleAfterLastName
	var b, c TitleAfterLastName

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPersonalInfoTitleAfterLastNameDBTypes, false, strmangle.SetComplement(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns, userPersonalInfoTitleAfterLastNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, titleAfterLastNameDBTypes, false, strmangle.SetComplement(titleAfterLastNamePrimaryKeyColumns, titleAfterLastNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, titleAfterLastNameDBTypes, false, strmangle.SetComplement(titleAfterLastNamePrimaryKeyColumns, titleAfterLastNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TitleAfterLastName{&b, &c} {
		err = a.SetTitleAfterLastName(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TitleAfterLastName != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserPersonalInfoTitleAfterLastNames[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TitleAfterLastNameID != x.TitleAfterLastNameID {
			t.Error("foreign key was wrong value", a.TitleAfterLastNameID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TitleAfterLastNameID))
		reflect.Indirect(reflect.ValueOf(&a.TitleAfterLastNameID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TitleAfterLastNameID != x.TitleAfterLastNameID {
			t.Error("foreign key was wrong value", a.TitleAfterLastNameID, x.TitleAfterLastNameID)
		}
	}
}

func testUserPersonalInfoTitleAfterLastNamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
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

func testUserPersonalInfoTitleAfterLastNamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPersonalInfoTitleAfterLastNameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserPersonalInfoTitleAfterLastNamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPersonalInfoTitleAfterLastNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userPersonalInfoTitleAfterLastNameDBTypes = map[string]string{`TitleAfterLastNameID`: `int8`, `TitleAfterLastNamePosition`: `int2`, `UserPersonalInfoID`: `int8`, `UserPersonalInfoTitleAfterLastNameID`: `int8`}
	_                                         = bytes.MinRead
)

func testUserPersonalInfoTitleAfterLastNamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userPersonalInfoTitleAfterLastNameColumns) == len(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserPersonalInfoTitleAfterLastNamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userPersonalInfoTitleAfterLastNameColumns) == len(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPersonalInfoTitleAfterLastNameDBTypes, true, userPersonalInfoTitleAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userPersonalInfoTitleAfterLastNameColumns, userPersonalInfoTitleAfterLastNamePrimaryKeyColumns) {
		fields = userPersonalInfoTitleAfterLastNameColumns
	} else {
		fields = strmangle.SetComplement(
			userPersonalInfoTitleAfterLastNameColumns,
			userPersonalInfoTitleAfterLastNamePrimaryKeyColumns,
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

	slice := UserPersonalInfoTitleAfterLastNameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserPersonalInfoTitleAfterLastNamesUpsert(t *testing.T) {
	t.Parallel()

	if len(userPersonalInfoTitleAfterLastNameColumns) == len(userPersonalInfoTitleAfterLastNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserPersonalInfoTitleAfterLastName{}
	if err = randomize.Struct(seed, &o, userPersonalInfoTitleAfterLastNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPersonalInfoTitleAfterLastName: %s", err)
	}

	count, err := UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userPersonalInfoTitleAfterLastNameDBTypes, false, userPersonalInfoTitleAfterLastNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitleAfterLastName struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPersonalInfoTitleAfterLastName: %s", err)
	}

	count, err = UserPersonalInfoTitleAfterLastNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}