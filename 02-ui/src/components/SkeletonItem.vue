<template>
  <div class="skeleton"></div>
</template>

<style scoped>
.skeleton {
  @apply flex w-full h-full;
  transition: 1.5s;
  animation: skeleton 1.5s infinite linear;
}

@keyframes skeleton {
  0% {
    opacity: 1;
    @apply bg-gray-300;
  }
  50% {
    opacity: 0.8;
    @apply bg-gray-200;
  }
  100% {
    @apply bg-gray-300;
    opacity: 1;
  }
}
</style>

<script lang="ts">
export default {
  name: 'SkeletonItem',
}
</script>
