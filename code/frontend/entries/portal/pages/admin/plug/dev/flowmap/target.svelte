<script lang="ts">
    export let data; 

//     {
// 	"0": {
// 		"id": 1,
// 		"name": "Documents",
// 		"target_type": "user_group_app",
// 		"target": "super_admin",
// 		"plug_id": "chc50lom4q7efu3enuq0",
// 		"agent_id": "default"
// 	}
// }

</script>


<div class="flex justify-center bg-blue-400 opacity-50">
  <div class="inline-flex items-start">
    <h3 class="uppercase">{data["name"] || ""}</h3>
    
    <span class="ml-1 mt-1 rounded bg-gray-100 p-0.5 text-xs text-gray-800"
      >{data["target_type"]}</span
    >
  </div>
</div>


