<script setup lang="ts">
import EmailList from '@/components/EmailList.vue'
</script>

<template>
  <section class="overflow-hidden flex gap-4 items-center m-4">
    <img src="/assets/icons/logo.png" alt="Logo" class="object-contain w-16 h-16" />

    <h1 class="font-bold text-4xl leading-none">MamuroEmail</h1>
  </section>

  <EmailList />
</template>

<script lang="ts">
export default {
  name: 'EmailHomeView',
}
</script>
