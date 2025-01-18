<script setup lang="ts">
import ImageIcon from './ImageIcon.vue'
</script>

<template>
  <button
    class="bg-gray-400 flex items-center justify-center rounded-full min-w-10 w-10 min-h-10 h-10"
  >
    <ImageIcon :icon="icon" alt="button" class-name="w-4 h-4" />
  </button>
</template>

<script lang="ts">
export default {
  name: 'ButtonCircle',
  props: {
    icon: {
      type: String,
      required: true,
    },
    onClick: {
      type: Function as () => void,
      required: false,
      default: () => () => {},
    },
  },
}
</script>
