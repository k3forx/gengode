package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dave/jennifer/jen"
)

func main() {
	// Generate struct
	f := jen.NewFile("model")
	f.ImportName("time", "time")
	f.Type().Id("User").Struct(
		jen.Id("ID").Int(),
		jen.Id("CreatedAt").Qual("time", "Time"),
		jen.Id("UpdatedAt").Qual("time", "Time"),
	)
	f.Func().Params(jen.Id("m").Id("*User")).Id("Exists").Call().Bool().Block(
		jen.Return(jen.Id("m.ID").Op(">").Lit(0)),
	)
	fmt.Printf("%#v", f)

	if err := f.Save("user.go"); err != nil {
		log.Fatal(err)
	}

	if err := os.Rename("./user.go", "./pkg/model/user.go"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------------------------")

	// Generate interface
	f = jen.NewFile("service")
	f.Comment("go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mock -mock_names=User=User")
	f.Type().Id("User").Interface()
	fmt.Printf("%#v", f)

	if err := f.Save("user.go"); err != nil {
		log.Fatal(err)
	}

	if err := os.Rename("./user.go", "./pkg/service/user.go"); err != nil {
		log.Fatal(err)
	}
}
