<script>
  import Autotable from "../../xcompo/autotable/autotable.svelte";
  import Layout from "./_layout.svelte";
</script>

<Layout>
  <Autotable
    action_key={"slug"}
    actions={[
      {
        Name: "Visit",
        Action: (id) => {},
        Class: "bg-green-400",
        icon: "link",
      },

      {
        Name: "Edit",
        Action: (id) => {
          location.hash = `/page/${id}`;
        },
        icon: "pencil-alt",
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: () => {},
        icon: "trash",
      },
    ]}
    datas={[
      { slug: "index", name: "Main Page" },
      { slug: "test", name: "Test page" },
    ]}
    key_names={[
      ["slug", "Slug"],
      ["name", "Name"],
    ]}
  />
</Layout>
