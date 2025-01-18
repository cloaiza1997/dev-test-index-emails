<template>
  <img :src="`/assets/icons/${icon}`" :alt="alt" :class="className" />
</template>

<script lang="ts">
export default {
  name: 'ImageIcon',
  props: {
    icon: {
      type: String,
      required: true,
    },
    alt: {
      type: String,
    },
    className: {
      type: String,
    },
  },
}
</script>
