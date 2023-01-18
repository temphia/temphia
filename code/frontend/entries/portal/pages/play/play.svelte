<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import SampleBuilder from "../admin/bprint/editor/data_schema/sample_builder.svelte";

  // import ExecutorPageForm from "../../../executor_pageform/index.svelte";

  const app: PortalService = getContext("__app__");
</script>

<SampleBuilder />

<!-- <ExecutorPageForm /> -->
