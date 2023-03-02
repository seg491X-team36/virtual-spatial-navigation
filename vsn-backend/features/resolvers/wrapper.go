package resolvers

func payloadWrapper[T any](t T, err error) (*T, *string) {
	if err != nil {
		msg := err.Error()
		return nil, &msg
	}
	return &t, nil
}
