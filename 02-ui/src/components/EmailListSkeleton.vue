<script setup lang="ts">
import SkeletonItem from './SkeletonItem.vue'
</script>

<template>
  <ul class="flex flex-1 flex-col h-full w-full bg-red-200">
    <li
      v-for="(_, index) in emails"
      :key="index"
      class="flex gap-2 p-2 w-full"
      :class="{ 'border-b': index + 1 < emails.length }"
    >
      <span class="overflow-hidden rounded-full min-w-10 min-h-10 max-w-10 max-h-10">
        <SkeletonItem />
      </span>

      <div class="flex flex-col justify-between w-full h-auto">
        <div class="flex items-center justify-between">
          <p class="overflow-hidden flex flex-1 rounded-full my-[1.25px] w-1/2 h-[14px]">
            <SkeletonItem />
          </p>

          <p class="overflow-hidden rounded-full ml-2 mr-1 w-[74px] h-[12px]">
            <SkeletonItem />
          </p>
        </div>

        <p class="overflow-hidden rounded-full w-3/5 my-[2px] h-[16px]">
          <SkeletonItem />
        </p>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
export default {
  name: 'EmailListSkeleton',
  data() {
    return {
      emails: new Array(20).fill(''),
    }
  },
}
</script>
