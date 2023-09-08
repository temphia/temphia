<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { CtypeIcons } from "../fields/field";

  export let ctype: string;
  export let classes = "h-5 w-5";
  export let solid = true;
</script>

<Icon name={CtypeIcons[ctype] || "hashtag"} class={classes} {solid} />
