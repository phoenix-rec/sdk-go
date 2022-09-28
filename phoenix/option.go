package phoenix

type Option func(opts *Options)

type Options struct {
	RequestId string
}

func Conv2Options(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func WithRequestId(requestId string) Option {
	return func(options *Options) {
		options.RequestId = requestId
	}
}
