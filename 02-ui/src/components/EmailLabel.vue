<script setup lang="ts">
import { replaceHighlight } from '@/utils/utils'
</script>

<template>
  <p v-if="text" class="text-sm sm:text-base">
    <span class="text-xs sm:text-sm font-bold">{{ label }}:</span> <span v-html="getText"></span>
  </p>
</template>

<script lang="ts">
export default {
  name: 'EmailLabel',
  props: {
    highlight: {
      type: Array as () => string[],
    },
    label: {
      type: String,
      required: true,
    },
    text: {
      type: String,
    },
  },
  computed: {
    getText() {
      return replaceHighlight(this.text, this.highlight)
    },
  },
}
</script>
