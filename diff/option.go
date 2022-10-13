package diff

import (
	"github.com/shogo82148/schemalex-deploy"
	"github.com/shogo82148/schemalex-deploy/internal/option"
)

type Option schemalex.Option

const (
	optkeyParser      = "parser"
	optkeyTransaction = "transaction"
	optkeyCurrent     = "current"
)

// WithParser specifies the parser instance to use when parsing
// the statements given to the diffing functions. If unspecified,
// a default parser will be used
func WithParser(p *schemalex.Parser) Option {
	return option.New(optkeyParser, p)
}

// WithTransaction specifies if statements to control transactions
// should be included in the diff.
func WithTransaction(b bool) Option {
	return option.New(optkeyTransaction, b)
}

// WithCurrentSchema specifies the current schema deployed in MySQL.
func WithCurrentSchema(schema string) Option {
	return option.New(optkeyCurrent, schema)
}
