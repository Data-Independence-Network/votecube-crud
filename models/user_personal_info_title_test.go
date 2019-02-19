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

func testUserPersonalInfoTitles(t *testing.T) {
	t.Parallel()

	query := UserPersonalInfoTitles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserPersonalInfoTitlesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
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

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitlesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserPersonalInfoTitles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitlesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPersonalInfoTitleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPersonalInfoTitlesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserPersonalInfoTitleExists(ctx, tx, o.UserPersonalInfoTitleID)
	if err != nil {
		t.Errorf("Unable to check if UserPersonalInfoTitle exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserPersonalInfoTitleExists to return true, but got false.")
	}
}

func testUserPersonalInfoTitlesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userPersonalInfoTitleFound, err := FindUserPersonalInfoTitle(ctx, tx, o.UserPersonalInfoTitleID)
	if err != nil {
		t.Error(err)
	}

	if userPersonalInfoTitleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserPersonalInfoTitlesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserPersonalInfoTitles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserPersonalInfoTitlesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserPersonalInfoTitles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserPersonalInfoTitlesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userPersonalInfoTitleOne := &UserPersonalInfoTitle{}
	userPersonalInfoTitleTwo := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, userPersonalInfoTitleOne, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}
	if err = randomize.Struct(seed, userPersonalInfoTitleTwo, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPersonalInfoTitleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPersonalInfoTitleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPersonalInfoTitles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserPersonalInfoTitlesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userPersonalInfoTitleOne := &UserPersonalInfoTitle{}
	userPersonalInfoTitleTwo := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, userPersonalInfoTitleOne, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}
	if err = randomize.Struct(seed, userPersonalInfoTitleTwo, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPersonalInfoTitleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPersonalInfoTitleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userPersonalInfoTitleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func userPersonalInfoTitleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPersonalInfoTitle) error {
	*o = UserPersonalInfoTitle{}
	return nil
}

func testUserPersonalInfoTitlesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserPersonalInfoTitle{}
	o := &UserPersonalInfoTitle{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle object: %s", err)
	}

	AddUserPersonalInfoTitleHook(boil.BeforeInsertHook, userPersonalInfoTitleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleBeforeInsertHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.AfterInsertHook, userPersonalInfoTitleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterInsertHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.AfterSelectHook, userPersonalInfoTitleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterSelectHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.BeforeUpdateHook, userPersonalInfoTitleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleBeforeUpdateHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.AfterUpdateHook, userPersonalInfoTitleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterUpdateHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.BeforeDeleteHook, userPersonalInfoTitleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleBeforeDeleteHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.AfterDeleteHook, userPersonalInfoTitleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterDeleteHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.BeforeUpsertHook, userPersonalInfoTitleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleBeforeUpsertHooks = []UserPersonalInfoTitleHook{}

	AddUserPersonalInfoTitleHook(boil.AfterUpsertHook, userPersonalInfoTitleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userPersonalInfoTitleAfterUpsertHooks = []UserPersonalInfoTitleHook{}
}

func testUserPersonalInfoTitlesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPersonalInfoTitlesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userPersonalInfoTitleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPersonalInfoTitleToOneUserPersonalInfoUsingUserPersonalInfo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPersonalInfoTitle
	var foreign UserPersonalInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
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

	slice := UserPersonalInfoTitleSlice{&local}
	if err = local.L.LoadUserPersonalInfo(ctx, tx, false, (*[]*UserPersonalInfoTitle)(&slice), nil); err != nil {
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

func testUserPersonalInfoTitleToOneTitleUsingTitle(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPersonalInfoTitle
	var foreign Title

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, titleDBTypes, false, titleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Title struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TitleID = foreign.TitleID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Title().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.TitleID != foreign.TitleID {
		t.Errorf("want: %v, got %v", foreign.TitleID, check.TitleID)
	}

	slice := UserPersonalInfoTitleSlice{&local}
	if err = local.L.LoadTitle(ctx, tx, false, (*[]*UserPersonalInfoTitle)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Title == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Title = nil
	if err = local.L.LoadTitle(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Title == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserPersonalInfoTitleToOneSetOpUserPersonalInfoUsingUserPersonalInfo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPersonalInfoTitle
	var b, c UserPersonalInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPersonalInfoTitleDBTypes, false, strmangle.SetComplement(userPersonalInfoTitlePrimaryKeyColumns, userPersonalInfoTitleColumnsWithoutDefault)...); err != nil {
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

		if x.R.UserPersonalInfoTitles[0] != &a {
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
func testUserPersonalInfoTitleToOneSetOpTitleUsingTitle(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPersonalInfoTitle
	var b, c Title

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPersonalInfoTitleDBTypes, false, strmangle.SetComplement(userPersonalInfoTitlePrimaryKeyColumns, userPersonalInfoTitleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, titleDBTypes, false, strmangle.SetComplement(titlePrimaryKeyColumns, titleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, titleDBTypes, false, strmangle.SetComplement(titlePrimaryKeyColumns, titleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Title{&b, &c} {
		err = a.SetTitle(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Title != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserPersonalInfoTitles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TitleID != x.TitleID {
			t.Error("foreign key was wrong value", a.TitleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TitleID))
		reflect.Indirect(reflect.ValueOf(&a.TitleID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TitleID != x.TitleID {
			t.Error("foreign key was wrong value", a.TitleID, x.TitleID)
		}
	}
}

func testUserPersonalInfoTitlesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
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

func testUserPersonalInfoTitlesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPersonalInfoTitleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserPersonalInfoTitlesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPersonalInfoTitles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userPersonalInfoTitleDBTypes = map[string]string{`TitleID`: `int8`, `TitlePosition`: `int2`, `UserPersonalInfoID`: `int8`, `UserPersonalInfoTitleID`: `int8`}
	_                            = bytes.MinRead
)

func testUserPersonalInfoTitlesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userPersonalInfoTitlePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userPersonalInfoTitleColumns) == len(userPersonalInfoTitlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserPersonalInfoTitlesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userPersonalInfoTitleColumns) == len(userPersonalInfoTitlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPersonalInfoTitleDBTypes, true, userPersonalInfoTitlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userPersonalInfoTitleColumns, userPersonalInfoTitlePrimaryKeyColumns) {
		fields = userPersonalInfoTitleColumns
	} else {
		fields = strmangle.SetComplement(
			userPersonalInfoTitleColumns,
			userPersonalInfoTitlePrimaryKeyColumns,
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

	slice := UserPersonalInfoTitleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserPersonalInfoTitlesUpsert(t *testing.T) {
	t.Parallel()

	if len(userPersonalInfoTitleColumns) == len(userPersonalInfoTitlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserPersonalInfoTitle{}
	if err = randomize.Struct(seed, &o, userPersonalInfoTitleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPersonalInfoTitle: %s", err)
	}

	count, err := UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userPersonalInfoTitleDBTypes, false, userPersonalInfoTitlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPersonalInfoTitle struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPersonalInfoTitle: %s", err)
	}

	count, err = UserPersonalInfoTitles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
