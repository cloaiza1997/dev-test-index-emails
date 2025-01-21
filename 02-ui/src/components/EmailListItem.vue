<script setup lang="ts">
import { getDate, replaceHighlight } from '@/utils/utils'
import AvatarImage from './AvatarImage.vue'
import type { EmailHighlightInterface } from '@/interfaces/email.interface'
</script>

<template>
  <div class="flex gap-2 p-2 w-full">
    <AvatarImage :text="data.email.from" />

    <div class="overflow-hidden flex flex-col justify-between w-full h-auto">
      <div class="flex items-center justify-between">
        <p class="text-sm truncate leading-tight" v-html="getFrom"></p>

        <p class="text-xs min-w-max ml-2 mr-1 leading-none">
          {{ getDate(data.email.date) }}
        </p>
      </div>

      <p class="text-left truncate leading-tight" v-html="getSubject"></p>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'EmailListItem',
  props: {
    data: {
      type: Object as () => EmailHighlightInterface,
      required: true,
    },
  },
  computed: {
    getFrom() {
      return replaceHighlight(this.data.email.from, this.data.highlight?.from)
    },

    getSubject() {
      return replaceHighlight(this.data.email.subject, this.data.highlight?.subject)
    },
  },
}
</script>
