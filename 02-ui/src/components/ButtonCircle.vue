<script setup lang="ts">
import ImageIcon from './ImageIcon.vue'
</script>

<template>
  <button
    class="flex items-center justify-center rounded-full min-w-10 w-10 min-h-10 h-10 shadow-2xl"
    :class="[
      disabled
        ? 'cursor-not-allowed pointer-events-none bg-gray-400 opacity-50'
        : pressed
          ? 'bg-orange-300 hover:bg-orange-400'
          : close
            ? 'bg-red-300 hover:bg-red-400'
            : 'bg-primary-400 hover:bg-primary-600',
    ]"
    :disabled="disabled"
    :title="title"
    @:click="_onClick"
  >
    <ImageIcon :icon="icon" alt="button" class-name="w-4 h-4" />
  </button>
</template>

<script lang="ts">
export default {
  name: 'ButtonCircle',
  props: {
    click: {
      type: Function,
      required: true,
      default: () => {},
    },
    disabled: {
      type: Boolean,
    },
    icon: {
      type: String,
      required: true,
    },
    pressed: {
      type: Boolean,
    },
    close: {
      type: Boolean,
    },
    title: {
      type: String,
      required: true,
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
