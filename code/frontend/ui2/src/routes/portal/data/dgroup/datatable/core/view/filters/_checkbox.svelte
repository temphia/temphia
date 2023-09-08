<script lang="ts">
  export let value = "false";
  export let column: object;
</script>

<select
  class="rounded p-1"
  value
  on:click={(ev) => {
    value = ev.target["value"];
  }}
>
  <option value="true">true</option>
  <option value="false">false</option>
</select>
