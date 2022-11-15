<script lang="ts">
    import App from "../../target/app/app_new.svelte";
    import { params } from "svelte-hash-router";
  
    export let source = $params.source;
    export let group = $params.group;
    export let table = $params.group;
  
   </script>
  
  <App ttype="data" target={`${source}/${group}/${table}`} />
  