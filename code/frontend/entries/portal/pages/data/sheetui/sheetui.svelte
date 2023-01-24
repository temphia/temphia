<script>
  import Renderer from "./_renderer.svelte";
</script>

<Renderer
  active_sheet={1}
  sheets={[
    { id: 1, name: "example 1" },
    { id: 2, name: "example 2" },
  ]}
  cells={{
    1: { 1: { value: "Ram" }, 2: { value: "This is some notes" } },
    2: { 1: { value: "Hari" }, 2: { value: "Something something" } },
    3: { 1: { value: "bin" }, 2: { value: "the mr bean ..." } },
  }}
  columns={[
    {
      ctype: "text",
      id: 1,
      name: "name",
      sheet_id: 1,
    },

    {
      ctype: "text",
      id: 2,
      name: "notes",
      sheet_id: 1,
    },
  ]}
  rows={[
    {
      id: 1,
      sheet_id: 1,
    },

    {
      id: 2,
      sheet_id: 1,
    },
    {
      id: 3,
      sheet_id: 1,
    },
  ]}
/>
