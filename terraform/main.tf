provider dgraph {
  server = "localhost:9080"
}

resource dgraph_predicate name {
  name = "name"
  type = "string"

  array = true
  edge_count = true
}
