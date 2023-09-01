<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import type { AutoForm, PortalService } from "$lib/core";

  const app = getContext("__app__") as PortalService;

  let done = false;
  let message = "";
  let data;

  const save = async (_data: any) => {
    done = true;
    console.log("@save_this", _data);
  };

  const options = app.api_manager.self_data.get_user_apps();
</script>

<div class="h-full p-4 bg-indigo-100">
  <div class="p-4 w-full  bg-white rounded">
    <div class="text-2xl text-indigo-900">Pair app device</div>
    <p class="text-red-500">{message}</p>

    {#if !done}
      <div class="flex-col flex py-3">
        <label for="field-0" class="pb-2 text-gray-700 font-semibold"
          >Name</label
        >
        <input
          id="field-0"
          type="text"
          list="field-0-datalist"
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>
      <div class="flex-col flex py-3">
        <label for="field-1" class="pb-2 text-gray-700 font-semibold">App</label
        >

        <div class="p-1 rounded flex flex-col gap-2">
          {#each options as opt}
            <div class="flex shadow p-2 text-blue-600 gap-1 flex-wrap">
              <input type="checkbox" />
              <img src="" alt="" class="h-10 w-10 bg-slate-400 rounded" />

              <p>
                Plug Id
                <span class="text-slate-800">{opt["plug_id"]}</span>
              </p>

              <p>
                Agent Id
                <span class="text-slate-800">{opt["agent_id"]}</span>
              </p>
            </div>
          {/each}
        </div>
      </div>
      <div class="flex justify-end py-3">
        <button
          on:click={save}
          class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:bg-blue-600 bg-blue-400"
        >
          <Icon name="ticket" class="h-5 w-5" />

          Register</button
        >
      </div>
    {:else}
      <div class="p-4">
        <img
          src="data:image/gif;base64,R0lGODlhZABkAJEAAAAAAP///wAAAAAAACH5BAEAAAIALAAAAABkAGQAAAL/jI+py+0Po5y02ouz3rz7D4biSJbmiabqCgLuC8fyjMD1bB/4zgNUD4zdXMNdMdj7IYPFI023BCqjvCa0eqXipr5LzvFthLlk7wsyXqQla/HZTHy0rZa5+l3Hu+NgfZmtJ5XFsDVoaCCI6CcXmHSYUKjIFxCWSLnYN2nZFnmp2Ygl2QX4Geo5CvkkSnfaaadQ6biaZcm6ibmHelvKdMfriksIavQYJfyrevrXmhyLdKwrqzzhTDwryjnsJMOKNgyMPGn73RzsG21anc27LfToHe5+zSxeTM79zlhPii4/DWsuEyp4A6nd6wcNw6tU+yKo03ZOYUCGBfm165bPokEt/+AQ2uOYjCDIauPigeQyEuK8hylRpvS48uBJFhHPLaT5YZ0/nCZ03uTJwedEFAtZ+utorKZGge2MNpWGkam+fkhlQs2Yi+DTeFutUZT4TGmuouUaujQVlSJZfPPA9gJoltbEjkGXJJy69ujQqcug6ez6U+rSsbiMdnLb12Zhq2wzBP5HuKHhkBth0qt4WatidnAsH5Ybty1oy8ucbsaclm7nkjtHa4arN3SFXV5N55F9Nnbtg4hzSqN9cbXvdL95CwcKm7Vy18iTf/RsvPnymJyx7VVB8jnrx4DtModdtfrsl7bVltXtlnxw8Oez3yZPdTHxxtyxvo7c2rrk61+P47cXuh9uYiXGXmPr9fdebwJ9ph9m9VG3m0ntXSXaYMCNFp5Xad13YYMZMghZacVJSJ93ziWIoUoIZoaXiQ4xRtpzeSVl4YgORvcfhVmlKN52c3Ek4nyoDaidfZV1tSGRECZZo5AEQuaeYzAyueJ0Ifr35Iz81SXgiVWexsKDV1ZoXpcdiPlXi2ZyCZ+VMwn2InzQkaihkjlq4SaQfFWmXpEe6kildIIOSmihhh6KaKKKLspoo45KVwAAOw=="
          alt=""
          class="w-64 mx-auto"
        />

        <div class="mx-auto">
          <details>
            <summary>Pair Code</summary>

            <pre>{"fixme"}</pre>
          </details>
        </div>
      </div>
    {/if}
  </div>
</div>
