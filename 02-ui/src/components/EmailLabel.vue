<template>
  <p v-if="text" class="text-sm sm:text-base">
    <span class="text-xs sm:text-sm font-bold">{{ label }}:</span> {{ text }}
  </p>
</template>

<script lang="ts">
export default {
  name: 'EmailLabel',
  props: {
    label: {
      type: String,
      required: true,
    },
    text: {
      type: String,
    },
  },
}
</script>
