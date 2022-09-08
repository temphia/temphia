<script lang="ts">
  import { LeafletMap, Marker, Popup, TileLayer } from "svelte-leafletjs";
  export let lat: number = 27.7116;
  export let lng: number = 85.3124;
  export let callback: (_lat: number, _lng: number) => void;


  const mapOptions = {
    center: [lat, lng],
    zoom: 11,
  };
  const tileUrl = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png";
  const tileLayerOptions = {
    minZoom: 0,
    maxZoom: 20,
    maxNativeZoom: 19,
    attribution: "© OpenStreetMap contributors",
  };
  const current = [lat, lng];
  const maybe = [...current];

  function updateCurrent(event) {
    const l = event.detail.latlng;
    maybe[0] = l.lat;
    maybe[1] = l.lng;
  }

  function mark() {
    current[0] = maybe[0];
    current[1] = maybe[1];
  }

  function select() {
    callback(current[0], current[1]);
    close();
  }
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
    integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
    crossorigin=""
  />
  <script
    src="https://unpkg.com/leaflet@1.6.0/dist/leaflet.js"
    integrity="sha512-gZwIG9x3wUXg2hdXF6+rVkLF/0Vi9U8D2Ntg4Ga5I5BZpVkVxlJWbSQtXPSiUTtC0TjtGOmxa1AJPuV0CPthew=="
    crossorigin="">
  </script>
</svelte:head>

<div class="w-full h-80 shadow rounded">
  <LeafletMap options={mapOptions} events={["click"]} on:click={updateCurrent}>
    <TileLayer url={tileUrl} options={tileLayerOptions} />
    <Marker latLng={current}>
      <Popup>Select a location</Popup>
    </Marker>
  </LeafletMap>
</div>

<div class="flex mt-2 justify-end">
  <button
    class="px-2 py-1 mr-2 shadow bg-blue-500 rounded text-white font-semibold"
    on:click={mark}
  >
    Mark
  </button>

  <button
    class="px-2 py-1 shadow bg-green-500 rounded text-white font-semibold"
    on:click={select}
  >
    Select
  </button>
</div>
