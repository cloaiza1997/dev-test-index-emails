<script setup lang="ts">
import EmailContent from './EmailContent.vue'
import EmailListItem from './EmailListItem.vue'

import type { EmailInterface } from '@/interfaces/email.interface'
import InputSearch from './InputSearch.vue'
import ButtonCircle from './ButtonCircle.vue'
</script>

<template>
  <div class="overflow-hidden flex flex-1 gap-3 w-full h-full">
    <section class="flex flex-col gap-4 h-full w-full md:w-72 lg:w-1/3">
      <InputSearch />

      <div class="overflow-hidden rounded-2xl h-full w-full">
        <ul class="overflow-auto flex flex-1 flex-col h-full w-full bg-red-200">
          <li
            v-for="(email, index) in emails"
            :key="email.messageId"
            :class="{ 'bg-red-400': email.messageId === emailSelected?.messageId }"
            class="hover:bg-red-300"
          >
            <button
              class="w-full"
              :class="{ 'border-b': index + 1 < emails.length }"
              @click="selectEmail(email)"
            >
              <EmailListItem :email="email" />
            </button>
          </li>
        </ul>
      </div>

      <div class="flex items-center justify-between w-full">
        <ButtonCircle icon="arrow-left.svg" />

        <p>1-50 de 200</p>

        <ButtonCircle icon="arrow-right.svg" />
      </div>
    </section>

    <section class="flex flex-1 h-full sm:block w-full" :class="{ hidden: true }">
      <EmailContent v-if="emailSelected" :email="emailSelected" />

      <div v-else class="flex items-center justify-center rounded-2xl w-full h-full bg-gray-400">
        <h3>Selecciona un correo para visualizar su contenido.</h3>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
export default {
  name: 'EmailList',
  props: {
    emails: {
      type: Array as () => EmailInterface[],
      required: true,
    },
  },
  data() {
    return {
      emailSelected: null as EmailInterface | null,
    }
  },
  methods: {
    selectEmail(email: EmailInterface) {
      this.emailSelected = email
    },
  },
}
</script>
