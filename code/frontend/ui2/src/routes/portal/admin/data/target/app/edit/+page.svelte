<script lang="ts">
    // import App from "../../target/app/app_new.svelte";
    import { params } from "$lib/params";
  
    export let source = $params["source"];
    export let group = $params["group"];
    export let table = $params["table"];
  
   </script>
  
  <!-- <App ttype="data" target={`${source}/${group}/${table}`} /> -->
  