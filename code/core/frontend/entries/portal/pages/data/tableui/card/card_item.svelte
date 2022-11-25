<script lang="ts">
  import type { Order } from "./order";

  export let order: Order;
  export let row: object;
  export let columns: { [_: string]: object };
</script>

<div class="relative mx-auto w-full">
  <div
    class="relative inline-block duration-300 ease-in-out transition-transform transform hover:-translate-y-2 w-full"
  >
    <div class="shadow p-4 rounded-lg bg-white">
      <div class="flex justify-center relative rounded-lg overflow-hidden h-52">
        <div
          class="transition-transform duration-500 transform ease-in-out hover:scale-110 w-full"
        >
          <div class="absolute inset-0 bg-black opacity-10" />
        </div>

        <div class="absolute flex justify-center bottom-0 mb-3">
          <div
            class="flex bg-white px-4 py-1 space-x-5 rounded-lg overflow-hidden shadow"
          >
            <p class="flex items-center font-medium text-gray-800">
              <svg
                class="w-5 h-5 fill-current mr-2"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 512 512"
                ><path
                  d="M480,226.15V80a48,48,0,0,0-48-48H80A48,48,0,0,0,32,80V226.15C13.74,231,0,246.89,0,266.67V472a8,8,0,0,0,8,8H24a8,8,0,0,0,8-8V416H480v56a8,8,0,0,0,8,8h16a8,8,0,0,0,8-8V266.67C512,246.89,498.26,231,480,226.15ZM64,192a32,32,0,0,1,32-32H208a32,32,0,0,1,32,32v32H64Zm384,32H272V192a32,32,0,0,1,32-32H416a32,32,0,0,1,32,32ZM80,64H432a16,16,0,0,1,16,16v56.9a63.27,63.27,0,0,0-32-8.9H304a63.9,63.9,0,0,0-48,21.71A63.9,63.9,0,0,0,208,128H96a63.27,63.27,0,0,0-32,8.9V80A16,16,0,0,1,80,64ZM32,384V266.67A10.69,10.69,0,0,1,42.67,256H469.33A10.69,10.69,0,0,1,480,266.67V384Z"
                /></svg
              >
              3 + 1
            </p>
          </div>
        </div>

        {#if order.tag}
          <span
            class="absolute top-0 left-0 inline-flex mt-3 ml-3 px-3 py-2 rounded-lg z-10 bg-red-500 text-sm font-medium text-white select-none"
          >
            {row[order.tag] || ""}
          </span>
        {/if}
      </div>

      <div class="mt-4">
        {#if order.name}
          <h2
            class="font-medium text-base md:text-lg text-gray-800 line-clamp-1"
            title={row[order.name]}
          >
            {row[order.name] || ""}
          </h2>
        {/if}

        {#if order.description}
          <p class="mt-2 text-sm text-gray-800 line-clamp-1">
            {row[order.description] || ""}
          </p>
        {/if}
      </div>

      <div class="grid grid-cols-2 grid-rows-2 gap-4 mt-8">
        {#each order.other as ocol}
          <p
            class="inline-flex flex-col xl:flex-row xl:items-center text-gray-800"
          >
            <svg
              class="inline-block w-5 h-5 xl:w-4 xl:h-4 mr-3 fill-current text-gray-800"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 576 512"
              ><path
                d="M570.53,242,512,190.75V48a16,16,0,0,0-16-16H400a16,16,0,0,0-16,16V78.75L298.53,4a16,16,0,0,0-21.06,0L5.47,242a16,16,0,0,0,21.07,24.09L64,233.27V464a48.05,48.05,0,0,0,48,48H464a48.05,48.05,0,0,0,48-48V233.27l37.46,32.79A16,16,0,0,0,570.53,242ZM480,464a16,16,0,0,1-16,16H112a16,16,0,0,1-16-16V205.27l192-168,192,168Zm0-301.25-64-56V64h64ZM208,218.67V325.34A26.75,26.75,0,0,0,234.66,352H341.3A26.76,26.76,0,0,0,368,325.34V218.67A26.75,26.75,0,0,0,341.3,192H234.66A26.74,26.74,0,0,0,208,218.67ZM240,224h96v96H240Z"
              /></svg
            >
            <span class="mt-2 xl:mt-0"> {columns[ocol]["name"] || ""} </span>
          </p>
          <p
            class="inline-flex flex-col xl:flex-row xl:items-center text-gray-800"
          >
            <svg
              class="inline-block w-5 h-5 xl:w-4 xl:h-4 mr-3 fill-current text-gray-800"
              viewBox="0 0 18 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M17.5883 7.872H16.4286L16.7097 8.99658H6.74743V10.1211H17.4309C17.5163 10.1211 17.6006 10.1017 17.6774 10.0642C17.7542 10.0267 17.8214 9.97222 17.874 9.90487C17.9266 9.83753 17.9631 9.75908 17.9808 9.6755C17.9986 9.59192 17.997 9.5054 17.9763 9.42251L17.5883 7.872ZM17.5883 4.49829H16.4286L16.7097 5.62286H6.74743V6.74743H17.4309C17.5163 6.74742 17.6006 6.72794 17.6774 6.69046C17.7542 6.65299 17.8214 6.59851 17.874 6.53116C17.9266 6.46381 17.9631 6.38537 17.9808 6.30179C17.9986 6.2182 17.997 6.13168 17.9763 6.04879L17.5883 4.49829ZM17.4309 0H0.562286C0.413158 0 0.270139 0.0592407 0.16469 0.16469C0.0592407 0.270139 0 0.413158 0 0.562286L0 2.81143C0 2.96056 0.0592407 3.10358 0.16469 3.20903C0.270139 3.31448 0.413158 3.37372 0.562286 3.37372H4.49829V5.62286H1.28271L1.56386 4.49829H0.404143L0.0175714 6.04879C-0.00313354 6.13162 -0.00470302 6.21808 0.012982 6.30161C0.0306671 6.38514 0.0671423 6.46355 0.119641 6.53088C0.172139 6.59822 0.239283 6.65271 0.315978 6.69023C0.392673 6.72775 0.476905 6.74731 0.562286 6.74743H4.49829V8.99658H1.28271L1.56386 7.872H0.404143L0.0175714 9.42251C-0.00313354 9.50534 -0.00470302 9.5918 0.012982 9.67533C0.0306671 9.75886 0.0671423 9.83727 0.119641 9.9046C0.172139 9.97193 0.239283 10.0264 0.315978 10.0639C0.392673 10.1015 0.476905 10.121 0.562286 10.1211H4.49829V14.7228C4.12312 14.8554 3.80693 15.1164 3.60559 15.4596C3.40424 15.8028 3.33072 16.2062 3.39801 16.5984C3.4653 16.9906 3.66907 17.3464 3.9733 17.6028C4.27754 17.8593 4.66265 18 5.06057 18C5.4585 18 5.84361 17.8593 6.14784 17.6028C6.45208 17.3464 6.65585 16.9906 6.72314 16.5984C6.79043 16.2062 6.7169 15.8028 6.51556 15.4596C6.31422 15.1164 5.99803 14.8554 5.62286 14.7228V3.37372H17.4309C17.58 3.37372 17.723 3.31448 17.8285 3.20903C17.9339 3.10358 17.9932 2.96056 17.9932 2.81143V0.562286C17.9932 0.413158 17.9339 0.270139 17.8285 0.16469C17.723 0.0592407 17.58 0 17.4309 0V0ZM5.06057 16.8686C4.94936 16.8686 4.84065 16.8356 4.74818 16.7738C4.65572 16.712 4.58365 16.6242 4.54109 16.5215C4.49853 16.4187 4.4874 16.3057 4.50909 16.1966C4.53079 16.0875 4.58434 15.9873 4.66298 15.9087C4.74162 15.8301 4.8418 15.7765 4.95088 15.7548C5.05995 15.7331 5.17301 15.7443 5.27575 15.7868C5.3785 15.8294 5.46631 15.9014 5.5281 15.9939C5.58988 16.0864 5.62286 16.1951 5.62286 16.3063C5.62286 16.4554 5.56362 16.5984 5.45817 16.7039C5.35272 16.8093 5.2097 16.8686 5.06057 16.8686ZM16.8686 2.24914H1.12457V1.12457H16.8686V2.24914Z"
              />
            </svg>
            <span class="mt-2 xl:mt-0"> {row[ocol] || ""} </span>
          </p>
        {/each}
      </div>

      {#if order.user}
        <div class="grid grid-cols-2 mt-8">
          <div class="flex items-center">
            <div class="relative">
              <div class="rounded-full w-6 h-6 md:w-8 md:h-8 bg-gray-200" />
              <span
                class="absolute top-0 right-0 inline-block w-3 h-3 bg-primary-red rounded-full"
              />
            </div>

            <p class="ml-2 text-gray-800 line-clamp-1">{row[order.user] || ""}</p>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
