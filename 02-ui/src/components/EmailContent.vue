<script setup lang="ts">
import { getDate } from '@/utils/utils'
import AvatarImage from './AvatarImage.vue'
import EmailLabel from './EmailLabel.vue'
import type { EmailInterface } from '@/interfaces/email.interface'
</script>

<template>
  <div class="overflow-hidden flex flex-col gap-4 w-full h-full">
    <section class="flex flex-col p-4 rounded-2xl text-sm sm:text-base bg-blue-400">
      <h3>{{ getDate(email.date) }}</h3>

      <h2 class="font-bold break-all">{{ email.subject || 'Sin asunto' }}</h2>
    </section>

    <section class="overflow-hidden flex flex-col flex-1 rounded-2xl h-full bg-blue-300">
      <div ref="emailContent" class="overflow-auto flex flex-col flex-1 h-full w-full">
        <div class="flex gap-2 p-4">
          <AvatarImage :text="email?.from" />

          <div class="flex flex-col gap-1">
            <EmailLabel label="De" :text="email.from" />
            <EmailLabel label="Para" :text="email.to" />
            <EmailLabel label="CC" :text="email.cc" />
            <EmailLabel label="BCC" :text="email.bcc" />
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
    email: {
      type: Object as () => EmailInterface,
      required: true,
    },
  },
  computed: {
    getEmailContent() {
      const emailContent = this.$refs.emailContent as HTMLElement

      emailContent?.scrollTo({ top: 0 })

      return this.email?.body.replace(/\n/gim, '</br>')
    },
  },
}
</script>
