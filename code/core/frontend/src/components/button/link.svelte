<script>
  import { Link, navigate } from "svelte-routing";
  export let url;
  export let title;
</script>

<Link
  class="text-center text-sm bg-blue-500 rounded py-2 p-2 text-white mt-2"
  to={url}>{title}</Link
>
