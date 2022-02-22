package stored

//DB_USERS, User{}, User{}.ID
type O1Key[T Object, K1 comparable] struct {
	ID ObjectID // internal object ID
}

func (o *O1Key[T, K1]) Get(key1 K1) (T, error) {
	var object T
	return object, nil
}

// creating function
func Object1Key[T Object, K1 comparable](id ObjectID, obj T, key1 K1) O1Key[T, K1] {
	return O1Key[T, K1]{ID: id}
}
