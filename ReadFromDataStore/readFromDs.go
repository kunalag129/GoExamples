package main

import (
	"golang.org/x/net/context"

	"fmt"
	"cloud.google.com/go/datastore"
)

type Abc struct {
	Name string
	Price string
	Content []string
}

func main()  {
	ctx := context.Background()

	client, _ := datastore.NewClient(ctx, "projectname")

	taskKey := datastore.IDKey("kindname", 103,nil)
	taskKey.Namespace = "namespace"

	abc := Abc{Name:"name", Price:"price"}


	var in Abc

	//q := datastore.NewQuery("Product")
	//q.Filter("__key__", "KEY(Product,1)")
	key, _ := client.Put(ctx, taskKey, &abc)
	fmt.Println(key)


	taskKey = datastore.IDKey("kind", 102, nil)
	taskKey.Namespace = "namespace"
	_ = client.Get(ctx, taskKey, &in)

	fmt.Println(in)

	taskKey = datastore.NameKey("kind", "entityid", nil)
	taskKey.Namespace = "namespace"
	_ = client.Get(ctx, taskKey, &in)
	fmt.Println(in.Content)
	fmt.Println(in.Content[2])
}
