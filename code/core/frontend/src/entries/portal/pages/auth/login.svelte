<script lang="ts">
  export let emailuser = "";
  export let password = "";
  export let doLogin = (user_ident, password) => {};
</script>

<div
  class="min-w-screen min-h-screen bg-blue-200 flex items-center justify-center px-5 py-5"
>
  <div
    class="bg-gray-50 text-gray-500 rounded-3xl shadow-xl w-full overflow-hidden"
    style="max-width:1000px"
  >
    <div class="md:flex w-full">
      <div class="hidden md:block w-1/2 bg-indigo-500 py-10 px-10">
            <svg data-name="Layer 1" viewBox="0 0 744.85 747.08" xmlns="http://www.w3.org/2000/svg">
            <path transform="translate(-227.58 -76.461)" d="m299.2 705.81-6.56-25.872a335.97 335.97 0 0 0-35.643-12.788l-0.828 12.024-3.358-13.247c-15.021-4.2939-25.24-6.183-25.24-6.183s13.8 52.489 42.754 92.617l33.734 5.926-26.207 3.779a135.93 135.93 0 0 0 11.719 12.422c42.115 39.092 89.024 57.028 104.77 40.06s-5.625-62.412-47.74-101.5c-13.056-12.119-29.457-21.844-45.875-29.5z" fill="#f2f2f2" data-name="Path 1"/>
            <path transform="translate(-227.58 -76.461)" d="m361.59 677.71 7.758-25.538a335.94 335.94 0 0 0-23.9-29.371l-6.924 9.865 3.972-13.076c-10.641-11.436-18.412-18.335-18.412-18.335s-15.315 52.067-11.275 101.38l25.815 22.51-24.392-10.312a135.92 135.92 0 0 0 3.614 16.694c15.846 55.234 46.731 94.835 68.983 88.451s27.446-56.335 11.6-111.57c-4.912-17.123-13.926-33.926-24.023-48.965z" fill="#f2f2f2" data-name="Path 2"/>
            <path transform="translate(-227.58 -76.461)" d="m747.33 253.44h-4.092v-112.1a64.883 64.883 0 0 0-64.883-64.883h-237.51a64.883 64.883 0 0 0-64.883 64.883v615a64.883 64.883 0 0 0 64.883 64.883h237.51a64.883 64.883 0 0 0 64.882-64.883v-423.1h4.092z" fill="#e6e6e6" data-name="Path 22"/>
            <path transform="translate(-227.58 -76.461)" d="m680.97 93.336h-31a23.02 23.02 0 0 1-21.316 31.714h-136.06a23.02 23.02 0 0 1-21.314-31.714h-28.956a48.454 48.454 0 0 0-48.454 48.454v614.11a48.454 48.454 0 0 0 48.454 48.454h238.65a48.454 48.454 0 0 0 48.454-48.454v-614.11a48.454 48.454 0 0 0-48.454-48.453z" fill="#fff" data-name="Path 23"/>
            <path transform="translate(-227.58 -76.461)" d="m531.23 337.96a24.437 24.437 0 0 1 12.23-21.174 24.45 24.45 0 1 0 0 42.345 24.434 24.434 0 0 1-12.23-21.171z" fill="#ccc" data-name="Path 6"/>
            <path transform="translate(-227.58 -76.461)" d="m561.97 337.96a24.436 24.436 0 0 1 12.23-21.174 24.45 24.45 0 1 0 0 42.345 24.434 24.434 0 0 1-12.23-21.171z" fill="#ccc" data-name="Path 7"/>
            <circle cx="364.43" cy="261.5" r="24.45" fill="#6c63ff" data-name="Ellipse 1"/>
            <path transform="translate(-227.58 -76.461)" d="m632.87 414.33h-142.5a5.123 5.123 0 0 1-5.117-5.117v-142.5a5.123 5.123 0 0 1 5.117-5.117h142.5a5.123 5.123 0 0 1 5.117 5.117v142.5a5.123 5.123 0 0 1-5.117 5.117zm-142.5-150.69a3.073 3.073 0 0 0-3.07 3.07v142.5a3.073 3.073 0 0 0 3.07 3.07h142.5a3.073 3.073 0 0 0 3.07-3.07v-142.5a3.073 3.073 0 0 0-3.07-3.07z" fill="#ccc" data-name="Path 8"/>
            <rect x="218.56" y="447.1" width="218.55" height="2.047" fill="#ccc" data-name="Rectangle 1"/>
            <circle cx="225.46" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 2"/>
            <rect x="218.56" y="516.12" width="218.55" height="2.047" fill="#ccc" data-name="Rectangle 2"/>
            <circle cx="225.46" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 3"/>
            <path transform="translate(-227.58 -76.461)" d="m660.69 671.17h-69.068a4.5049 4.5049 0 0 1-4.5-4.5v-24.208a4.5049 4.5049 0 0 1 4.5-4.5h69.068a4.5049 4.5049 0 0 1 4.5 4.5v24.208a4.5049 4.5049 0 0 1-4.5 4.5z" fill="#6c63ff"/>
            <circle cx="247.98" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 7"/>
            <circle cx="270.49" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 8"/>
            <circle cx="247.98" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 9"/>
            <circle cx="270.49" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 10"/>
            <path transform="translate(-227.58 -76.461)" d="m969.64 823.54h-717.99c-1.537 0-2.782-0.546-2.782-1.218s1.245-1.219 2.782-1.219h717.99c1.536 0 2.782 0.546 2.782 1.219s-1.246 1.218-2.782 1.218z" fill="#3f3d56" data-name="Path 88"/>
            <path transform="translate(-227.58 -76.461)" d="m792.25 565.92a10.094 10.094 0 0 1 1.4108 0.78731l44.852-19.143 1.6009-11.815 17.922-0.10956-1.0587 27.098-59.2 15.656a10.608 10.608 0 0 1-0.44749 1.2084 10.235 10.235 0 1 1-5.0795-13.682z" fill="#ffb8b8"/>
            <polygon points="636.98 735.02 624.72 735.02 618.89 687.73 636.98 687.73" fill="#ffb8b8"/>
            <path d="m615.96 731.52h23.644v14.887h-38.531a14.887 14.887 0 0 1 14.887-14.887z" fill="#2f2e41"/>
            <polygon points="684.66 731.56 672.46 732.76 662.02 686.27 680.02 684.5" fill="#ffb8b8"/>
            <path transform="translate(-303.01 15.291) rotate(-5.6253)" d="m891.69 806.13h23.644v14.887h-38.531a14.887 14.887 0 0 1 14.887-14.887z" fill="#2f2e41"/>
            <circle cx="640.39" cy="384.57" r="24.561" fill="#ffb8b8"/>
            <path transform="translate(-227.58 -76.461)" d="m849.56 801.92a4.4709 4.4709 0 0 1-4.415-3.6973c-6.3457-35.226-27.088-150.41-27.584-153.6a1.4268 1.4268 0 0 1-0.01562-0.22168v-8.5879a1.489 1.489 0 0 1 0.27929-0.87207l2.7402-3.8379a1.4784 1.4784 0 0 1 1.1436-0.625c15.622-0.73242 66.784-2.8789 69.256 0.209 2.4824 3.1035 1.6055 12.507 1.4043 14.36l0.00977 0.19336 22.985 147a4.5124 4.5124 0 0 1-3.7148 5.1348l-14.356 2.3652a4.5213 4.5213 0 0 1-5.0254-3.0928c-4.4404-14.188-19.329-61.918-24.489-80.387a0.49922 0.49922 0 0 0-0.98047 0.13868c0.25781 17.605 0.88086 62.523 1.0957 78.037l0.02344 1.6709a4.5181 4.5181 0 0 1-4.0928 4.5361l-13.844 1.2578c-0.14066 0.01268-0.28131 0.01854-0.41995 0.01854z" fill="#2f2e41"/>
            <path transform="translate(-227.58 -76.461)" d="m852.38 495.25c-4.2863 2.548-6.8512 7.2304-8.3228 11.995a113.68 113.68 0 0 0-4.8844 27.159l-1.5555 27.6-19.255 73.17c16.689 14.121 26.315 10.912 48.78-0.63879s25.032 3.8512 25.032 3.8512l4.4924-62.258 6.4184-68.032a30.164 30.164 0 0 0-4.8614-4.6742 49.658 49.658 0 0 0-42.442-8.9954z" fill="#fff" data-name="Path 99"/>
            <path transform="translate(-227.58 -76.461)" d="m846.13 580.7a10.526 10.526 0 0 1 1.5006 0.70389l44.348-22.197 0.736-12.026 18.294-1.2613 0.98041 27.413-59.266 19.599a10.496 10.496 0 1 1-6.5933-12.232z" fill="#ffb8b8"/>
            <path transform="translate(-227.58 -76.461)" d="m902.77 508.41c10.912 3.8512 12.834 45.574 12.834 45.574-12.837-7.0604-28.241 4.4932-28.241 4.4932s-3.2092-10.912-7.0603-25.032a24.53 24.53 0 0 1 5.1344-23.106s6.4223-5.7819 17.334-1.9284z" fill="#fff" data-name="Path 101"/>
            <path transform="translate(-227.58 -76.461)" d="m889.99 467.53c-3.06-2.4484-7.2352 2.0017-7.2352 2.0017l-2.4484-22.033s-15.301 1.8329-25.094-0.61161-11.323 8.8751-11.323 8.8751a78.58 78.58 0 0 1-0.30582-13.771c0.61158-5.5084 8.5684-11.017 22.645-14.689s21.421 12.241 21.421 12.241c9.7936 4.8958 5.3994 30.436 2.3394 27.988z" fill="#2f2e41" data-name="Path 102"/>
            </svg>
      </div>
      <div class="w-full md:w-1/2 py-10 px-5 md:px-10">
        <div class="w-full px-6 py-8">
          <h2
            class="text-2xl font-semibold text-center text-gray-700 dark:text-white"
          >
            Temphia
          </h2>

          <p class="text-xl text-center text-gray-600 dark:text-gray-200">
            User Login
          </p>

          <div class="mt-4">
            <label
              class="block mb-2 text-sm font-medium text-gray-600 dark:text-gray-200"
              for="LoggingEmailAddress">Email Address / Username</label
            >
            <input
              id="LoggingEmailAddress"
              class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring"
              bind:value={emailuser}
              type="email"
            />
          </div>

          <div class="mt-4">
            <div class="flex justify-between">
              <label
                class="block mb-2 text-sm font-medium text-gray-600 dark:text-gray-200"
                for="loggingPassword">Password</label
              >
              <a
                href="#"
                class="text-xs text-gray-500 dark:text-gray-300 hover:underline"
                >Forget Password?</a
              >
            </div>

            <input
              id="loggingPassword"
              class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring"
              type="password"
              bind:value={password}
            />
          </div>

          <div class="mt-8">
            <button
              on:click={() => doLogin(emailuser, password)}
              class="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-gray-700 rounded hover:bg-gray-600 focus:outline-none focus:bg-gray-600"
            >
              Login
            </button>
          </div>

          <div class="flex items-center justify-between mt-4">
            <span class="w-1/5 border-b dark:border-gray-600 lg:w-1/4" />

            <a
              href="#"
              class="text-xs text-center text-gray-500 uppercase dark:text-gray-400 hover:underline"
              >or</a
            >

            <span class="w-1/5 border-b dark:border-gray-400 lg:w-1/4" />
          </div>

          <a
            href="#"
            class="flex items-center justify-center mt-4 text-gray-600 rounded-lg shadow-md dark:bg-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-600"
          >
            <div class="px-4 py-3" />

            <span class="w-5/6 px-4 py-3 font-bold text-center"
              >Other Sign In Options
            </span>
          </a>

          <div class="flex items-center justify-between mt-4">
            <span class="w-1/5 border-b dark:border-gray-600 md:w-1/4" />

            <a
              href="#"
              class="text-xs text-gray-500 uppercase dark:text-gray-400 hover:underline"
              >or sign up</a
            >

            <span class="w-1/5 border-b dark:border-gray-600 md:w-1/4" />
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

