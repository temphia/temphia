<script>
  export let user = "Susan";
  export let url = "https://i.pravatar.cc/?img=5";
</script>

<figure
  class="avatar flex aspect-square text-surface-50 font-semibold justify-center items-center overflow-hidden isolate bg-surface-400-500-token w-8  rounded-full   "
  data-testid="avatar"
>
  <img
    class="avatar-image w-full h-full object-cover"
    style=""
    src={url}
    alt=""
  />
</figure>
<span class="flex-auto">{user}</span>
