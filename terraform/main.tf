provider dgraph {
  server = "localhost:9080"
}

resource dgraph_type person {
  name = "Person"
  fields = {
    name = dgraph_predicate.name.type
    prename = dgraph_predicate.prename.type
    age = dgraph_predicate.age.type
  }
}

resource dgraph_predicate name {
  name = "name"
  type = "string"
  index = true
  tokenizer = "exact"
}

resource dgraph_predicate prename {
  name = "name"
  type = "string"
  index = true
  tokenizer = "exact"
}

resource dgraph_predicate age {
  name = "age"
  type = "int"
}

output "person_id" {
  value = dgraph_type.person.id
  description = "ID of person type"
}
