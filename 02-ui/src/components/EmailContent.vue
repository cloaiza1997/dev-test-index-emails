<script setup lang="ts">
import { getDate, replaceHighlight } from '@/utils/utils'
import AvatarImage from './AvatarImage.vue'
import EmailLabel from './EmailLabel.vue'
import type { EmailHighlightInterface } from '@/interfaces/email.interface'
</script>

<template>
  <div class="overflow-hidden flex flex-col gap-4 w-full h-full">
    <section class="flex flex-col p-4 rounded-2xl text-sm sm:text-base bg-blue-400">
      <h3>{{ getDate(data.email.date) }}</h3>

      <h2 class="font-bold break-all" v-html="getSubject"></h2>
    </section>

    <section class="overflow-hidden flex flex-col flex-1 rounded-2xl h-full bg-blue-300">
      <div ref="emailContent" class="overflow-auto flex flex-col flex-1 h-full w-full">
        <div class="flex gap-2 p-4">
          <AvatarImage :text="data.email.from" />

          <div class="flex flex-col gap-1">
            <EmailLabel label="De" :text="data.email.from" :highlight="data.highlight?.from" />
            <EmailLabel label="Para" :text="data.email.to" :highlight="data.highlight?.to" />
            <EmailLabel label="CC" :text="data.email.cc" :highlight="data.highlight?.cc" />
            <EmailLabel label="BCC" :text="data.email.bcc" :highlight="data.highlight?.bcc" />
          </div>
        </div>

        <p v-html="getEmailContent" class="border-t p-4 text-sm sm:text-base break-all w-full"></p>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
export default {
  name: 'EmailContent',
  props: {
    data: {
      type: Object as () => EmailHighlightInterface,
      required: true,
    },
  },
  computed: {
    getEmailContent() {
      const emailContent = this.$refs.emailContent as HTMLElement

      emailContent?.scrollTo({ top: 0 })

      return replaceHighlight(this.data.email.body, this.data.highlight?.body)
    },

    getSubject() {
      return replaceHighlight(this.data.email.subject, this.data.highlight?.subject) || 'Sin asunto'
    },
  },
}
</script>
