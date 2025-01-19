<script setup lang="ts">
import ImageIcon from './ImageIcon.vue'
</script>

<template>
  <button
    class="flex items-center justify-center rounded-full min-w-10 w-10 min-h-10 h-10"
    :class="[
      disabled
        ? 'cursor-not-allowed pointer-events-none bg-gray-300'
        : 'bg-blue-300 hover:bg-blue-500',
    ]"
    :disabled="disabled"
    @:click="_onClick"
  >
    <ImageIcon :icon="icon" alt="button" class-name="w-4 h-4" />
  </button>
</template>

<script lang="ts">
export default {
  name: 'ButtonCircle',
  props: {
    disabled: {
      type: Boolean,
    },
    icon: {
      type: String,
      required: true,
    },
    click: {
      type: Function,
      required: true,
      default: () => {},
    },
  },
  methods: {
    _onClick() {
      if (!this.disabled) {
        this.click()
      }
    },
  },
}
</script>
