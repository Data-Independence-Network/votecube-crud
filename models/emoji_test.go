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

func testEmojis(t *testing.T) {
	t.Parallel()

	query := Emojis()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmojisDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
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

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmojisQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Emojis().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmojisSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmojiSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmojisExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmojiExists(ctx, tx, o.EmojiID)
	if err != nil {
		t.Errorf("Unable to check if Emoji exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmojiExists to return true, but got false.")
	}
}

func testEmojisFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	emojiFound, err := FindEmoji(ctx, tx, o.EmojiID)
	if err != nil {
		t.Error(err)
	}

	if emojiFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmojisBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Emojis().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmojisOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Emojis().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmojisAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	emojiOne := &Emoji{}
	emojiTwo := &Emoji{}
	if err = randomize.Struct(seed, emojiOne, emojiDBTypes, false, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}
	if err = randomize.Struct(seed, emojiTwo, emojiDBTypes, false, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emojiOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emojiTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Emojis().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmojisCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	emojiOne := &Emoji{}
	emojiTwo := &Emoji{}
	if err = randomize.Struct(seed, emojiOne, emojiDBTypes, false, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}
	if err = randomize.Struct(seed, emojiTwo, emojiDBTypes, false, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = emojiOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emojiTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func emojiBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func emojiAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Emoji) error {
	*o = Emoji{}
	return nil
}

func testEmojisHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Emoji{}
	o := &Emoji{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, emojiDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Emoji object: %s", err)
	}

	AddEmojiHook(boil.BeforeInsertHook, emojiBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	emojiBeforeInsertHooks = []EmojiHook{}

	AddEmojiHook(boil.AfterInsertHook, emojiAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	emojiAfterInsertHooks = []EmojiHook{}

	AddEmojiHook(boil.AfterSelectHook, emojiAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	emojiAfterSelectHooks = []EmojiHook{}

	AddEmojiHook(boil.BeforeUpdateHook, emojiBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	emojiBeforeUpdateHooks = []EmojiHook{}

	AddEmojiHook(boil.AfterUpdateHook, emojiAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	emojiAfterUpdateHooks = []EmojiHook{}

	AddEmojiHook(boil.BeforeDeleteHook, emojiBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	emojiBeforeDeleteHooks = []EmojiHook{}

	AddEmojiHook(boil.AfterDeleteHook, emojiAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	emojiAfterDeleteHooks = []EmojiHook{}

	AddEmojiHook(boil.BeforeUpsertHook, emojiBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	emojiBeforeUpsertHooks = []EmojiHook{}

	AddEmojiHook(boil.AfterUpsertHook, emojiAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	emojiAfterUpsertHooks = []EmojiHook{}
}

func testEmojisInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmojisInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(emojiColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmojiToManyDirections(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c Direction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, directionDBTypes, false, directionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, directionDBTypes, false, directionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.EmojiID, a.EmojiID)
	queries.Assign(&c.EmojiID, a.EmojiID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	direction, err := a.Directions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range direction {
		if queries.Equal(v.EmojiID, b.EmojiID) {
			bFound = true
		}
		if queries.Equal(v.EmojiID, c.EmojiID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EmojiSlice{&a}
	if err = a.L.LoadDirections(ctx, tx, false, (*[]*Emoji)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Directions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Directions = nil
	if err = a.L.LoadDirections(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Directions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", direction)
	}
}

func testEmojiToManyPollsDimensionsDirections(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c PollsDimensionsDirection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsDimensionsDirectionDBTypes, false, pollsDimensionsDirectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsDimensionsDirectionDBTypes, false, pollsDimensionsDirectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.EmojiID, a.EmojiID)
	queries.Assign(&c.EmojiID, a.EmojiID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsDimensionsDirection, err := a.PollsDimensionsDirections().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsDimensionsDirection {
		if queries.Equal(v.EmojiID, b.EmojiID) {
			bFound = true
		}
		if queries.Equal(v.EmojiID, c.EmojiID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EmojiSlice{&a}
	if err = a.L.LoadPollsDimensionsDirections(ctx, tx, false, (*[]*Emoji)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsDimensionsDirections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsDimensionsDirections = nil
	if err = a.L.LoadPollsDimensionsDirections(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsDimensionsDirections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsDimensionsDirection)
	}
}

func testEmojiToManyAddOpDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e Direction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Direction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, directionDBTypes, false, strmangle.SetComplement(directionPrimaryKeyColumns, directionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Direction{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDirections(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.EmojiID, first.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID, first.EmojiID)
		}
		if !queries.Equal(a.EmojiID, second.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID, second.EmojiID)
		}

		if first.R.Emoji != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Emoji != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Directions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Directions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Directions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testEmojiToManySetOpDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e Direction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Direction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, directionDBTypes, false, strmangle.SetComplement(directionPrimaryKeyColumns, directionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetDirections(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Directions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetDirections(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Directions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.EmojiID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.EmojiID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.EmojiID, d.EmojiID) {
		t.Error("foreign key was wrong value", a.EmojiID, d.EmojiID)
	}
	if !queries.Equal(a.EmojiID, e.EmojiID) {
		t.Error("foreign key was wrong value", a.EmojiID, e.EmojiID)
	}

	if b.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Emoji != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Emoji != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Directions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Directions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testEmojiToManyRemoveOpDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e Direction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Direction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, directionDBTypes, false, strmangle.SetComplement(directionPrimaryKeyColumns, directionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddDirections(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Directions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveDirections(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Directions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.EmojiID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.EmojiID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Emoji != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Emoji != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Directions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Directions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Directions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testEmojiToManyAddOpPollsDimensionsDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e PollsDimensionsDirection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsDimensionsDirection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsDimensionsDirectionDBTypes, false, strmangle.SetComplement(pollsDimensionsDirectionPrimaryKeyColumns, pollsDimensionsDirectionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PollsDimensionsDirection{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsDimensionsDirections(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.EmojiID, first.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID, first.EmojiID)
		}
		if !queries.Equal(a.EmojiID, second.EmojiID) {
			t.Error("foreign key was wrong value", a.EmojiID, second.EmojiID)
		}

		if first.R.Emoji != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Emoji != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsDimensionsDirections[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsDimensionsDirections[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsDimensionsDirections().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testEmojiToManySetOpPollsDimensionsDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e PollsDimensionsDirection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsDimensionsDirection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsDimensionsDirectionDBTypes, false, strmangle.SetComplement(pollsDimensionsDirectionPrimaryKeyColumns, pollsDimensionsDirectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetPollsDimensionsDirections(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PollsDimensionsDirections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPollsDimensionsDirections(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PollsDimensionsDirections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.EmojiID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.EmojiID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.EmojiID, d.EmojiID) {
		t.Error("foreign key was wrong value", a.EmojiID, d.EmojiID)
	}
	if !queries.Equal(a.EmojiID, e.EmojiID) {
		t.Error("foreign key was wrong value", a.EmojiID, e.EmojiID)
	}

	if b.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Emoji != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Emoji != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.PollsDimensionsDirections[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.PollsDimensionsDirections[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testEmojiToManyRemoveOpPollsDimensionsDirections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Emoji
	var b, c, d, e PollsDimensionsDirection

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, emojiDBTypes, false, strmangle.SetComplement(emojiPrimaryKeyColumns, emojiColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsDimensionsDirection{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsDimensionsDirectionDBTypes, false, strmangle.SetComplement(pollsDimensionsDirectionPrimaryKeyColumns, pollsDimensionsDirectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddPollsDimensionsDirections(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PollsDimensionsDirections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePollsDimensionsDirections(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PollsDimensionsDirections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.EmojiID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.EmojiID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Emoji != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Emoji != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Emoji != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.PollsDimensionsDirections) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.PollsDimensionsDirections[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.PollsDimensionsDirections[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testEmojisReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
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

func testEmojisReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmojiSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmojisSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Emojis().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	emojiDBTypes = map[string]string{`CSSClass`: `varchar`, `EmojiID`: `int8`}
	_            = bytes.MinRead
)

func testEmojisUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(emojiPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(emojiColumns) == len(emojiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmojisSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(emojiColumns) == len(emojiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Emoji{}
	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emojiDBTypes, true, emojiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(emojiColumns, emojiPrimaryKeyColumns) {
		fields = emojiColumns
	} else {
		fields = strmangle.SetComplement(
			emojiColumns,
			emojiPrimaryKeyColumns,
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

	slice := EmojiSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmojisUpsert(t *testing.T) {
	t.Parallel()

	if len(emojiColumns) == len(emojiPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Emoji{}
	if err = randomize.Struct(seed, &o, emojiDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Emoji: %s", err)
	}

	count, err := Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, emojiDBTypes, false, emojiPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Emoji struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Emoji: %s", err)
	}

	count, err = Emojis().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
