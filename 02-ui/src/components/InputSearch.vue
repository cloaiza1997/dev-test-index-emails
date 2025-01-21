<script setup lang="ts">
import ButtonCircle from './ButtonCircle.vue'
import ImageIcon from './ImageIcon.vue'
</script>

<template>
  <div
    class="flex gap-2 items-center border border-gray-400 focus:border-gray-700 rounded-full overflow-hidden pl-5 pr-1 py-1 w-full min-h-[50px]"
  >
    <ImageIcon icon="search.svg" alt="search" class-name="w-4 h-4" />

    <input
      ref="inputSearch"
      type="text"
      placeholder="Buscar correos"
      class="w-full h-10"
      :autofocus="true"
      v-model="term"
      @keyup.enter="onSearch"
    />

    <ButtonCircle v-if="term.trim()" icon="close.svg" :click="clean" :disabled="disabled" />
  </div>
</template>

<script lang="ts">
export default {
  name: 'InputSearch',
  props: {
    disabled: Boolean,
  },
  data() {
    return {
      term: '',
    }
  },
  methods: {
    clean() {
      this.term = ''
      const inputSearch = this.$refs.inputSearch as HTMLInputElement

      if (inputSearch) {
        inputSearch.focus()
      }

      this.onSearch()
    },
    onSearch() {
      this.$emit('initSearch', this.term.trim())
    },
  },
}
</script>
