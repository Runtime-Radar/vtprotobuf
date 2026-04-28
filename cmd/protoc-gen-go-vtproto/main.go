package main

import (
	"flag"
	"strings"

	_ "github.com/runtime-radar/vtprotobuf/features/clone"
	_ "github.com/runtime-radar/vtprotobuf/features/equal"
	_ "github.com/runtime-radar/vtprotobuf/features/grpc"
	_ "github.com/runtime-radar/vtprotobuf/features/marshal"
	_ "github.com/runtime-radar/vtprotobuf/features/pool"
	_ "github.com/runtime-radar/vtprotobuf/features/size"
	_ "github.com/runtime-radar/vtprotobuf/features/unmarshal"
	"github.com/runtime-radar/vtprotobuf/generator"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var cfg generator.Config
	var features string
	var f flag.FlagSet

	f.BoolVar(&cfg.AllowEmpty, "allow-empty", false, "allow generation of empty files")
	cfg.Poolable = generator.NewObjectSet()
	cfg.PoolableExclude = generator.NewObjectSet()
	cfg.IgnoreUnknownFields = generator.NewObjectSet()
	f.Var(&cfg.Poolable, "pool", "use memory pooling for this object")
	f.Var(&cfg.PoolableExclude, "pool-exclude", "do not use memory pooling for this object")
	f.Var(&cfg.IgnoreUnknownFields, "ignoreUnknownFields", "ignore unknown fields instead of saving them")
	f.BoolVar(&cfg.Wrap, "wrap", false, "generate wrapper types")
	f.StringVar(&features, "features", "all", "list of features to generate (separated by '+')")
	f.StringVar(&cfg.BuildTag, "buildTag", "", "the go:build tag to set on generated files")
	f.BoolVar(&cfg.WKTImportRewrite, "wktImportRewrite", false, "respect rewritten GoIdent.GoImportPath for well-known types instead of using the bundled fork's WKT package paths")

	protogen.Options{ParamFunc: f.Set}.Run(func(plugin *protogen.Plugin) error {
		gen, err := generator.NewGenerator(plugin, strings.Split(features, "+"), &cfg)
		if err != nil {
			return err
		}
		gen.Generate()
		return nil
	})
}
