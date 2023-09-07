<script>
  export let user = "Susan";
  export let url = "https://i.pravatar.cc/?img=6";
</script>

<figure
  class="avatar flex aspect-square text-surface-50 font-semibold justify-center items-center overflow-hidden isolate bg-surface-400-500-token w-8 rounded-full bg-gray-100"
  data-testid="avatar"
>
  <img
    class="avatar-image w-full h-full object-cover"
    style=""
    src={url}
    alt=""
  />
</figure>