<template>
  <span
    class="flex items-center justify-center leading-none font-bold text-lg border rounded-full min-w-10 min-h-10 max-w-10 max-h-10 bg-gray-300"
  >
    {{ text?.[0]?.toUpperCase() }}
  </span>
</template>

<script lang="ts">
export default {
  name: 'AvatarImage',
  props: {
    text: {
      type: String as () => string | undefined,
      required: true,
    },
  },
}
</script>
