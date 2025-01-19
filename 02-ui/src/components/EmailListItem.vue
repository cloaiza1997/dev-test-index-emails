<script setup lang="ts">
import { getDate } from '@/utils/utils'
import type { EmailInterface } from '@/interfaces/email.interface'

import AvatarImage from './AvatarImage.vue'
</script>

<template>
  <div class="flex gap-2 p-2 w-full">
    <AvatarImage :text="email.from" />

    <div class="overflow-hidden flex flex-col gap-2 justify-between w-full h-auto">
      <div class="flex items-center justify-between">
        <p class="text-sm truncate leading-none">{{ email.from }}</p>
        <p class="text-xs min-w-max pl-2 pr-1 leading-none">{{ getDate(email.date) }}</p>
      </div>

      <p class="text-left truncate leading-tight">{{ email.subject }}</p>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'EmailListItem',
  props: {
    email: {
      type: Object as () => EmailInterface,
      required: true,
    },
  },
  data() {
    return {}
  },
}
</script>
