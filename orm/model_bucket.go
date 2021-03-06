package orm

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

// TODO
// - migrations
// - do not use Bucket but directly access KVStore
// - register for queries

// Model is implemented by any entity that can be stored using ModelBucket.
//
// This is the same interface as CloneableData. Using the right type names
// provides an easier to read API.
type Model interface {
	weave.Persistent
	Validate() error
}

// ModelSlicePtr represents a pointer to a slice of models. Think of it as
// *[]Model Because of Go type system, using []Model would not work for us.
// Instead we use a placeholder type and the validation is done during the
// runtime.
type ModelSlicePtr interface{}

// ModelBucket is implemented by buckets that operates on Models rather than
// Objects.
type ModelBucket interface {
	// One query the database for a single model instance. Lookup is done
	// by the primary index key. Result is loaded into given destination
	// model.
	// This method returns ErrNotFound if the entity does not exist in the
	// database.
	// If given model type cannot be used to contain stored entity, ErrType
	// is returned.
	One(db weave.ReadOnlyKVStore, key []byte, dest Model) error

	// ByIndex returns all objects that secondary index with given name and
	// given key. Main index is always unique but secondary indexes can
	// return more than one value for the same key.
	// All matching entities are appended to given destination slice. If no
	// result was found, no error is returned and destination slice is not
	// modified.
	ByIndex(db weave.ReadOnlyKVStore, indexName string, key []byte, dest ModelSlicePtr) (keys [][]byte, err error)

	// Index returns the index with given name that is maintained for this
	// bucket. This function can return ErrInvalidIndex if an index with
	// requested name does not exist.
	//
	// This method allows for lower level access and more control than
	// ByIndex method provides. When iterating over huge native indexes
	// where you do not want to load everything into memory, using an index
	// directly is a way to go.
	Index(name string) (Index, error)

	// Put saves given model in the database. Before inserting into
	// database, model is validated using its Validate method.
	// If the key is nil or zero length then a sequence generator is used
	// to create a unique key value.
	// Using a key that already exists in the database cause the value to
	// be overwritten.
	Put(db weave.KVStore, key []byte, m Model) ([]byte, error)

	// Delete removes an entity with given primary key from the database.
	// It returns ErrNotFound if an entity with given key does not exist.
	Delete(db weave.KVStore, key []byte) error

	// Has returns nil if an entity with given primary key value exists. It
	// returns ErrNotFound if no entity can be found.
	// Has is a cheap operation that that does not read the data and only
	// checks the existence of it.
	Has(db weave.KVStore, key []byte) error

	// Register registers this buckets content to be accessible via query
	// requests under the given name.
	Register(name string, r weave.QueryRouter)
}

// IterAll returns an iterator instance that loops through all entities kept by
// given bucket.
// Returned iterator does not support schema migration. It always returns the
// entity in version it is stored in the database.
func IterAll(bucketName string) *ModelBucketIterator {
	// This is how Bucket.DBKey is implemented.
	dbkey := []byte(bucketName + ":")

	start := []byte(bucketName + ":")
	end := append([]byte(bucketName+":"), 255, 255, 255, 255, 255, 255, 255)

	return &ModelBucketIterator{
		dbprefix: len(dbkey),
		cursor:   start,
		end:      end,
	}
}

// ModelBucketIterator allows for iteration over all entities of a single
// bucket.
type ModelBucketIterator struct {
	dbprefix int
	cursor   []byte
	end      []byte
	err      error
}

// Next returns the next item key. Loads the value of the item into the
// provided model. For each call a new iterator is created so that a database
// modification can be done between this method calls.
func (it *ModelBucketIterator) Next(db weave.ReadOnlyKVStore, dest Model) ([]byte, error) {
	if it.err != nil {
		return nil, it.err
	}

	if bytes.Compare(it.cursor, it.end) >= 0 {
		return nil, errors.ErrIteratorDone
	}

	iter, err := db.Iterator(it.cursor, it.end)
	if err != nil {
		return nil, errors.Wrap(err, "new iterator")
	}
	// Use iterator as a one time fetch, so that other database operations
	// can be performed between Next method calls.
	defer iter.Release()

	key, value, err := iter.Next()
	if err != nil {
		return nil, errors.Wrap(err, "iterator next")
	}

	if err := dest.Unmarshal(value); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal model value")
	}

	// Key was consumed. Iterator is inclusive, so we must use the very
	// next possible key. Which is this very key with zero appended.
	it.cursor = make([]byte, len(key)+1)
	copy(it.cursor, key)

	return key[it.dbprefix:], nil
}

// NewModelBucket returns a ModelBucket instance. This implementation relies on
// a bucket instance. Final implementation should operate directly on the
// KVStore instead.
func NewModelBucket(name string, m Model, opts ...ModelBucketOption) ModelBucket {
	b := NewBucket(name, m)

	tp := reflect.TypeOf(m)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	mb := &modelBucket{
		b:     b,
		idSeq: b.Sequence("id"),
		model: tp,
	}
	for _, fn := range opts {
		fn(mb)
	}
	return mb
}

// ModelBucketOption is implemented by any function that can configure
// ModelBucket during creation.
type ModelBucketOption func(mb *modelBucket)

