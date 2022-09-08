<script>
  const data = {
    main_node: {
      name: "one",
      sub_node: ["xyz1", "pqr2"],
    },
    links: [{ from: ["one", "xyz1"], to: ["two", "pqr2"] }],
    nodes: [
      {
        name: "two",
        sub_node: ["xyz1", "pqr2"],
      },
    ],
  };
</script>

<div>Flow Map</div>
