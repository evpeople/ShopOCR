<template>
    <div>
        <ComHeaderQuery></ComHeaderQuery>
        <SeekMain></SeekMain>
    </div>
</template>

<script>
import ComHeaderQuery from '@/components/HeaderFollow.vue'
import SeekMain from '@/components/SeekMain.vue'

export default {
  name: 'HomePage',
  components: {
    ComHeaderQuery,
    SeekMain
  }
}
</script>

<style>

</style>