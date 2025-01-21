<script setup lang="ts">
import EmailListSkeleton from './EmailListSkeleton.vue'
import SkeletonItem from './SkeletonItem.vue'
</script>

<template>
  <div class="overflow-hidden flex flex-1 gap-3 w-full h-full">
    <section class="flex flex-col gap-4 h-full w-full md:w-72 lg:w-1/3">
      <div class="overflow-hidden rounded-full w-full min-h-[50px]">
        <SkeletonItem />
      </div>

      <div class="overflow-hidden rounded-2xl h-full w-full">
        <EmailListSkeleton />
      </div>

      <div class="overflow-hidden rounded-full w-full min-h-11">
        <SkeletonItem />
      </div>
    </section>

    <section class="hidden sm:flex flex-1 h-full w-full">
      <div class="overflow-hidden rounded-2xl w-full h-full">
        <SkeletonItem />
      </div>
    </section>
  </div>
</template>

<script lang="ts">
export default {
  name: 'EmailViewSkeleton',
  data() {
    return {
      emails: new Array(10).fill(''),
    }
  },
}
</script>
