<script lang="ts">
  import { time_ago } from "../../../../../../../../lib/vendor/timeago";
  import Jsonview from "../../../../../../../xcompo/jsonview/jsonview.svelte";

  const CREATE_RECORD = "insert";
  const UPDATE_RECORD = "update";
  const DELETE_RECORD = "delete";
  const COMMENT_RECORD = "comment";
  export let data;
  export let expanded = false;
  export let onClick;

  $: _text = {
    [CREATE_RECORD]: "created new a record",
    [UPDATE_RECORD]: "updated a record",
    [DELETE_RECORD]: "deleted a record",
    [COMMENT_RECORD]: "commented ",
  }[data.type];
</script>

<div class="w-full p-3 mt-4 bg-white rounded shadow flex flex-shrink-0 border">
  <div
    tabindex="0"
    aria-label="group icon"
    role="img"
    class="focus:outline-none w-8 h-8 border rounded-full border-gray-200 flex flex-shrink-0 items-center justify-center"
  >
    {#if data.type === CREATE_RECORD}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6 text-green-500"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        stroke-width="2"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    {:else if data.type === UPDATE_RECORD}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5 text-blue-500"
        viewBox="0 0 20 20"
        fill="currentColor"
      >
        <path
          d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
        />
        <path
          fill-rule="evenodd"
          d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
          clip-rule="evenodd"
        />
      </svg>
    {:else if data.type === DELETE_RECORD}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5 text-red-500"
        viewBox="0 0 20 20"
        fill="currentColor"
      >
        <path
          fill-rule="evenodd"
          d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
          clip-rule="evenodd"
        />
      </svg>
    {:else if data.type === COMMENT_RECORD}
      <svg
        class="h-5 w-5"
        viewBox="0 0 16 16"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        ><path
          d="M4.30325 12.6667L1.33325 15V2.66667C1.33325 2.48986 1.40349 2.32029 1.52851 2.19526C1.65354 2.07024 1.82311 2 1.99992 2H13.9999C14.1767 2 14.3463 2.07024 14.4713 2.19526C14.5963 2.32029 14.6666 2.48986 14.6666 2.66667V12C14.6666 12.1768 14.5963 12.3464 14.4713 12.4714C14.3463 12.5964 14.1767 12.6667 13.9999 12.6667H4.30325ZM5.33325 6.66667V8H10.6666V6.66667H5.33325Z"
          fill="#4338CA"
        /></svg
      >
    {/if}
  </div>
  <div class="pl-3 w-full cursor-pointer" on:click={onClick}>
    <div class="flex items-center justify-between w-full">
      <p tabindex="0" class="focus:outline-none leading-none">
        <span class="text-indigo-700">{data.user_id}</span>
        {_text}

        <!-- <span class="text-indigo-700">UX Designers</span> -->
      </p>
      <div
        tabindex="0"
        aria-label="close icon"
        role="button"
        class="focus:outline-none cursor-pointer rounded-full border p-2"
      >
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M1.33325 14.6667C1.33325 13.2522 1.89516 11.8956 2.89535 10.8954C3.89554 9.89523 5.2521 9.33333 6.66659 9.33333C8.08107 9.33333 9.43763 9.89523 10.4378 10.8954C11.438 11.8956 11.9999 13.2522 11.9999 14.6667H1.33325ZM6.66659 8.66666C4.45659 8.66666 2.66659 6.87666 2.66659 4.66666C2.66659 2.45666 4.45659 0.666664 6.66659 0.666664C8.87659 0.666664 10.6666 2.45666 10.6666 4.66666C10.6666 6.87666 8.87659 8.66666 6.66659 8.66666ZM11.5753 10.1553C12.595 10.4174 13.5061 10.9946 14.1788 11.8046C14.8515 12.6145 15.2515 13.6161 15.3219 14.6667H13.3333C13.3333 12.9267 12.6666 11.3427 11.5753 10.1553ZM10.2266 8.638C10.7852 8.13831 11.232 7.52622 11.5376 6.84183C11.8432 6.15743 12.0008 5.41619 11.9999 4.66666C12.0013 3.75564 11.7683 2.85958 11.3233 2.06466C12.0783 2.21639 12.7576 2.62491 13.2456 3.2208C13.7335 3.81668 14.0001 4.56315 13.9999 5.33333C14.0001 5.80831 13.8987 6.27784 13.7027 6.71045C13.5066 7.14306 13.2203 7.52876 12.863 7.84169C12.5056 8.15463 12.0856 8.38757 11.6309 8.52491C11.1762 8.66224 10.6974 8.7008 10.2266 8.638Z"
            fill="#047857"
          />
        </svg>
      </div>
    </div>

    {#if data.type === COMMENT_RECORD}
      <p class="italic text-gray-700">{data.payload}</p>
    {/if}

    <p
      tabindex="0"
      class="focus:outline-none text-xs leading-3 pt-1 text-gray-500"
    >
      {time_ago(data.created_at)}
    </p>

    {#if expanded && data.type !== COMMENT_RECORD}
      <div class="p-2 border rounded bg-gray-50">
        <Jsonview json={JSON.parse(data.payload || "{}")} />
      </div>
    {/if}
  </div>
</div>