// WithIndex configures the bucket to build an index with given name. All
// entities stored in the bucket are indexed using value returned by the
// indexer function. If an index is unique, there can be only one entity
// referenced per index value.
// Indexer value must be a function that implements either Indexer or
// MultiKeyIndexer interface.
func WithIndex(name string, indexer interface{}, unique bool) ModelBucketOption {
	var idx MultiKeyIndexer
	switch fn := indexer.(type) {
	case MultiKeyIndexer:
		idx = fn
	case func(Object) ([][]byte, error):
		idx = fn
	case Indexer:
		// Indexer is a subset of a MultiKeyIndexer but to be backward
		// compatible, we must support this type as well.
		idx = asMultiKeyIndexer(fn)
	case func(Object) ([]byte, error):
		// This is indexer interface
		idx = asMultiKeyIndexer(fn)
	default:
		text := fmt.Sprintf("indexer must implement either Indexer or MultiKeyIndexer interface, got %T", indexer)
		panic(text)
	}

	return func(mb *modelBucket) {
		mb.b = mb.b.WithMultiKeyIndex(name, idx, unique)
	}
}

// WithNativeIndex configures a bucket to maintain an index. Used index
// implementation is using the underlying storage mechanisms to maintain and
// iterate over indexed data.
// This implementation should be used to maintain an index for big collections.
// For small collections, use WithIndex function that configures a compact
// index implementation.
func WithNativeIndex(name string, indexer MultiKeyIndexer) ModelBucketOption {
	return func(mb *modelBucket) {
		mb.b = mb.b.WithNativeIndex(name, indexer)
	}
}

// WithIDSequence configures the bucket to use the given sequence instance for
// generating ID.
func WithIDSequence(s Sequence) ModelBucketOption {
	return func(mb *modelBucket) {
		mb.idSeq = s
	}
}

type modelBucket struct {
	b     Bucket
	idSeq Sequence

	// model is referencing the structure type. Event if the structure
	// pointer is implementing Model interface, this variable references
	// the structure directly and not the structure's pointer type.
	model reflect.Type
}

func (mb *modelBucket) Register(name string, r weave.QueryRouter) {
	mb.b.Register(name, r)
}

func (mb *modelBucket) One(db weave.ReadOnlyKVStore, key []byte, dest Model) error {
	obj, err := mb.b.Get(db, key)
	if err != nil {
		return err
	}
	if obj == nil || obj.Value() == nil {
		return errors.Wrapf(errors.ErrNotFound, "%T not in the store", dest)
	}
	res := obj.Value()

	if !reflect.TypeOf(res).AssignableTo(reflect.TypeOf(dest)) {
		return errors.Wrapf(errors.ErrType, "%T cannot be represented as %T", res, dest)
	}

	reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(res).Elem())
	return nil
}

func (mb *modelBucket) Index(name string) (Index, error) {
	return mb.b.Index(name)
}

func (mb *modelBucket) ByIndex(db weave.ReadOnlyKVStore, indexName string, key []byte, destination ModelSlicePtr) ([][]byte, error) {
	objs, err := mb.b.GetIndexed(db, indexName, key)
	if err != nil {
		return nil, err
	}
	if len(objs) == 0 {
		return nil, nil
	}

	dest := reflect.ValueOf(destination)
	if dest.Kind() != reflect.Ptr {
		return nil, errors.Wrap(errors.ErrType, "destination must be a pointer to slice of models")
	}
	if dest.IsNil() {
		return nil, errors.Wrap(errors.ErrImmutable, "got nil pointer")
	}
	dest = dest.Elem()
	if dest.Kind() != reflect.Slice {
		return nil, errors.Wrap(errors.ErrType, "destination must be a pointer to slice of models")
	}

	// It is allowed to pass destination as both []MyModel and []*MyModel
	sliceOfPointers := dest.Type().Elem().Kind() == reflect.Ptr

	allowed := dest.Type().Elem()
	if sliceOfPointers {
		allowed = allowed.Elem()
	}
	if mb.model != allowed {
		return nil, errors.Wrapf(errors.ErrType, "this bucket operates on %s model and cannot return %s", mb.model, allowed)
	}

	keys := make([][]byte, 0, len(objs))
	for _, obj := range objs {
		if obj == nil || obj.Value() == nil {
			continue
		}
		val := reflect.ValueOf(obj.Value())
		if !sliceOfPointers {
			val = val.Elem()
		}
		dest.Set(reflect.Append(dest, val))
		keys = append(keys, obj.Key())
	}
	return keys, nil

}

func (mb *modelBucket) Put(db weave.KVStore, key []byte, m Model) ([]byte, error) {
	mTp := reflect.TypeOf(m)
	if mTp.Kind() != reflect.Ptr {
		return nil, errors.Wrap(errors.ErrType, "model destination must be a pointer")
	}
	if mb.model != mTp.Elem() {
		return nil, errors.Wrapf(errors.ErrType, "cannot store %T type in this bucket", m)
	}

	if err := m.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid model")
	}

	if len(key) == 0 {
		var err error
		key, err = mb.idSeq.NextVal(db)
		if err != nil {
			return nil, errors.Wrap(err, "ID sequence")
		}
	}

	obj := NewSimpleObj(key, m)
	if err := mb.b.Save(db, obj); err != nil {
		return nil, errors.Wrap(err, "cannot store in the database")
	}
	return key, nil
}

func (mb *modelBucket) Delete(db weave.KVStore, key []byte) error {
	if err := mb.Has(db, key); err != nil {
		return err
	}
	return mb.b.Delete(db, key)
}

func (mb *modelBucket) Has(db weave.KVStore, key []byte) error {
	if key == nil {
		// nil key is a special case that would cause the store API to panic.
		return errors.ErrNotFound
	}

	// As long as we rely on the Bucket implementation to access the
	// database, we must refine the key.
	key = mb.b.DBKey(key)

	ok, err := db.Has(key)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrNotFound
	}
	return nil
}

var _ ModelBucket = (*modelBucket)(nil)
