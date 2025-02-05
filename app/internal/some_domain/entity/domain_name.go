package entity

type SomeEntity struct {
}

func CreateSomeEntity(
	attr1 string,
	attr2 string,
	attr3 string,
) (SomeEntity, error) {
	// Some validation logic here.
	return NewSomeEntity(attr1, attr2, attr3), nil
}

func NewSomeEntity(
	attr1 string,
	attr2 string,
	attr3 string,
) SomeEntity {
	return SomeEntity{}
}

func (s *SomeEntity) ProcessSomeEvent() (Event, error) {
	// Some logic here
	return Event{}, nil
}

type Event struct {
}
