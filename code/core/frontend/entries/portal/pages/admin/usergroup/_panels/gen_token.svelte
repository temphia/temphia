<div class="w-full flex flex-col">
  <label>
    Token Type
    <select class="p-2 mb-8 rounded">
      <option>Login</option>
      <option>Signup</option>
    </select>
  </label>
  <hr>

  <label>
    Bind Domain
    <input
      type="text"
      placeholder="optional"
      class="mb-3 py-3 px-4 w-full border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
    />
  </label>

  <label>
    Expiry

    <select class="p-2 m-1 rounded">
      <option>Week</option>
      <option>2 Weeks</option>
      <option>3 Weeks</option>
      <option>Month</option>
      <option>2 Month</option>
      <option>3 Month</option>
      <option>6 Month</option>
    </select>
  </label>

  <button
    class="w-full bg-green-400 mt-8 mb-4 text-white p-3 rounded-lg font-semibold text-lg"
    >Generate</button
  >
</div>
