<script>
import { ref } from 'vue'

export default {
  setup() {
    const data = ref(null)
    const isLoading = ref(false)

    async function fetchData() {
      isLoading.value = true
      const response = await fetch('https://api-node-uqw7du2n4q-uc.a.run.app/logs')
      const json = await response.json()
      data.value = json.map(log => ({
        ...log,
        createdat: removeAfterUTC(log.createdat)
      }))
      isLoading.value = false
    }

    function removeAfterUTC(dateString) {
      const index = dateString.indexOf('+0000');
      if (index !== -1) {
        return dateString.substring(0, index).trim();
      }
      return dateString;
    }

    return { data, isLoading, fetchData }
  }
}
</script>

<template>
  <div class="flex flex-col p-16 gap-14 w-screen h-screen bg-slate-200">
    <h1 class="mx-auto font-mono font-semibold text-3xl">Proyecto 2 - SO1 - Alvaro García - 202109567</h1>

    <div class="flex flex-row gap-9 h-full ">
      <div class="bg-black text-white p-4 overflow-auto h-full w-full rounded-lg overflow-y-auto">
        <pre id="logs">
          Logs: 
        </pre>

        <pre v-for="(log, index) in data" :key="log.createdat" >
          {{ index + 1 }}.  {{ log.createdat }} - {{ log.data.name }} {{ log.data.album }} {{ log.data.year }} {{ log.data.rank }}
        </pre> 
      </div>

      <div class="basis-1/4">
        <button 
        @click="fetchData"
        class="bg-blue-400 hover:bg-blue-500 text-slate-100 font-bold py-2 px-4 rounded-lg">
          <span v-if="isLoading">
            <i class="fa fa-spinner fa-spin"></i> <!-- Spinner icon -->
          </span>
          <span v-else>
            Actualizar
          </span>
        </button>
      </div>
    </div>
  </div>
</template>