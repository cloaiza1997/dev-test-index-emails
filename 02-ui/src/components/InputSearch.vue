<script setup lang="ts">
import ButtonCircle from './ButtonCircle.vue'
import ImageIcon from './ImageIcon.vue'
</script>

<template>
  <div
    class="flex gap-2 items-center border border-gray-400 focus:border-gray-700 rounded-full pl-5 pr-1 py-1 w-full"
  >
    <ImageIcon icon="search.svg" alt="search" class-name="w-4 h-4" />

    <input type="text" placeholder="Buscar correos" class="w-full h-full" autofocus="true" />

    <ButtonCircle icon="close.svg" />
  </div>
</template>

<script lang="ts">
export default {
  name: 'InputSearch',
}
</script>
