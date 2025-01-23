<template>
  <span
    class="overflow-hidden flex items-center justify-center leading-none font-bold text-lg border rounded-full min-w-12 min-h-12 max-w-12 max-h-12 bg-white"
  >
    <img
      v-if="text"
      :src="'https://robohash.org/' + text"
      :alt="text"
      class="object-contain w-full h-full"
    />
  </span>
</template>

<script lang="ts">
export default {
  name: 'AvatarImage',
  props: {
    text: {
      type: [String, undefined],
      required: true,
    },
  },
}
</script>
