# terraform provider for Dgraph

This is a terraform provider to manage predicates and types within a dgraph instance

For information about dgraph go to [Dgraph](https://dgraph.io/)

For information about terraform go to [terraform](https://www.terraform.io/)

## Provider specific information

### ID conventions used

The ID convention used is either `type_<name>` or `predicate_<name>` depending on the resource being created.

### Dependencies between predicates and types

You get an error if you want to create a type with that has not been created before as a predicate. Lets assume you want to create a person type with a `name: string`:

    type Person {
      name
    }

And you create your terraform file like this:

    provider dgraph {
      server = "localhost:9080"
    }
    
    resource dgraph_type person {
      name = "Person"
      fields = {
        name
      }
    }

The resulting error would be `rpc error: code = Unknown desc = Schema does not contain a matching predicate for field name in type Person`. Let's change the file by adding the following snippet:

    resource dgraph_predicate name {
      name = "name"
      type = "string"
      index = true
      tokenizer = "exact"
    }

For the first run you still get the error, the second one works out well. Why is this? There is no visible dependency between the type and the predicate that terraform can know. Let's change the person type slightly:

    resource dgraph_type person {
      name = "Person"
      fields = {
        name = dgraph_predicate.name.type
      }
    }

As there is now a reference to the predicate it runs well.

## History

|Version|Description|
|---|---|