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

func testHonors(t *testing.T) {
	t.Parallel()

	query := Honors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHonorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
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

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHonorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Honors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHonorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HonorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHonorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HonorExists(ctx, tx, o.HonorID)
	if err != nil {
		t.Errorf("Unable to check if Honor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HonorExists to return true, but got false.")
	}
}

func testHonorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	honorFound, err := FindHonor(ctx, tx, o.HonorID)
	if err != nil {
		t.Error(err)
	}

	if honorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHonorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Honors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHonorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Honors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHonorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	honorOne := &Honor{}
	honorTwo := &Honor{}
	if err = randomize.Struct(seed, honorOne, honorDBTypes, false, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}
	if err = randomize.Struct(seed, honorTwo, honorDBTypes, false, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = honorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = honorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Honors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHonorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	honorOne := &Honor{}
	honorTwo := &Honor{}
	if err = randomize.Struct(seed, honorOne, honorDBTypes, false, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}
	if err = randomize.Struct(seed, honorTwo, honorDBTypes, false, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = honorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = honorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func honorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func honorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Honor) error {
	*o = Honor{}
	return nil
}

func testHonorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Honor{}
	o := &Honor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, honorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Honor object: %s", err)
	}

	AddHonorHook(boil.BeforeInsertHook, honorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	honorBeforeInsertHooks = []HonorHook{}

	AddHonorHook(boil.AfterInsertHook, honorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	honorAfterInsertHooks = []HonorHook{}

	AddHonorHook(boil.AfterSelectHook, honorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	honorAfterSelectHooks = []HonorHook{}

	AddHonorHook(boil.BeforeUpdateHook, honorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	honorBeforeUpdateHooks = []HonorHook{}

	AddHonorHook(boil.AfterUpdateHook, honorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	honorAfterUpdateHooks = []HonorHook{}

	AddHonorHook(boil.BeforeDeleteHook, honorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	honorBeforeDeleteHooks = []HonorHook{}

	AddHonorHook(boil.AfterDeleteHook, honorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	honorAfterDeleteHooks = []HonorHook{}

	AddHonorHook(boil.BeforeUpsertHook, honorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	honorBeforeUpsertHooks = []HonorHook{}

	AddHonorHook(boil.AfterUpsertHook, honorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	honorAfterUpsertHooks = []HonorHook{}
}

func testHonorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHonorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(honorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHonorToManyUserPersonalInfoHonors(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Honor
	var b, c UserPersonalInfoHonor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userPersonalInfoHonorDBTypes, false, userPersonalInfoHonorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userPersonalInfoHonorDBTypes, false, userPersonalInfoHonorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HonorID = a.HonorID
	c.HonorID = a.HonorID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	userPersonalInfoHonor, err := a.UserPersonalInfoHonors().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userPersonalInfoHonor {
		if v.HonorID == b.HonorID {
			bFound = true
		}
		if v.HonorID == c.HonorID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HonorSlice{&a}
	if err = a.L.LoadUserPersonalInfoHonors(ctx, tx, false, (*[]*Honor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserPersonalInfoHonors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserPersonalInfoHonors = nil
	if err = a.L.LoadUserPersonalInfoHonors(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserPersonalInfoHonors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userPersonalInfoHonor)
	}
}

func testHonorToManyAddOpUserPersonalInfoHonors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Honor
	var b, c, d, e UserPersonalInfoHonor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, honorDBTypes, false, strmangle.SetComplement(honorPrimaryKeyColumns, honorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserPersonalInfoHonor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userPersonalInfoHonorDBTypes, false, strmangle.SetComplement(userPersonalInfoHonorPrimaryKeyColumns, userPersonalInfoHonorColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserPersonalInfoHonor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserPersonalInfoHonors(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.HonorID != first.HonorID {
			t.Error("foreign key was wrong value", a.HonorID, first.HonorID)
		}
		if a.HonorID != second.HonorID {
			t.Error("foreign key was wrong value", a.HonorID, second.HonorID)
		}

		if first.R.Honor != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Honor != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserPersonalInfoHonors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserPersonalInfoHonors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserPersonalInfoHonors().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testHonorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
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

func testHonorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HonorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHonorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Honors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	honorDBTypes = map[string]string{`HonorID`: `int8`, `HonorName`: `varchar`}
	_            = bytes.MinRead
)

func testHonorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(honorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(honorColumns) == len(honorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, honorDBTypes, true, honorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHonorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(honorColumns) == len(honorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Honor{}
	if err = randomize.Struct(seed, o, honorDBTypes, true, honorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, honorDBTypes, true, honorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(honorColumns, honorPrimaryKeyColumns) {
		fields = honorColumns
	} else {
		fields = strmangle.SetComplement(
			honorColumns,
			honorPrimaryKeyColumns,
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

	slice := HonorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHonorsUpsert(t *testing.T) {
	t.Parallel()

	if len(honorColumns) == len(honorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Honor{}
	if err = randomize.Struct(seed, &o, honorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Honor: %s", err)
	}

	count, err := Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, honorDBTypes, false, honorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Honor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Honor: %s", err)
	}

	count, err = Honors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
