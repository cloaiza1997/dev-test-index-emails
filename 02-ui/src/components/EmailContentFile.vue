<script setup lang="ts">
import { replaceHighlight } from '@/utils/utils'
import type { EmailHighlightInterface } from '@/interfaces/email.interface'
</script>

<template>
  <section class="overflow-hidden flex flex-1 rounded-2xl w-full h-full bg-blue-300">
    <p
      ref="emailContent"
      v-html="getFileContent"
      class="overflow-auto p-4 font-serif text-sm sm:text-base break-all w-full"
    ></p>
  </section>
</template>

<script lang="ts">
export default {
  name: 'EmailContentFile',
  props: {
    data: {
      type: Object as () => EmailHighlightInterface | null,
      required: true,
    },
  },
  computed: {
    getFileContent() {
      const emailContent = this.$refs.emailContent as HTMLElement

      emailContent?.scrollTo({ top: 0 })

      const { email, highlight } = this.data ?? {}

      const fileContent = [
        this.getHeader('Message ID', email?.messageId),
        this.getHeader('Date', email?.date),
        this.getHeader('From', email?.from),
        this.getHeader('To', email?.to),
        this.getHeader('Cc', email?.cc),
        this.getHeader('Bcc', email?.bcc),
        this.getHeader('Subject', email?.subject),
        this.getHeader('X-From', email?.xFrom),
        this.getHeader('X-To', email?.xTo),
        this.getHeader('X-Cc', email?.xCc),
        this.getHeader('X-Bcc', email?.xBcc),
        this.getHeader('X-Folder', email?.xFolder),
        this.getHeader('X-Origin', email?.xOrigin),
        this.getHeader('X-FileName', email?.xFileName),
        '',
        email?.body,
      ]

      const allHighlight: string[] = []

      Object.values(highlight ?? {}).forEach((value) => {
        allHighlight.push(...value)
      })

      return replaceHighlight(fileContent.join('</br>'), allHighlight)
    },
  },
  methods: {
    getHeader(header: string, value = '') {
      return `<strong>${header}:</strong> ${value}`
    },
  },
}
</script>
