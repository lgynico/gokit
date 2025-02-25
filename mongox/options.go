package mongox

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	// uri  string
	opts options.ClientOptions
}

func defaultOptions() Options {
	return Options{
		// uri:  "mongodb://localhost:27017",
		opts: *options.Client(),
	}
}

// func WithURI(uri string) func(opts *Options) {
// 	return func(opts *Options) {
// 		opts.opts.ApplyURI(uri)
// 		opts.uri = uri
// 	}
// }

func WithAuth(username, password string) func(opts *Options) {
	return func(opts *Options) {
		opts.opts.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}
}
