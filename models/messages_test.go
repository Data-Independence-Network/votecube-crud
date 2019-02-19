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

func testMessages(t *testing.T) {
	t.Parallel()

	query := Messages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Messages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MessageExists(ctx, tx, o.MessageID)
	if err != nil {
		t.Errorf("Unable to check if Message exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MessageExists to return true, but got false.")
	}
}

func testMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	messageFound, err := FindMessage(ctx, tx, o.MessageID)
	if err != nil {
		t.Error(err)
	}

	if messageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Messages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Messages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func messageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func testMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Message{}
	o := &Message{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, messageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Message object: %s", err)
	}

	AddMessageHook(boil.BeforeInsertHook, messageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterInsertHook, messageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	messageAfterInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterSelectHook, messageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	messageAfterSelectHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpdateHook, messageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpdateHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpdateHook, messageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	messageAfterUpdateHooks = []MessageHook{}

	AddMessageHook(boil.BeforeDeleteHook, messageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	messageBeforeDeleteHooks = []MessageHook{}

	AddMessageHook(boil.AfterDeleteHook, messageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	messageAfterDeleteHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpsertHook, messageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpsertHook, messageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	messageAfterUpsertHooks = []MessageHook{}
}

func testMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(messageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessageOneToOneMessageUsingParentMessageMessage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign Message
	var local Message

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.ParentMessageID = local.MessageID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ParentMessageMessage().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ParentMessageID != foreign.ParentMessageID {
		t.Errorf("want: %v, got %v", foreign.ParentMessageID, check.ParentMessageID)
	}

	slice := MessageSlice{&local}
	if err = local.L.LoadParentMessageMessage(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentMessageMessage == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ParentMessageMessage = nil
	if err = local.L.LoadParentMessageMessage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentMessageMessage == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMessageOneToOneSetOpMessageUsingParentMessageMessage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Message{&b, &c} {
		err = a.SetParentMessageMessage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ParentMessageMessage != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.ParentMessage != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.MessageID != x.ParentMessageID {
			t.Error("foreign key was wrong value", a.MessageID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.ParentMessageID))
		reflect.Indirect(reflect.ValueOf(&x.ParentMessageID)).Set(zero)

		if err = x.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MessageID != x.ParentMessageID {
			t.Error("foreign key was wrong value", a.MessageID, x.ParentMessageID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testMessageToManyMessagesLinks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c MessagesLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, messagesLinkDBTypes, false, messagesLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messagesLinkDBTypes, false, messagesLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MessageID = a.MessageID
	c.MessageID = a.MessageID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	messagesLink, err := a.MessagesLinks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range messagesLink {
		if v.MessageID == b.MessageID {
			bFound = true
		}
		if v.MessageID == c.MessageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MessageSlice{&a}
	if err = a.L.LoadMessagesLinks(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessagesLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MessagesLinks = nil
	if err = a.L.LoadMessagesLinks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessagesLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", messagesLink)
	}
}

func testMessageToManyPollsMessages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c PollsMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pollsMessageDBTypes, false, pollsMessageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pollsMessageDBTypes, false, pollsMessageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MessageID = a.MessageID
	c.MessageID = a.MessageID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	pollsMessage, err := a.PollsMessages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range pollsMessage {
		if v.MessageID == b.MessageID {
			bFound = true
		}
		if v.MessageID == c.MessageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MessageSlice{&a}
	if err = a.L.LoadPollsMessages(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsMessages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PollsMessages = nil
	if err = a.L.LoadPollsMessages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PollsMessages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", pollsMessage)
	}
}

func testMessageToManyAddOpMessagesLinks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c, d, e MessagesLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MessagesLink{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, messagesLinkDBTypes, false, strmangle.SetComplement(messagesLinkPrimaryKeyColumns, messagesLinkColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*MessagesLink{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessagesLinks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.MessageID != first.MessageID {
			t.Error("foreign key was wrong value", a.MessageID, first.MessageID)
		}
		if a.MessageID != second.MessageID {
			t.Error("foreign key was wrong value", a.MessageID, second.MessageID)
		}

		if first.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MessagesLinks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MessagesLinks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MessagesLinks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testMessageToManyAddOpPollsMessages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c, d, e PollsMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PollsMessage{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pollsMessageDBTypes, false, strmangle.SetComplement(pollsMessagePrimaryKeyColumns, pollsMessageColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PollsMessage{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPollsMessages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.MessageID != first.MessageID {
			t.Error("foreign key was wrong value", a.MessageID, first.MessageID)
		}
		if a.MessageID != second.MessageID {
			t.Error("foreign key was wrong value", a.MessageID, second.MessageID)
		}

		if first.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PollsMessages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PollsMessages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PollsMessages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testMessageToOneUserAccountUsingUserAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Message
	var foreign UserAccount

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

	slice := MessageSlice{&local}
	if err = local.L.LoadUserAccount(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
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

func testMessageToOneMessageUsingParentMessage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Message
	var foreign Message

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ParentMessageID = foreign.MessageID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ParentMessage().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.MessageID != foreign.MessageID {
		t.Errorf("want: %v, got %v", foreign.MessageID, check.MessageID)
	}

	slice := MessageSlice{&local}
	if err = local.L.LoadParentMessage(ctx, tx, false, (*[]*Message)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentMessage == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ParentMessage = nil
	if err = local.L.LoadParentMessage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentMessage == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMessageToOneSetOpUserAccountUsingUserAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c UserAccount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
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

		if x.R.Messages[0] != &a {
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
func testMessageToOneSetOpMessageUsingParentMessage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Message
	var b, c Message

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, messageDBTypes, false, strmangle.SetComplement(messagePrimaryKeyColumns, messageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Message{&b, &c} {
		err = a.SetParentMessage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ParentMessage != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentMessageMessage != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ParentMessageID != x.MessageID {
			t.Error("foreign key was wrong value", a.ParentMessageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ParentMessageID))
		reflect.Indirect(reflect.ValueOf(&a.ParentMessageID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ParentMessageID != x.MessageID {
			t.Error("foreign key was wrong value", a.ParentMessageID, x.MessageID)
		}
	}
}

func testMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
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

func testMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	messageDBTypes = map[string]string{`CreatedAt`: `timestamptz`, `MessageID`: `int8`, `MessageSubject`: `varchar`, `MessageText`: `varchar`, `MessageType`: `varchar`, `ParentMessageID`: `int8`, `UserAccountID`: `int8`}
	_              = bytes.MinRead
)

func testMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(messageColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(messageColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(messageColumns, messagePrimaryKeyColumns) {
		fields = messageColumns
	} else {
		fields = strmangle.SetComplement(
			messageColumns,
			messagePrimaryKeyColumns,
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

	slice := MessageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(messageColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Message{}
	if err = randomize.Struct(seed, &o, messageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Message: %s", err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, messageDBTypes, false, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Message: %s", err)
	}

	count, err = Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
