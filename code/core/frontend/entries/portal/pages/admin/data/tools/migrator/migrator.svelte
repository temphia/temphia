<script>
  const data = {
    extends: "temphia.schema.todo",
    check_pattern: {
      has: [{ ctype: "json", slug: "extra_meta" }],
    },
    steps: [
      {
        name: "add end date column",
        type: "add_column",
        data: { name: "end_date", ctype: "datetime" },
      },

      {
        name: "change column",
        type: "soft_change_column",
        data: { name: "ratings", from: "number", to: "rangenumber" },
      },

      {
        name: "remove extra column",
        type: "remove_column",
        data: { name: "extra_meta" },
      },
    ],
  };
</script>

<div>Migrator</div>
