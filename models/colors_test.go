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

func testColors(t *testing.T) {
	t.Parallel()

	query := Colors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testColorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
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

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Colors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ColorExists(ctx, tx, o.ColorID)
	if err != nil {
		t.Errorf("Unable to check if Color exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ColorExists to return true, but got false.")
	}
}

func testColorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	colorFound, err := FindColor(ctx, tx, o.ColorID)
	if err != nil {
		t.Error(err)
	}

	if colorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testColorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Colors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testColorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Colors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testColorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	colorOne := &Color{}
	colorTwo := &Color{}
	if err = randomize.Struct(seed, colorOne, colorDBTypes, false, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}
	if err = randomize.Struct(seed, colorTwo, colorDBTypes, false, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Colors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testColorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	colorOne := &Color{}
	colorTwo := &Color{}
	if err = randomize.Struct(seed, colorOne, colorDBTypes, false, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}
	if err = randomize.Struct(seed, colorTwo, colorDBTypes, false, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func colorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func colorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Color) error {
	*o = Color{}
	return nil
}

func testColorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Color{}
	o := &Color{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, colorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Color object: %s", err)
	}

	AddColorHook(boil.BeforeInsertHook, colorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	colorBeforeInsertHooks = []ColorHook{}

	AddColorHook(boil.AfterInsertHook, colorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	colorAfterInsertHooks = []ColorHook{}

	AddColorHook(boil.AfterSelectHook, colorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	colorAfterSelectHooks = []ColorHook{}

	AddColorHook(boil.BeforeUpdateHook, colorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	colorBeforeUpdateHooks = []ColorHook{}

	AddColorHook(boil.AfterUpdateHook, colorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	colorAfterUpdateHooks = []ColorHook{}

	AddColorHook(boil.BeforeDeleteHook, colorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	colorBeforeDeleteHooks = []ColorHook{}

	AddColorHook(boil.AfterDeleteHook, colorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	colorAfterDeleteHooks = []ColorHook{}

	AddColorHook(boil.BeforeUpsertHook, colorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	colorBeforeUpsertHooks = []ColorHook{}

	AddColorHook(boil.AfterUpsertHook, colorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	colorAfterUpsertHooks = []ColorHook{}
}

func testColorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(colorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColorToManyPollsFactorsPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Color
	var b, c PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsFactorsPositionDBTypes, false, pollsFactorsPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsFactorsPositionDBTypes, false, pollsFactorsPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ColorID = a.ColorID
	c.ColorID = a.ColorID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsFactorsPosition, err := a.PollsFactorsPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsFactorsPosition {
		if v.ColorID == b.ColorID {
			bFound = true
		}
		if v.ColorID == c.ColorID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ColorSlice{&a}
	if err = a.L.LoadPollsFactorsPositions(ctx, tx, false, (*[]*Color)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsFactorsPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsFactorsPositions = nil
	if err = a.L.LoadPollsFactorsPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsFactorsPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsFactorsPosition)
	}
}

func testColorToManyAddOpPollsFactorsPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Color
	var b, c, d, e PollsFactorsPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorDBTypes, false, strmangle.SetComplement(colorPrimaryKeyColumns, colorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsFactorsPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsFactorsPositionDBTypes, false, strmangle.SetComplement(pollsFactorsPositionPrimaryKeyColumns, pollsFactorsPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PollsFactorsPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsFactorsPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ColorID != first.ColorID {
			t.Error("foreign key was wrong value", a.ColorID, first.ColorID)
		}
		if a.ColorID != second.ColorID {
			t.Error("foreign key was wrong value", a.ColorID, second.ColorID)
		}

		if first.R.Color != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Color != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsFactorsPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsFactorsPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsFactorsPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testColorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
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

func testColorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testColorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Colors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	colorDBTypes = map[string]string{`ColorID`: `int8`, `RGBHexValue`: `varchar`}
	_            = bytes.MinRead
)

func testColorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(colorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(colorColumns) == len(colorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colorDBTypes, true, colorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testColorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(colorColumns) == len(colorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Color{}
	if err = randomize.Struct(seed, o, colorDBTypes, true, colorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colorDBTypes, true, colorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(colorColumns, colorPrimaryKeyColumns) {
		fields = colorColumns
	} else {
		fields = strmangle.SetComplement(
			colorColumns,
			colorPrimaryKeyColumns,
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

	slice := ColorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testColorsUpsert(t *testing.T) {
	t.Parallel()

	if len(colorColumns) == len(colorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Color{}
	if err = randomize.Struct(seed, &o, colorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Color: %s", err)
	}

	count, err := Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, colorDBTypes, false, colorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Color struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Color: %s", err)
	}

	count, err = Colors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
