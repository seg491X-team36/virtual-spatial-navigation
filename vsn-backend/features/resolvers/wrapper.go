package resolvers

func payloadWrapper[T any](t T, err error) (*T, *string) {
	if err != nil {
		msg := err.Error()
		return nil, &msg
	}
	return &t, nil
}

func nullable[T any](t T, err error) *T {
	if err != nil {
		return nil
	}
	return &t
}
