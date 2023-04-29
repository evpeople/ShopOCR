<template>
    <div>
        <ComHeaderQuery></ComHeaderQuery>
        <ComMain></ComMain>
    </div>
</template>

<script>
import ComHeaderQuery from '@/components/HeaderFollow.vue'
import ComMain from '@/components/HomeMain.vue'

export default {
  name: 'HomePage',
  components: {
    ComHeaderQuery,
    ComMain
  }
}
</script>

<style>

</style>