<script setup lang="ts">
import EmailContent from './EmailContent.vue'
import EmailListItem from './EmailListItem.vue'

import { searchEmails } from '@/services/email.service'
import ButtonCircle from './ButtonCircle.vue'
import InputSearch from './InputSearch.vue'
import type { EmailInterface, PaginationInterface } from '@/interfaces/email.interface'
</script>

<template>
  <div class="overflow-hidden flex flex-1 gap-3 w-full h-full">
    <section
      class="flex-col gap-4 h-full w-full md:w-72 lg:w-1/3"
      :class="[emailSelected ? 'hidden md:flex' : 'flex']"
    >
      <InputSearch @init-search="onSearchEmails" />

      <div class="overflow-hidden rounded-2xl h-full w-full">
        <ul ref="emailList" class="overflow-auto flex flex-1 flex-col h-full w-full bg-red-200">
          <li
            v-for="(email, index) in emails"
            :key="email.messageId"
            :class="{ 'bg-red-400 email-selected': email.messageId === emailSelected?.messageId }"
            class="hover:bg-red-300"
          >
            <button
              class="w-full"
              :class="{ 'border-b': index + 1 < emails.length }"
              @click="setEmailSelected(index, email)"
            >
              <EmailListItem :email="email" />
            </button>
          </li>
        </ul>
      </div>

      <div class="flex items-center justify-between w-full">
        <template v-if="pagination.pages > 1">
          <ButtonCircle
            icon="arrow-left.svg"
            :disabled="pagination.prev == 0"
            :click="() => changePage(-1)"
          />

          <div class="flex flex-col gap-1 items-center justify-center text-sm text-center">
            <p>Página {{ paginationConfig.page }} de {{ pagination.pages }}</p>

            <p>
              {{ paginationConfig.itemsProcessed + 1 }}
              -
              {{ paginationConfig.itemsProcessed + pagination.count }}
              de
              {{ pagination.total }}
              emails
            </p>
          </div>

          <ButtonCircle
            icon="arrow-right.svg"
            :disabled="pagination.next == 0"
            :click="() => changePage(1)"
          />
        </template>

        <template v-else>
          <p class="text-sm text-center w-full">{{ pagination.total }} emails</p>
        </template>
      </div>
    </section>

    <section
      class="sm:flex flex-col flex-1 gap-4 h-full w-full"
      :class="[emailSelected ? 'flex' : 'hidden']"
    >
      <div v-if="emailSelected" class="flex items-center justify-between gap-2 py-1 my-[1px]">
        <p class="text-lg font-bold">
          Email:
          {{ paginationConfig.itemsProcessed + emailSelectedIdex + 1 }}
          /
          {{ pagination.total }}
        </p>

        <div class="flex items-center gap-2">
          <ButtonCircle
            icon="arrow-left.svg"
            :disabled="emailSelectedIdex == 0 && pagination.prev == 0"
            :click="() => handleEmailChange(-1)"
          />

          <ButtonCircle
            icon="arrow-right.svg"
            :disabled="emailSelectedIdex == emails.length - 1 && pagination.next == 0"
            :click="() => handleEmailChange(1)"
          />

          <ButtonCircle icon="close.svg" :click="() => setEmailSelected(0, null)" />
        </div>
      </div>

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
  data() {
    return {
      paginationConfig: {
        page: 1,
        limit: 100,
        itemsProcessed: 0,
      },
      emailSelectedIdex: 0,
      emailSelected: null as EmailInterface | null,
      emails: [] as EmailInterface[],
      skeleton: true,
      loading: false,
      pagination: {
        total: 0,
        count: 0,
        pages: 0,
        next: 0,
        prev: 0,
      } as PaginationInterface,
      term: '',
    }
  },
  mounted() {
    this.initEmails()
  },
  methods: {
    async changePage(page: number) {
      const newPage = this.paginationConfig.page + page

      if (newPage >= 1 || newPage <= this.pagination.pages) {
        this.paginationConfig.page = newPage

        await this.getEmails()

        this.scrollToTop()
      }
    },

    async getEmails() {
      this.loading = true

      const { page, limit } = this.paginationConfig

      const response = await searchEmails(this.term.toLowerCase(), page, limit)

      const { data } = response || {}

      if (data?.success) {
        this.emails = data.data.items
        this.pagination = data.data.pagination
      } else {
        this.emails = []
        this.pagination = {
          total: 0,
          count: 0,
          pages: 0,
          next: 0,
          prev: 0,
        }
      }

      this.setEmailSelected(0, null)
      this.paginationConfig.itemsProcessed = limit * page - limit
      this.loading = false
    },

    async handleEmailChange(index: number) {
      let newIndex = this.emailSelectedIdex + index

      if (newIndex < 0 || newIndex >= this.emails.length) {
        await this.changePage(index)
        newIndex = index > 0 ? 0 : this.emails.length - 1
      }

      const newEmail = this.emails[newIndex]
      this.emailSelectedIdex = newIndex
      this.setEmailSelected(newIndex, newEmail)

      setTimeout(() => {
        const emailList = this.$refs.emailList as HTMLUListElement | null
        const selectedItem = emailList?.querySelector('.email-selected') as HTMLLIElement | null

        selectedItem?.scrollIntoView({ behavior: 'smooth', block: 'center' })
      }, 0)
    },

    async initEmails() {
      this.skeleton = true

      await this.getEmails()

      this.skeleton = false
    },

    async onSearchEmails(term: string) {
      this.paginationConfig.page = 1
      this.term = term

      await this.getEmails()

      this.scrollToTop()
    },

    scrollToTop() {
      const emailList = this.$refs.emailList as HTMLUListElement | null

      if (emailList) {
        emailList.scrollTo({ top: 0, behavior: 'smooth' })
      }
    },

    setEmailSelected(emailSelectedIdex: number, email: EmailInterface | null) {
      this.emailSelectedIdex = emailSelectedIdex
      this.emailSelected = email
    },
  },
}
</script>
