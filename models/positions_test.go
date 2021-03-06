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

func testPositions(t *testing.T) {
	t.Parallel()

	query := Positions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPositionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
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

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Positions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PositionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PositionExists(ctx, tx, o.PositionID)
	if err != nil {
		t.Errorf("Unable to check if Position exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PositionExists to return true, but got false.")
	}
}

func testPositionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	positionFound, err := FindPosition(ctx, tx, o.PositionID)
	if err != nil {
		t.Error(err)
	}

	if positionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPositionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Positions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPositionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Positions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPositionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	positionOne := &Position{}
	positionTwo := &Position{}
	if err = randomize.Struct(seed, positionOne, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err = randomize.Struct(seed, positionTwo, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = positionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = positionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Positions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPositionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	positionOne := &Position{}
	positionTwo := &Position{}
	if err = randomize.Struct(seed, positionOne, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err = randomize.Struct(seed, positionTwo, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = positionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = positionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func positionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func testPositionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Position{}
	o := &Position{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, positionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Position object: %s", err)
	}

	AddPositionHook(boil.BeforeInsertHook, positionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	positionBeforeInsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterInsertHook, positionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	positionAfterInsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterSelectHook, positionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	positionAfterSelectHooks = []PositionHook{}

	AddPositionHook(boil.BeforeUpdateHook, positionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	positionBeforeUpdateHooks = []PositionHook{}

	AddPositionHook(boil.AfterUpdateHook, positionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	positionAfterUpdateHooks = []PositionHook{}

	AddPositionHook(boil.BeforeDeleteHook, positionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	positionBeforeDeleteHooks = []PositionHook{}

	AddPositionHook(boil.AfterDeleteHook, positionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	positionAfterDeleteHooks = []PositionHook{}

	AddPositionHook(boil.BeforeUpsertHook, positionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	positionBeforeUpsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterUpsertHook, positionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	positionAfterUpsertHooks = []PositionHook{}
}

func testPositionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPositionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(positionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPositionToManyFactorPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c FactorPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, factorPositionDBTypes, false, factorPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, factorPositionDBTypes, false, factorPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PositionID = a.PositionID
	c.PositionID = a.PositionID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	factorPosition, err := a.FactorPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range factorPosition {
		if v.PositionID == b.PositionID {
			bFound = true
		}
		if v.PositionID == c.PositionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PositionSlice{&a}
	if err = a.L.LoadFactorPositions(ctx, tx, false, (*[]*Position)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FactorPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FactorPositions = nil
	if err = a.L.LoadFactorPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FactorPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", factorPosition)
	}
}

func testPositionToManyAddOpFactorPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c, d, e FactorPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*FactorPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, factorPositionDBTypes, false, strmangle.SetComplement(factorPositionPrimaryKeyColumns, factorPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*FactorPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFactorPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PositionID != first.PositionID {
			t.Error("foreign key was wrong value", a.PositionID, first.PositionID)
		}
		if a.PositionID != second.PositionID {
			t.Error("foreign key was wrong value", a.PositionID, second.PositionID)
		}

		if first.R.Position != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Position != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FactorPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FactorPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FactorPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPositionToOneUserAccountUsingUserAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Position
	var foreign UserAccount

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userAccountDBTypes, false, userAccountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserAccount struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserAccountID = foreign.UserAccountID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UserAccount().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserAccountID != foreign.UserAccountID {
		t.Errorf("want: %v, got %v", foreign.UserAccountID, check.UserAccountID)
	}

	slice := PositionSlice{&local}
	if err = local.L.LoadUserAccount(ctx, tx, false, (*[]*Position)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserAccount == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UserAccount = nil
	if err = local.L.LoadUserAccount(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserAccount == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPositionToOneEmojiUsingEmoji(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Position
	var foreign Emoji

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, emojiDBTypes, false, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.EmojiID, foreign.EmojiID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Emoji().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.EmojiID, foreign.EmojiID) {
		t.Errorf("want: %v, got %v", foreign.EmojiID, check.EmojiID)
	}

	slice := PositionSlice{&local}
	if err = local.L.LoadEmoji(ctx, tx, false, (*[]*Position)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Emoji == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Emoji = nil
	if err = local.L.LoadEmoji(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Emoji == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPositionToOneDesignPatternUsingDesignPattern(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Position
	var foreign DesignPattern

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, designPatternDBTypes, false, designPatternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DesignPattern struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DesignPatternID, foreign.DesignPatternID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DesignPattern().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.DesignPatternID, foreign.DesignPatternID) {
		t.Errorf("want: %v, got %v", foreign.DesignPatternID, check.DesignPatternID)
	}

	slice := PositionSlice{&local}
	if err = local.L.LoadDesignPattern(ctx, tx, false, (*[]*Position)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DesignPattern == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DesignPattern = nil
	if err = local.L.LoadDesignPattern(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DesignPattern == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPositionToOneSetOpUserAccountUsingUserAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c UserAccount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userAccountDBTypes, false, strmangle.SetComplement(userAccountPrimaryKeyColumns, userAccountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userAccountDBTypes, false, strmangle.SetComplement(userAccountPrimaryKeyColumns, userAccountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserAccount{&b, &c} {
		err = a.SetUserAccount(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UserAccount != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Positions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserAccountID != x.UserAccountID {
			t.Error("foreign key was wrong value", a.UserAccountID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserAccountID))
		reflect.Indirect(reflect.ValueOf(&a.UserAccountID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserAccountID != x.UserAccountID {
			t.Error("foreign key was wrong value", a.UserAccountID, x.UserAccountID)
		}
	}
}
func testPositionToOneSetOpEmojiUsingEmoji(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c Emoji

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Emoji{&b, &c} {
		err = a.SetEmoji(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Emoji != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Positions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.EmojiID, x.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EmojiID))
		reflect.Indirect(reflect.ValueOf(&a.EmojiID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.EmojiID, x.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID, x.EmojiID)
		}
	}
}

func testPositionToOneRemoveOpEmojiUsingEmoji(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b Emoji

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetEmoji(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveEmoji(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Emoji().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Emoji != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.EmojiID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Positions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPositionToOneSetOpDesignPatternUsingDesignPattern(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c DesignPattern

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DesignPattern{&b, &c} {
		err = a.SetDesignPattern(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DesignPattern != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Positions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DesignPatternID, x.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DesignPatternID))
		reflect.Indirect(reflect.ValueOf(&a.DesignPatternID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DesignPatternID, x.DesignPatternID) {
			t.Error("foreign key was wrong value", a.DesignPatternID, x.DesignPatternID)
		}
	}
}

func testPositionToOneRemoveOpDesignPatternUsingDesignPattern(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b DesignPattern

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, designPatternDBTypes, false, strmangle.SetComplement(designPatternPrimaryKeyColumns, designPatternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDesignPattern(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDesignPattern(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.DesignPattern().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.DesignPattern != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DesignPatternID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Positions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPositionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
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

func testPositionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PositionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPositionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Positions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	positionDBTypes = map[string]string{`CreatedAt`: `timestamptz`, `DesignPatternID`: `int8`, `EmojiID`: `int8`, `ParentPositionID`: `int8`, `PositionDescription`: `varchar`, `PositionID`: `int8`, `UserAccountID`: `int8`}
	_               = bytes.MinRead
)

func testPositionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(positionColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, positionDBTypes, true, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPositionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(positionColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, positionDBTypes, true, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(positionColumns, positionPrimaryKeyColumns) {
		fields = positionColumns
	} else {
		fields = strmangle.SetComplement(
			positionColumns,
			positionPrimaryKeyColumns,
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

	slice := PositionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPositionsUpsert(t *testing.T) {
	t.Parallel()

	if len(positionColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Position{}
	if err = randomize.Struct(seed, &o, positionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Position: %s", err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, positionDBTypes, false, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Position: %s", err)
	}

	count, err = Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
