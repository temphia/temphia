<script>
  import FormEdit from "../admin/core/form/auto_form.svelte";
</script>

<div class="w-full h-full">
  <!-- <iframe
    height="100%"
    width="100%"
    title="aa"
    src="http://localhost:8000"
    frameborder="0"
  /> -->

  <FormEdit
    schema={{
      name: "TEST1",
      fields: [
        {
          disabled: false,
          ftype: "TEXT",
          key_name: "field1",
          name: "Field 1",
          options: ["apple", "banana"],
        },
      ],
    }}
    data={{}}
    message=""
  />
</div>
