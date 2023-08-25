<script lang="ts">
  import { escape } from "svelte/internal"; // fixme => check this, maybe move to dompurify
  import { sleep } from "yootils";

  export const success = (message: string) => open_toast(message, "green");
  export const error = (message: string) => open_toast(message, "red");

  let toast_root: HTMLElement;

  const open_toast = async (message: string, color: string) => {
    const node = document.createElement("div");
    node.classList.add(
      "flex",
      "items-center",
      "space-x-2",
      "transition",
      `bg-${color}-500`,
      `hover:bg-${color}-600`,
      "p-2"
    );
    node.innerHTML = `<svg
          class="h-7 w-7"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >

        ${
          color === "red"
            ? `<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />`
            : `<path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />`
        }

          
        </svg>
        <p class="font-bold">${escape(message)}</p>`;

    toast_root.appendChild(node);
    await sleep(1000);
    toast_root.removeChild(node);
  };
</script>



<div
  class="fixed top-4 right-4 z-50 cursor-pointer rounded-md text-white transition"
>
  <div class="flex flex-col" bind:this={toast_root} />
</div>
